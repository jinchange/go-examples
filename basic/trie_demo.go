package basic

import "fmt"

// Trie 节点结构定义
type Trie struct {
	children map[rune]*Trie // 子节点（用 map 存储，支持 Unicode 字符）
	isEnd    bool           // 标记是否为单词结尾
}

// 初始化 Trie 树
func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
		isEnd:    false,
	}
}

// 插入单词
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word { // 遍历单词的每个字符
		if node.children[ch] == nil { // 如果字符路径不存在
			node.children[ch] = NewTrie() // 创建新节点
		}
		node = node.children[ch] // 移动到子节点
	}
	node.isEnd = true // 标记单词结尾
}

// 搜索完整单词是否存在
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			return false // 路径中断，单词不存在
		}
		node = node.children[ch]
	}
	return node.isEnd // 检查是否为单词结尾
}

// 检查前缀是否存在
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true // 前缀路径存在即可
}

// 示例：查找所有以某前缀开头的单词（递归实现）
func (t *Trie) SearchPrefix(prefix string) []string {
	node := t
	// 先定位到前缀的最后一个节点
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	// 收集所有以该前缀开头的单词
	var results []string
	t.dfs(node, prefix, &results)
	return results
}

// 深度优先遍历收集单词
func (t *Trie) dfs(node *Trie, currentWord string, results *[]string) {
	if node.isEnd {
		*results = append(*results, currentWord)
	}
	for ch, child := range node.children {
		t.dfs(child, currentWord+string(ch), results)
	}
}

func TestTrieDemo() {
	trie := NewTrie()

	// 插入示例单词
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("application")
	trie.Insert("banana")

	// 测试搜索
	fmt.Println("Search 'app':", trie.Search("app"))     // true
	fmt.Println("Search 'apple':", trie.Search("apple")) // true
	fmt.Println("Search 'appl':", trie.Search("appl"))   // false

	// 测试前缀匹配
	fmt.Println("StartsWith 'app':", trie.StartsWith("app")) // true
	fmt.Println("StartsWith 'ba':", trie.StartsWith("ba"))   // true

	// 查找所有以 "app" 开头的单词
	fmt.Println("SearchPrefix 'app':", trie.SearchPrefix("app"))
	// 输出: [app apple application]

	//Search 'app': true
	//Search 'apple': true
	//Search 'appl': false
	//StartsWith 'app': true
	//StartsWith 'ba': true
	//SearchPrefix 'app': [app apple application]
}
