package main
// 时间复杂度O(J+S)，空间复杂度O(J)
func numJewelsInStones(J string, S string) int {
	hash := make(map[byte]int)
	for i:=0; i<len(J); i++ {
		hash[J[i]] = 1
	}
	length:=0
	for j:=0; j<len(S); j++ {
		if _, ok := hash[S[j]]; ok {
			length++
		}
	}
	return length
}
