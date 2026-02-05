package ds

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string) { // insert at the Tail
	newNode := &Node{data: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *LinkedList) InsertAt(position int, value string) error { //inserts at a position, returns true if success or false if not, like if pos doesn't exist
	if position < 0 || position > l.Size {
		return errors.New("position out of range")
	}
	if position == 0 {
		newNode := &Node{data: value, Next: l.Head}
		l.Head = newNode
	} else {
		curr := l.Head
		for i := 0; i < position-1; i++ {
			curr = curr.Next
		}
		newNode := &Node{data: value, Next: curr.Next}
		curr.Next = newNode
	}
	l.Size++
	return nil
}

func (l *LinkedList) Remove(value string) error { // remove first occurrence of the value
	if l.Head == nil {
		return errors.New("list is empty")
	}
	if l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		return nil
	}
	prev := l.Head
	curr := l.Head.Next
	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next
			l.Size--
			return nil
		}
		prev = curr
		curr = curr.Next
	}
	return errors.New("value not found")
}

func (l *LinkedList) RemoveAll(value string) error { //remove all occurrences of a value
	removed := false
	for l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		removed = true
	}

	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.Next.data == value {
			curr.Next = curr.Next.Next
			l.Size--
			removed = true
		} else {
			curr = curr.Next
		}
	}
	if !removed {
		return errors.New("value not found")
	}
	return nil
}
func (l *LinkedList) RemoveAt(pos int) error { // remove at a position, if index exists
	if pos < 0 || pos >= l.Size {
		return errors.New("position out of range")
	}
	if pos == 0 {
		l.Head = l.Head.Next
	} else {
		curr := l.Head
		for i := 0; i < pos-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
	}
	l.Size--
	return nil
}

func (l *LinkedList) IsEmpty() bool { // checks if the linked list is empty
	return l.Head == nil
}

func (l *LinkedList) GetSize() int { // get Size of ll
	return l.Size
}

func (l *LinkedList) Reverse() { //reverse the list
	var prev *Node
	curr := l.Head

	for curr != nil {
		Next := curr.Next
		curr.Next = prev
		prev = curr
		curr = Next
	}
	l.Head = prev
}

func (l *LinkedList) PrintList() { //print the list
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.data)
		if curr.Next != nil {
			fmt.Print("->")
		}
		curr = curr.Next
	}
	fmt.Println()
}
