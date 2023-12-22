// Copyright (C) 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mali

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/gapid/core/log"
	"github.com/google/gapid/core/os/device"
	"github.com/google/gapid/gapis/api/sync"
	"github.com/google/gapid/gapis/perfetto"
	"github.com/google/gapid/gapis/service"
	"github.com/google/gapid/gapis/trace/android/profile"
)

var (
	queueSubmitQuery = "" +
		"SELECT submission_id, command_buffer FROM gpu_slice s JOIN track t ON s.track_id = t.id WHERE s.name = 'vkQueueSubmit' AND t.name = 'Vulkan Events' ORDER BY submission_id"
	counterTracksQuery = "" +
		"SELECT id, name, unit, description FROM gpu_counter_track ORDER BY id"
	countersQueryFmt = "" +
		"SELECT ts, value FROM counter c WHERE c.track_id = %d ORDER BY ts"
)

func ProcessProfilingData(ctx context.Context, processor *perfetto.Processor,
	desc *device.GpuCounterDescriptor, handleMapping map[uint64][]service.VulkanHandleMappingItem,
	syncData *sync.Data, data *profile.ProfilingData) error {

	err := processGpuSlices(ctx, processor, handleMapping, syncData, data)
	if err != nil {
		log.Err(ctx, err, "Failed to get GPU slices")
	}
	data.Counters, err = processCounters(ctx, processor, desc)
	if err != nil {
		log.Err(ctx, err, "Failed to get GPU counters")
	}
	data.ComputeCounters(ctx)
	updateCounterGroups(ctx, data)
	return nil
}

func processGpuSlices(ctx context.Context, processor *perfetto.Processor,
	handleMapping map[uint64][]service.VulkanHandleMappingItem, syncData *sync.Data,
	data *profile.ProfilingData) (err error) {
	data.Slices, err = profile.ExtractSliceData(ctx, processor)
	if err != nil {
		return log.Errf(ctx, err, "Extracting slice data failed")
	}

	queueSubmitQueryResult, err := processor.Query(queueSubmitQuery)
	if err != nil {
		return log.Errf(ctx, err, "SQL query failed: %v", queueSubmitQuery)
	}
	queueSubmitColumns := queueSubmitQueryResult.GetColumns()
	queueSubmitIds := queueSubmitColumns[0].GetLongValues()
	queueSubmitCommandBuffers := queueSubmitColumns[1].GetLongValues()
	submissionOrdering := make(map[int64]int)

	order := 0
	for i, v := range queueSubmitIds {
		if queueSubmitCommandBuffers[i] == 0 {
			// This is a spurious submission. See b/150854367
			log.W(ctx, "Spurious vkQueueSubmit slice with submission id %v", v)
			continue
		}
		submissionOrdering[v] = order
		order++
	}

	data.Slices.MapIdentifiers(ctx, handleMapping)

	groupID := int32(-1)
	for i := range data.Slices {
		slice := &data.Slices[i]
		subOrder, ok := submissionOrdering[slice.Submission]
		if ok {
			cb := uint64(slice.CommandBuffer)
			key := sync.RenderPassKey{
				subOrder, cb, uint64(slice.Renderpass), uint64(slice.RenderTarget),
			}
			// Create a new group for each main renderPass slice.
			name := slice.Name
			indices := syncData.RenderPassLookup.Lookup(ctx, key)
			if !indices.IsNil() && (name == "vertex" || name == "fragment") {
				slice.Name = fmt.Sprintf("%v-%v %v", indices.From, indices.To, name)
				groupID = data.Groups.GetOrCreateGroup(
					fmt.Sprintf("RenderPass %v, RenderTarget %v", uint64(slice.Renderpass), uint64(slice.RenderTarget)),
					indices,
				)
			}
		} else {
			log.W(ctx, "Encountered submission ID mismatch %v", slice.Submission)
		}

		if groupID < 0 {
			log.W(ctx, "Group missing for slice %v at submission %v, commandBuffer %v, renderPass %v, renderTarget %v",
				slice.Name, slice.Submission, slice.CommandBuffer, slice.Renderpass, slice.RenderTarget)
		}
		slice.GroupID = groupID
	}

	return nil
}

func processCounters(ctx context.Context, processor *perfetto.Processor, desc *device.GpuCounterDescriptor) ([]*service.ProfilingData_Counter, error) {
	counterTracksQueryResult, err := processor.Query(counterTracksQuery)
	if err != nil {
		return nil, log.Errf(ctx, err, "SQL query failed: %v", counterTracksQuery)
	}
	// t.id, name, unit, description, ts, value
	tracksColumns := counterTracksQueryResult.GetColumns()
	numTracksRows := counterTracksQueryResult.GetNumRecords()
	counters := make([]*service.ProfilingData_Counter, numTracksRows)
	// Grab all the column values. Depends on the order of columns selected in countersQuery
	trackIds := tracksColumns[0].GetLongValues()
	names := tracksColumns[1].GetStringValues()
	units := tracksColumns[2].GetStringValues()
	descriptions := tracksColumns[3].GetStringValues()

	//log.W(ctx, "mali processCounters names is %v, tracksColumns is %d, numTracksRows is %d", names, len(tracksColumns), (numTracksRows))

	nameToSpec := map[string]*device.GpuCounterDescriptor_GpuCounterSpec{}
	if desc != nil {
		for _, spec := range desc.Specs {
			nameToSpec[spec.Name] = spec
		}
	}

	for i := uint64(0); i < numTracksRows; i++ {
		countersQuery := fmt.Sprintf(countersQueryFmt, trackIds[i])
		countersQueryResult, err := processor.Query(countersQuery)
		countersColumns := countersQueryResult.GetColumns()
		if err != nil {
			return nil, log.Errf(ctx, err, "SQL query failed: %v", counterTracksQuery)
		}
		timestampsLong := countersColumns[0].GetLongValues()
		timestamps := make([]uint64, len(timestampsLong))
		for i, t := range timestampsLong {
			timestamps[i] = uint64(t)
		}
		values := countersColumns[1].GetDoubleValues()

		spec, _ := nameToSpec[names[i]]
		// TODO(apbodnar) Populate the `default` field once the trace processor supports it (b/147432390)
		counters[i] = &service.ProfilingData_Counter{
			Id:          uint32(trackIds[i]),
			Name:        names[i],
			Unit:        units[i],
			Description: descriptions[i],
			Spec:        spec,
			Timestamps:  timestamps,
			Values:      values,
		}

		//log.W(ctx, "processCounter names is %s, Id is %d, Unit is %d,description is %v,spc is %v", names[i], uint32(trackIds[i]), units[i], descriptions[i], spec)

	}
	return counters, nil
}

const (
	summaryCounterGroup   uint32 = 1 //性能参数
	defaultCounterGroup   uint32 = 2 //custom参数
	textureCounterGroup   uint32 = 3
	loadStoreCounterGroup uint32 = 4
	memoryCounterGroup    uint32 = 5
	cacheCounterGroup     uint32 = 6
)

type name_key_value struct {
	name string
	des  string
}

///暂时移除一些性能参数的分类
func updateCounterGroups(ctx context.Context, data *profile.ProfilingData) {
	data.CounterGroups = append(data.CounterGroups,
		&service.ProfilingData_CounterGroup{
			Id:    summaryCounterGroup,
			Label: "Performance",
		},
		// &service.ProfilingData_CounterGroup{
		// 	Id:    defaultCounterGroup,
		// 	Label: "Overview",
		// },
		// &service.ProfilingData_CounterGroup{
		// 	Id:    textureCounterGroup,
		// 	Label: "Texture",
		// },
		// &service.ProfilingData_CounterGroup{
		// 	Id:    loadStoreCounterGroup,
		// 	Label: "Load/Store",
		// },
		// &service.ProfilingData_CounterGroup{
		// 	Id:    memoryCounterGroup,
		// 	Label: "Memory",
		// },
		// &service.ProfilingData_CounterGroup{
		// 	Id:    cacheCounterGroup,
		// 	Label: "Cache",
		// },
	)
	//一些重要的性能参数用来做赛选
	// summaryCounters := map[string]struct{}{
	// 	"gpu time":                                   struct{}{},
	// 	"gpu active cycles":                          struct{}{},
	// 	"fragment queue active cycles":               struct{}{},
	// 	"non-fragment queue active cycles":           struct{}{},
	// 	"tiler active cycles":                        struct{}{},
	// 	"texture filtering cycles":                   struct{}{},
	// 	"cycles per pixel":                           struct{}{},
	// 	"pixels":                                     struct{}{},
	// 	"output external write bytes":                struct{}{},
	// 	"output external read bytes":                 struct{}{},
	// 	"gpu utilization":                            struct{}{},
	// 	"fragment queue utilization":                 struct{}{},
	// 	"non-fragment queue utilization":             struct{}{},
	// 	"tiler utilization":                          struct{}{},
	// 	"texture unit utilization":                   struct{}{},
	// 	"varying unit utilization":                   struct{}{},
	// 	"load/store unit utilization":                struct{}{},
	// 	"load/store read bytes from l2 cache":        struct{}{},
	// 	"load/store read bytes from external memory": struct{}{},
	// 	"texture read bytes from l2 cache":           struct{}{},
	// 	"texture read bytes from external memory":    struct{}{},
	// 	"front-end read bytes from l2 cache":         struct{}{},
	// 	"front-end read bytes from external memory":  struct{}{},
	// 	"load/store write bytes":                     struct{}{},
	// 	"tile buffer write bytes":                    struct{}{},
	// }

	summaryCounters := map[string]name_key_value{

		"gpu time":                                   {"gpu 耗时 ", "gpu总耗时"}, //Total time spent on the GPU, computed by summing the duration of all the GPU activity slices.
		"gpu active cycles":                          {"gpu负载", "该计数器反映应用程序对GPU的总体处理负载\n在GPU的任何处理队列中存在未完成工作时递增\n该计数器在任何处理队列中存在工作负载时递增，即使GPU由于等待外部存储器返回数据而停滞时，仍然增加\n"},
		"fragment queue active cycles":               {"ragment queue active cycles", ""},
		"non-fragment queue active cycles":           {"non-fragment queue active cycles", ""},
		"tiler active cycles":                        {"块处理活动周期数", "这个计数器在图块处理单元（tiler）的处理队列中有工作负载时递增\n图块处理单元可以与顶点着色和片段着色同时运行\n这里的高周期计数不一定意味着瓶颈，除非着色核心中的Non-fragment active cycles计数相对于此非常低"},
		"texture filtering cycles":                   {"纹理过滤周期数", "由于多周期数据访问和过滤操作，某些指令可能需要多个周期\n2D 双线性过滤需要一个周期\n2D 三线性过滤需要两个周期\n3D双线性过滤需要两个周期\n3D三线性过滤需要四个周期\n"},
		"cycles per pixel":                           {"每像素渲染花费的GPU周期数", "每像素渲染花费的GPU周期的平均值，包括任何顶点着色的消耗\n这个指标对于对每个渲染过程设置一个循环预算是一个有用的做法，这个预算基于目标分辨率和帧率\n在大众市场设备上，以1080p分辨率和60FPS的帧率进行渲染是可能的，但你可用于每个像素的循环周期数可能很有限\n为了实现60FPS的性能目标，必须明智地利用这些周期\n因此，建议开发者要仔细考虑每个像素的计算成本，确保在有限的循环周期内达到所期望的性能水平"},
		"pixels":                                     {"渲染pass总像素数", "任何渲染通道中着色的像素总数\n请注意，这个数据有可能比实际值大，因为底层硬件计数器会将渲染表面的宽度和高度值四舍五入为32像素对齐\n即使这些像素在着色过程中实际上没有被处理，比如在活动视口和/或剪刀区域之外"},
		"output external write bytes":                {"写带宽", "GPU 将数据写入到外部设备（如主存储器或其他外部目标的数据)"},
		"output external read bytes":                 {"读带宽", "GPU 从外部设备（如主存储器或其他外部源）读取数据的速率"},
		"gpu utilization":                            {"GPU利用率", "该计数器定义了 GPU 相对于设备支持的最大 GPU 频率的总使用率\n受GPU限制的应用程序应该达到大约98%的利用率\n低于这个值的利用率通常表示以下情况之一：\n内容受到垂直同步（vsync）相关限制\n受到 CPU 限制，GPU 已经没有足够的工作要处理\n流水线处理效果不佳，因此繁忙时期在CPU和GPU之间波动"},
		"fragment queue utilization":                 {"fragment queue utilization", ""},
		"non-fragment queue utilization":             {"non-fragment queue utilization", ""},
		"tiler utilization":                          {"图块处理单元的利用率", "图块处理单元，相对于总GPU活动周期的利用率\n请注意，这衡量的是基于索引驱动的顶点着色（IDVS）工作负载的总体处理时间，除了固定功能的图块处理过程之外\n它不一定反映了固定功能图块处理过程本身的运行时间\n"},
		"texture unit utilization":                   {"纹理单元利用率百分比", "在图形渲染中，纹理单元用于处理纹理映射，即将纹理贴图应用到物体表面上\n纹理单元的利用率百分比表示在一定时间内纹理单元实际用于处理纹理的比例"},
		"varying unit utilization":                   {"varying单元利用率", "在图形渲染中，varying单元负责处理着色器中使用的变量数据"},
		"load/store unit utilization":                {"加载/存储单元的利用率", "在图形渲染中，加载/存储单元负责处理着色器中的加载和存储操作，例如读取纹理数据或将数据写入内存\n加载/存储单元的利用率百分比表示在一定时间内该单元实际用于执行加载和存储操作的比例"},
		"load/store read bytes from l2 cache":        {"加载/存储单元L2 cache", "L2缓存通常是一个高速缓存层，用于存储临时数据，以加速处理器对内存的访问\n这个参数指示了加载/存储单元从L2缓存读取数据的总量\n是评估内存访问效率和性能的一个关键指标\n"},
		"load/store read bytes from external memory": {"加载/存储单元从外部存储系统读取的总字节数", "这表示加载/存储单元从外部内存系统(而非缓存中)读取数据的总字节数\n这个度量是评估系统内存访问效率和性能的一个关键指标\n因为从外部内存读取数据的速率直接影响图形渲染和其他计算任务的性能"},
		"texture read bytes from l2 cache":           {"纹理单元从L2存储系统读取的总字节数\n", "纹理单元负责处理图形渲染中的纹理贴图，它从L2缓存中读取纹理数据以供渲染使用\n这个度量指示了纹理单元从L2缓存中读取纹理数据的总量"},
		"texture read bytes from external memory":    {"由纹理单元从外部存储系统读取的总字节数\n", "纹理单元负责处理图形渲染中的纹理贴图,它从外部内存系统中读取纹理数据以供渲染使用\n这个度量指示了纹理单元从外部内存读取纹理数据的总量\n对于评估图形渲染性能和内存访问效率非常重要\n纹理单元高效地读取纹理数据有助于提高图形渲染速度和质量"},
		"front-end read bytes from l2 cache":         {"片段前端单元从 L2 存储系统读取的总字节数", "片段前端单元是图形渲染中的一个组件，负责处理片段（即屏幕上的像素）的前处理阶段\n它从 L2 缓存中读取所需的数据以进行处理\n这个度量表示片段前端单元从 L2 缓存中读取数据的总量，是评估图形渲染性能的一个关键指标\n有效的片段前端数据读取有助于提高渲染效率"},
		"front-end read bytes from external memory":  {"片段前端单元从外部存储系统读取的总字节数", "片段前端单元是图形渲染中的一个组件，负责处理片段（即屏幕上的像素）的前处理阶段。\n它从外部内存系统中读取所需的数据以进行处理\n这个度量表示片段前端单元从外部内存读取数据的总量，是评估图形渲染性能和内存访问效率的一个关键指标。\n有效的片段前端数据读取有助于提高渲染效率"},
		"load/store write bytes":                     {"加载/存储单元写入到 L2 存储系统的总字节数", "加载/存储单元写入数据到L2存储系统的总量，是评估内存写入性能和系统带宽使用的一个关键指标"},
		"tile buffer write bytes":                    {"tile buffer写入L2存储系统的总字节数", "tile buffer写回单元负责将瓦片缓冲中的数据写回到L2存储系统，以便进行进一步的处理和存储\n这个度量表示tile buffer写回单元写入数据到L2存储系统的总量，是评估内存写入性能和系统带宽使用的一个关键指标"},
	}

	//性能参数分类
	//default
	for _, counter := range data.GpuCounters.Metrics {
		name := strings.ToLower(counter.Name)
		if value, ok := summaryCounters[name]; ok {
			counter.CounterGroupIds = append(counter.CounterGroupIds, summaryCounterGroup)
			counter.Description = value.des
			counter.Name = value.name
			delete(summaryCounters, name)
		}
		//暂时移除一些不太重要的性能参数
		//  else if strings.Index(name, "iterator active") >= 0 ||
		// 	strings.Index(name, "iterator utilization") >= 0 {
		// 	counter.CounterGroupIds = append(counter.CounterGroupIds, summaryCounterGroup)
		// }
		if counter.SelectByDefault {
			counter.CounterGroupIds = append(counter.CounterGroupIds, defaultCounterGroup)
		}
		if strings.Index(name, "texture") >= 0 {
			counter.CounterGroupIds = append(counter.CounterGroupIds, textureCounterGroup)
		}
		if strings.Index(name, "load/store") >= 0 {
			counter.CounterGroupIds = append(counter.CounterGroupIds, loadStoreCounterGroup)
		}
		if strings.Index(name, "memory") >= 0 ||
			strings.Index(name, "mmu") >= 0 ||
			strings.Index(name, "read bytes") >= 0 ||
			strings.Index(name, "write bytes") >= 0 {
			counter.CounterGroupIds = append(counter.CounterGroupIds, memoryCounterGroup)
		}
		if strings.Index(name, "cache") >= 0 {
			counter.CounterGroupIds = append(counter.CounterGroupIds, cacheCounterGroup)
		}
	}
	//输出没有找到的参数信息
	for counter := range summaryCounters {
		log.W(ctx, "Summary counter %v not found", counter)
	}
}
