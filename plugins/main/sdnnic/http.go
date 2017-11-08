package main

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
func (plugin *HTTPSDNPlugin) AllocateNic() (int, error) {
	return 0, nil
}

//DeleteNic invoke service to delete nic
func (plugin *HTTPSDNPlugin) DeleteNic(int) error {
	return nil
}
