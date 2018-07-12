
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		colsFlag  string
		delimFlag string
		cols      []int
	)
	flag.StringVar(&colsFlag, "cols", "", "Comma separated indexes of column to display. Default is all columns. First column is 0.  Order is preserved.")
	flag.StringVar(&delimFlag, "delim", ",", "Delimiter. Default is comma.")
	flag.Parse()
	if colsFlag != "" {
		for _, cs := range strings.Split(colsFlag, ",") {
			c, err := strconv.Atoi(cs)
			if err != nil {
				panic("invalid columns:" + colsFlag)
			}
			cols = append(cols, c)
		}
	}
	r := csv.NewReader(os.Stdin)
	var sel []string
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if colsFlag == "" {
			sel = row
		} else {
			sel = sel[:0]
			for _, c := range cols {
				sel = append(sel, row[c])
			}
		}
		fmt.Println(strings.Join(sel, delimFlag))
	}
}
