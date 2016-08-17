package plantcyc

import (
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

	r := csv.NewReader(dat)
	r.Comma = '\t'
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		p := new(Pathway)
		value := 0
		if record[0] == "UNIQUE-ID" {
			value = 179
		}
		switch value {
		case 0:
			p.ID = record[value]
			// fmt.Printf("	%v\n", record[value])
			value++
			fallthrough
		case 1:
			p.Name = record[value]
			value++
			fallthrough
		case 2:
			for record[value] != "" && value < 91 {
				p.GeneName = append(p.GeneName, record[value])
				value++
			}
			value = 91
			fallthrough
		case 91:
			for record[value] != "" && value < 179 {
				p.GeneID = append(p.GeneID, record[value])
				value++
			}
			p.Valid = true
			Pathways = append(Pathways, p)
		}

	}
	return Pathways
}
