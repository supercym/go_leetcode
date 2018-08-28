func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	tmp := []byte{}
	for i := 0; i < len(s); i++ {
		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) {
			tmp = append(tmp, s[i])
		} else if s[i] >= 97 && s[i] <= 122 {
			low_char := s[i] - 32
			tmp = append(tmp, low_char)
		}
	}
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}
	return true
}
