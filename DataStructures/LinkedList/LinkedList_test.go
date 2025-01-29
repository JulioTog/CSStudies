package linkedlist

import (
	"math/rand"
	"testing"
)

func BenchmarkPushFront(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list
		newList.PushFront(rand.Int())
	}
}

func BenchmarkTopFront(b *testing.B) {

	list := LinkedList{Head: nil}

	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list
		newList.TopFront()
	}

}

func BenchmarkPopFront(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list
		newList.PopFront()

	}
}

func BenchmarkPushBack(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list

		newList.PushBack(rand.Int())
	}
}

func BenchmarkTopBack(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	for i := 0; i < b.N; i++ {
		newList := list

		newList.TopBack()
	}

}

func BenchmarkPopBack(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	for i := 0; i < b.N; i++ {
		newList := list

		newList.PopBack()
	}

}

func BenchmarkFind(b *testing.B) {
	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list

		newList.Find(999998)
	}

}

func BenchmarkErase(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newList := list

		newList.Erase(999990)
	}

}

func BenchmarkEmpty(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newList := list

		newList.Empty()
	}
}

func BenchmarkAddBefore(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}
	nodo := list.Head
	for i := 0; i < 3000; i++ {
		nodo = nodo.next
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		newList := list

		newList.AddBefore(nodo, 1)
	}
}

func BenchmarkAddAfter(b *testing.B) {

	list := LinkedList{Head: nil}
	for i := 0; i < 999999; i++ {
		list.PushFront(rand.Int())
	}
	nodo := list.Head
	for i := 0; i < 3000; i++ {
		nodo = nodo.next
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		newList := list
		newList.AddAfter(nodo, 1)

	}
}
