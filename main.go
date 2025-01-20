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

	// Eingabe von Domains durch den Benutzer
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,darcRecord\n")
	fmt.Println(" write a Email Domain: ")
	for scanner.Scan() {

		checkDomain(scanner.Text()) //checkDomain Funktion wird aufgerufen
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)

	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool // MX-Einträge geben an, welche Mailserver für den Empfang von E-Mails für die Domäne zuständig sind. SPF (Sender Policy Framework),DMARC (Domain-based Message Authentication, Reporting & Conformance) Einträge.
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain) //  DNS-Abfrage für TXT-Einträge (Text Records)
	if err != nil {
		log.Printf("Error:%v\n", err)

	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") { // prüfen, ob ein gegebener String (record) mit einem bestimmten Präfix beginnt "v=spf1".
			hasSPF = true
			spfRecord = record
			break
		}

	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
