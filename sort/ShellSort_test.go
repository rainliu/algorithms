package sort

import (
	"algorithms"
    "testing"
    "fmt"
)

func TestShellSort(t *testing.T) {
	var b =[]string{"Sorter", "Orange", "Red", "Tank", "eneny", "X-man", "Angela", "Moon", "Peter", "Linkedin", "Extern"}
	var s ShellSort
	
	a := algorithms.NewStringSlice(b);
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%s ", a[i]);
	}
	fmt.Printf("\n")
	
	s.Sort(a)
	
	for i:=0; i<len(a); i++{
		fmt.Printf("%s ", a[i]);
	}
	fmt.Printf("\n")
	fmt.Printf("a[] is sored? %v\n", s.IsSorted(a))
}

