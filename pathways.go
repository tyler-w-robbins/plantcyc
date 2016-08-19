package plantcyc

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Pathway struct {
	ID       string
	Name     string
	GeneName []string
	GeneID   []string
	Valid    bool
}

func ParsePathways(path string) []*Pathway {
	Pathways := []*Pathway{}

	dat, err := os.Open(path)
	check(err)

	// Initialize csv reader
	r := csv.NewReader(dat)
	r.Comma = '\t'
	r.Comment = '#'
	r.LazyQuotes = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		p := new(Pathway)
		nameLoc := 0
		idLoc := 0
		value := 0

		// Examine header and store locations of column names
		if record[0] == "UNIQUE-ID" {
			for i := range record {
				if record[i] == "GENE-NAME" {
					nameLoc = i
				}
				if record[i] == "GENE-ID" {
					idLoc = i
				}
			}
			value = len(record)
		}

		// Switch to iterate through each column in a record
		switch value {
		case 0:
			p.ID = record[value]
			value++
			fallthrough
		case 1:
			p.Name = record[value]
			value++
			fallthrough
			// This case begins w
		case 2:
			for record[value] != "" && value <= nameLoc {
				p.GeneName = append(p.GeneName, record[value])
				value++
			}
			value = nameLoc + 1
			fallthrough
			// Final case begins where the GENE-NAME columns end in file to be parsed
		case nameLoc + 1:
			for record[value] != "" && value <= idLoc {
				p.GeneID = append(p.GeneID, record[value])
				value++
			}
			p.Valid = true
			Pathways = append(Pathways, p)
		}

	}
	return Pathways
}

func WritePathways(path string, w *bufio.Writer, p []*Pathway) error {
	for i := range p {
		_, err := w.WriteString("PCYC:" + p[i].ID + "|" + p[i].Name + "|PlantCyc_Pathways|Pathway\n")
		check(err)
	}
	return nil
}
