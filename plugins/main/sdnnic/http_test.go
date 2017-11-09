package main_test

import (
	. "github.com/containernetworking/plugins/plugins/main/sdnnic"

	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Http", func() {
	var (
		sdnplugin SDNPlugin
		server    *httptest.Server
	)

	BeforeEach(func() {
		sdnplugin = &HTTPSDNPlugin{}
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp, _ := json.Marshal(NICServiceResponse{
				0, "4a:00:03:b5:69:b1", "192.168.0.1", "fe80::107c:b6cf:1580:cf9c",
			})
			fmt.Fprintln(w, resp)
		}))
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Testing http setup", func() {
		Context("plugin configuration is in wrong format", func() {
			config := make(map[string]interface{})
			config["url"] = server.URL
			Expect(sdnplugin.Setup(config)).NotTo(HaveOccurred())
		})
	})

	Describe("Testing http function", func() {
		Context("service returned correct result", func() {
		})
	})
})
