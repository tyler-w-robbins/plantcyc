package plantcyc

import (
	"bufio"
	"os"
	"strings"
)

type Compound struct {
	ID         string
	Name       string
	Comment    []string
	DBLinks    []string
	Synonyms   []string
	SMILes     string
	SystemName string
}

func ParseCompounds(path string) []*Compound {
	Compounds := []*Compound{}

	dat, err := os.Open(path)
	check(err)
	c := new(Compound)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "UNIQUE-ID") {
			c = new(Compound)
			c.ID = strings.TrimPrefix(scanner.Text(), "UNIQUE-ID - ")
		}
		if strings.HasPrefix(scanner.Text(), "COMMON-NAME") {
			c.Name = strings.TrimPrefix(scanner.Text(), "COMMON-NAME - ")
		}
		if strings.HasPrefix(scanner.Text(), "COMMENT") {
			c.Comment = append(c.Comment, strings.TrimPrefix(scanner.Text(), "COMMENT - "))
		}
		if strings.HasPrefix(scanner.Text(), "/") && !strings.HasPrefix(scanner.Text(), "//") {
			c.Comment = append(c.Comment, strings.TrimPrefix(scanner.Text(), "/"))
		}
		if strings.HasPrefix(scanner.Text(), "DBLINKS") {
			c.DBLinks = append(c.DBLinks, strings.TrimPrefix(scanner.Text(), "DBLINKS - "))
		}
		if strings.HasPrefix(scanner.Text(), "SYNONYMS") {
			c.Synonyms = append(c.Synonyms, strings.TrimPrefix(scanner.Text(), "SYNONYMS - "))
		}
		if strings.HasPrefix(scanner.Text(), "SMILES") {
			c.SMILes = strings.TrimPrefix(scanner.Text(), "SMILES - ")
		}
		if strings.HasPrefix(scanner.Text(), "SYSTEMATIC-NAME") {
			c.SystemName = strings.TrimPrefix(scanner.Text(), "SYSTEMATIC-NAME - ")
		}
		if scanner.Text() == "//" {
			Compounds = append(Compounds, c)
		}
	}
	return Compounds
}

func WriteCompounds(path string, w *bufio.Writer, c []*Compound) error {
	for i := range c {
		_, err := w.WriteString("PCYC:" + c[i].ID + "|" + c[i].Name + "|PlantCyc_Chemicals|")
		check(err)
		for _, com := range c[i].Comment {
			_, err = w.WriteString(strings.Replace(com, "|", "", -1))
			check(err)
		}
		_, err = w.WriteString("|")
		check(err)
		for length, syn := range c[i].Synonyms {
			if length > 0 {
				_, err = w.WriteString(";")
				check(err)
			}
			_, err = w.WriteString(strings.Replace(syn, "|", "", -1))
			check(err)
		}
		_, err = w.WriteString("|Chemical\n")
		check(err)
	}
	return nil
}
