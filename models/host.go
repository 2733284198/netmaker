package models

import (
	"net"

	"github.com/google/uuid"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Host - represents a host on the network
type Host struct {
	ID               uuid.UUID        `json:"id" yaml:"id"`
	Verbosity        int              `json:"verbosity" yaml:"verbosity"`
	FirewallInUse    string           `json:"firewallinuse" yaml:"firewallinuse"`
	Version          string           `json:"version" yaml:"version"`
	IPForwarding     bool             `json:"ipforwarding" yaml:"ipforwarding"`
	DaemonInstalled  bool             `json:"daemoninstalled" yaml:"daemoninstalled"`
	HostPass         string           `json:"hostpass" yaml:"hostpass"`
	Name             string           `json:"name" yaml:"name"`
	OS               string           `json:"os" yaml:"os"`
	Interface        string           `json:"interface" yaml:"interface"`
	Debug            bool             `json:"debug" yaml:"debug"`
	ListenPort       int              `json:"listenport" yaml:"listenport"`
	LocalAddress     net.IPNet        `json:"localaddress" yaml:"localaddress"`
	LocalRange       net.IPNet        `json:"localrange" yaml:"localrange"`
	LocalListenPort  int              `json:"locallistenport" yaml:"locallistenport"`
	ProxyListenPort  int              `json:"proxy_listen_port" yaml:"proxy_listen_port"`
	MTU              int              `json:"mtu" yaml:"mtu"`
	PublicKey        wgtypes.Key      `json:"publickey" yaml:"publickey"`
	MacAddress       net.HardwareAddr `json:"macaddress" yaml:"macaddress"`
	TrafficKeyPublic []byte           `json:"traffickeypublic" yaml:"trafficekeypublic"`
	InternetGateway  net.UDPAddr      `json:"internetgateway" yaml:"internetgateway"`
	Nodes            []string         `json:"nodes" yaml:"nodes"`
}
