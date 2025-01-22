This tool checks email domains for various DNS records to assess the configuration and security of the domain.
It checks for MX, SPF, and DMARC records and outputs the results in CSV format.

Features
MX Records: Checks for the presence of MX records that specify which mail servers are responsible for receiving emails for the domain.
SPF Records: Checks for the presence of SPF (Sender Policy Framework) records that help prevent email spoofing.
DMARC Records: Checks for the presence of DMARC (Domain-based Message Authentication, Reporting & Conformance) records that provide an additional layer of security.

Usage
Run the tool.
Enter the email domain you want to check.
The tool outputs the results in CSV format: domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord.

Example
$ go run main.go
 write a Email Domain: 
yahoo.com
yahoo.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
