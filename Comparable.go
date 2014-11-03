package algorithms

import (
	"bytes"
)

type Comparable interface {
	CompareTo(Comparable) int
}

type Integer int

func NewIntegerSlice(a []int) []Comparable{
	b := make([]Comparable, len(a))
	for i:=0; i<len(a); i++ {
		b[i] = Integer(a[i])
	}
	return b
}

func (this Integer) CompareTo(c Comparable) int {
	if that, ok := c.(Integer); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	} else {
		panic("Input type is not Integer type\n")

		return 0
	}
}

type Double float64

func NewDoubleSlice(a []float64) []Comparable{
	b := make([]Comparable, len(a))
	for i:=0; i<len(a); i++ {
		b[i] = Double(a[i])
	}
	return b
}

func (this Double) CompareTo(c Comparable) int {
	if that, ok := c.(Double); ok {
		if this < that {
			return -1
		} else if this == that {
			return 0
		} else {
			return 1
		}
	} else {
		panic("Input type is not Double type\n")
		return 0
	}
}

type String string

func NewStringSlice(a []string) []Comparable{
	b := make([]Comparable, len(a))
	for i:=0; i<len(a); i++ {
		b[i] = String(a[i])
	}
	return b
}

func (this String) CompareTo(c Comparable) int {
	if that, ok := c.(String); ok {
		thisBytes := []byte(this)
		thatBytes := []byte(that)

		return bytes.Compare(thisBytes, thatBytes)
	} else {
		panic("Input type is not String type\n")
		return 0
	}
}
