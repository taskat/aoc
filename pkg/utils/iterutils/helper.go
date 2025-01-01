package iterutils

import "iter"

type iterator[K comparable, V any] interface {
	iter.Seq[V] | iter.Seq2[K, V]
}

const (
	emptyIterPanicMsg   = "using empty iterator"
	indexOOBPanicMsg    = "index out of bounds"
	keyNotFoundPanicMsg = "key not found"
	nilIterPanicMsg     = "using nil iterator"
)

func panicIfNil[K comparable, V any, ITER iterator[K, V]](iter ITER) {
	if iter == nil {
		panic(nilIterPanicMsg)
	}
}
