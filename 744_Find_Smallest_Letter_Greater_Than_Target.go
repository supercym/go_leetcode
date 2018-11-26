func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)-1
	for l <= h {
		m := int(l + (h-l)/2)
		if letters[m] == target {
			if m == len(letters)-1 {
				return letters[0]
			}
			for i := m+1; i < len(letters); i++ {
				if letters[i] > target {
					return letters[i]
				}
			}
			return letters[0]
		} else if letters[m] > target {
			h = m-1
		} else {
			l = m+1
		}
	}
	if h == len(letters)-1 {
		return letters[0]
	}
	return letters[h+1]
}