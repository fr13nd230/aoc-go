package main

import (
	"math"
	"sort"
)

// solve will take a list pointer and try to
// sort both right and left side (samllest to smallest)
// than measure the total distance between them.
func solve(l *Lists, size int) int {
    if size == 0 {
        size = LIST_SIZE
    }
    sort.Slice(l.Right, func(i, j int) bool {
        return l.Right[i] < l.Right[j]
    })
    sort.Slice(l.Left, func(i, j int) bool {
        return l.Left[i] < l.Left[j]
    })
    
    var result int = 0
    for i := range size {
        result += int(math.Abs(float64(l.Right[i] - l.Left[i])))
    }
    
    return result
} 
