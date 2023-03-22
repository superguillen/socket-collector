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
	globalConnStadistics, diagErr := sockstats.GetConnStatistics(sockstats.BASIC_METRICS, sockstats.BASIC_STATS)

	if diagErr != nil {
		panic(diagErr)
	}

	fmt.Printf("=== Incoming connections\n")
	for key, status := range globalConnStadistics.IncomingConns {
		fmt.Printf("%+v %v %+v\n", key, status.Local_Addr, status.TCPInfoStats)
	}
	fmt.Printf("=== Outgoing connections\n")
	for key, status := range globalConnStadistics.OutgoingConns {
		fmt.Printf("%+v %+v\n", key, status.TCPInfoStats)
	}
}
