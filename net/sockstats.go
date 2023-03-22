package sockstats

import (
	"math"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/slices"
)

type SockStat struct {
	Local_Port                uint16
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

type GlobalConnStatistics struct {
	IncomingConns map[Port]ConnStatistics
	OutgoingConns map[Port]ConnStatistics
}

type ConnStatistics struct {
	TotalConnections uint64
	NetStats         NetMetrics
	TCPInfoStats     map[string]interface{}
	TCPInfoStatsAcum map[string][]float64
}

type NetMetrics struct {
	Established int
	SynSent     int
	SynRecv     int
	FinWait1    int
	FinWait2    int
	TimeWait    int
	Close       int
	CloseWait   int
	LastAck     int
	Listen      int
	Closing     int
	None        int
}

type Port struct {
	PortNum uint16
	Addr    string
}

type StatType int

const (
	BASIC_STATS StatType = 0
	FULL_STATS           = 1
)

var FullStatsMetrics = map[string][]string{
	"Ca_state":                  {"max", "min", "mean"},
	"Retransmits":               {"max", "min", "mean", "p75", "p95", "p99"},
	"Probes":                    {"max", "min", "mean"},
	"Backoff":                   {"max", "min", "mean"},
	"Options":                   {"max", "min", "mean"},
	"Snd_wscale":                {"max", "min", "mean"},
	"Rcv_wscale":                {"max", "min", "mean"},
	"Delivery_rate_app_limited": {"max", "min", "mean"},
	"Rto":                       {"max", "min", "mean"},
	"Ato":                       {"max", "min", "mean"},
	"Snd_mss":                   {"max", "min", "mean"},
	"Rcv_mss":                   {"max", "min", "mean"},
	"Unacked":                   {"max", "min", "mean"},
	"Sacked":                    {"max", "min", "mean"},
	"Lost":                      {"max", "min", "mean"},
	"Retrans":                   {"max", "min", "mean", "p75", "p95", "p99"},
	"Fackets":                   {"max", "min", "mean"},
	"Last_data_sent":            {"max", "min", "mean"},
	"Last_ack_sent":             {"max", "min", "mean"},
	"Last_data_recv":            {"max", "min", "mean"},
	"Last_ack_recv":             {"max", "min", "mean"},
	"Pmtu":                      {"max", "min", "mean"},
	"Rcv_ssthresh":              {"max", "min", "mean"},
	"Rtt":                       {"max", "min", "mean", "p75", "p95", "p99"},
	"Rttvar":                    {"max", "min", "mean"},
	"Snd_ssthresh":              {"max", "min", "mean"},
	"Snd_cwnd":                  {"max", "min", "mean"},
	"Advmss":                    {"max", "min", "mean"},
	"Reordering":                {"max", "min", "mean"},
	"Rcv_rtt":                   {"max", "min", "mean"},
	"Rcv_space":                 {"max", "min", "mean"},
	"Total_retrans":             {"max", "min", "mean"},
	"Pacing_rate":               {"max", "min", "mean"},
	"Max_pacing_rate":           {"max", "min", "mean"},
	"Bytes_acked":               {"max", "min", "mean"},
	"Bytes_received":            {"max", "min", "mean", "p75", "p95", "p99"},
	"Segs_out":                  {"max", "min", "mean"},
	"Segs_in":                   {"max", "min", "mean"},
	"Notsent_bytes":             {"max", "min", "mean"},
	"Min_rtt":                   {"max", "min", "mean"},
	"Data_segs_in":              {"max", "min", "mean"},
	"Data_segs_out":             {"max", "min", "mean"},
	"Delivery_rate":             {"max", "min", "mean", "p75", "p95", "p99"},
	"Busy_time":                 {"max", "min", "mean"},
	"Rwnd_limited":              {"max", "min", "mean"},
	"Sndbuf_limited":            {"max", "min", "mean"},
	"Delivered":                 {"max", "min", "mean"},
	"Delivered_ce":              {"max", "min", "mean"},
	"Bytes_sent":                {"max", "min", "mean", "p75", "p95", "p99"},
	"Bytes_retrans":             {"max", "min", "mean"},
	"Dsack_dups":                {"max", "min", "mean"},
	"Reord_seen":                {"max", "min", "mean"},
	"Rcv_ooopack":               {"max", "min", "mean"},
	"Snd_wnd":                   {"max", "min", "mean"},
}

var BasicStatsMetrics = map[string][]string{
	"Ca_state":                  {"mean"},
	"Retransmits":               {"mean", "p99"},
	"Probes":                    {"mean"},
	"Backoff":                   {"mean"},
	"Options":                   {"mean"},
	"Snd_wscale":                {"mean"},
	"Rcv_wscale":                {"mean"},
	"Delivery_rate_app_limited": {"mean"},
	"Rto":                       {"mean"},
	"Ato":                       {"mean"},
	"Snd_mss":                   {"mean"},
	"Rcv_mss":                   {"mean"},
	"Unacked":                   {"mean"},
	"Sacked":                    {"mean"},
	"Lost":                      {"mean"},
	"Retrans":                   {"mean", "p99"},
	"Fackets":                   {"mean"},
	"Last_data_sent":            {"mean"},
	"Last_ack_sent":             {"mean"},
	"Last_data_recv":            {"mean"},
	"Last_ack_recv":             {"mean"},
	"Pmtu":                      {"mean"},
	"Rcv_ssthresh":              {"mean"},
	"Rtt":                       {"mean", "p99"},
	"Rttvar":                    {"mean"},
	"Snd_ssthresh":              {"mean"},
	"Snd_cwnd":                  {"mean"},
	"Advmss":                    {"mean"},
	"Reordering":                {"mean"},
	"Rcv_rtt":                   {"mean"},
	"Rcv_space":                 {"mean"},
	"Total_retrans":             {"mean"},
	"Pacing_rate":               {"mean"},
	"Max_pacing_rate":           {"mean"},
	"Bytes_acked":               {"mean"},
	"Bytes_received":            {"mean", "p99"},
	"Segs_out":                  {"mean"},
	"Segs_in":                   {"mean"},
	"Notsent_bytes":             {"mean"},
	"Min_rtt":                   {"mean"},
	"Data_segs_in":              {"mean"},
	"Data_segs_out":             {"mean"},
	"Delivery_rate":             {"mean", "p99"},
	"Busy_time":                 {"mean"},
	"Rwnd_limited":              {"mean"},
	"Sndbuf_limited":            {"mean"},
	"Delivered":                 {"mean"},
	"Delivered_ce":              {"mean"},
	"Bytes_sent":                {"mean", "p99"},
	"Bytes_retrans":             {"mean"},
	"Dsack_dups":                {"mean"},
	"Reord_seen":                {"mean"},
	"Rcv_ooopack":               {"mean"},
	"Snd_wnd":                   {"mean"},
}

var FULL_METRICS = []string{"Ca_state", "Retransmits", "Probes", "Backoff", "Options", "Snd_wscale", "Rcv_wscale",
	"Delivery_rate_app_limited", "Rto", "Ato", "Snd_mss", "Rcv_mss", "Unacked", "Sacked", "Lost", "Retrans", "Fackets",
	"Last_data_sent", "Last_ack_sent", "Last_data_recv", "Last_ack_recv", "Pmtu", "Rcv_ssthresh", "Rtt", "Rttvar",
	"Snd_ssthresh", "Snd_cwnd", "Advmss", "Reordering", "Rcv_rtt", "Rcv_space", "Total_retrans", "Pacing_rate", "Max_pacing_rate",
	"Bytes_acked", "Bytes_received", "Segs_out", "Segs_in", "Notsent_bytes", "Min_rtt", "Data_segs_in", "Data_segs_out", "Delivery_rate",
	"Busy_time", "Rwnd_limited", "Sndbuf_limited", "Delivered", "Delivered_ce", "Bytes_sent", "Bytes_retrans", "Dsack_dups", "Reord_seen", "Rcv_ooopack", "Snd_wnd"}

var BASIC_METRICS = []string{"Rtt", "Rttvar", "Min_rtt", "Delivery_rate", "Retransmits", "Total_retrans", "Bytes_sent", "Bytes_received", "Bytes_retrans"}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetConnStatistics(metrics []string, statType StatType) (GlobalConnStatistics, error) {

	if metrics == nil {
		metrics = BASIC_METRICS
	}

	globalConnStadistics := GlobalConnStatistics{
		IncomingConns: map[Port]ConnStatistics{},
		OutgoingConns: map[Port]ConnStatistics{},
	}

	var status_port ConnStatistics
	sockstats_list, listen_ports, diagErr := GetConnections()

	if diagErr != nil {
		return globalConnStadistics, diagErr
	}

	for _, record := range sockstats_list {
		var port Port
		var ok bool
		// isListen := false
		if slices.Contains(listen_ports, record.Local_Port) {
			// isListen = true
			port = Port{
				PortNum: record.Local_Port,
			}

			status_port, ok = globalConnStadistics.IncomingConns[port]
			if !ok {
				globalConnStadistics.IncomingConns[port] = ConnStatistics{
					TCPInfoStats:     map[string]interface{}{},
					TCPInfoStatsAcum: map[string][]float64{},
				}

				status_port = globalConnStadistics.IncomingConns[port]
			}
		} else {
			port = Port{
				PortNum: record.Remote_Port,
				Addr:    record.Remote_Addr,
			}
			status_port, ok = globalConnStadistics.OutgoingConns[port]
			if !ok {
				globalConnStadistics.OutgoingConns[port] = ConnStatistics{
					TCPInfoStats:     map[string]interface{}{},
					TCPInfoStatsAcum: map[string][]float64{},
				}
				status_port = globalConnStadistics.OutgoingConns[port]
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

		status_port.TCPInfoStatsAcum["Ca_state"] = append(status_port.TCPInfoStatsAcum["Ca_state"], float64(record.Ca_state))
		status_port.TCPInfoStatsAcum["Retransmits"] = append(status_port.TCPInfoStatsAcum["Retransmits"], float64(record.Retransmits))
		status_port.TCPInfoStatsAcum["Probes"] = append(status_port.TCPInfoStatsAcum["Probes"], float64(record.Probes))
		status_port.TCPInfoStatsAcum["Backoff"] = append(status_port.TCPInfoStatsAcum["Backoff"], float64(record.Backoff))
		status_port.TCPInfoStatsAcum["Options"] = append(status_port.TCPInfoStatsAcum["Options"], float64(record.Options))
		status_port.TCPInfoStatsAcum["Snd_wscale"] = append(status_port.TCPInfoStatsAcum["Snd_wscale"], float64(record.Snd_wscale))
		status_port.TCPInfoStatsAcum["Rcv_wscale"] = append(status_port.TCPInfoStatsAcum["Rcv_wscale"], float64(record.Rcv_wscale))
		status_port.TCPInfoStatsAcum["Delivery_rate_app_limited"] = append(status_port.TCPInfoStatsAcum["Delivery_rate_app_limited"], float64(record.Delivery_rate_app_limited))
		status_port.TCPInfoStatsAcum["Rto"] = append(status_port.TCPInfoStatsAcum["Rto"], float64(record.Rto))
		status_port.TCPInfoStatsAcum["Ato"] = append(status_port.TCPInfoStatsAcum["Ato"], float64(record.Ato))
		status_port.TCPInfoStatsAcum["Snd_mss"] = append(status_port.TCPInfoStatsAcum["Snd_mss"], float64(record.Snd_mss))
		status_port.TCPInfoStatsAcum["Rcv_mss"] = append(status_port.TCPInfoStatsAcum["Rcv_mss"], float64(record.Rcv_mss))
		status_port.TCPInfoStatsAcum["Unacked"] = append(status_port.TCPInfoStatsAcum["Unacked"], float64(record.Unacked))
		status_port.TCPInfoStatsAcum["Sacked"] = append(status_port.TCPInfoStatsAcum["Sacked"], float64(record.Sacked))
		status_port.TCPInfoStatsAcum["Lost"] = append(status_port.TCPInfoStatsAcum["Lost"], float64(record.Lost))
		status_port.TCPInfoStatsAcum["Retrans"] = append(status_port.TCPInfoStatsAcum["Retrans"], float64(record.Retrans))
		status_port.TCPInfoStatsAcum["Fackets"] = append(status_port.TCPInfoStatsAcum["Fackets"], float64(record.Fackets))
		status_port.TCPInfoStatsAcum["Last_data_sent"] = append(status_port.TCPInfoStatsAcum["Last_data_sent"], float64(record.Last_data_sent))
		status_port.TCPInfoStatsAcum["Last_ack_sent"] = append(status_port.TCPInfoStatsAcum["Last_ack_sent"], float64(record.Last_ack_sent))
		status_port.TCPInfoStatsAcum["Last_data_recv"] = append(status_port.TCPInfoStatsAcum["Last_data_recv"], float64(record.Last_data_recv))
		status_port.TCPInfoStatsAcum["Last_ack_recv"] = append(status_port.TCPInfoStatsAcum["Last_ack_recv"], float64(record.Last_ack_recv))
		status_port.TCPInfoStatsAcum["Pmtu"] = append(status_port.TCPInfoStatsAcum["Pmtu"], float64(record.Pmtu))
		status_port.TCPInfoStatsAcum["Rcv_ssthresh"] = append(status_port.TCPInfoStatsAcum["Rcv_ssthresh"], float64(record.Rcv_ssthresh))
		status_port.TCPInfoStatsAcum["Rtt"] = append(status_port.TCPInfoStatsAcum["Rtt"], float64(record.Rtt))
		status_port.TCPInfoStatsAcum["Rttvar"] = append(status_port.TCPInfoStatsAcum["Rttvar"], float64(record.Rttvar))
		status_port.TCPInfoStatsAcum["Snd_ssthresh"] = append(status_port.TCPInfoStatsAcum["Snd_ssthresh"], float64(record.Snd_ssthresh))
		status_port.TCPInfoStatsAcum["Snd_cwnd"] = append(status_port.TCPInfoStatsAcum["Snd_cwnd"], float64(record.Snd_cwnd))
		status_port.TCPInfoStatsAcum["Advmss"] = append(status_port.TCPInfoStatsAcum["Advmss"], float64(record.Advmss))
		status_port.TCPInfoStatsAcum["Reordering"] = append(status_port.TCPInfoStatsAcum["Reordering"], float64(record.Reordering))
		status_port.TCPInfoStatsAcum["Rcv_rtt"] = append(status_port.TCPInfoStatsAcum["Rcv_rtt"], float64(record.Rcv_rtt))
		status_port.TCPInfoStatsAcum["Rcv_space"] = append(status_port.TCPInfoStatsAcum["Rcv_space"], float64(record.Rcv_space))
		status_port.TCPInfoStatsAcum["Total_retrans"] = append(status_port.TCPInfoStatsAcum["Total_retrans"], float64(record.Total_retrans))
		status_port.TCPInfoStatsAcum["Pacing_rate"] = append(status_port.TCPInfoStatsAcum["Pacing_rate"], float64(record.Pacing_rate))
		status_port.TCPInfoStatsAcum["Max_pacing_rate"] = append(status_port.TCPInfoStatsAcum["Max_pacing_rate"], float64(record.Max_pacing_rate))
		status_port.TCPInfoStatsAcum["Bytes_acked"] = append(status_port.TCPInfoStatsAcum["Bytes_acked"], float64(record.Bytes_acked))
		status_port.TCPInfoStatsAcum["Bytes_received"] = append(status_port.TCPInfoStatsAcum["Bytes_received"], float64(record.Bytes_received))
		status_port.TCPInfoStatsAcum["Segs_out"] = append(status_port.TCPInfoStatsAcum["Segs_out"], float64(record.Segs_out))
		status_port.TCPInfoStatsAcum["Segs_in"] = append(status_port.TCPInfoStatsAcum["Segs_in"], float64(record.Segs_in))
		status_port.TCPInfoStatsAcum["Notsent_bytes"] = append(status_port.TCPInfoStatsAcum["Notsent_bytes"], float64(record.Notsent_bytes))
		status_port.TCPInfoStatsAcum["Min_rtt"] = append(status_port.TCPInfoStatsAcum["Min_rtt"], float64(record.Min_rtt))
		status_port.TCPInfoStatsAcum["Data_segs_in"] = append(status_port.TCPInfoStatsAcum["Data_segs_in"], float64(record.Data_segs_in))
		status_port.TCPInfoStatsAcum["Data_segs_out"] = append(status_port.TCPInfoStatsAcum["Data_segs_out"], float64(record.Data_segs_out))
		status_port.TCPInfoStatsAcum["Delivery_rate"] = append(status_port.TCPInfoStatsAcum["Delivery_rate"], float64(record.Delivery_rate))
		status_port.TCPInfoStatsAcum["Busy_time"] = append(status_port.TCPInfoStatsAcum["Busy_time"], float64(record.Busy_time))
		status_port.TCPInfoStatsAcum["Rwnd_limited"] = append(status_port.TCPInfoStatsAcum["Rwnd_limited"], float64(record.Rwnd_limited))
		status_port.TCPInfoStatsAcum["Sndbuf_limited"] = append(status_port.TCPInfoStatsAcum["Sndbuf_limited"], float64(record.Sndbuf_limited))
		status_port.TCPInfoStatsAcum["Delivered"] = append(status_port.TCPInfoStatsAcum["Delivered"], float64(record.Delivered))
		status_port.TCPInfoStatsAcum["Delivered_ce"] = append(status_port.TCPInfoStatsAcum["Delivered_ce"], float64(record.Delivered_ce))
		status_port.TCPInfoStatsAcum["Bytes_sent"] = append(status_port.TCPInfoStatsAcum["Bytes_sent"], float64(record.Bytes_sent))
		status_port.TCPInfoStatsAcum["Bytes_retrans"] = append(status_port.TCPInfoStatsAcum["Bytes_retrans"], float64(record.Bytes_retrans))
		status_port.TCPInfoStatsAcum["Dsack_dups"] = append(status_port.TCPInfoStatsAcum["Dsack_dups"], float64(record.Dsack_dups))
		status_port.TCPInfoStatsAcum["Reord_seen"] = append(status_port.TCPInfoStatsAcum["Reord_seen"], float64(record.Reord_seen))
		status_port.TCPInfoStatsAcum["Rcv_ooopack"] = append(status_port.TCPInfoStatsAcum["Rcv_ooopack"], float64(record.Rcv_ooopack))
		status_port.TCPInfoStatsAcum["Snd_wnd"] = append(status_port.TCPInfoStatsAcum["Snd_wnd"], float64(record.Snd_wnd))
		status_port.TotalConnections = 100
		// fmt.Printf("====================================================================================\n")
		// json_str,_ := json.Marshal(port)
		// fmt.Printf("Status Port: %+v\n", string(json_str))
		// json_str,_ = json.Marshal(status_port)
		// fmt.Printf("%+v\n", string(json_str))
		// if isListen {
		// 	fmt.Printf("++++++++++++++++++++++INCOMING+++++++++++++++++++\n")
		// 	json_str,_ = json.Marshal(globalConnStadistics.IncomingConns[port].NetStats)
		// 	fmt.Printf("======= NetStats: %+v\n", string(json_str))
		// 	json_str,_ = json.Marshal(globalConnStadistics.IncomingConns[port].TCPInfoStatsAcum)
		// 	fmt.Printf("======= TCPInfoStatsAcum: %+v\n", string(json_str))
		// } else {
		// 	fmt.Printf("++++++++++++++++++++++OUTGOING+++++++++++++++++++\n")
		// 	json_str,_ = json.Marshal(globalConnStadistics.OutgoingConns[port].NetStats)
		// 	fmt.Printf("======= NetStats: %+v\n", string(json_str))
		// 	json_str,_ = json.Marshal(globalConnStadistics.OutgoingConns[port].TCPInfoStatsAcum)
		// 	fmt.Printf("======= TCPInfoStatsAcum: %+v\n", string(json_str))

		// }
	}

	var sockMetricList map[string][]string
	if statType == BASIC_STATS {
		sockMetricList = BasicStatsMetrics
	} else {
		sockMetricList = FullStatsMetrics
	}

	for _, port_stat := range globalConnStadistics.IncomingConns {
		for _, metric_name := range metrics {
			stat_list := sockMetricList[metric_name]
			for _, stat := range stat_list {
				var metric_value float64
				switch stat {
				case "mean":
					metric_value, _ = stats.Mean(port_stat.TCPInfoStatsAcum[metric_name])
				case "max":
					metric_value, _ = stats.Max(port_stat.TCPInfoStatsAcum[metric_name])
				case "min":
					metric_value, _ = stats.Min(port_stat.TCPInfoStatsAcum[metric_name])
				case "p75":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.75)
				case "p95":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.95)
				case "p99":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.99)
				}
				port_stat.TCPInfoStats[metric_name+"_"+stat] = roundFloat(metric_value, 2)
			}
		}
		// fmt.Printf("====================================================================================\n")
		// json_str, _ := json.Marshal(port)
		// fmt.Printf("Status Port: %+v\n", string(json_str))
		// fmt.Printf("++++++++++++++++++++++INCOMING+++++++++++++++++++\n")
		// fmt.Printf("======= NetStats: %+v\n", port_stat.NetStats)
		// fmt.Printf("======= TCPInfoStats: %+v\n", port_stat.TCPInfoStats)
		// fmt.Printf("======= TCPInfoStatsAcum: %+v\n", port_stat.TCPInfoStatsAcum)
		// fmt.Printf("======= TotalConnections: %+v\n", port_stat.TotalConnections)
	}

	for _, port_stat := range globalConnStadistics.OutgoingConns {
		for _, metric_name := range metrics {
			stat_list := sockMetricList[metric_name]
			for _, stat := range stat_list {
				var metric_value float64
				switch stat {
				case "mean":
					metric_value, _ = stats.Mean(port_stat.TCPInfoStatsAcum[metric_name])
				case "max":
					metric_value, _ = stats.Max(port_stat.TCPInfoStatsAcum[metric_name])
				case "min":
					metric_value, _ = stats.Min(port_stat.TCPInfoStatsAcum[metric_name])
				case "p75":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.75)
				case "p95":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.95)
				case "p99":
					metric_value, _ = stats.PercentileNearestRank(port_stat.TCPInfoStatsAcum[metric_name], 0.99)
				}
				port_stat.TCPInfoStats[metric_name+"_"+stat] = roundFloat(metric_value, 2)
			}
		}
		// fmt.Printf("====================================================================================\n")
		// json_str, _ := json.Marshal(port)
		// fmt.Printf("Status Port: %+v\n", string(json_str))
		// fmt.Printf("++++++++++++++++++++++OUTGOING+++++++++++++++++++\n")
		// fmt.Printf("======= NetStats: %+v\n", port_stat.NetStats)
		// fmt.Printf("======= TCPInfoStats: %+v\n", port_stat.TCPInfoStats)
		// fmt.Printf("======= TCPInfoStatsAcum: %+v\n", port_stat.TCPInfoStatsAcum)

	}
	// for _,port_stat := range globalConnStadistics.OutgoingConns {
	// 	data := stats.LoadRawData(port_stat.TCPInfoStatsAcum.Rtt)
	// 	port_stat.TCPInfoStatsAvg.Rtt,_ = stats.Mean(data)
	// 	// fmt.Printf("%+v\n", port_stat.TCPInfoStatsAvg.Rtt)
	// }

	return globalConnStadistics, nil
}
