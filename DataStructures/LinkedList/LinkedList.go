package linkedlist

import "fmt"

// Define estructura de un nodo donde cada nodo tiene el dato y la ref al siguiente nodo
type Node struct {
	data int
	next *Node
}

// Define la estructura del LinkedList que contiene un Head que apunta al primer nodo
type LinkedList struct {
	Head *Node
}

//Operaciones

// PushFront inserta un item al inicio de la lista
func (list *LinkedList) PushFront(data int) {
	newNode := &Node{data: data, next: list.Head}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	list.Head = newNode
}

// TopFront retorna el primer elemento de la lista
func (list *LinkedList) TopFront() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}

	return list.Head.data, nil
}

// PopFront elmina el primer item de la lista
func (list *LinkedList) PopFront() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	list.Head = list.Head.next
	return nil
}

//Pushback inserta un nuevo nodo al final de la lista

func (list *LinkedList) PushBack(data int) {
	newNode := &Node{data: data, next: list.Head}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	nextNode := list.Head.next
	for nextNode.next != nil {
		nextNode = nextNode.next
	}
	nextNode.next = newNode
}

// retorna el ultimo item de la lista
func (list *LinkedList) TopBack() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}
	currentNode := list.Head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}

// Elimina el ultimo item de la lista
func (list *LinkedList) PopBack() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	currentNode := list.Head
	nextNode := currentNode.next

	if nextNode == nil {
		list.Head = nil
		return nil
	}

	for nextNode.next != nil {
		currentNode = nextNode
		nextNode = nextNode.next
	}

	currentNode.next = nil

	return nil
}
