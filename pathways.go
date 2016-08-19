package plantcyc

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Pathway struct {
	ID       string
	Name     string
	GeneName []string
	GeneID   []string
	Gene     []*Gene
	Valid    bool
}

func ParsePathways(path string) []*Pathway {
	Pathways := []*Pathway{}

	dat, err := os.Open(path)
	check(err)

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
			fmt.Println(value)
		}
		switch value {
		case 0:
			p.ID = record[value]
			value++
			fallthrough
		case 1:
			p.Name = record[value]
			value++
			fallthrough
		case 2:
			for record[value] != "" && value <= nameLoc {
				p.GeneName = append(p.GeneName, record[value])
				value++
			}
			value = nameLoc
			fallthrough
		case 91:
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
