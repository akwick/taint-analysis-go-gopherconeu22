# README

Example for a CWE-522 vulnerability

0. Install Go Flow Levee
1. `$ go vet -vettool=$(which levee) -config=./analyzer_configuration.yaml ./cwe522.go`

Output:
```
... (panic logs)
./cwe522.go:17:13: a source has reached a sink
 source: ./cwe522.go:14:19
```
