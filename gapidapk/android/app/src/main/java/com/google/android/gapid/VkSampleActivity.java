/*
 * Copyright (C) 2019 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.google.android.gapid;

import android.app.NativeActivity;
import android.content.pm.PackageManager;
import android.content.pm.PermissionInfo;
import android.os.Bundle;

// This class exists to disambiguate activity names between native activities inside the GAPID
// APK. It just needs to extend NativeActivity.
public class VkSampleActivity extends NativeActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        try {
            PermissionInfo permissionInfo =  getApplicationContext().getPackageManager().getPermissionInfo("com.android.permission.GET_INSTALLED_APPS", 0);
            if (permissionInfo != null && permissionInfo.packageName.equals("com.lbe.security.miui")) {//MIUI 系统支持动态申请该权限
                if (checkSelfPermission("com.android.permission.GET_INSTALLED_APPS") != PackageManager.PERMISSION_GRANTED) {
                    //没有权限，需要申请
                    requestPermissions(new String[]{"com.android.permission.GET_INSTALLED_APPS"}, 999);
                }
            }
        } catch (PackageManager.NameNotFoundException e) {
            // not support
        }
    }

}
