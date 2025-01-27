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
// Esta operacion tiene una complejidad const O(1) ya que se puede acceder al primer item por el HEAD
func (list *LinkedList) PushFront(data int) {
	newNode := &Node{data: data, next: list.Head}
	if list.Head == nil {
		list.Head = newNode
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

	list.Head = list.Head.next
	return nil
}

//Pushback inserta un nuevo nodo al final de la lista
//En el caso de linked list sin tails se debe reccorrer toda la lista
//Por lo que la complejidad es lineal O(n)

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
// En el caso de linked list sin tails se debe reccorrer toda la lista
// Por lo que la complejidad es lineal O(n)
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
// En el caso de linked list sin tails se debe reccorrer toda la lista
// Por lo que la complejidad es lineal O(n)
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

// Find encuentra el y devuelve el item que tiene determinado key
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *LinkedList) Find(key int) (Node, error) {
	currentNode := list.Head

	for currentNode.data != key {
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return Node{}, fmt.Errorf("key not found")
	}

	return *currentNode, nil
}

// la funcion Erase elimina el item que cumple con la key enviada.
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *LinkedList) Erase(key int) error {
	currentNode := list.Head

	// en caso de que la key este en el primer item cambio el Head al siguiente item
	if currentNode.data == key {
		list.Head = currentNode.next
		return nil
	}

	// recorro todo el listado busando el key, si no encuentro salgo por error
	for currentNode.next.data != key {
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return fmt.Errorf("key not found")
	}

	currentNode.next = currentNode.next.next
	return nil

}

// Funcion empty elimina el linked list
// Simplemente se borra la referencia del HEAD lo que es una operacion de complejidad O(1)
func (list *LinkedList) Empty() error {
	list.Head = nil
	return nil
}

// Funcion AddBefore agrega un nuevo nodo despues del nodo indicado
// Esta operacion es de complejidad O(1) ya que se puede acceder al next del nodo indicado
func (list *LinkedList) AddBefore(nodo *Node, key int) (Node, error) {
	currentNode := nodo
	newNode := Node{data: key, next: currentNode.next}

	currentNode.next = &newNode
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
