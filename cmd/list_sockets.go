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
	// sockstats_list,listen_ports := sockstats.GetSockStats()
	status_listen_ports,status_remote_ports := sockstats.GetSockStatsAcum()
	

	// status_listen_ports := map[sockstats.Port]*sockstats.SockStatAcumm{}
	// status_remote_ports := map[sockstats.Port]*sockstats.SockStatAcumm{}

	// for _,record := range sockstats_list {
	// 	var status_port *sockstats.SockStatAcumm
	// 	var ok bool
	// 	if slices.Contains(listen_ports,record.Local_Port) {
	// 		port := sockstats.Port{
	// 			PortNum: record.Local_Port,
	// 		}			
	// 		status_port, ok = status_listen_ports[port]
	// 		if !ok {
	// 			status_port = &sockstats.SockStatAcumm{
	// 				Local_Port: record.Local_Port,
	// 				Local_Addr: record.Local_Addr,
	// 			}
	// 			status_listen_ports[port] = status_port
	// 		}
	// 	} else {
	// 		port := sockstats.Port{
	// 			PortNum: record.Remote_Port,
	// 			Addr: record.Remote_Addr,
	// 		}
	// 		status_port, ok = status_remote_ports[port]
	// 		if !ok {
	// 			status_port = &sockstats.SockStatAcumm{
	// 				Remote_Port: record.Remote_Port,
	// 				Remote_Addr: record.Remote_Addr,
	// 			}
	// 			status_remote_ports[port] = status_port
	// 		}
	// 	}

	// 	switch record.Status {
	// 	case "LISTEN":
	// 		status_port.Listen += 1
	// 	case "ESTABLISHED":
	// 		status_port.Established += 1
	// 	case "SYN_SENT":
	// 		status_port.SynSent += 1
	// 	case "SYN_RECV":
	// 		status_port.SynRecv += 1
	// 	case "FIN_WAIT1":
	// 		status_port.FinWait1 += 1
	// 	case "FIN_WAIT2":
	// 		status_port.FinWait2 += 1
	// 	case "TIME_WAIT":
	// 		status_port.TimeWait += 1
	// 	case "CLOSE":
	// 		status_port.Close += 1
	// 	case "CLOSE_WAIT":
	// 		status_port.CloseWait += 1
	// 	case "LAST_ACK":
	// 		status_port.LastAck += 1
	// 	case "CLOSING":
	// 		status_port.Closing += 1
	// 	case "NONE":
	// 		status_port.None += 1
	// 	}			

	// 	// json_str,_ := json.Marshal(record)
	// 	// fmt.Printf("%+v\n", string(json_str))
	// }



	fmt.Printf("=== Incoming connections\n")
	for _,status := range status_listen_ports {
		json_str,_ := json.Marshal(status)
		fmt.Printf("%+v\n", string(json_str))
	}
	fmt.Printf("=== Outgoing connections\n")
	for _,status := range status_remote_ports {
		json_str,_ := json.Marshal(status)
		fmt.Printf("%+v\n", string(json_str))
	}	
}
