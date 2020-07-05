package main

type Trie struct {
	is_end bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	var node *Trie = this
	for _, ch := range word {
		if node.next[ch - 'a'] == nil {
			node.next[ch - 'a'] = new(Trie)
		}
		node = node.next[ch - 'a']
	}
	node.is_end = true
}

// 判断一个单词是否为trie树
/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var node *Trie = this
	for _, ch := range word {
		node = node.next[ch - 'a']
		if node == nil {
			return false
		}
	}
	return node.is_end
}

// 判断一个前缀是否在Trie树中
/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node *Trie = this
	for _, ch := range prefix {
			node = node.next[ch - 'a']
		if node == nil {
			return false
		}
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