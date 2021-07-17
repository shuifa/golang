package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("+-132  1ssdas"))
}

func myAtoi(s string) int {

	var symbol int = 1
	var str strings.Builder
	var flag bool

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !flag {
			continue
		}
		if !flag && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				symbol = -1
			}
			flag = true
		} else if s[i] >= '0' && s[i] <= '9' {
			str.WriteByte(s[i])
			flag = true
		} else {
			break
		}
	}

	i, _ := strconv.Atoi(str.String())

	ret := i * symbol

	if ret < math.MinInt32 {
		return math.MinInt64
	}

	if ret > math.MaxInt64 {
		return math.MaxInt64
	}
	return ret
}
