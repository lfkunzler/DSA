package lc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constraints:
//
//	0 <= len(strs) < 100
//	0 <= len(strs[i]) < 200
func encode(strs []string) string {
	var res strings.Builder

	for _, str := range strs {
		length := len(str)
		if length > 999 {
			continue
		}

		res.WriteString(fmt.Sprintf("%03s", strconv.Itoa(length)))
		res.WriteString(str)
	}

	return res.String()
}

func decode(str string) []string {
	res := []string{}

	for i := 0; i < len(str); {
		lengthStr := str[i : i+3]
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error converting length field=<%s> at position=<%d>\n", lengthStr, i)
			continue
		}

		i += 3
		if i+length > len(str) {
			fmt.Fprintf(os.Stderr, "declared length=<%d> exceeds the string length=<%d>\n", i+length, len(str))
			continue
		}

		aux := str[i : i+length]
		res = append(res, aux)

		i += length
	}

	return res
}
