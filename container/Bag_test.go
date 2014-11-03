package container

import (
    "testing"
    "math"
    "fmt"
)

func TestBag(t *testing.T) {
	var numbers Bag
	var A = []float64{100, 99, 101, 120, 98, 107, 109, 81, 101, 90}
	
	for i:=0; i<len(A); i++{
		numbers.Push(A[i])
	}
	
	N := float64(numbers.Size())
	
	sum := float64(0.0)
	iter := numbers.Iterator();
	for iter.HasNext() {
		item := iter.Next()
		sum += item.Value.(float64)
	}
	mean := sum/N
	
	sum = 0.0;
	iter = numbers.Iterator();
	for iter.HasNext() {
		item := iter.Next()
		sum +=(item.Value.(float64)-mean)*(item.Value.(float64)-mean)
	}
	std := math.Sqrt(sum/(N-1))
	
	fmt.Printf("Mean: %.2f\n", mean);
	fmt.Printf("Std dev: %.2f\n", std);
}

