// Copyright 2017 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

const ShellPlugin = "shell"

func init() {
	PluginRegistry[ShellPlugin] = &ShellSDNPlugin{}
}

//ShellSDNPlugin invoke shell script to attach nic
type ShellSDNPlugin struct {
	script string
}

func (plugin *ShellSDNPlugin) Setup(map[string]interface{}) error {
	return nil
}
func (plugin *ShellSDNPlugin) AllocateNic() (*NetworkInterfaceCard, error) {
	return nil, nil
}
func (plugin *ShellSDNPlugin) DeleteNic(int) error {
	return nil
}
