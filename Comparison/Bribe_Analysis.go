package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func minimumBribes(q []int32) {
	is_swapped := false
	var j int32
	bribe := make(map[int32]int32)
	for i, _ := range q {
		for j = 0; j < (int32)(len(q)-i-1); j++ {
			if q[j] > q[j+1] {
				q[j], q[j+1] = q[j+1], q[j]
				bribe[q[j+1]]++
				is_swapped = true
			}
		}
		if !is_swapped {
			break
		}
	}

	count := 1
	for _, v := range bribe {
		if v > 2 {
			fmt.Println("Too chaotic")
			os.Exit(0)
		}
		if v > 0 && v <= 2 {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
