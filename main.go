package main

import (
	"fmt"

	util "github.com/raufhm/contract-match/util/v1"
)

func main() {
	result := util.ContractMatch("001", "aaaaabaab")
	fmt.Println(result)
}
