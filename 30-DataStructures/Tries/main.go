package main

import "fmt"

type Node struct {
	isTerminal bool
	character  byte
	arrChars   [26]*Node
}

type Trie struct {
	root *Node
}

func (trie *Trie) Insert(word string) {
	curr := trie.root

	for i := 0; i < len(word); i++ {
		char := word[i]

		if curr.arrChars[char-'a'] == nil {
			newNode := &Node{isTerminal: false, character: char}
			curr.arrChars[char-'a'] = newNode
		}

		curr = curr.arrChars[char-'a']
	}

	curr.isTerminal = true
}

func (trie *Trie) Search(word string) bool {
	curr := trie.root

	for i := 0; i < len(word); i++ {
		char := word[i]

		if curr.arrChars[char-'a'] == nil {
			return false
		}

		curr = curr.arrChars[char-'a']
	}

	return curr.isTerminal
}

func main() {

	trie := &Trie{root: &Node{isTerminal: false, character: 'a'}}

	trie.Insert("apple")
	trie.Insert("apply")
	trie.Insert("pen")
	trie.Insert("pencil")

	fmt.Println(trie.Search("nikhil"))
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("apply"))
	fmt.Println(trie.Search("appl"))
	fmt.Println(trie.Search("pen"))
	fmt.Println(trie.Search("pencil"))
	fmt.Println(trie.Search("penci"))

}
