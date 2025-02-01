package linkedlist

import (
	"math/rand"
	"testing"
)

func TestPushFront(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushFront(1)
	emptyList.PushFront(2)

	if emptyList.Head.data != 2 {
		t.Errorf("el primer item debe ser 1")
	}

	if emptyList.Tail.data != 1 {
		t.Errorf("el ultimo item debe ser 2")
	}

}

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

func TestTopFront(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushFront(1)
	emptyList.PushFront(2)

	got, err := emptyList.TopFront()

	if err != nil {
		t.Errorf("error: %q", err)
	}

	if got != 2 {
		t.Errorf("el resultado deberia ser 2")
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

func TestPopFront(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushFront(1)
	emptyList.PushFront(2)

	err := emptyList.PopFront()

	if err != nil {
		t.Errorf("error: %q", err)
	}

	if emptyList.Head.data != 1 {
		t.Errorf("error: header debe ser 1")
	}
	if emptyList.Tail.data != 1 {
		t.Errorf("el Tail deberia ser 1")
	}

	//test con 1 unico item
	err = emptyList.PopFront()

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if emptyList.Head != nil {
		t.Errorf("error: header debe ser 1")
	}
	if emptyList.Tail != nil {
		t.Errorf("el Tail deberia ser 1")
	}

	//test con empty list
	err = emptyList.PopFront()
	if err == nil {
		t.Errorf("list empty debe dar error")
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

func TestPushBack(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushBack(1)

	if emptyList.Head.data != 1 || emptyList.Head.next != nil {
		t.Errorf("el primer item debe ser 1")
	}
	if emptyList.Tail.data != 1 || emptyList.Tail.next != nil {
		t.Errorf("el ultimo item debe ser 1")
	}

	//Test con segundo item
	emptyList.PushBack(2)

	if emptyList.Head.data != 1 || emptyList.Head.next.data != 2 {
		t.Errorf("el primer item debe ser 1")
	}
	if emptyList.Tail.data != 2 || emptyList.Tail.next != nil {
		t.Errorf("el ultimo item debe ser 2")
	}

	emptyList.PushBack(3)

	if emptyList.Head.data != 1 || emptyList.Head.next.data != 2 {
		t.Errorf("el primer item debe ser 1")
	}
	if emptyList.Tail.data != 3 || emptyList.Tail.next != nil {
		t.Errorf("el ultimo item debe ser 2")
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

func TestTopBack(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}

	_, err := emptyList.TopBack()
	if err == nil {
		t.Errorf("list es empty debe dar error")
	}

	emptyList.PushBack(1)
	emptyList.PushBack(2)

	got, err := emptyList.TopBack()

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if got != 2 {
		t.Errorf("debe dar 2")
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

func TestPopBack(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}

	err := emptyList.PopBack()
	if err == nil {
		t.Errorf("list is empty should give error")
	}

	// test unico item
	emptyList.PushFront(1)
	err = emptyList.PopBack()

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if emptyList.Head != nil {
		t.Errorf("el head debe ser nil")
	}
	if emptyList.Tail != nil {
		t.Errorf("el tail debe ser nil")
	}

	// test multi item
	emptyList.PushFront(1)
	emptyList.PushFront(2)

	err = emptyList.PopBack()
	if err != nil {
		t.Errorf("error: %q", err)
	}
	if emptyList.Head.data != 2 {
		t.Errorf("el head debe ser 2")
	}
	if emptyList.Tail.data != 2 {
		t.Errorf("debe ser 2")
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

func TestFind(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	//test nil find
	_, err := emptyList.Find(4)

	if err == nil {
		t.Errorf("empty should return err")
	}

	//find hit
	emptyList.PushBack(1)
	emptyList.PushBack(2)
	emptyList.PushBack(3)
	emptyList.PushBack(4)
	emptyList.PushBack(5)

	nodo, err := emptyList.Find(3)

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if nodo.data != 3 {
		t.Errorf("debe ser 3")
	}

	// find miss
	nodo, err = emptyList.Find(7)
	if err == nil {
		t.Errorf("debe retornar que no encontro item")
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

func TestErase(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	//test nil find
	err := emptyList.Erase(4)

	if err == nil {
		t.Errorf("empty should return err")
	}

	//erase hit
	emptyList.PushBack(1)
	emptyList.PushBack(2)
	emptyList.PushBack(3)
	emptyList.PushBack(4)
	emptyList.PushBack(5)

	//erase first
	err = emptyList.Erase(1)
	if err != nil {
		t.Errorf("err: %q", err)
	}
	if emptyList.Head.data != 2 {
		t.Errorf("el head debe ser 2")
	}

	err = emptyList.Erase(3)

	if err != nil {
		t.Errorf("err: %q", err)
	}
	if _, err := emptyList.Find(3); err == nil {
		t.Errorf("should return not found")
	}

	//erase last
	err = emptyList.Erase(5)
	if err != nil {
		t.Errorf("err: %q", err)
	}
	if emptyList.Tail.data != 4 {
		t.Errorf("el head debe ser 4")
	}

	// erase miss
	err = emptyList.Erase(7)
	if err == nil {
		t.Errorf("debe retornar que no encontro item")
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

func TestEmpty(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushBack(1)
	emptyList.PushBack(2)
	emptyList.PushBack(3)
	emptyList.PushBack(4)
	emptyList.PushBack(5)

	err := emptyList.Empty()
	if err != nil {
		t.Errorf("error: %q", err)
	}
	if emptyList.Head != nil {
		t.Errorf("head debe ser nil")
	}
	if emptyList.Tail != nil {
		t.Errorf("tail debe ser nil")
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

func TestAddBefore(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushBack(1)
	emptyList.PushBack(2)
	emptyList.PushBack(3)
	emptyList.PushBack(4)
	emptyList.PushBack(5)

	nodo, _ := emptyList.Find(3)

	addNode, err := emptyList.AddBefore(nodo, 123)

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if addNode.data != 123 {
		t.Errorf("el nodo agregado debe ser 123")
	}
	if addNode.next.data != 4 {
		t.Errorf("el nodo siguiente debe ser 4")
	}
	if _, err := emptyList.Find(123); err != nil {
		t.Errorf("error: %q", err)
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

func TestAddAfter(t *testing.T) {
	var emptyList = LinkedList{Head: nil, Tail: nil}
	emptyList.PushBack(1)
	emptyList.PushBack(2)
	emptyList.PushBack(3)
	emptyList.PushBack(4)
	emptyList.PushBack(5)

	nodo, _ := emptyList.Find(3)

	addNode, err := emptyList.AddAfter(nodo, 123)

	if err != nil {
		t.Errorf("error: %q", err)
	}
	if _, err := emptyList.Find(123); err != nil {
		t.Errorf("error: %q", err)
	}
	if addNode.next != nodo {
		t.Errorf("next item must be 3")
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
