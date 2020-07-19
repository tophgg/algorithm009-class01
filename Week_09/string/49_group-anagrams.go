package set

// 解法1：先将每个字符串转化为byte 排序，然后用一个map映射字符串和他所在的数组位置
// 解法2：生成26个质数，分别于转化为byte的单词字母相乘，记录到map中，
// 解法3：使用字母对应26个数字，先把每个字符排序，再把每个字母使用#或其他字符来做间隔合并，保证每个字母卫衣

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	m := map[string]int{}
	res := [][]string{}
	for _, str := range strs {
		kbyte := bytes(str)
		sort.Sort(kbyte)
		kstr := string(kbyte)
		if idx, ok := m[kstr]; !ok {
			// 数组的第几个位置
			m[kstr] = len(res)
			res = append(res, []string{str})
		} else {
			res[idx] = append(res[idx], str)
		}
	}
	return res
}
