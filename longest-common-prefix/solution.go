package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if !areEntryStringsValid(strs) {
		return ""
	}

	var commonPrefix string

	for i := 0; i <= len(strs[0]); i++ {
		var (
			newPrefix     string
			newPrefixHits int
		)

		for strIndex, str := range strs {
			if len(str) < i {
				continue
			}

			if strIndex == 0 {
				newPrefix = str[:i]
			}

			if str[:i] == newPrefix {
				newPrefixHits++
			}
		}

		if newPrefixHits == len(strs) {
			commonPrefix = newPrefix
		}
	}

	return commonPrefix
}

func areEntryStringsValid(strs []string) bool {
	if len(strs) < 1 || len(strs) > 200 {
		return false
	}

	for _, str := range strs {
		if len(str) > 200 {
			return false
		}

		for _, char := range str {
			if char < 97 || char > 122 {
				return false
			}
		}
	}

	return true
}
