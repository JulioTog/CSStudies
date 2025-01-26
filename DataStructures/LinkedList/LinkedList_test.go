package linkedlist

import (
	"math/rand"
	"testing"
)

func BenchmarkPushFront(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}

		b.StartTimer()
		list.PushFront(rand.Int())
	}
}

func BenchmarkTopFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}

		b.StartTimer()
		setvar, err := list.TopFront()
		b.StopTimer()

		//just to eliminate error when var is not used
		if err != nil {
			setvar += 1
		}
		b.StartTimer()
	}

}

func BenchmarkPopFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}

		b.StartTimer()
		list.PopFront()

	}
}

func BenchmarkPushBack(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		b.StartTimer()

		list.PushBack(rand.Int())
	}
}

func BenchmarkTopBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		b.StartTimer()

		list.TopBack()
	}

}

func BenchmarkPopBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		b.StartTimer()

		list.PopBack()
	}

}
