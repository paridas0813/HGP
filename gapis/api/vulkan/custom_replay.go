// Copyright (C) 2017 Google Inc.
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

package vulkan

import (
	"context"
	"strings"

	"github.com/google/gapid/core/log"
	"github.com/google/gapid/gapis/api"
	"github.com/google/gapid/gapis/memory"
	"github.com/google/gapid/gapis/replay/builder"
	"github.com/google/gapid/gapis/replay/value"
)

const virtualSwapchainStruct = 0xFFFFFFAA

func (i VkInstance) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkPhysicalDevice) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDevice) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkQueue) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkCommandBuffer) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkSemaphore) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkFence) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDeviceMemory) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkBuffer) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkImage) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkEvent) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkQueryPool) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkBufferView) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkImageView) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkShaderModule) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkPipelineCache) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkPipelineLayout) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkRenderPass) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkPipeline) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDescriptorSetLayout) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkSampler) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDescriptorPool) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDescriptorSet) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkFramebuffer) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkCommandPool) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkSurfaceKHR) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkSwapchainKHR) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDisplayKHR) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDisplayModeKHR) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDebugReportCallbackEXT) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkSamplerYcbcrConversion) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDebugUtilsMessengerEXT) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (i VkDescriptorUpdateTemplate) remap(api.Cmd, *api.GlobalState) (key interface{}, remap bool) {
	if i != 0 {
		key, remap = i, true
	}
	return
}

func (a *VkCreateInstance) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	cb := CommandBuilder{Thread: a.Thread()}
	// Hijack VkCreateInstance's Mutate() method entirely with our ReplayCreateVkInstance's Mutate().

	// As long as we guarantee that the synthetic replayCreateVkInstance API function has the same
	// logic as the real vkCreateInstance API function, we can do observation correctly. Additionally,
	// ReplayCreateVkInstance's Mutate() will invoke our custom wrapper function replayCreateVkInstance()
	// in vulkan_gfx_api_extras.cpp, which modifies VkInstanceCreateInfo to enable virtual swapchain
	// layer before delegating the real work back to the normal flow.

	hijack := cb.ReplayCreateVkInstance(a.PCreateInfo(), a.PAllocator(), a.PInstance(), a.Result())
	hijack.Extras().MustClone(a.Extras().All()...)
	err := hijack.Mutate(ctx, id, s, b, w)

	if b == nil || err != nil {
		return err
	}

	// Call the replayRegisterVkInstance() synthetic API function.
	instance, err := a.PInstance().Read(ctx, a, s, b)
	if err != nil {
		return err
	}
	return cb.ReplayRegisterVkInstance(instance).Mutate(ctx, id, s, b, nil)
}

func (a *VkDestroyInstance) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	cb := CommandBuilder{Thread: a.Thread()}
	hijack := cb.ReplayDestroyVkInstance(a.Instance(), a.PAllocator())
	hijack.Extras().MustClone(a.Extras().All()...)
	err := hijack.Mutate(ctx, id, s, b, w)

	if b == nil || err != nil {
		return err
	}
	// Call the replayUnregisterVkInstance() synthetic API function.
	return cb.ReplayUnregisterVkInstance(a.Instance()).Mutate(ctx, id, s, b, nil)
}

func EnterRecreate(ctx context.Context, s *api.GlobalState) func() {
	GetState(s).SetIsRebuilding(true)
	return func() { GetState(s).SetIsRebuilding(false) }
}

func (a *VkCreateDevice) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Hijack VkCreateDevice's Mutate() method entirely with our
	// ReplayCreateVkDevice's Mutate(). Similar to VkCreateInstance's Mutate()
	// above.
	// And we need to strip off the VK_EXT_debug_marker extension name when
	// building instructions for replay.
	createInfoPtr := a.PCreateInfo()
	allocated := []*api.AllocResult{}
	if b != nil {
		a.Extras().Observations().ApplyReads(s.Memory.ApplicationPool())
		createInfo, err := a.PCreateInfo().Read(ctx, a, s, nil)
		if err != nil {
			return err
		}
		defer func() {
			for _, d := range allocated {
				d.Free()
			}
		}()
		extensionCount := uint64(createInfo.EnabledExtensionCount())
		newExtensionNames := []memory.Pointer{}
		enabledExtensionNames, err := createInfo.PpEnabledExtensionNames().Slice(0, extensionCount, s.MemoryLayout).Read(ctx, a, s, nil)
		if err != nil {
			return err
		}
		for _, e := range enabledExtensionNames {
			rawExtensionName, err := e.StringSlice(ctx, s).Read(ctx, a, s, nil)
			if err != nil {
				return err
			}
			extensionName := string(memory.CharToBytes(rawExtensionName))
			if !strings.Contains(extensionName, "VK_EXT_debug_marker") {
				nameSliceData := s.AllocDataOrPanic(ctx, extensionName)
				allocated = append(allocated, &nameSliceData)
				newExtensionNames = append(newExtensionNames, nameSliceData.Ptr())
			}
		}
		newExtensionNamesData := s.AllocDataOrPanic(ctx, newExtensionNames)
		allocated = append(allocated, &newExtensionNamesData)
		createInfo.SetEnabledExtensionCount(uint32(len(newExtensionNames)))
		createInfo.SetPpEnabledExtensionNames(NewCharᶜᵖᶜᵖ(newExtensionNamesData.Ptr()))

		newCreateInfoData := s.AllocDataOrPanic(ctx, createInfo)
		allocated = append(allocated, &newCreateInfoData)
		createInfoPtr = NewVkDeviceCreateInfoᶜᵖ(newCreateInfoData.Ptr())

		cb := CommandBuilder{Thread: a.Thread()}
		hijack := cb.ReplayCreateVkDevice(a.PhysicalDevice(), createInfoPtr, a.PAllocator(), a.PDevice(), a.Result())
		hijack.Extras().MustClone(a.Extras().All()...)

		for _, d := range allocated {
			hijack.AddRead(d.Data())
		}

		if err := hijack.Mutate(ctx, id, s, b, w); err != nil {
			return err
		}
		// Call the replayRegisterVkDevice() synthetic API function.
		device, err := a.PDevice().Read(ctx, a, s, b)
		if err != nil {
			return err
		}
		return cb.ReplayRegisterVkDevice(a.PhysicalDevice(), device, a.PCreateInfo()).Mutate(ctx, id, s, b, nil)
	}

	return a.mutate(ctx, id, s, b, w)
}

func (a *VkDestroyDevice) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Call the underlying vkDestroyDevice() and do the observation.
	cb := CommandBuilder{Thread: a.Thread()}
	err := a.mutate(ctx, id, s, b, w)
	if b == nil || err != nil {
		return err
	}
	// Call the replayUnregisterVkDevice() synthetic API function.
	return cb.ReplayUnregisterVkDevice(a.Device()).Mutate(ctx, id, s, b, nil)
}

func (a *VkAllocateCommandBuffers) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Call the underlying vkAllocateCommandBuffers() and do the observation.
	cb := CommandBuilder{Thread: a.Thread()}
	err := a.mutate(ctx, id, s, b, w)
	if b == nil || err != nil {
		return err
	}
	// Call the replayRegisterVkCommandBuffers() synthetic API function to link these command buffers to the device.
	allocateInfo, err := a.PAllocateInfo().Read(ctx, a, s, b)
	if err != nil {
		return err
	}
	count := allocateInfo.CommandBufferCount()
	return cb.ReplayRegisterVkCommandBuffers(a.Device(), count, a.PCommandBuffers()).Mutate(ctx, id, s, b, nil)
}

func (a *VkFreeCommandBuffers) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Call the underlying vkFreeCommandBuffers() and do the observation.
	cb := CommandBuilder{Thread: a.Thread()}
	err := a.mutate(ctx, id, s, b, w)
	if b == nil || err != nil {
		return err
	}
	// Call the replayUnregisterVkCommandBuffers() synthetic API function to discard the link of these command buffers.
	count := a.CommandBufferCount()
	return cb.ReplayUnregisterVkCommandBuffers(count, a.PCommandBuffers()).Mutate(ctx, id, s, b, nil)
}

func (a *VkCreateSwapchainKHR) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return a.mutate(ctx, id, s, b, w)
	}

	cb := CommandBuilder{Thread: a.Thread()}
	hijack := cb.ReplayCreateSwapchain(a.Device(), a.PCreateInfo(), a.PAllocator(), a.PSwapchain(), a.Result())
	hijack.Extras().MustClone(a.Extras().All()...)

	a.Extras().Observations().ApplyReads(s.Memory.ApplicationPool())
	info, err := a.PCreateInfo().Read(ctx, a, s, nil)
	if err != nil {
		return err
	}
	pNext := NewVirtualSwapchainPNext(
		VkStructureType_VK_STRUCTURE_TYPE_VIRTUAL_SWAPCHAIN_PNEXT, // sType
		info.PNext(), // pNext
		0,            // surfaceCreateInfo
	)
	for _, extra := range a.Extras().All() {
		if d, ok := extra.(*DisplayToSurface); ok {
			log.D(ctx, "Activating display to surface")
			sType, _ := d.SurfaceTypes[uint64(info.Surface())]
			sTypeData := s.AllocDataOrPanic(ctx, sType)
			defer sTypeData.Free()
			pNext.SetSurfaceCreateInfo(NewVoidᶜᵖ(sTypeData.Ptr()))
			hijack.AddRead(sTypeData.Data())
		}
	}
	pNextData := s.AllocDataOrPanic(ctx, pNext)
	defer pNextData.Free()

	info.SetPNext(NewVoidᶜᵖ(pNextData.Ptr()))
	infoData := s.AllocDataOrPanic(ctx, info)
	defer infoData.Free()
	hijack.SetPCreateInfo(NewVkSwapchainCreateInfoKHRᶜᵖ(infoData.Ptr()))

	hijack.AddRead(pNextData.Data())
	hijack.AddRead(infoData.Data())

	return hijack.Mutate(ctx, id, s, b, w)
}

func (a *VkAcquireNextImageKHR) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Do the mutation, including applying memory write observations, before having the replay device call the vkAcquireNextImageKHR() command.
	// This is to pass the returned image index value captured in the trace, into the replay device to acquire for the specific image.
	// Note that this is only necessary for building replay instructions
	err := a.mutate(ctx, id, s, nil, w)
	if err != nil {
		return err
	}
	if b != nil {
		l := s.MemoryLayout
		// Ensure that the builder reads pImageIndex (which points to the correct image index at this point).
		a.PImageIndex().Slice(0, 1, l).OnRead(ctx, a, s, b)
		a.Call(ctx, s, b)
	}
	return err
}

func (a *VkAcquireNextImage2KHR) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	// Do the mutation, including applying memory write observations, before having the replay device call the vkAcquireNextImageKHR() command.
	// This is to pass the returned image index value captured in the trace, into the replay device to acquire for the specific image.
	// Note that this is only necessary for building replay instructions
	err := a.mutate(ctx, id, s, nil, w)
	if err != nil {
		return err
	}
	if b != nil {
		l := s.MemoryLayout
		// Ensure that the builder reads pImageIndex (which points to the correct image index at this point).
		a.PImageIndex().Slice(0, 1, l).OnRead(ctx, a, s, b)
		a.PAcquireInfo().Slice(0, 1, l).OnRead(ctx, a, s, b)
		a.Call(ctx, s, b)
	}
	return err
}

type structWithPNext interface {
	PNext() Voidᶜᵖ
	SetPNext(v Voidᶜᵖ)
}

func insertVirtualSwapchainPNext(ctx context.Context, cmd api.Cmd, id api.CmdID,
	info structWithPNext, g *api.GlobalState) (api.AllocResult, api.AllocResult, error) {
	pNextData := g.AllocDataOrPanic(ctx, NewVulkanStructHeader(
		virtualSwapchainStruct, // sType
		0,                      // pNext
	))
	if info.PNext().IsNullptr() {
		info.SetPNext(NewVoidᶜᵖ(pNextData.Ptr()))
	} else {
		pNext := NewVoidᵖ(info.PNext())
		for !pNext.IsNullptr() {
			structHeader, err := NewVulkanStructHeaderᵖ(pNext).Read(ctx, cmd, g, nil)
			if err != nil {
				return api.AllocResult{}, api.AllocResult{}, err
			}
			if !structHeader.PNext().IsNullptr() {
				structHeader.SetPNext(NewVoidᵖ(pNextData.Ptr()))
				break
			}
			pNext = structHeader.PNext()
		}
	}
	newInfoData := g.AllocDataOrPanic(ctx, info)
	return newInfoData, pNextData, nil
}

func (c *VkCreateXlibSurfaceKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateXlibSurfaceKHR(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}
	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_XLIB)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkCreateXcbSurfaceKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateXcbSurfaceKHR(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}

	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_XCB)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkCreateWaylandSurfaceKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateWaylandSurfaceKHR(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}
	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_WAYLAND)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkCreateWin32SurfaceKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateWin32SurfaceKHR(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}
	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_WIN32)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkCreateAndroidSurfaceKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateAndroidSurfaceKHR(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}
	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_ANDROID)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkCreateMacOSSurfaceMVK) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, insert a pNext struct to enable the
	// virtual surface on the replay device.
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	createInfo, err := c.PCreateInfo().Read(ctx, c, g, nil)
	if err != nil {
		return err
	}
	newInfoData, pNextData, err := insertVirtualSwapchainPNext(ctx, c, id, createInfo, g)
	if err != nil {
		return err
	}
	defer newInfoData.Free()
	defer pNextData.Free()
	cb := CommandBuilder{Thread: c.Thread()}
	hijack := cb.VkCreateMacOSSurfaceMVK(
		c.Instance(), newInfoData.Ptr(), c.PAllocator(), c.PSurface(), c.Result(),
	).AddRead(newInfoData.Data()).AddRead(pNextData.Data())
	for _, r := range c.Extras().Observations().Reads {
		hijack.AddRead(r.Range, r.ID)
	}
	for _, w := range c.Extras().Observations().Writes {
		hijack.AddWrite(w.Range, w.ID)
	}
	hijack.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	info, err := hijack.PCreateInfo().Read(ctx, hijack, g, b)
	if err != nil {
		return err
	}
	if (info.PNext()) != (Voidᶜᵖ(0)) {
		numPNext := (externs{ctx, hijack, id, g, b, nil}.numberOfPNext(info.PNext()))
		next := NewMutableVoidPtr(Voidᵖ(info.PNext()))
		for i := uint32(0); i < numPNext; i++ {
			if _, err := VkStructureTypeᶜᵖ(next.Ptr()).Read(ctx, hijack, g, b); err != nil {
				return err
			}
			header, err := VulkanStructHeaderᵖ(next.Ptr()).Read(ctx, hijack, g, b)
			if err != nil {
				return err
			}
			next.SetPtr(header.PNext())
		}
	}
	surface := NewSurfaceObjectʳ(
		VkInstance(0), VkSurfaceKHR(0), SurfaceType(0),
		NilVulkanDebugMarkerInfoʳ, NewVkPhysicalDeviceːQueueFamilySupportsʳᵐ())
	surface.SetInstance(hijack.Instance())
	surface.SetType(SurfaceType_SURFACE_TYPE_MACOS_MVK)

	hijack.Call(ctx, g, b)

	hijack.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	handle, err := hijack.PSurface().Read(ctx, hijack, g, nil)
	if err != nil {
		return err
	}
	if err := hijack.PSurface().Write(ctx, handle, hijack, g, b); err != nil {
		return err
	}
	surface.SetVulkanHandle(handle)
	GetState(g).Surfaces().Add(handle, surface)
	hijack.Result()
	return nil
}

func (c *VkGetPhysicalDeviceSurfaceFormatsKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	// When building replay instructions, apply the write observations so that
	// the returned surface format count and formats, which are captured in
	// the trace, will be passed to the virtual swapchain. This is to cheat
	// the validation layers, as the returned surface formats will always match
	// with the format used in the trace.
	l := g.MemoryLayout
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	c.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	givenCount, err := c.PSurfaceFormatCount().Read(ctx, c, g, b)
	if err != nil {
		return err
	}
	if (c.PSurfaceFormats()) != (VkSurfaceFormatKHRᵖ(0)) {
		c.PSurfaceFormats().Slice(0, uint64(givenCount), l).OnRead(ctx, c, g, b)
	}
	c.Call(ctx, g, b)
	c.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	if (c.PSurfaceFormats()) == (VkSurfaceFormatKHRᵖ(0)) {
		count, err := c.PSurfaceFormatCount().Read(ctx, c, g, nil)
		if err != nil {
			return err
		}
		if err := c.PSurfaceFormatCount().Write(ctx, count, c, g, b); err != nil {
			return err
		}
	} else {
		count, err := c.PSurfaceFormatCount().Read(ctx, c, g, nil)
		if err != nil {
			return err
		}
		formats := c.PSurfaceFormats().Slice(0, uint64(count), l)
		for i := uint32(0); i < count; i++ {
			surfaceFmt, err := c.PSurfaceFormats().Slice(uint64(uint32(0)), uint64(count), l).Index(uint64(i)).Read(ctx, c, g, nil)
			if err != nil {
				return err
			}
			if _, err := formats.Index(uint64(i)).Write(ctx, []VkSurfaceFormatKHR{surfaceFmt[0]}, c, g, b); err != nil {
				return err
			}
		}
		if err := c.PSurfaceFormatCount().Write(ctx, count, c, g, b); err != nil {
			return err
		}
	}
	return nil
}

func (c *VkGetPhysicalDeviceSurfacePresentModesKHR) Mutate(ctx context.Context, id api.CmdID, g *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if b == nil {
		return c.mutate(ctx, id, g, b, w)
	}
	l := g.MemoryLayout
	c.Extras().Observations().ApplyReads(g.Memory.ApplicationPool())
	c.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	givenCount, err := c.PPresentModeCount().Read(ctx, c, g, b)
	if err != nil {
		return err
	}
	if (c.PPresentModes()) != (VkPresentModeKHRᵖ(0)) {
		c.PPresentModes().Slice(0, uint64(givenCount), l).OnRead(ctx, c, g, b)
	}
	c.Call(ctx, g, b)
	c.Extras().Observations().ApplyWrites(g.Memory.ApplicationPool())
	if (c.PPresentModes()) == (VkPresentModeKHRᵖ(0)) {
		count, err := c.PPresentModeCount().Read(ctx, c, g, nil)
		if err != nil {
			return err
		}
		if err := c.PPresentModeCount().Write(ctx, count, c, g, b); err != nil {
			return err
		}
	} else {
		count, err := c.PPresentModeCount().Read(ctx, c, g, nil)
		if err != nil {
			return err
		}
		modes := c.PPresentModes().Slice(0, uint64(count), l)
		for i := uint32(0); i < count; i++ {
			mode, err := c.PPresentModes().Slice(0, uint64(count), l).Index(uint64(i)).Read(ctx, c, g, nil)
			if err != nil {
				return err
			}
			if _, err := modes.Index(uint64(i)).Write(ctx, []VkPresentModeKHR{mode[0]}, c, g, b); err != nil {
				return err
			}
		}
		if err := c.PPresentModeCount().Write(ctx, count, c, g, b); err != nil {
			return err
		}
	}
	return nil
}

func (a *VkGetFenceStatus) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	cb := CommandBuilder{Thread: a.Thread()}
	err := a.mutate(ctx, id, s, b, w)
	if b == nil || err != nil {
		return err
	}

	return cb.ReplayGetFenceStatus(a.Device(), a.Fence(), a.Result(), a.Result()).Mutate(ctx, id, s, b, nil)
}

func (a *VkGetEventStatus) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	cb := CommandBuilder{Thread: a.Thread()}
	err := a.mutate(ctx, id, s, b, w)
	if b == nil || err != nil {
		return err
	}
	var wait bool
	switch a.Result() {
	case VkResult_VK_EVENT_SET:
		wait = GetState(s).Events().Get(a.Event()).Signaled() == true
	case VkResult_VK_EVENT_RESET:
		wait = GetState(s).Events().Get(a.Event()).Signaled() == false
	default:
		wait = false
	}

	return cb.ReplayGetEventStatus(a.Device(), a.Event(), a.Result(), wait, a.Result()).Mutate(ctx, id, s, b, nil)
}

func (a *ReplayAllocateImageMemory) Mutate(ctx context.Context, id api.CmdID, s *api.GlobalState, b *builder.Builder, w api.StateWatcher) error {
	if err := a.mutate(ctx, id, s, b, w); err != nil {
		return err
	}
	l := s.MemoryLayout
	c := GetState(s)
	pMemory, err := a.PMemory().Slice(0, 1, l).Read(ctx, a, s, nil)
	if err != nil {
		return err
	}
	memory := pMemory[0]
	imageObject := c.Images().Get(a.Image())
	imageWidth := imageObject.Info().Extent().Width()
	imageHeight := imageObject.Info().Extent().Height()
	imageFormat, err := getImageFormatFromVulkanFormat(imageObject.Info().Fmt())
	if err != nil {
		return err
	}
	imageSize := VkDeviceSize(imageFormat.Size(int(imageWidth), int(imageHeight), 1))
	memoryObject := NewDeviceMemoryObjectʳ(
		a.Device(),                        // Device
		memory,                            // VulkanHandle
		imageSize,                         // AllocationSize
		NewU64ːVkDeviceSizeᵐ(),            // BoundObjects
		0,                                 // MappedOffset
		0,                                 // MappedSize
		0,                                 // MappedLocation
		0,                                 // MemoryTypeIndex
		MakeU8ˢ(uint64(imageSize), s),     // Data
		NilVulkanDebugMarkerInfoʳ,         // DebugInfo
		NilMemoryDedicatedAllocationInfoʳ, // DedicatedAllocationNV
		NilMemoryDedicatedAllocationInfoʳ, // DedicatedAllocationKHR
		NilMemoryAllocateFlagsInfoʳ,       // MemoryAllocateFlagsInfo
		0,                                 // ExternalHandleTypeFlags
		0,                                 // AndroidHardwareBuffer
		0,                                 // DeviceMemoryOpaqueCaptureAddress
	)

	c.DeviceMemories().Add(memory, memoryObject)
	if _, err = a.PMemory().Slice(0, 1, l).Write(ctx, []VkDeviceMemory{memory}, a, s, b); err != nil {
		return err
	}
	return nil
}

func (i AllocationCallbacks) value(b *builder.Builder, cmd api.Cmd, s *api.GlobalState) value.Value {
	// Return 0 (nullptr) here. We don't have an allocator set up for replay. Since we cannot use the
	// application's allocator. If we pass in null for all allocator calls, then it will use the default
	// allocator.
	return value.AbsolutePointer(0)
}

func (cmd *VkWaitForFences) Mutate(ctx context.Context, id api.CmdID, inputState *api.GlobalState, builder *builder.Builder, watcher api.StateWatcher) error {
	if err := cmd.mutate(ctx, id, inputState, builder, watcher); err != nil {
		return err
	}

	if builder == nil {
		return nil
	}

	fenceState := findFenceState(cmd.Extras())
	if fenceState == nil {
		return nil
	}

	cb := CommandBuilder{Thread: cmd.Thread()}

	allocated := []*api.AllocResult{}
	defer func() {
		for _, d := range allocated {
			d.Free()
		}
	}()

	var waitAll bool
	if cmd.WaitAll() == 0 {
		waitAll = false
	} else {
		waitAll = true
	}

	fencesData := inputState.AllocDataOrPanic(ctx, fenceState.fences)
	allocated = append(allocated, &fencesData)
	statusesData := inputState.AllocDataOrPanic(ctx, fenceState.statuses)
	allocated = append(allocated, &statusesData)
	hijack := cb.ReplayWaitForFences(cmd.Device(),
		uint64(len(fenceState.fences)),
		NewVkFenceᵖ(fencesData.Ptr()),
		NewU64ᵖ(statusesData.Ptr()),
		waitAll,
		cmd.Timeout(),
		cmd.Result())

	for _, d := range allocated {
		hijack.AddRead(d.Data())
	}
	return hijack.Mutate(ctx, id, inputState, builder, nil)
}
