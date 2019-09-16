package problem0003

// support ASCII code only
func lengthOfLongestSubstringAscii(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中没有出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	//locMap := make(map[string]int, 256)
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

// Support UNICODE
func lengthOfLongestSubstringUnicode(s string) int {

	// convert s to slice of rune for all Unicode runes
	runes := []rune(s)

	// location map, key is string , value is the location of which the key is appeared in s.
	// the max length of the map is the len of s.
	location := make(map[int32]int, len(runes))

	// init map, -1 means no appearance yet.
	for _, v := range s {
		if _, ok := location[v]; !ok {
			location[v] = -1
		}
	}

	maxLen, left := 0, 0

	for k, v := range runes {
		// char already exist in s[left:k]
		if location[v] >= left {
			left = location[v] + 1
		} else if curLen := k - left + 1; curLen > maxLen {
			maxLen = curLen
		}
		location[v] = k
	}

	return maxLen
}
