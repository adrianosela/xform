package xform

import "iter"

// Seq is a transformer over sequences of individual values.
type Seq[V1 any, V2 any] func(V1) V2

// XSeq applies a transform function to an iter.Seq.
func XSeq[V1 any, V2 any](
	iter iter.Seq[V1],
	fn Seq[V1, V2],
) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for v := range iter {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
