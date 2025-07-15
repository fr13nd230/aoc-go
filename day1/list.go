package main

// total number of lines.
const LIST_SIZE = 1000

// Lists is a struct contains two fields
// Right and Left are slices of int64.
type Lists struct {
    Right   []int64
    Left    []int64
} 

// NewLists will generate new empty slices for each
// left and right parts and returns a pointer *Lists.
func NewLists() *Lists {
    return &Lists{
        Right: []int64{},
        Left: []int64{},
    }
}

// AppendRight reciever function of *Lists 
// that will append value of int64 to the Right end of a list.
func (l *Lists) AppendRight(value int64) {
    l.Right = append(l.Right, value)
    return
}

// AppendLeft reciever function of *Lists 
// that will append value of int64 to the Left end of a list.
func (l *Lists) AppendLeft(value int64) {
    l.Left = append(l.Left, value)
    return
}