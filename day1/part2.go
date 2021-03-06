package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []int
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("could not parse %q: %w", str, err)
		}
		lines = append(lines, number)
	}

	for i := range lines[:len(lines)-2] {
		for j := range lines[i+1 : len(lines)-1] {
			for k := range lines[j+1:] {
				if lines[i]+lines[j]+lines[k] == 2020 {
					log.Printf("%d * %d * %d= %d", lines[i], lines[j], lines[k], lines[i]*lines[j]*lines[k])
					os.Exit(0)
				}
			}
		}
	}
}
