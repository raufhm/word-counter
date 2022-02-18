package util

func ContractMatch(pattern string, s string) int {
	patternArr := []string{}
	for _, p := range pattern {
		if string(p) == "0" {
			patternArr = append(patternArr, string(p))
		} else {
			patternArr = append(patternArr, string(p))
		}
	}

	contractLen := len(patternArr)
	dataArr := []string{}
	for _, data := range s {
		dataArr = append(dataArr, string(data))
	}

	result := 0
	for idx := range dataArr {
		if len(dataArr)-idx < contractLen {
			break
		}
		matchCount := checkMatchdData(patternArr, dataArr[idx:idx+contractLen])
		result += matchCount
	}

	return result
}

func checkMatchdData(contract []string, rawData []string) int {
	matchCounter := 0
	for idx, d := range rawData {
		for i, c := range contract {
			if idx != i {
				continue
			}
			response := checkExistVowel(d)
			if response == c {
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
