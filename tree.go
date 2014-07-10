package main

/*type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}*/


import (
    "code.google.com/p/go-tour/tree"
    "fmt"
    "reflect"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    l := reflect.ValueOf(t.Left)
    r := reflect.ValueOf(t.Right)
    if l.Pointer() != 0 {
    	Walk(t.Left, ch)
    }
    ch <- t.Value;
    if r.Pointer() != 0 {
    	Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    var l1, l2 int
    var v1, v2 [10]int
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    done := func() bool {
    	return  l1==10 && l2==10
    }
    check := func() bool {
        for i, v := range v1 {
    		fmt.Printf("%d - %d %d \n", i, v, v2[i])
            if v != v2[i] {
            	return false
            }
        }
        return true
    }

    for {
        select {
        case v := <- ch1:
            v1[l1] = v
            l1++
            if done() {
            	return check()
            }
        case v := <- ch2:
            v2[l2] = v
            l2++
            if done() {
            	return check()
            }
        }
    }

}

func main() {
    v := Same(tree.New(1), tree.New(2))
    fmt.Printf("%t", v)
}
