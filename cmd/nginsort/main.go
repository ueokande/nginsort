package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/ueokande/nginsort"
)

func main() {
	var logs []nginsort.AccessLog
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		log, err := nginsort.Parse(scanner.Text())
		if err != nil {
			panic(err)
		}
		logs = append(logs, *log)
	}

	sort.Sort(nginsort.ByDate(logs))
	for _, log := range logs {
		fmt.Println(log.Origin)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
