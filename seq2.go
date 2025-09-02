package xform

import "iter"

// Seq2 is a transformer over sequences of pairs of values, most commonly key-value pairs.
type Seq2[K1 any, V1 any, K2 any, V2 any] func(K1, V1) (K2, V2)

// XSeq2 applies a transform function to an iter.Seq2.
func XSeq2[K1 any, V1 any, K2 any, V2 any](
	iter iter.Seq2[K1, V1],
	fn Seq2[K1, V1, K2, V2],
) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range iter {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}
