package plantcyc

import (
	"bufio"
	"os"
	"strings"
)

type Protein struct {
	ID        string
	Types     string
	Name      []string
	Catalyzes []string
	Gene      []string
	Synonyms  []string
	GOTerms   []string
	Citations string
}

func ParseProteins(path string) []*Protein {
	Proteins := []*Protein{}

	dat, err := os.Open(path)
	check(err)
	p := new(Protein)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "UNIQUE-ID") {
			p = new(Protein)
			p.ID = strings.TrimPrefix(scanner.Text(), "UNIQUE-ID - ")
		}
		if strings.HasPrefix(scanner.Text(), "TYPES") {
			p.Types = strings.TrimPrefix(scanner.Text(), "TYPES - ")
		}
		if strings.HasPrefix(scanner.Text(), "COMMON-NAME") {
			p.Name = append(p.Name, strings.TrimPrefix(scanner.Text(), "COMMON-NAME - "))
		}
		if strings.HasPrefix(scanner.Text(), "CATALYZES") {
			p.Catalyzes = append(p.Catalyzes, strings.TrimPrefix(scanner.Text(), "CATALYZES - "))
		}
		if strings.HasPrefix(scanner.Text(), "GENE") {
			p.Gene = append(p.Gene, strings.TrimPrefix(scanner.Text(), "GENE - "))
		}
		if strings.HasPrefix(scanner.Text(), "SYNONYMS") {
			p.Synonyms = append(p.Synonyms, strings.TrimPrefix(scanner.Text(), "SYNONYMS - "))
		}
		if strings.HasPrefix(scanner.Text(), "GO-TERMS") {
			p.GOTerms = append(p.GOTerms, strings.TrimPrefix(scanner.Text(), "GO-TERMS - "))
		}
		if strings.HasPrefix(scanner.Text(), "CITATIONS") {
			p.Citations = strings.TrimPrefix(scanner.Text(), "CITATIONS - ")
		}
		if scanner.Text() == "//" {
			Proteins = append(Proteins, p)
		}
	}
	return Proteins
}
