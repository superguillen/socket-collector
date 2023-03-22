package slicetools

// Based of: https://itnext.io/generic-map-filter-and-reduce-in-go-3845781a591c

// ==========================================
// Iterator implementation
// ==========================================

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

// ==========================================
// Map implementation
// ==========================================


type mapIterator[T any] struct {
	source Iterator[T]
	mapper func(T) T
}

// advance to next element
func (iter *mapIterator[T]) Next() bool {
	return iter.source.Next()
}

func (iter *mapIterator[T]) Value() T {
	value := iter.source.Value()
	return iter.mapper(value)
}

func Map[T any](iter Iterator[T], f func(T) T) Iterator[T] {
	return &mapIterator[T]{
		iter, f,
	}
}

// ==========================================
// Filter implementation
// ==========================================

type filterIterator[T any] struct {
	source Iterator[T]
	pred   func(T) bool
}

func (iter *filterIterator[T]) Next() bool {
	for iter.source.Next() {
		if iter.pred(iter.source.Value()) {
			return true
		}
	}
	return false
}

func (iter *filterIterator[T]) Value() T {
	return iter.source.Value()
}

func Filter[T any](iter Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIterator[T]{
		iter, pred,
	}
}

// ==========================================
// Collect implementation
// ==========================================

func Collect[T any](iter Iterator[T]) []T {
	var xs []T

	for iter.Next() {
		xs = append(xs, iter.Value())
	}

	return xs
}

// ==========================================
// Reducer implementation
// ==========================================

type Reducer[T, V any] func(accum T, value V) T

// Reduce values iterated over to a single value
func Reduce[T, V any](iter Iterator[V], f Reducer[T, V]) T {
	var accum T
	for iter.Next() {
		accum = f(accum, iter.Value())
	}
	return accum
}


//Example:
// //Get socket list
// sockstats_list,listen_ports := sockstats.GetSockStats()

// // Create iterator over a slice of integers
// iter := slicetools.NewSliceIterator(sockstats_list)

// listen_acum := map[uint16]map[string]sockstats.SockStatAcumm{}

// // Get Listen Ports
// ports := slicetools.Filter(iter, func(item sockstats.SockStat) bool {
// 	return item.Status == "LISTEN"
// })

// listen_ports := slicetools.Collect(ports)