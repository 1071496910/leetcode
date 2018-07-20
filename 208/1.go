//Implement a trie with insert, search, and startsWith methods.
//
//Example:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // returns true
//trie.search("app");     // returns false
//trie.startsWith("app"); // returns true
//trie.insert("app");
//trie.search("app");     // returns true

package main

//import (
//	"fmt"
//)

type Trie struct {
	child [26]*Trie
	wd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		wd: false,
	}

}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}
	index := ([]byte(word[0:1]))[0] - byte('a')
	if this.child[index] == nil {
		this.child[index] = &Trie{
			wd: false,
		}
	}

	if word[1:] == "" {
		this.child[index].wd = true
		return
	} else {
		this.child[index].Insert(word[1:])
	}

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" || this == nil {
		return false
	}
	index := ([]byte(word[0:1]))[0] - byte('a')
	if this.child[index] == nil {
		return false
	}
	if word[1:] == "" {
		return this.child[index].wd == true
	}
	return this.child[index].Search(word[1:])

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	if this == nil {
		return false
	}
	index := ([]byte(prefix[0:1]))[0] - byte('a')
	if this.child[index] == nil {
		return false
	}
	if prefix[1:] == "" {
		return true
	}
	return this.child[index].StartsWith(prefix[1:])

}

/**
 * Your Trie object will be instantiated and called as such:
 */

//func main() {
//	obj := Constructor()
//	obj.Insert("hello")
//	fmt.Println(obj.Search("hello"))
//	fmt.Println(obj.StartsWith("hel"))
//	fmt.Println(obj)
//}
