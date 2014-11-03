package sort

import (
	"algorithms"
    "testing"
    "fmt"
)

func TestQuickSort(t *testing.T) {
	var b =[]int{'S', 'O', 'R', 'T', 'e', 'X', 'A', 'M', 'P', 'L', 'E'}
	var s QuickSort
	
	a := algorithms.NewIntegerSlice(b);
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%c ", a[i]);
	}
	fmt.Printf("\n")
	
	s.Sort(a)
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%c ", a[i]);
	}
	fmt.Printf("\n")	
	fmt.Printf("a[] is sored? %v\n", s.IsSorted(a))
}
