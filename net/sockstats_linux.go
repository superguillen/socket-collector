//+build linux

package sockstats

import (
	//"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"strconv"
)

// http://students.mimuw.edu.pl/lxr/source/include/net/tcp_states.h
var tcpStatuses = map[uint8]string{
	1: "ESTABLISHED",
	2: "SYN_SENT",
	3: "SYN_RECV",
	4: "FIN_WAIT1",
	5: "FIN_WAIT2",
	6: "TIME_WAIT",
	7: "CLOSE",
	8: "CLOSE_WAIT",
	9: "LAST_ACK",
	10: "LISTEN",
	11: "CLOSING",
}

func GetSockStats() map[string]SockStat {
	res, diagErr := netlink.SocketDiagTCPInfo(unix.AF_INET)
	if diagErr != nil {
		panic(diagErr)
	}

	//records := make(map[string]SockStat)
	records := make(map[string]SockStat)
	//var sockstat SockStat
	for idx, _ := range res {
		var sockstat SockStat
		//record := res[idx]
		if record := res[idx]; record != nil {
			
			var tcp_status string
			if record.TCPInfo.State <= 11 {
				tcp_status = tcpStatuses[record.TCPInfo.State]
			} else {
				tcp_status = "NONE"
			}

			sockstat = SockStat{
				Local_Port:                record.InetDiagMsg.ID.SourcePort,
				Remote_Port:               record.InetDiagMsg.ID.DestinationPort,
				Local_Addr:                record.InetDiagMsg.ID.Source.String(),
				Remote_Addr:               record.InetDiagMsg.ID.Destination.String(),
				Status:                    tcp_status,
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