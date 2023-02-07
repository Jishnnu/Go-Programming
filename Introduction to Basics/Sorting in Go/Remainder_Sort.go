package main

import (
	"fmt"
	"sort"
)

func stringSort(strArr []string) []string {
	arr := make(map[string]int)

	for _, v := range strArr {
		arr[v] = len(v) % 3
	}

	keys := make([]string, 0, len(arr))

	for key := range arr {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return arr[keys[i]] < arr[keys[j]]
	})

	i := 0
	for _, k := range keys {
		strArr[i] = k
		i++
	}
	return strArr
}

func main() {
	var str = []string{"abc", "ab", "bc", "a"}
	str = stringSort(str)
	fmt.Println(str)
}
