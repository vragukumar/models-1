package models

import (
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

type ConfigObj interface {
	UnmarshalHTTP(*http.Request) error
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	DestinationNw string
	NetworkMask string
	Cost int
	NextHopIp string
	OutgoingInterface  string
	Protocol string
}

func (obj *IPV4Route) UnmarshalHTTP(r *http.Request) error {
	var v4Route IPV4Route
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
	 return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v4Route); err != nil  {
		fmt.Println("### IPV4Route Create is called", v4Route)
	}
	return nil
}

func (obj *Vlan) UnmarshalHTTP(r *http.Request) error {
	return nil
}

type GlobalConfig struct {
	AS uint32
	RouterId string
}

type GlobalState struct {
	AS uint32
	RouterId string
	TotalPaths uint32
	TotalPrefixes uint32
}

type PeerType int

const (
	PeerTypeInternal PeerType = iota
	PeerTypeExternal
)

type BgpCounters struct {
	Update uint64
	Notification uint64
}

type Messages struct {
	Sent BgpCounters
	Received BgpCounters
}

type Queues struct {
	Input uint32
	Output uint32
}

type NeighborConfig struct {
	PeerAS uint32
	LocalAS uint32
	AuthPassword string
	Description string
	NeighborAddress string
}

type NeighborState struct {
	PeerAS uint32
	LocalAS uint32
	PeerType PeerType
	AuthPassword string
	Description string
	NeighborAddress string
	SessionState uint32
	Messages Messages
	Queues Queues
}

/* Start - Asicd objects */
type Vlan struct {
    VlanId      int
    PortMap     string
    PortTagType string
}

/* FIXME : RouterIf needs to be replaced by generic
 * layer 2 object name e.x Port-21 or Vlan-5 etc.
 * Internally this l2 object name can be translated
 * into appropriate key.
 */

type IPv4Intf {
    IpAddr   string
    RouterIf int32
}

type IPv4Neighbor struct {
    IpAddr   string
    MacAddr  string
    VlanId   int32
    RouterIf int32
}

func (obj *Vlan) UnmarshalHTTP(r *http.Request) error {
	var vlanobj Vlan
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
	 return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err = json.Unmarshal(body, &vlanobj); err != nil  {
		fmt.Println("# Vlan create called, unmarshal failed", vlanobj)
	}
	return nil
}

func (obj *IPv4Intf) UnmarshalHTTP(r *http.Request) error {
	var v4Intf IPv4Intf
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
	 return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v4Intf); err != nil  {
		fmt.Println("# IPv4Intf create called, unmarshal failed", v4Intf)
	}
	return nil
}

func (obj *IPv4Neighbor) UnmarshalHTTP(r *http.Request) error {
	var v4Nbr v4Nbr
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
	 return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err = json.Unmarshal(body, &v4Nbr); err != nil  {
		fmt.Println("# IPv4Neighbor create called, unmarshal failed", v4Nbr)
	}
	return nil
}

/* End - Asicd objects*/
