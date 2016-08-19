package plantcyc

import (
	"bufio"
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

func WriteGenes(path string, w *bufio.Writer, g []*Gene) error {
	// Parse gene nodes
	for i := range g {
		_, err := w.WriteString("PCYC:" + g[i].ID + "|" + g[i].Name)
		check(err)
		// Sometimes this field is blank
		if g[i].SwissProtID != "" {
			_, err = w.WriteString(";" + g[i].SwissProtID)
			check(err)
		}
		// Synonyms are stored as a string array, so appends a string for each synonym
		for _, syn := range g[i].Synonyms {
			_, err := w.WriteString(";" + syn)
			check(err)
		}
		_, err = w.WriteString("|" + g[i].Product + "|PlantCyc_Gene|Gene\n")
		check(err)
	}
	err := w.Flush()
	check(err)
	return nil
}
