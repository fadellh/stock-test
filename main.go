package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	ana := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagramSplit(ana)
	str := findFirstStringInBracket("(kita)(test)))()")

	fmt.Println(str)

}

func anagramSplit(ana []string) [][]string {
	saver := map[string][]string{}
	var orderStr string
	var result [][]string

	for _, item := range ana {
		orderStr = orderString(item)
		saver[orderStr] = append(saver[orderStr], item)
	}

	for _, value := range saver {
		result = append(result, value)
	}

	return result
}

func orderString(item string) string {
	var nums []int

	for _, char := range item {
		nums = append(nums, int(char))
	}

	sort.Ints(nums)
	var output string

	for _, number := range nums {
		output += string(rune(number))
	}

	return output
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")

		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")

			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			}
		}
	}

	return ""
}