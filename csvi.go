package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const head = 40

func main() {
	data := make(map[string]map[string]int)
	var headers []string
	sums := make(map[string]int)
	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]
		f, err := os.Open(fn)
		if err != nil {
			panic("error opening file:" + os.Args[i])
		}
		r := bufio.NewScanner(f)
		for r.Scan() {
			var col string
			var val int
			fmt.Sscanf(strings.Trim(r.Text(), " "), "%d %s", &val, &col)
			if _, ok := data[col]; !ok {
				data[col] = make(map[string]int)
				headers = append(headers, col)
			}
			data[col][fn] = val
			sums[col] += val
		}
		f.Close()
	}
	sort.Slice(headers, func(i, j int) bool {
		return sums[headers[i]] > sums[headers[j]]
	})
	fmt.Print(";")
	fmt.Println(strings.Join(headers[:head], ";"))
	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]
		fmt.Print(fn)
		fmt.Print(";")
		for _, h := range headers[:head] {
			fmt.Print(data[h][fn])
			fmt.Print(";")
		}
		fmt.Println()
	}
}
