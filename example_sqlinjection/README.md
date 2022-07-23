# README

Example for SQL Injection and two-file flow

Execution ` % go vet -vettool=$(which levee) -config=./analyzer_configuration.yaml ./...`

Analysis results

```
# gopherconeu22/taint/injection
./authentication.go:12:17: a source has reached a sink
 source: ./authentication.go:10:22	
```