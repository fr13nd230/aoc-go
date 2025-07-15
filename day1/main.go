package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main is the go entry program here
// you will find only the setups.
func main() {    
    fset, err := os.OpenFile("set.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        panic("Set file is not found.")
    }
    defer fset.Close()

    r := bufio.NewReader(fset)
    lists := NewLists()
    
    for range LIST_SIZE {
        line, _, err := r.ReadLine()
        if err != nil {
            panic(err)
        }
        parts := strings.Split(string(line), "  ")
        if len(parts) >= 2 {
            left, _ := strconv.ParseInt(parts[0], 10, 64)
            right, _ := strconv.ParseInt(parts[1], 10, 64)
            lists.AppendLeft(left)
            lists.AppendRight(right)
        }
    }
    
    fmt.Printf("Final result is: %d", solve(lists, LIST_SIZE))
    return
}