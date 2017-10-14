package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	err := printValidLabels(path)
	if err != nil {
		fmt.Printf("Error printing labels: %v\n", err)
		os.Exit(1)
	}
}

func printValidLabels(path string) error {
	labels, err := NewFromCsv(path)
	if err != nil {
		return fmt.Errorf("could not parse labels from csv: %v", err)
	}

	for _, l := range labels {
		err := l.Validate()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n\n", l)
	}

	return nil
}

type Label struct {
	Names string
	Line1 string
	Line2 string
	City  string
	State string
	Zip   string
}

func NewLabel(fields []string) (*Label, error) {
	if len(fields) != 6 {
		return nil, fmt.Errorf("expected 6 fields, got %d", len(fields))
	}
	return &Label{
		fields[0],
		fields[1],
		fields[2],
		fields[3],
		fields[4],
		fields[5],
	}, nil
}

func NewFromCsv(path string) ([]*Label, error) {
	records, err := readCsv(path)
	if err != nil {
		return nil, err
	}

	labels := []*Label{}
	for _, rec := range records {
		l, err := NewLabel(rec)
		if err != nil {
			return nil, err
		}
		labels = append(labels, l)
	}

	return labels, nil
}

func (l *Label) String() string {
	str := l.Names
	str += "\n" + l.Line1
	if l.Line2 != "" {
		str += "\n" + l.Line2
	}
	str += "\n" + l.City + ", " + l.State + " " + l.Zip
	return str
}

func (l *Label) Validate() error {
	if l.Names == "" {
		return fmt.Errorf("Names field is empty")
	} else if l.Line1 == "" {
		return fmt.Errorf("Line1 field is empty")
	} else if l.City == "" {
		return fmt.Errorf("City field is empty")
	} else if l.State == "" {
		return fmt.Errorf("State field is empty")
	} else if l.Zip == "" {
		return fmt.Errorf("Zip field is empty")
	}
	return nil
}

func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
