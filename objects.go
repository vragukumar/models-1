package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	CreateDBTable(dbHdl *sql.DB) error
	StoreObjectInDb(dbHdl *sql.DB) (int64, error)
	DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error
	GetKey() (string, error)
	GetSqlKeyStr(string) (string, error)
	GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl *sql.DB) error
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	BaseObj
	DestinationNw     string `SNAPROUTE: "KEY"`
	NetworkMask       string `SNAPROUTE: "KEY"`
	Cost              uint32
	NextHopIp         string
	OutgoingIntfType  string
	OutgoingInterface string
	Protocol          string
}

func (obj IPV4Route) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Route IPV4Route
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Route); err != nil {
			fmt.Println("### Trouble in unmarshaling IPV4Route from Json", body)
		}
	}
	return v4Route, err
}

type BGPGlobalConfig struct {
	BaseObj
	ASNum    uint32
	RouterId string `SNAPROUTE: "KEY"`
}

func (obj BGPGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BGPGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

type BGPGlobalState struct {
	BaseObj
	AS            uint32
	RouterId      string
	TotalPaths    uint32
	TotalPrefixes uint32
}

func (obj BGPGlobalState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gState BGPGlobalState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gState); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gState, err
}

type PeerType int

const (
	PeerTypeInternal PeerType = iota
	PeerTypeExternal
)

type BgpCounters struct {
	Update       uint64
	Notification uint64
}

type BGPMessages struct {
	Sent     BgpCounters
	Received BgpCounters
}

type BGPQueues struct {
	Input  uint32
	Output uint32
}

type BGPNeighborConfig struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	AuthPassword            string
	Description             string
	NeighborAddress         string `SNAPROUTE: "KEY"`
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
}

func (obj BGPNeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nConf BGPNeighborConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nConf); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborConfig from Json", body)
		}
	}
	return nConf, err
}

type BGPNeighborState struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	PeerType                PeerType
	AuthPassword            string
	Description             string
	NeighborAddress         string
	SessionState            uint32
	Messages                BGPMessages
	Queues                  BGPQueues
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
}

func (obj BGPNeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nState BGPNeighborState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nState); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborConfig from Json", body)
		}
	}
	return nState, err
}

/* Start - Asicd objects */
type Vlan struct {
	BaseObj
	VlanId      int32 `SNAPROUTE: "KEY"`
	Ports       string
	PortTagType string
}

/* FIXME : RouterIf needs to be replaced by generic
 * layer 2 object name e.x Port-21 or Vlan-5 etc.
 * Internally this l2 object name can be translated
 * into appropriate key.
 */
type IPv4Intf struct {
	BaseObj
	IpAddr   string `SNAPROUTE: "KEY"`
	RouterIf int32
	IfType   int32
}

type IPv4Neighbor struct {
	BaseObj
	IpAddr   string `SNAPROUTE: "KEY"`
	MacAddr  string
	VlanId   int32
	RouterIf int32
}

func (obj Vlan) UnmarshalObject(body []byte) (ConfigObj, error) {
	var vlanObj Vlan
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &vlanObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Vlan from Json", body)
		}
	}
	return vlanObj, err
}

func (obj IPv4Intf) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Intf IPv4Intf
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Intf); err != nil {
			fmt.Println("### Trouble in unmarshaling IPv4Intf from Json", body)
		}
	}
	return v4Intf, err
}

func (obj IPv4Neighbor) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Nbr IPv4Neighbor
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Nbr); err != nil {
			fmt.Println("### Trouble in unmarshaling IPv4Neighbor from Json", body)
		}
	}
	return v4Nbr, err
}

/* ARP */
type ArpConfig struct {
	BaseObj
	// placeholder to create a key
	ArpConfigKey string `SNAPROUTE: "KEY"`
	Timeout      int32
}

func (obj ArpConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpConfigObj ArpConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}

	return arpConfigObj, err
}

type ArpEntry struct {
	BaseObj
	IpAddr         string `SNAPROUTE: "KEY"`
	MacAddr        string
	Vlan           uint32
	Intf           string
	ExpiryTimeLeft string
}

func (obj ArpEntry) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpEntryObj ArpEntry
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpEntryObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}
	return arpEntryObj, err
}

/* PortInterface */
type PortIntfConfig struct {
	BaseObj
	PortNum     int32 `SNAPROUTE: "KEY"`
	Name        string
	Description string
	Type        string
	AdminState  string
	OperState   string
	MacAddr     string
	Speed       int32
	Duplex      string
	Autoneg     string
	MediaType   string
	Mtu         int32
}

func (obj PortIntfConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var portIntfConfigObj PortIntfConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &portIntfConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling PortIntfConfig from Json", body)
		}
	}

	return portIntfConfigObj, err
}

type PortIntfState struct {
	BaseObj
	PortNum   int32
	PortStats []int64
}

func (obj PortIntfState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var portIntfStateObj PortIntfState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &portIntfStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling PortIntfState from Json", body)
		}
	}

	return portIntfStateObj, err
}

type UserConfig struct {
	BaseObj
	UserName    string `SNAPROUTE: "KEY"`
	Password    string
	Description string
	Previledge  string
}

func (obj UserConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userConfigObj UserConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserConfig from Json", body)
		}
	}

	return userConfigObj, err
}

type UserState struct {
	BaseObj
	UserName      string
	LastLoginTime time.Time
	LastLoginIp   string
	NumAPICalled  uint32
}

func (obj UserState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userStateObj UserState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserState from Json", body)
		}
	}

	return userStateObj, err
}

type Login struct {
	BaseObj
	UserName string
	Password string
}

func (obj Login) UnmarshalObject(body []byte) (ConfigObj, error) {
	var loginObj Login
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &loginObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Login from Json", body)
		}
	}

	return loginObj, err
}

type Logout struct {
	BaseObj
	UserName  string
	SessionId uint32
}

func (obj Logout) UnmarshalObject(body []byte) (ConfigObj, error) {
	var logoutObj Logout
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &logoutObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Logout from Json", body)
		}
	}

	return logoutObj, err
}
