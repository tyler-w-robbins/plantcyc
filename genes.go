package plantcyc

import (
	"encoding/csv"
	"io"
	"os"
)

type Gene struct {
	ID          string
	Name        string
	Product     string
	SwissProtID string
	Synonyms    []string
	Valid       bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseGenes(path string) []*Gene {
	Genes := []*Gene{}

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

		g := new(Gene)
		value := 0

		if record[0] == "UNIQUE-ID" {
			value = 179
		}
		switch value {
		case 0:
			g.ID = record[value]
			value++
			fallthrough
		case 1:
			g.Name = record[value]
			value++
			fallthrough
		case 2:
			g.Product = record[value]
			value++
			fallthrough
		case 3:
			g.SwissProtID = record[value]
			value += 4
			fallthrough
		case 7:
			if record[value] != "" {
				g.Synonyms = append(g.Synonyms, record[value])
				value++
				if record[value] != "" {
					g.Synonyms = append(g.Synonyms, record[value])
					value++
					if record[value] != "" {
						g.Synonyms = append(g.Synonyms, record[value])
						value++
						if record[value] != "" {
							g.Synonyms = append(g.Synonyms, record[value])
						}
					}
				}
			}
			g.Valid = true
			Genes = append(Genes, g)
		}
	}
	return Genes
}
