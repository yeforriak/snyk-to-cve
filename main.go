package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Usage: ./snyk-cve <snyk-output>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
	}
	defer f.Close()

	snykVulns := getSnykVulns(bufio.NewScanner(f))

	requests := make(chan string)
	go attachCVEs(requests)

	for _, v := range snykVulns {
		requests <- v
	}
	close(requests)
}

func attachCVEs(requests chan string) {
	for job := range requests {
		fmt.Printf("%s -> %s\n", job, getCVEFrom(job))
	}
}
func getCVEFrom(snykVuln string) string {
	resp, err := http.Get(snykVuln)

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	r, _ := regexp.Compile("CVE-\\d{4}-\\d{4,7}")
	match := r.FindString(string(b))
	return match
}

func getSnykVulns(scanner *bufio.Scanner) []string {
	var snykVulns []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Info: ") {
			snykVulns = append(snykVulns, line[8:])
		}
	}

	if err := scanner.Err(); err != nil {
		//TODO: handle error
	}

	return snykVulns
}
