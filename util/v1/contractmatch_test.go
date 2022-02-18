package util_test

import (
	"testing"

	"github.com/raufhm/contract-match/util/v1"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	contract string
	rawData  string
	expected int
}

func TestContractMatch(t *testing.T) {

	raw := []*testData{
		{
			contract: "000",
			rawData:  "aaababbuuu",
			expected: 2,
		},
		{
			contract: "100",
			rawData:  "babbuuu",
			expected: 1,
		},
		{
			contract: "011",
			rawData:  "aiibbocceccdihh",
			expected: 4,
		},
		{
			contract: "00",
			rawData:  "aabbuu",
			expected: 2,
		},
	}

	for _, r := range raw {
		result := util.ContractMatch(r.contract, r.rawData)
		assert.Equal(t, result, r.expected)
	}
}
