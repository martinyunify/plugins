package main

const ShellPlugin = "shell"

func init() {
	PluginRegistry[ShellPlugin] = &ShellSDNPlugin{}
}

type ShellSDNPlugin struct {
	script string
}

func (plugin *ShellSDNPlugin) Setup(map[string]interface{}) error {
	return nil
}
func (plugin *ShellSDNPlugin) AllocateNic() (int, error) {
	return 0, nil
}
func (plugin *ShellSDNPlugin) DeleteNic(int) error {
	return nil
}
