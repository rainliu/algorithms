package container

import (
    "testing"
    "fmt"
)

func TestStack(t *testing.T) {
	var A = []int{ 0, 1, 2, 3, 4}
	var s Stack
	
	for i:=0; i<len(A); i++{
		s.Push(A[i])
	}
	
	N:= s.Size()
	
	fmt.Printf("Stack Pop\n");
	for i:=0; i<N; i++{
		item := s.Pop()
		fmt.Printf("%d ", item.Value.(int))
	}
}
