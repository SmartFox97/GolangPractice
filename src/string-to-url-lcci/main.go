package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ",13))
	fmt.Println(replaceSpaces("               ", 5))
}

func replaceSpaces(S string, length int) string {
	str := S[:length]
	return strings.Replace(str, " ", "%20", -1)
}
