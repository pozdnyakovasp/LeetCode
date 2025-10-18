package solutions

func LengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	length := 1
	charMap := make(map[byte]int)
	currentCharIndex := 0
	lastLeft := 0
	for ; currentCharIndex < len(s); currentCharIndex++ {
		currentChar := s[currentCharIndex]
		if index, ok := charMap[currentChar]; ok {
			ln := currentCharIndex - lastLeft
			length = max(length, ln)
			lastLeft = max(lastLeft, index+1)
		}
		charMap[currentChar] = currentCharIndex
	}

	return max(length, currentCharIndex-lastLeft)
}
