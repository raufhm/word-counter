package util

func ContractMatch(pattern string, s string) int {
	contractLen := len([]rune(pattern))
	dataArr := []rune(s)
	result := 0
	for idx := range dataArr {
		if len(dataArr)-idx < contractLen {
			break
		}
		matchCount := checkMatchdData([]rune(pattern), dataArr[idx:idx+contractLen])
		result += matchCount
	}
	return result
}

func checkMatchdData(contract []rune, rawData []rune) int {
	matchCounter := 0
	for idx, d := range rawData {
		for i, c := range contract {
			if idx != i {
				continue
			}
			response := checkExistVowel(string(d))
			if response == string(c) {
				matchCounter++
				break
			}
		}
	}
	if matchCounter == len(contract) {
		return 1
	}
	return 0

}

func checkExistVowel(d string) string {
	vowel := []string{"a", "e", "i", "u", "o"}
	for _, v := range vowel {
		if v == d {
			return "0"
		}
	}
	return "1"
}
