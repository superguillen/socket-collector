package main

import (
	"fmt"
	//"strconv"
	"encoding/json"
	"github.com/superguillen/socket-collector/net"
	// "github.com/superguillen/socket-collector/slice"
	// "golang.org/x/exp/slices"
)

func main() {
	//Get socket list
	globalConnStadistics := sockstats.GetConnStatistics()

	fmt.Printf("=== Incoming connections\n")
	for key,status := range globalConnStadistics.IncomingConns {
		json_str,_ := json.Marshal(status)
		fmt.Printf("%+v %+v\n", key,string(json_str))
	}
	fmt.Printf("=== Outgoing connections\n")
	for key,status := range globalConnStadistics.OutgoingConns {
		json_str,_ := json.Marshal(status)
		fmt.Printf("%+v %+v\n", key,string(json_str))
	}	
}
