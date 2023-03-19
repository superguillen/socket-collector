package main

import (
	"fmt"
	//"strconv"
	"encoding/json"
	"github.com/superguillen/socket-collector/net"
)

func main() {
	records := sockstats.GetSockStats()
	for _,record := range records {
		json_str,_ := json.Marshal(record)
		fmt.Printf("%+v\n", string(json_str))
	}
}
