package sockstats

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

// https://itnext.io/generic-map-filter-and-reduce-in-go-3845781a591c

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type SliceIterator[T any] struct {
	Elements []T
	value    T
	index    int
}

// Create an iterator over the slice xs
func NewSliceIterator[T any](xs []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: xs,
	}
}

// Move to next value in collection
func (iter *SliceIterator[T]) Next() bool {
	if iter.index < len(iter.Elements) {
		iter.value = iter.Elements[iter.index]
		iter.index += 1
		return true
	}

	return false
}

// Get current element
func (iter *SliceIterator[T]) Value() T {
	return iter.value
}