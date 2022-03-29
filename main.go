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
	f, err := os.Open("snyk-output.txt")
	if err != nil {
	}

	defer f.Close()

	snykVulns := getSnykVulns(bufio.NewScanner(f))
	cves := getCVEsFrom(snykVulns)
	fmt.Println(cves)

}

func getCVEsFrom(snykVulns []string) []string {
	var cves []string
	for _, snykVuln := range snykVulns {
		cves = append(cves, getCVEFrom(snykVuln))
	}

	return cves
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

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
