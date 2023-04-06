package zhourui_test

import (
	"fmt"
	"strings"
)

func GetAnStr() string {
	str1 := "a,b,c"
	arr1 := strings.Split(str1, ",")
	fmt.Println(arr1)
	return "dddd"
}
