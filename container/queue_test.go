package container

import (
    "testing"
    "fmt"
)

func TestQueue(t *testing.T) {
	var A = []int{ 0, 1, 2, 3, 4}
	var q Queue
	
	for i:=0; i<len(A); i++{
		q.Push(A[i])
	}
	
	N:= q.Size()
	
	fmt.Printf("Queue Pop\n");
	for i:=0; i<N; i++{
		item := q.Pop()
		fmt.Printf("%d ", item.Value.(int))
	}
}

