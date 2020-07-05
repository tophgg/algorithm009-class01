package main

// 前缀树+dfs，将words中的单词写入Trie中，然后遍历board，
// dp
type Trie struct{
	letter string
	// is_end bool
	next [26]*Trie
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		node := root
		for _, ch := range word {
			if node.next[ch-'a'] == nil {
				node.next[ch-'a'] = &Trie{}
			}
			node = node.next[ch-'a']
		}
		node.letter = word
	}

	col := len(board)
	if col == 0 { return []string{} }
	row := len(board[0])
	if row == 0 { return []string{} }
	result := []string{}
	for i:=0; i<col; i++ {
		for j:=0; j<row; j++ {
			dfs(i, j, &result, board, root)
		}
	}
	return result
}

func dfs(i, j int, result *[]string, board [][]byte, t *Trie) {
	// 定义结束条件
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	ch := board[i][j]

	// 当前层处理,查找前缀树
	if ch == '#' || t.next[ch-'a'] == nil { // 已经访问过 则不处理
		return
	}
	node := t.next[ch-'a']
	if node.letter != "" {
		*result = append(*result, node.letter)
		// 防止重复写入单词，所以置为空
		node.letter = ""
	}
	// 标记当前dfs为已访问
	board[i][j] = '#'
	// 下探下一层
	dfs(i+1, j, result, board, node)
	dfs(i, j+1, result, board, node)
	dfs(i-1, j, result, board, node)
	dfs(i, j-1, result, board, node)
	// 遍历完周围4个点后，还原当前点为原始值
	board[i][j] = ch
}


// func (this *Trie) insert (word string) {
//     for _, ch := range word {
//         if this.next[ch-'a'] == nil {
//             this.next[ch-'a'] = new(Trie)
//         }
//         this = this.next[ch-'a']
//     }
//     this.letter = word
//     this.is_end = true
// }

// func (this *Trie) search (word string) bool {
//     for _, ch := range word {
//         if this.next[ch-'a'] == nil {
//             return false
//         }
//         this = this.next[ch-'a']
//     }
//     return this.is_end
// }