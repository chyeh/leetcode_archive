package solution

type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			t := Constructor()
			curr.children[c] = &t
		}
		curr = curr.children[c]
	}
	curr.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	if curr.isWord {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
