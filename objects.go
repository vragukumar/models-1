package models
import ("fmt"
		  "encoding/json"
		  "io"
		  "io/ioutil"
		  "net/http")

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

type Vlan struct {
	 VlanId  int
	 PortMap  string
}

func (obj *Vlan) UnmarshalHTTP(r *http.Request) error {
	 return nil
}

