package str

// TrieNode 代表 Trie 树的节点
type TrieNode struct {
	children map[byte]*TrieNode // 子节点映射
	isEnd    bool               // 是否是完整单词的结尾
}

// Trie 结构
type Trie struct {
	root *TrieNode
}

// NewTrie 初始化 Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}}
}

// Insert 插入操作
func (t *Trie) Insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := current.children[word[i]]; !ok {
			// 节点不存在，创建一个新节点
			current.children[word[i]] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		current = current.children[word[i]]
	}
}

// Find 查找操作
func (t *Trie) Find(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := current.children[word[i]]; !ok {
			// 节点不存在，返回 false
			return false
		}
		current = current.children[word[i]]
	}
	return current.isEnd
}

// StartWith 查找所有相同前缀
func (t *Trie) StartWith(word string) []string {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[byte(char)]; !ok {
			return []string{}
		}
		node = node.children[byte(char)]
	}
	return t.findAllWords(node, word)
}

func (t *Trie) findAllWords(node *TrieNode, word string) []string {
	var words []string
	if node.isEnd {
		words = append(words, word)
	}
	for char, childNode := range node.children {
		words = append(words, t.findAllWords(childNode, word+string(char))...)
	}
	return words
}
