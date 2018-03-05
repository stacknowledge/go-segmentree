# go-segmentree

Implementation of Segment tree with lazy propagation in Golang

## Basic Usage

```
package main

import "github.com/stacknowledge/go-segmentree"

func main(){
    tree := segment.NewTree([]int{1,3,5,7,9,11})

    q1 := tree.Query(0, 5)
    q2 := tree.Query(0, 2)
    q3 := tree.Query(3, 5)

    fmt.Println(q1, q2, q3)

    tree.Update(0, 2, 13)

    q1 := tree.Query(0, 5)
    q2 := tree.Query(0, 2)
    q3 := tree.Query(3, 5)

    fmt.Println(q1, q2, q3)
}
```