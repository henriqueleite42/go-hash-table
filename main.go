package main

import "log"

//
//
// Types
//
//

// Size of the HashTable
const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

// Linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// Linked list node that holds the key
type bucketNode struct {
	key string
	next *bucketNode
}

//
//
// Hash Function
//
//

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//
//
// Hash Table Functions
//
//

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//
//
// Bucket Functions
//
//

func (b *bucket) insert(key string) {
	if b.search(key) {
		return
	}

	newNode := &bucketNode{
		key: key,
	}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

//
//
// Init Function
//
//

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

//
//
// Main
//
//

func main() {
	log.Println("hash RAZAL ", hash("RAZAL"))
	hashTable := Init()
	log.Println("hashTable ", hashTable)

	list := []string{
		"RAZAL",
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
		"RANDY",
		"RAZAL",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	log.Println("RAZAL exists ", hashTable.Search("RAZAL"))
	hashTable.Delete("RAZAL")
	log.Println("RAZAL exists (deleted) ", hashTable.Search("RAZAL"))
}
