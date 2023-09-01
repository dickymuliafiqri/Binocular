package slice

func RemoveAllByString(arrayString []string, removedString string) []string {
	for i, s := range arrayString {
		if s == removedString {
			return RemoveAllByString(append(arrayString[:i], arrayString[i+1:]...), removedString)
		}
	}

	return arrayString
}
