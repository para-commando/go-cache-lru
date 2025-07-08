package main

import (
	"fmt"
)

const SIZE = 5 // maximum size of the cache

type Node struct {
	Left  *Node
	Right *Node
	Value string
}
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	// initializing an empty queue with only 2 nodes, head and tail
	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func (cache *Cache) Check(word string) {
	node := &Node{}
	if val, ok := cache.Hash[word]; ok {
		// if the word is already in the cache, move it to the front

		node = cache.Remove(val)
	} else {
		// if the word is not in the cache, create a new node
		node = &Node{Value: word}
	}

	cache.Add(node)
	cache.Hash[word] = node
}

func (cache *Cache) Remove(node *Node) *Node {
	if node.Left != nil {
		node.Left.Right = node.Right
	}
	if node.Right != nil {
		node.Right.Left = node.Left
	}
	cache.Queue.Length--
	delete(cache.Hash, node.Value)
	return node
}

func (cache *Cache) Add(node *Node) {
	fmt.Printf("Adding %s to cache\n", node.Value)
	tmp := cache.Queue.Head.Right         // store the current first node after head
	cache.Queue.Head.Right = node   // insert the new node to right after head
	node.Left = cache.Queue.Head     // set the new node's left pointer to head as head is always at the first
	node.Right = tmp       // set the new node's right pointer to the previous first node
	tmp.Left = node          // update the previous first node's left pointer to the new node
	cache.Queue.Length++
	cache.Hash[node.Value] = node
	if cache.Queue.Length > SIZE {
		// if the cache is full, remove the least recently used item
		oldest := cache.Queue.Tail.Left
		cache.Remove(oldest)
		delete(cache.Hash, oldest.Value)
	}
}

func (cache *Cache) Display() {
	cache.Queue.Display()
	fmt.Printf("\nCache size: %d\n", cache.Queue.Length)
}

func (queue *Queue) Display() {
	current := queue.Head.Right
	fmt.Print("Cache: ")
	// bcs we had initialize head.right = tail while creating an empty queue, hence using it as a check

	for i := range queue.Length {

		fmt.Printf("{%s}", current.Value)
		if i < queue.Length-1 {
			fmt.Print("<-->")
		}

		current = current.Right
	}

}
func main() {
	cache := NewCache()
	for _, word := range []string{"Radio-Check", "marine", "commando", "Do you copy ?", "Over"} {
		cache.Check(word)
		cache.Display()

	}
}
