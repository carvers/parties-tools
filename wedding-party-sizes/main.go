package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type person struct {
	name  string
	id    string
	party string
}

func importPeople(fpath string) (map[string]person, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	people := map[string]person{}
	_, err = r.Read() // skip column headers
	if err != nil {
		return nil, err
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		people[record[11]] = person{
			id:    record[11],
			party: record[5],
			name:  strings.Join(record[0:2], " "),
		}
	}
	return people, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./wedding-party-sizes {path to guest list.csv}")
		os.Exit(1)
	}
	people, err := importPeople(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing guest list: %+v\n", err)
		os.Exit(1)
	}
	parties := map[string]int{}
	sizes := map[int]int{}
	for _, person := range people {
		parties[person.party] = parties[person.party] + 1
	}
	for party, count := range parties {
		fmt.Println(party, "=", count)
		sizes[count] = sizes[count] + 1
	}
	for size, count := range sizes {
		fmt.Println("Parties with", size, "people:", count)
	}
}
