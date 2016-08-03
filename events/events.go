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

package events

import (
	"time"
)

type OwnerId uint8
type EventId uint32

type EventBase struct {
	OwnerId     OwnerId
	OwnerName   string
	EvtId       EventId
	EventName   string
	TimeStamp   time.Time
	Description string
	SrcObjName  string
}

type Event struct {
	EventBase
	SrcObjKey interface{}
}

type KeyMap map[string]interface{}

var EventKeyMap map[string]KeyMap = map[string]KeyMap{
	"ASICD": AsicdEventKeyMap,
	"ARPD":  ArpdEventKeyMap,
	"BGPD":  BGPdEventKeyMap,
	"LLDP":  LLDPEventKeyMap,
}

type EventObject struct {
	OwnerName   string
	EventName   string
	TimeStamp   string
	SrcObjName  string
	SrcObjKey   string
	Description string
}