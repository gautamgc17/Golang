package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// NewScanner returns a new Scanner to read from r.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Email Validator Golang Tool")
	fmt.Printf("Domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord \n")

	// Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method.
	// It returns false when the scan stops, either by reaching the end of the input or an error.
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input %v \n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	fmt.Printf("MX Records: %v \n", mxRecords)
	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	fmt.Printf("SPF Records: %v \n", spfRecords)
	for _, record := range spfRecords {
		fmt.Println(record)
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	fmt.Printf("DMARC Records: %v \n", dmarcRecords)
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Final Result: %v %v %v %v %v %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
