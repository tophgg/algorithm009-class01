package main

// 解法1：bfs-dfs-dfs，先找第一层起始变一个，能在库里的，每变一个就下沉一次，次数+1，直到变的结果等于end，匹配一个基因库的基因可以把它删除，因为要最少变化次数
// 遍历基因的每个碱基，然后将其进行替换，如果在基因库中有候选，那就完成一次突变，并从基因库中将对应的基因删除
func minMutation(start string, end string, bank []string) int {
	used:=make(map[string]bool)
	for _, str := range bank {
		used[str] = true
	}
	if _,ok := used[end]; !ok {
		return -1
	}
	bstr := []rune{'A','C','G','T'}
	queue:=[]string{start}
	l := len(start)
	nums:=0
	for len(queue) > 0 {
		nums++
		tmp := []string{}
		for _,k := range queue {
			chs := []rune(k)
			for i:=0; i<l; i++ {
				ch := chs[i]
				for _,c := range bstr {
					// for c:='A'; c<='Z'; c++ {
					chs[i] = c
					t := string(chs)
					// 比较
					if t == end {
						return nums
					}
					if _,ok := used[t]; !ok {
						continue
					}
					delete(used, t)
					tmp = append(tmp, t)
				}
				chs[i] = ch
			}
		}
		queue = tmp
	}
	return -1
}