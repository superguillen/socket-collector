package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"strconv"
)

type SockStat struct {
	Local_Port                uint16
	Remote_Port               uint16
	Local_Addr                string
	Remote_Addr               string
	Status                    uint8
	Ca_state                  uint8
	Retransmits               uint8
	Probes                    uint8
	Backoff                   uint8
	Options                   uint8
	Snd_wscale                uint8
	Rcv_wscale                uint8
	Delivery_rate_app_limited uint8
	Rto                       uint32
	Ato                       uint32
	Snd_mss                   uint32
	Rcv_mss                   uint32
	Unacked                   uint32
	Sacked                    uint32
	Lost                      uint32
	Retrans                   uint32
	Fackets                   uint32
	Last_data_sent            uint32
	Last_ack_sent             uint32
	Last_data_recv            uint32
	Last_ack_recv             uint32
	Pmtu                      uint32
	Rcv_ssthresh              uint32
	Rtt                       uint32
	Rttvar                    uint32
	Snd_ssthresh              uint32
	Snd_cwnd                  uint32
	Advmss                    uint32
	Reordering                uint32
	Rcv_rtt                   uint32
	Rcv_space                 uint32
	Total_retrans             uint32
	Pacing_rate               uint64
	Max_pacing_rate           uint64
	Bytes_acked               uint64
	Bytes_received            uint64
	Segs_out                  uint32
	Segs_in                   uint32
	Notsent_bytes             uint32
	Min_rtt                   uint32
	Data_segs_in              uint32
	Data_segs_out             uint32
	Delivery_rate             uint64
	Busy_time                 uint64
	Rwnd_limited              uint64
	Sndbuf_limited            uint64
	Delivered                 uint32
	Delivered_ce              uint32
	Bytes_sent                uint64
	Bytes_retrans             uint64
	Dsack_dups                uint32
	Reord_seen                uint32
	Rcv_ooopack               uint32
	Snd_wnd                   uint32
}

func main() {
	records := GetSockStats()
	fmt.Printf("%+v\n", records)
}
func GetSockStats() map[string]SockStat {
	res, diagErr := netlink.SocketDiagTCPInfo(unix.AF_INET)
	if diagErr != nil {
		panic(diagErr)
	}

	records := make(map[string]SockStat)
	//var sockstat SockStat
	for idx, _ := range res {
		var sockstat SockStat
		//record := res[idx]
		if record := res[idx]; record != nil {
		sockstat = SockStat{
			Local_Port:                record.InetDiagMsg.ID.SourcePort,
			Remote_Port:               record.InetDiagMsg.ID.DestinationPort,
			Local_Addr:                record.InetDiagMsg.ID.Source.String(),
			Remote_Addr:               record.InetDiagMsg.ID.Destination.String(),
			Status:                    record.TCPInfo.State,
			Ca_state:                  record.TCPInfo.Ca_state,
			Retransmits:               record.TCPInfo.Retransmits,
			Probes:                    record.TCPInfo.Probes,
			Backoff:                   record.TCPInfo.Backoff,
			Options:                   record.TCPInfo.Options,
			Snd_wscale:                record.TCPInfo.Snd_wscale,
			Rcv_wscale:                record.TCPInfo.Rcv_wscale,
			Delivery_rate_app_limited: record.TCPInfo.Delivery_rate_app_limited,
			Rto:                       record.TCPInfo.Rto,
			Ato:                       record.TCPInfo.Ato,
			Snd_mss:                   record.TCPInfo.Snd_mss,
			Rcv_mss:                   record.TCPInfo.Rcv_mss,
			Unacked:                   record.TCPInfo.Unacked,
			Sacked:                    record.TCPInfo.Sacked,
			Lost:                      record.TCPInfo.Lost,
			Retrans:                   record.TCPInfo.Retrans,
			Fackets:                   record.TCPInfo.Fackets,
			Last_data_sent:            record.TCPInfo.Last_data_sent,
			Last_ack_sent:             record.TCPInfo.Last_ack_sent,
			Last_data_recv:            record.TCPInfo.Last_data_recv,
			Last_ack_recv:             record.TCPInfo.Last_ack_recv,
			Pmtu:                      record.TCPInfo.Pmtu,
			Rcv_ssthresh:              record.TCPInfo.Rcv_ssthresh,
			Rtt:                       record.TCPInfo.Rtt,
			Rttvar:                    record.TCPInfo.Rttvar,
			Snd_ssthresh:              record.TCPInfo.Snd_ssthresh,
			Snd_cwnd:                  record.TCPInfo.Snd_cwnd,
			Advmss:                    record.TCPInfo.Advmss,
			Reordering:                record.TCPInfo.Reordering,
			Rcv_rtt:                   record.TCPInfo.Rcv_rtt,
			Rcv_space:                 record.TCPInfo.Rcv_space,
			Total_retrans:             record.TCPInfo.Total_retrans,
			Pacing_rate:               record.TCPInfo.Pacing_rate,
			Max_pacing_rate:           record.TCPInfo.Max_pacing_rate,
			Bytes_acked:               record.TCPInfo.Bytes_acked,
			Bytes_received:            record.TCPInfo.Bytes_received,
			Segs_out:                  record.TCPInfo.Segs_out,
			Segs_in:                   record.TCPInfo.Segs_in,
			Notsent_bytes:             record.TCPInfo.Notsent_bytes,
			Min_rtt:                   record.TCPInfo.Min_rtt,
			Data_segs_in:              record.TCPInfo.Data_segs_in,
			Data_segs_out:             record.TCPInfo.Data_segs_out,
			Delivery_rate:             record.TCPInfo.Delivery_rate,
			Busy_time:                 record.TCPInfo.Busy_time,
			Rwnd_limited:              record.TCPInfo.Rwnd_limited,
			Sndbuf_limited:            record.TCPInfo.Sndbuf_limited,
			Delivered:                 record.TCPInfo.Delivered,
			Delivered_ce:              record.TCPInfo.Delivered_ce,
			Bytes_sent:                record.TCPInfo.Bytes_sent,
			Bytes_retrans:             record.TCPInfo.Bytes_retrans,
			Dsack_dups:                record.TCPInfo.Dsack_dups,
			Reord_seen:                record.TCPInfo.Reord_seen,
			Rcv_ooopack:               record.TCPInfo.Rcv_ooopack,
			Snd_wnd:                   record.TCPInfo.Snd_wnd,
		}
		}
		//fmt.Printf("%+v\n", record.InetDiagMsg)
		//fmt.Printf("%+v\n", record.TCPInfo)
		//fmt.Printf("%+v\n", sockstat)
		id := sockstat.Local_Addr + "_" + strconv.Itoa(int(sockstat.Local_Port)) + "_" + sockstat.Remote_Addr + "_" + strconv.Itoa(int(sockstat.Remote_Port))
		records[id] = sockstat
	}
	return records
}
