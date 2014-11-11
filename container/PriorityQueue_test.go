package container

import (
	"algorithms"
    "testing"
    "fmt"
)

func TestPriorityQueue(t *testing.T) {	
	// insert a bunch of strings
    var strs = []string{ "it", "was", "the", "best", "of", "times", "it", "was", "the", "worst" };

    pq := NewPriorityQueue(len(strs));
    for i := 0; i < len(strs); i++ {
        pq.Push(&PriorityQueueItem{i, algorithms.String(strs[i])});
    }

    // print each key using the iterator
    iter := pq.Iterator()    
    for iter.HasNext() {
    	i := iter.Next().Value.(*PriorityQueueItem).Index
        fmt.Printf("%d %s ", i, strs[i]);
    }
    fmt.Printf("\n");

    // increase or decrease the key
    for i := 0; i < len(strs); i++ {
        if i%2==0{
            pq.IncreaseKey(i, algorithms.String(strs[i] + strs[i]));
        }else{
            pq.DecreaseKey(i, algorithms.String(strs[i][0:1]));
        }
    }

    // delete and print each key
    for !pq.IsEmpty() {
    	pqi := pq.Pop().Value.(*PriorityQueueItem)
        key := pqi.Key;
        i   := pqi.Index;
        fmt.Printf("%d %s ", i, key);
    }
    fmt.Printf("\n");

/*
    // reinsert the same strings
    for (int i = 0; i < strings.length; i++) {
        pq.insert(i, strings[i]);
    }

    // delete them in random order
    int[] perm = new int[strings.length];
    for (int i = 0; i < strings.length; i++)
        perm[i] = i;
    StdRandom.shuffle(perm);
    for (int i = 0; i < perm.length; i++) {
        String key = pq.keyOf(perm[i]);
        pq.delete(perm[i]);
        StdOut.println(perm[i] + " " + key);
    }
 */
}

