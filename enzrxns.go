package plantcyc

import (
	"bufio"
	"os"
	"strings"
)

type Enzrxn struct {
	ID        string
	Types     []string
	Name      string
	Citations []string
	Enzyme    string
	Reaction  string
}

func ParseEnzrxns(path string) []*Enzrxn {
	Enzrxns := []*Enzrxn{}

	dat, err := os.Open(path)
	check(err)
	er := new(Enzrxn)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "UNIQUE-ID") {
			er = new(Enzrxn)
			er.ID = strings.TrimPrefix(scanner.Text(), "UNIQUE-ID - ")
		}
		if strings.HasPrefix(scanner.Text(), "TYPES") {
			er.Types = append(er.Types, strings.TrimPrefix(scanner.Text(), "TYPES - "))
		}
		if strings.HasPrefix(scanner.Text(), "COMMON-NAME") {
			er.Name = strings.TrimPrefix(scanner.Text(), "COMMON-NAME - ")
		}
		if strings.HasPrefix(scanner.Text(), "CITATIONS") {
			er.Citations = append(er.Citations, strings.TrimPrefix(scanner.Text(), "CITATIONS - "))
		}
		if strings.HasPrefix(scanner.Text(), "ENZYME") {
			er.Enzyme = strings.TrimPrefix(scanner.Text(), "ENZYME - ")
		}
		if strings.HasPrefix(scanner.Text(), "REACTION") {
			er.Reaction = strings.TrimPrefix(scanner.Text(), "REACTION - ")
		}
		if scanner.Text() == "//" {
			Enzrxns = append(Enzrxns, er)
		}
	}
	return Enzrxns
}

func WritePathways(path string, w *bufio.Writer, e []*Enzrxn) error {
	for i := range e {
		_, err := w.WriteString("PCYC:" + e[i].ID + "|" + e[i].Name + "|PlantCyc_Pathways|Pathway\n")
		check(err)
	}
	return nil
}
