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

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(i)
		}
		b.StartTimer()

		list.Find(999998)
	}

}

func BenchmarkErase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(i)
		}
		b.StartTimer()

		list.Erase(999990)
	}

}

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		b.StartTimer()

		list.Empty()
	}
}

func BenchmarkAddBefore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		nodo := list.Head
		for i := 0; i < 3000; i++ {
			nodo = nodo.next
		}
		b.StartTimer()

		list.AddBefore(nodo, 1)
	}
}

func BenchmarkAddAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := LinkedList{Head: nil}
		for i := 0; i < 999999; i++ {
			list.PushFront(rand.Int())
		}
		nodo := list.Head
		for i := 0; i < 3000; i++ {
			nodo = nodo.next
		}
		b.StartTimer()

		list.AddAfter(nodo, 1)

	}
}
