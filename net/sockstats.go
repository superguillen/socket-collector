package sockstats

import (
	"golang.org/x/exp/slices"
)

type SockStat struct {
	Local_Port                uint16  `json: "local_port" `
	Remote_Port               uint16
	Local_Addr                string
	Remote_Addr               string
	Status                    string
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

type SockStatAcumm struct {
	Local_Port                uint16 
	Remote_Port               uint16
	Local_Addr                string
	Remote_Addr               string
	NetStats                NetMetrics
	TCPInfoStats			SockStatAcumMetrics
}

type NetMetrics struct{
	Established             int
	SynSent                 int
	SynRecv                 int
	FinWait1                int
	FinWait2                int
	TimeWait                int
	Close                   int
	CloseWait               int
	LastAck                 int
	Listen                  int
	Closing                 int
	None                    int
}

type SockStatAcumMetrics struct{
	Ca_state                  []uint8
	Retransmits               []uint8
	Probes                    []uint8
	Backoff                   []uint8
	Options                   []uint8
	Snd_wscale                []uint8
	Rcv_wscale                []uint8
	Delivery_rate_app_limited []uint8
	Rto                       []uint32
	Ato                       []uint32
	Snd_mss                   []uint32
	Rcv_mss                   []uint32
	Unacked                   []uint32
	Sacked                    []uint32
	Lost                      []uint32
	Retrans                   []uint32
	Fackets                   []uint32
	Last_data_sent            []uint32
	Last_ack_sent             []uint32
	Last_data_recv            []uint32
	Last_ack_recv             []uint32
	Pmtu                      []uint32
	Rcv_ssthresh              []uint32
	Rtt                       []uint32
	Rttvar                    []uint32
	Snd_ssthresh              []uint32
	Snd_cwnd                  []uint32
	Advmss                    []uint32
	Reordering                []uint32
	Rcv_rtt                   []uint32
	Rcv_space                 []uint32
	Total_retrans             []uint32
	Pacing_rate               []uint64
	Max_pacing_rate           []uint64
	Bytes_acked               []uint64
	Bytes_received            []uint64
	Segs_out                  []uint32
	Segs_in                   []uint32
	Notsent_bytes             []uint32
	Min_rtt                   []uint32
	Data_segs_in              []uint32
	Data_segs_out             []uint32
	Delivery_rate             []uint64
	Busy_time                 []uint64
	Rwnd_limited              []uint64
	Sndbuf_limited            []uint64
	Delivered                 []uint32
	Delivered_ce              []uint32
	Bytes_sent                []uint64
	Bytes_retrans             []uint64
	Dsack_dups                []uint32
	Reord_seen                []uint32
	Rcv_ooopack               []uint32
	Snd_wnd                   []uint32
}

type Port struct {
	PortNum uint16
	Addr 	string 
}

func GetSockStatsAcum() (map[Port]*SockStatAcumm,map[Port]*SockStatAcumm){
	sockstats_list,listen_ports := GetSockStats()

	status_listen_ports := map[Port]*SockStatAcumm{}
	status_remote_ports := map[Port]*SockStatAcumm{}

	for _,record := range sockstats_list {
		var status_port *SockStatAcumm
		var ok bool
		if slices.Contains(listen_ports,record.Local_Port) {
			port := Port{
				PortNum: record.Local_Port,
			}			
			status_port, ok = status_listen_ports[port]
			if !ok {
				status_port = &SockStatAcumm{
					Local_Port: record.Local_Port,
					Local_Addr: record.Local_Addr,
				}
				status_listen_ports[port] = status_port
			}
		} else {
			port := Port{
				PortNum: record.Remote_Port,
				Addr: record.Remote_Addr,
			}
			status_port, ok = status_remote_ports[port]
			if !ok {
				status_port = &SockStatAcumm{
					Remote_Port: record.Remote_Port,
					Remote_Addr: record.Remote_Addr,
				}
				status_remote_ports[port] = status_port
			}
		}

		switch record.Status {
		case "LISTEN":
			status_port.NetStats.Listen += 1
		case "ESTABLISHED":
			status_port.NetStats.Established += 1
		case "SYN_SENT":
			status_port.NetStats.SynSent += 1
		case "SYN_RECV":
			status_port.NetStats.SynRecv += 1
		case "FIN_WAIT1":
			status_port.NetStats.FinWait1 += 1
		case "FIN_WAIT2":
			status_port.NetStats.FinWait2 += 1
		case "TIME_WAIT":
			status_port.NetStats.TimeWait += 1
		case "CLOSE":
			status_port.NetStats.Close += 1
		case "CLOSE_WAIT":
			status_port.NetStats.CloseWait += 1
		case "LAST_ACK":
			status_port.NetStats.LastAck += 1
		case "CLOSING":
			status_port.NetStats.Closing += 1
		case "NONE":
			status_port.NetStats.None += 1
		}

		status_port.TCPInfoStats.Ca_state = append(status_port.TCPInfoStats.Ca_state,record.Ca_state)
		status_port.TCPInfoStats.Retransmits = append(status_port.TCPInfoStats.Retransmits,record.Retransmits)
		status_port.TCPInfoStats.Probes = append(status_port.TCPInfoStats.Probes,record.Probes)
		status_port.TCPInfoStats.Backoff = append(status_port.TCPInfoStats.Backoff,record.Backoff)
		status_port.TCPInfoStats.Options = append(status_port.TCPInfoStats.Options,record.Options)
		status_port.TCPInfoStats.Snd_wscale = append(status_port.TCPInfoStats.Snd_wscale,record.Snd_wscale)
		status_port.TCPInfoStats.Rcv_wscale = append(status_port.TCPInfoStats.Rcv_wscale,record.Rcv_wscale)
		status_port.TCPInfoStats.Delivery_rate_app_limited = append(status_port.TCPInfoStats.Delivery_rate_app_limited,record.Delivery_rate_app_limited)
		status_port.TCPInfoStats.Rto = append(status_port.TCPInfoStats.Rto,record.Rto)
		status_port.TCPInfoStats.Ato = append(status_port.TCPInfoStats.Ato,record.Ato)
		status_port.TCPInfoStats.Snd_mss = append(status_port.TCPInfoStats.Snd_mss,record.Snd_mss)
		status_port.TCPInfoStats.Rcv_mss = append(status_port.TCPInfoStats.Rcv_mss,record.Rcv_mss)
		status_port.TCPInfoStats.Unacked = append(status_port.TCPInfoStats.Unacked,record.Unacked)
		status_port.TCPInfoStats.Sacked = append(status_port.TCPInfoStats.Sacked,record.Sacked)
		status_port.TCPInfoStats.Lost = append(status_port.TCPInfoStats.Lost,record.Lost)
		status_port.TCPInfoStats.Retrans = append(status_port.TCPInfoStats.Retrans,record.Retrans)
		status_port.TCPInfoStats.Fackets = append(status_port.TCPInfoStats.Fackets,record.Fackets)
		status_port.TCPInfoStats.Last_data_sent = append(status_port.TCPInfoStats.Last_data_sent,record.Last_data_sent)
		status_port.TCPInfoStats.Last_ack_sent = append(status_port.TCPInfoStats.Last_ack_sent,record.Last_ack_sent)
		status_port.TCPInfoStats.Last_data_recv = append(status_port.TCPInfoStats.Last_data_recv,record.Last_data_recv)
		status_port.TCPInfoStats.Last_ack_recv = append(status_port.TCPInfoStats.Last_ack_recv,record.Last_ack_recv)
		status_port.TCPInfoStats.Pmtu = append(status_port.TCPInfoStats.Pmtu,record.Pmtu)
		status_port.TCPInfoStats.Rcv_ssthresh = append(status_port.TCPInfoStats.Rcv_ssthresh,record.Rcv_ssthresh)
		status_port.TCPInfoStats.Rtt = append(status_port.TCPInfoStats.Rtt,record.Rtt)
		status_port.TCPInfoStats.Rttvar = append(status_port.TCPInfoStats.Rttvar,record.Rttvar)
		status_port.TCPInfoStats.Snd_ssthresh = append(status_port.TCPInfoStats.Snd_ssthresh,record.Snd_ssthresh)
		status_port.TCPInfoStats.Snd_cwnd = append(status_port.TCPInfoStats.Snd_cwnd,record.Snd_cwnd)
		status_port.TCPInfoStats.Advmss = append(status_port.TCPInfoStats.Advmss,record.Advmss)
		status_port.TCPInfoStats.Reordering = append(status_port.TCPInfoStats.Reordering,record.Reordering)
		status_port.TCPInfoStats.Rcv_rtt = append(status_port.TCPInfoStats.Rcv_rtt,record.Rcv_rtt)
		status_port.TCPInfoStats.Rcv_space = append(status_port.TCPInfoStats.Rcv_space,record.Rcv_space)
		status_port.TCPInfoStats.Total_retrans = append(status_port.TCPInfoStats.Total_retrans,record.Total_retrans)
		status_port.TCPInfoStats.Pacing_rate = append(status_port.TCPInfoStats.Pacing_rate,record.Pacing_rate)
		status_port.TCPInfoStats.Max_pacing_rate = append(status_port.TCPInfoStats.Max_pacing_rate,record.Max_pacing_rate)
		status_port.TCPInfoStats.Bytes_acked = append(status_port.TCPInfoStats.Bytes_acked,record.Bytes_acked)
		status_port.TCPInfoStats.Bytes_received = append(status_port.TCPInfoStats.Bytes_received,record.Bytes_received)
		status_port.TCPInfoStats.Segs_out = append(status_port.TCPInfoStats.Segs_out,record.Segs_out)
		status_port.TCPInfoStats.Segs_in = append(status_port.TCPInfoStats.Segs_in,record.Segs_in)
		status_port.TCPInfoStats.Notsent_bytes = append(status_port.TCPInfoStats.Notsent_bytes,record.Notsent_bytes)
		status_port.TCPInfoStats.Min_rtt = append(status_port.TCPInfoStats.Min_rtt,record.Min_rtt)
		status_port.TCPInfoStats.Data_segs_in = append(status_port.TCPInfoStats.Data_segs_in,record.Data_segs_in)
		status_port.TCPInfoStats.Data_segs_out = append(status_port.TCPInfoStats.Data_segs_out,record.Data_segs_out)
		status_port.TCPInfoStats.Delivery_rate = append(status_port.TCPInfoStats.Delivery_rate,record.Delivery_rate)
		status_port.TCPInfoStats.Busy_time = append(status_port.TCPInfoStats.Busy_time,record.Busy_time)
		status_port.TCPInfoStats.Rwnd_limited = append(status_port.TCPInfoStats.Rwnd_limited,record.Rwnd_limited)
		status_port.TCPInfoStats.Sndbuf_limited = append(status_port.TCPInfoStats.Sndbuf_limited,record.Sndbuf_limited)
		status_port.TCPInfoStats.Delivered = append(status_port.TCPInfoStats.Delivered,record.Delivered)
		status_port.TCPInfoStats.Delivered_ce = append(status_port.TCPInfoStats.Delivered_ce,record.Delivered_ce)
		status_port.TCPInfoStats.Bytes_sent = append(status_port.TCPInfoStats.Bytes_sent,record.Bytes_sent)
		status_port.TCPInfoStats.Bytes_retrans = append(status_port.TCPInfoStats.Bytes_retrans,record.Bytes_retrans)
		status_port.TCPInfoStats.Dsack_dups = append(status_port.TCPInfoStats.Dsack_dups,record.Dsack_dups)
		status_port.TCPInfoStats.Reord_seen = append(status_port.TCPInfoStats.Reord_seen,record.Reord_seen)
		status_port.TCPInfoStats.Rcv_ooopack = append(status_port.TCPInfoStats.Rcv_ooopack,record.Rcv_ooopack)
		status_port.TCPInfoStats.Snd_wnd = append(status_port.TCPInfoStats.Snd_wnd,record.Snd_wnd)
		// json_str,_ := json.Marshal(record)
		// fmt.Printf("%+v\n", string(json_str))
	}
	return status_listen_ports,status_remote_ports
}