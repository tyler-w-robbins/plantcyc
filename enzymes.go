package plantcyc

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Enzyme struct {
	ID    string
	Name  string
	Path  []string
	Act   []string
	Inhib []string
	Valid bool
}

func ParseEnzymes(path string) []*Enzyme {
	Enzymes := []*Enzyme{}

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
		e := new(Enzyme)
		value := 0
		if record[0] == "UNIQUE-ID" {
			value = 179
		}
		switch value {
		case 0:
			e.ID = record[value]
			value++
			fallthrough
		case 1:
			e.Name = record[value]
			value += 2
			fallthrough
		case 3:
			if record[value] != "" {
				e.Path = append(e.Path, record[value])
				value++
				if record[value] != "" {
					e.Path = append(e.Path, record[value])
					value++
					if record[value] != "" {
						e.Path = append(e.Path, record[value])
						value++
						if record[value] != "" {
							e.Path = append(e.Path, record[value])
							value += 5
						} else {
							value += 5
						}
					} else {
						value += 6
					}
				} else {
					value += 7
				}
			} else {
				value += 8
			}
			fallthrough
		case 11:
			if record[value] != "" {
				e.Act = append(e.Act, record[value])
				value++
				//fmt.Println(value)
				if record[value] != "" {
					e.Act = append(e.Act, record[value])
					value++
					if record[value] != "" {
						e.Act = append(e.Act, record[value])
						value++
						if record[value] != "" {
							e.Act = append(e.Act, record[value])
							value++
						} else {
							value++
						}
					} else {
						value += 2
					}
				} else {
					value += 3
				}
			} else {
				value += 4
			}
			fallthrough
		case 14:
			if record[value] != "" {
				e.Inhib = append(e.Inhib, record[value])
				value++
				if record[value] != "" {
					e.Inhib = append(e.Inhib, record[value])
					value++
					if record[value] != "" {
						e.Inhib = append(e.Inhib, record[value])
						value++
					}
				}
			}
			e.Valid = true
			Enzymes = append(Enzymes, e)
		}
	}
	return Enzymes
}

func WriteEnzymes(path string, w *bufio.Writer, e []*Enzyme) error {
	for i := range e {
		_, err := w.WriteString("PCYC:" + e[i].ID + "|" + e[i].Name + "|PlantCyc_Pathways|Pathway\n")
		check(err)
	}
	return nil
}
