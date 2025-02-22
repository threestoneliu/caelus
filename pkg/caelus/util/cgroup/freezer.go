/*
 * Copyright (c) 2021 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 *
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cgroup

import (
	"github.com/opencontainers/runc/libcontainer/cgroups"
	"path"
)

const (
	freezerCgroup = "freezer"
)

// GetFreezerPids will get pids from "cgroup.procs"
func GetFreezerPids(pathInRoot string) ([]int, error) {
	rootPath := GetRoot()
	fullCgPath := path.Join(rootPath, freezerCgroup, pathInRoot)
	return cgroups.GetPids(fullCgPath)
}
