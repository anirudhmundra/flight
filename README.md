
# shahaj-flight-assignment

This solutions fetches tickets records from a csv file, validates them and generate valid and invalid tickets csv files respectively.



## Solution

The current solution is really simple, which loads all the entries from the file and processes them altogether. 
It uses dependency injection, loose coupling and interfaces in order to accomodate different ways of reading and writing the data. 
We can later add JSON/XML/protobuf implementations and inject them.

## Future Actions

It can further be enchanced using concurrency, thread-safety and batch processing. 
Multiple goroutines could process 50/100 feeds at a time and dump them into the respective files. 

## Commands

Test

`go test ./...`

Run

`inputFile - testData/tickets.csv`

`outputFiles - testData/output/valid_tickets.csv, testData/output/invalid_tickets.csv`

`go run main.go` 
