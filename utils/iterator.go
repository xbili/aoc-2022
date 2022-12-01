package utils

// An iterator helps with reading input data from any underlying source for the
// challenges.
type Iterator[T any] interface {
	Next() (value T, ok bool)
}
