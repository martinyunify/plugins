package main

import "net"

//SDNPlugin invoke sdn service to attach nic
type SDNPlugin interface {
	Setup(map[string]interface{}) error
	AllocateNic() (*NetworkInterfaceCard, error)
	DeleteNic(int) error
}

//NetworkInterfaceCard internal nic object
type NetworkInterfaceCard struct {
	index    int
	MacAddr  net.HardwareAddr
	Ipv4Addr net.IP
	Ipv6Addr net.IP
}

//NICServiceResponse nic service response
type NICServiceResponse struct {
	Index    int    `json:"nic_index"`
	MacAddr  string `json:"mac_addr"`
	Ipv4Addr string `json:"ipv4_addr,omitempty"`
	Ipv6Addr string `json:"ipv6_addr,omitempty"`
}
