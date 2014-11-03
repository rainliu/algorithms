package sort

import (
	"algorithms"
    "testing"
    "fmt"
)

func TestInsertionSort(t *testing.T) {
	var b =[]float64{'S', 'O', 'R', 'T', 'e', 'X', 'A', 'M', 'P', 'L', 'E'}
	var s InsertionSort
	
	a := algorithms.NewDoubleSlice(b);
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%f ", a[i]);
	}
	fmt.Printf("\n")
	
	s.Sort(a)
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%f ", a[i]);
	}
	fmt.Printf("\n")
	fmt.Printf("a[] is sored? %v\n", s.IsSorted(a))
}

