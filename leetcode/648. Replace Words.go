package main

import "strings"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: [26]*TrieNode{},
			isEnd:    false,
		},
	}
}

func (trie *Trie) getShortestRoot(word string) string {
	cur := trie.root

	for i, ch := range word {
		if cur.isEnd {
			return word[:i]
		}

		if cur.children[ch-'a'] == nil {
			return word
		}

		cur = cur.children[ch-'a']
	}

	return word
}

func (trie *Trie) push(word string) {
	cur := trie.root

	for _, ch := range word {
		if cur.children[ch-'a'] == nil {
			cur.children[ch-'a'] = &TrieNode{
				children: [26]*TrieNode{},
				isEnd:    false,
			}
		}

		cur = cur.children[ch-'a']
	}

	cur.isEnd = true
}

func replaceWords(dictionary []string, sentence string) string {
	trie := NewTrie()
	for _, word := range dictionary {
		trie.push(word)
	}

	result := make([]string, 0)

	for _, word := range strings.Split(sentence, " ") {
		result = append(result, trie.getShortestRoot(word))
	}

	return strings.Join(result, " ")
}
