package main_test

import (
	. "github.com/containernetworking/plugins/plugins/main/sdnnic"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell plugin unit test", func() {
	var (
		sdnplugin SDNPlugin
	)

	BeforeEach(func() {
		sdnplugin = &ShellSDNPlugin{}
	})

	AfterEach(func() {

	})

	Describe("Testing shell plugin setup", func() {
		Context("plugin configuration is in wrong format", func() {
			config := make(map[string]interface{})
			resp := NICServiceResponse{
				0, "4a:00:03:b5:69:b1", "192.168.0.1", "fe80::107c:b6cf:1580:cf9c",
			}
			respText, _ := json.Marshal(resp)
			config["script"] = "echo '" + string(respText) + "'"
			Expect(sdnplugin.Setup(config)).NotTo(HaveOccurred())
		})
	})

	Describe("Testing shell plugin function", func() {
		Context("service returned correct result", func() {
		})
	})
})
