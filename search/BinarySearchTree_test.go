package search

import (
	"algorithms"
    "testing"
    "fmt"
)

func TestBinarySearchTree(t *testing.T) {
	var b =[]int{'S', 'E', 'A', 'R', 'C', 'H', 'E', 'X', 'A', 'M', 'P', 'L', 'E'}
	var st	SymbolTable
		
	a := algorithms.NewIntegerSlice(b);
	
	st = &BinarySearchTree{}
	
	for i:=0; i<len(a); i++ {
		st.Set(a[i], i);
	}
	
	iter := st.Keys();
	for iter.HasNext() {
		key := iter.Next().Value.(algorithms.Integer)
		fmt.Printf("%c %d\n", key, st.Get(key));
	}
}


func TestBinarySearchTree2(t *testing.T) {
	var b =[]string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst", "of", "times", 
					"it", "was", "the", "age", "of", "wisdom", "it", "was", "the", "age", "of", "foolishess",
					"it", "was", "the", "epoch", "of", "belief", "it", "was", "the", "epoch", "of", "incredulity",
					"it", "was", "the", "season", "of", "light", "it", "was", "the", "season", "of", "darkness",
					"it", "was", "the", "spring", "of", "hope", "it", "was", "the", "winter", "of", "despair"}
					
	var st	SymbolTable
	
	a := algorithms.NewStringSlice(b);
	
	st = &BinarySearchTree{}
	
	minlen := 1;
	
	for i:=0; i<len(a); i++ {
		word := a[i]
		if len(word.(algorithms.String)) < minlen {
			continue
		}
		if !st.Contains(word) {
			st.Set(word, int(1));
		}else{
			st.Set(word, st.Get(word).(int)+int(1));
		}
	}
	
	max := algorithms.String(" ")
	st.Set(max, 0);
	iter := st.Keys()
	for iter.HasNext() {
		word := iter.Next().Value.(algorithms.String)
		if st.Get(word).(int) > st.Get(max).(int) {
			max = word;
		}
	}
	
	fmt.Printf("%s %d\n", max, st.Get(max));
}
