package linkedlist

import "fmt"

//Una doubleLinkedList tiene en cada nodo una referencia al nodo siguiente y el nodo anteriro

type DNode struct {
	prev *DNode
	data int
	next *DNode
}

type DoubleLinkedList struct {
	Head *DNode
	Tail *DNode
}

//Operaciones

// PushFront inserta un item al inicio de la lista
// Esta operacion tiene una complejidad const O(1) ya que se puede acceder al primer item por el HEAD
func (list *DoubleLinkedList) PushFront(data int) {
	newNode := &DNode{prev: nil, data: data, next: list.Head}
	list.Head = newNode
}

// TopFront retorna el primer elemento de la lista
// Complejidad const O(1) idem PushFront
func (list *DoubleLinkedList) TopFront() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("doubleLinked list is empty")
	}

	return list.Head.data, nil
}

// PopFront elmina el primer item de la lista
// Complejidad const O(1) idem PushFront
func (list *DoubleLinkedList) PopFront() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	list.Head = list.Head.next
	list.Head.prev = nil
	return nil
}

// Pushback inserta un nuevo nodo al final de la lista
// En el caso de linked list sin tails se debe reccorrer toda la lista
// Por lo que la complejidad es lineal O(n)
func (list *DoubleLinkedList) PushBack(data int) {
	newNode := &DNode{prev: list.Tail, data: data, next: nil}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	list.Tail.next = newNode
	list.Tail = newNode

}

// retorna el ultimo item de la lista
// En el caso de linked list sin tails se debe reccorrer toda la lista
// Por lo que la complejidad es lineal O(n)
func (list *DoubleLinkedList) TopBack() (data int, err error) {
	if list.Head == nil {
		return 0, fmt.Errorf("linked list is empty")
	}

	return list.Tail.data, nil
}

// Elimina el ultimo item de la lista
// En un double linked list es una oper constante del tipo O(1) ya que podemos acceder
// al nodo anterior al ultimo
func (list *DoubleLinkedList) PopBack() error {
	if list.Head == nil {
		return fmt.Errorf("linked list is empty")
	}

	//Si el linked list tiene un solo item lo elimino
	if list.Head == list.Tail {
		list.Head = nil
		list.Tail = nil
		return nil
	}

	// modifico el next del anteultimo item para eliminar el ultimo
	prevToLastNode := list.Tail.prev
	prevToLastNode.next = nil

	//Actualizo el Tail
	list.Tail = prevToLastNode

	return nil
}

// Find encuentra el y devuelve el item que tiene determinado key
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *DoubleLinkedList) Find(key int) (DNode, error) {
	currentNode := list.Head

	for currentNode.data != key {
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return DNode{}, fmt.Errorf("key not found")
	}

	return *currentNode, nil
}

// la funcion Erase elimina el item que cumple con la key enviada.
// se debe recorrer todo el list por lo que presenta ona complejidad O(n)
func (list *DoubleLinkedList) Erase(key int) error {
	currentNode := list.Head

	// en caso de que la key este en el primer item cambio el Head al siguiente item
	if currentNode.data == key {
		list.Head = currentNode.next
		list.Head.prev = nil
		return nil
	}

	// recorro todo el listado busando el key, si no encuentro salgo por error
	for currentNode.data != key {
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return fmt.Errorf("key not found")
	}

	if currentNode.prev == nil { //first item
		list.Head = currentNode.next
		list.Head.prev = nil
	}

	if currentNode.next == nil { //last item
		list.Tail = currentNode.prev
		list.Tail.next = nil
	}
	prevToEraseNode := currentNode.prev
	nexToEraseNode := currentNode.next

	//Salteo el nodo a borrar
	prevToEraseNode.next = nexToEraseNode
	//actualizo el prev del next
	nexToEraseNode.prev = prevToEraseNode
	return nil

}

// Funcion empty elimina el linked list
// Simplemente se borra la referencia del HEAD lo que es una operacion de complejidad O(1)
func (list *DoubleLinkedList) Empty() error {
	list.Head = nil
	list.Tail = nil
	return nil
}

// Funcion AddBefore agrega un nuevo nodo despues del nodo indicado
// Esta operacion es de complejidad O(1) ya que se puede acceder al next del nodo indicado
func (list *DoubleLinkedList) AddBefore(nodo *DNode, key int) (DNode, error) {
	currentNode := nodo
	newNode := DNode{prev: currentNode, data: key, next: currentNode.next}

	currentNode.next.prev = &newNode
	currentNode.next = &newNode
	return newNode, nil
}

// Func AddAfter agrega un nodo antes del nodo indicado
// Complejidad O(1)
func (list *DoubleLinkedList) AddAfter(nodo *DNode, key int) (DNode, error) {
	NodeAfter := nodo
	newNode := DNode{prev: NodeAfter.prev, data: key, next: NodeAfter}

	NodeAfter.prev = &newNode
	NodeAfter.prev.next = &newNode
	return newNode, nil
}
