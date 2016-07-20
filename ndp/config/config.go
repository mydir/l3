//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package config

import (
	"github.com/google/gopacket/pcap"
)

const (
	STATE_UP      = "UP"
	STATE_DOWN    = "DOWN"
	CONFIG_CREATE = "CREATE"
	CONFIG_DELETE = "DELETE"
	CONFIG_UPDATE = "UPDATE"
)

type PcapBase struct {
	// Pcap Handler for Each Port
	PcapHandle *pcap.Handle
	PcapCtrl   chan bool
}

type PortInfo struct {
	PcapBase
	IntfRef     string
	IfIndex     int32
	Name        string
	OperState   string
	MacAddr     string
	Description string
}

type PortState struct {
	IfIndex int32
	IfState string
}

type IPv6IntfInfo struct {
	PcapBase
	IntfRef     string
	IfIndex     int32
	IpAddr      string
	MsgType     string
	OperState   string
	LinkLocalIp string
}

type IPIntfNotification struct {
	IfIndex   int32
	IpAddr    string
	Operation string
}

type StateNotification struct {
	IfIndex int32
	State   string
	IpAddr  string
}

type NeighborInfo struct {
	MacAddr        string
	VlanId         int32
	IfIndex        int32
	Intf           string
	IpAddr         string
	LinkLocalIp    string
	ExpiryTimeLeft string
}

type VlanInfo struct {
	IfIndex       int32        // vlan ifIndex generated by the system
	Name          string       // vlan name
	UntagPortsMap map[int]bool // key is port ifIndex
	OperState     string
}

type VlanNotification struct {
	VlanId     int32
	VlanName   string
	Operation  string
	UntagPorts []int32
}
