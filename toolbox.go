package walk

func CharIsAlphaNumeric(char rune) bool {
	return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func CharIsAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func CharIsNumeric(char rune) bool {
	return (char >= '0' && char <= '9')
}
func StringIsAlphaNumeric(chars string) bool {
	for _, char := range chars {
		if !CharIsAlphaNumeric(char) {
			return false
		}
	}
	return len(chars) > 0
}

func StringIsAlpha(chars string) bool {
	for _, char := range chars {
		if !CharIsAlpha(char) {
			return false
		}
	}
	return len(chars) > 0
}

func StringIsNumeric(chars string) bool {
	for _, char := range chars {
		if !CharIsNumeric(char) {
			return false
		}
	}
	return len(chars) > 0
}

// IsAbsolutePath returns true if path is not relative (starts from root)
func IsAbsolutePath(name string) bool {
	return len(name) > 2 && name[1] == ':' && CharIsAlpha(rune(name[0]))
}
