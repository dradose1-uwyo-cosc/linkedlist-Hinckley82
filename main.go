package main

import (
	"fmt"
	"hw-linked/ds" // replace with your actual module path
)

func main() {
	fmt.Println("=== LINKED LIST TESTS ===")
	testLinkedList()

	fmt.Println("\n=== STACK TESTS ===")
	testStack()

	fmt.Println("\n=== QUEUE TESTS ===")
	testQueue()
}

func testLinkedList() {
	var list ds.LinkedList

	fmt.Println("IsEmpty (should be true):", list.IsEmpty())
	fmt.Println("Size (should be 0):", list.GetSize())

	list.Insert("A")
	list.Insert("B")
	list.Insert("C")
	fmt.Print("After inserting A, B, C: ")
	list.PrintList()

	list.Reverse()
	fmt.Print("After Reverse(): ")
	list.PrintList()

	list.Reverse()
	fmt.Print("After Reverse() back to normal: ")
	list.PrintList()

	list.InsertAt(1, "X")
	fmt.Print("After InsertAt(1, X): ")
	list.PrintList()

	list.Remove("B")
	fmt.Print("After Remove(B): ")
	list.PrintList()

	list.Insert("A")
	fmt.Print("Adding another A to the end of the list: ")
	list.PrintList()

	list.RemoveAll("A")
	fmt.Print("After RemoveAll(A): ")
	list.PrintList()

	list.RemoveAt(1)
	fmt.Print("After RemoveAt(1): ")
	list.PrintList()

	fmt.Println("Final Size:", list.GetSize())
}

func testStack() {
	var stack ds.Stack

	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three")

	fmt.Println("Pop (should be Three):")
	val, ok := stack.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop (should be Two):")
	val, ok = stack.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop (should be One):")
	val, ok = stack.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop from empty stack (should fail):")
	val, ok = stack.Pop()
	fmt.Println(val, ok)
}

func testQueue() {
	var queue ds.Queue

	queue.Push("First")
	queue.Push("Second")
	queue.Push("Third")

	fmt.Println("Pop (should be First):")
	val, ok := queue.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop (should be Second):")
	val, ok = queue.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop (should be Third):")
	val, ok = queue.Pop()
	fmt.Println(val, ok)

	fmt.Println("Pop from empty queue (should fail):")
	val, ok = queue.Pop()
	fmt.Println(val, ok)
}
