package main

import (
	"fmt"
	//"strconv"

	sockstats "github.com/superguillen/socket-collector/net"
	// "github.com/superguillen/socket-collector/slice"
	// "golang.org/x/exp/slices"
)

func main() {
	//Get socket list
	globalConnStadistics := sockstats.GetConnStatistics()

	fmt.Printf("=== Incoming connections\n")
	for key, status := range globalConnStadistics.IncomingConns {
		//json_str,_ := json.Marshal(status.TCPInfoStats)
		fmt.Printf("%+v %+v\n", key, status.TCPInfoStats)
	}
	fmt.Printf("=== Outgoing connections\n")
	for key, status := range globalConnStadistics.OutgoingConns {
		// json_str, _ := json.Marshal(status.TCPInfoStats)
		fmt.Printf("%+v %+v\n", key, status.TCPInfoStats)
	}
}
