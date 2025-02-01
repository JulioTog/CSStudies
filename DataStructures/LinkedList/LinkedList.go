package linkedlist

import (
	"fmt"
)

// Define estructura de un nodo donde cada nodo tiene el dato y la ref al siguiente nodo
type Node struct {
	data int
	next *Node
}

// Define la estructura del LinkedList que contiene un Head que apunta al primer nodo
// Agregando ademas el Tail las operaciones de PushBack y TopBack se vuelve de complejidad O(1)
type LinkedList struct {
	Head *Node
	Tail *Node
}

//Operaciones

// PushFront inserta un item al inicio de la lista
// Esta operacion tiene una complejidad const O(1) ya que se puede acceder al primer item por el HEAD
func (list *LinkedList) PushFront(data int) {
	newNode := &Node{data: data, next: list.Head}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Head = newNode
}

// TopFront retorna el primer elemento de la lista
// Complejidad const O(1) idem PushFront
func (list *LinkedList) TopFront() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}

	return list.Head.data, nil
}

// PopFront elmina el primer item de la lista
// Complejidad const O(1) idem PushFront
func (list *LinkedList) PopFront() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	//en caso de que sea un unico nodo
	if list.Head == list.Tail {
		list.Tail = nil
	}
	list.Head = list.Head.next

	return nil
}

//Pushback inserta un nuevo nodo al final de la lista
//En el caso de linked list sin tails se debe reccorrer toda la lista
//Por lo que la complejidad es lineal O(n)

func (list *LinkedList) PushBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Tail.next = newNode
	list.Tail = newNode

}

// retorna el ultimo item de la lista
// En el caso de linked list sin tails se debe reccorrer toda la lista
// Por lo que la complejidad es lineal O(n)
func (list *LinkedList) TopBack() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}

	return list.Tail.data, nil
}

// Elimina el ultimo item de la lista
func (list *LinkedList) PopBack() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	currentNode := list.Head
	nextNode := currentNode.next

	//unico item
	if nextNode == nil {
		list.Head = nil
		list.Tail = nil
		return nil
	}

	for nextNode.next != nil {
		currentNode = nextNode
		nextNode = nextNode.next
	}

	currentNode.next = nil
	list.Tail = currentNode

	return nil
}

// Find encuentra el y devuelve el item que tiene determinado key
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *LinkedList) Find(key int) (*Node, error) {

	//lista vacia
	if list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	currentNode := list.Head

	for currentNode.data != key {

		currentNode = currentNode.next
		if currentNode == nil {
			break
		}
	}

	if currentNode == nil {
		return nil, fmt.Errorf("key not found")
	}

	return currentNode, nil
}

// la funcion Erase elimina el item que cumple con la key enviada.
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *LinkedList) Erase(key int) error {
	currentNode := list.Head

	// Caso list emtpy
	if currentNode == nil {
		return fmt.Errorf("list is empty")
	}
	// en caso de que la key este en el primer item cambio el Head al siguiente item
	if currentNode.data == key {
		list.Head = currentNode.next
		return nil
	}

	// recorro todo el listado busando el key, si no encuentro salgo por error
	for currentNode.next.data != key {
		currentNode = currentNode.next
		if currentNode.next == nil {
			break

		}
	}
	if currentNode.next == nil {
		return fmt.Errorf("key not found")
	}

	// si es ultimo
	if currentNode.next == list.Tail {
		list.Tail = currentNode
		return nil
	}
	currentNode.next = currentNode.next.next
	return nil

}

// Funcion empty elimina el linked list
// Simplemente se borra la referencia del HEAD lo que es una operacion de complejidad O(1)
func (list *LinkedList) Empty() error {
	list.Head = nil
	list.Tail = nil

	return nil
}

// Funcion AddBefore agrega un nuevo nodo despues del nodo indicado
// Esta operacion es de complejidad O(1) ya que se puede acceder al next del nodo indicado
func (list *LinkedList) AddBefore(nodo *Node, key int) (Node, error) {

	newNode := Node{data: key, next: nodo.next}

	nodo.next = &newNode
	return newNode, nil
}

// Func AddAfter agrega un nodo antes del nodo indicado
// En este caso de SimpleLinkedList se debe recorrer todo el list para modificar el item Anteriror
// Complejidad O(n) (Esto mejora en los doubleLinkedList)
func (list *LinkedList) AddAfter(nodo *Node, key int) (Node, error) {
	NodeAfter := nodo
	newNode := Node{data: key, next: NodeAfter}

	head := list.Head
	if head == NodeAfter {
		list.Head = &newNode
	}
	for *head.next != *NodeAfter {
		head = head.next
	}

	if head == nil {
		return Node{}, fmt.Errorf("node not found")
	}
	head.next = &newNode
	return newNode, nil
}
