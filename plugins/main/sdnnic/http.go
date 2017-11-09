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

//HTTPPLUGIN plugin id
const HTTPPLUGIN = "http"

func init() {
	PluginRegistry[HTTPPLUGIN] = &HTTPSDNPlugin{}
}

//HTTPSDNPlugin invoke http endpoint to attach nic
type HTTPSDNPlugin struct {
	url string
}

//Setup setup plugin
func (plugin *HTTPSDNPlugin) Setup(map[string]interface{}) error {
	return nil
}

//AllocateNic invoke backedn service to allocate nic
func (plugin *HTTPSDNPlugin) AllocateNic() (*NetworkInterfaceCard, error) {
	return nil, nil
}

//DeleteNic invoke service to delete nic
func (plugin *HTTPSDNPlugin) DeleteNic(int) error {
	return nil
}
