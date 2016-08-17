package plantcyc

import (
	"bufio"
	"os"
	"strings"
)

type Reaction struct {
	ID        string
	Types     []string
	Name      string
	Citations []string
	Sources   []string
	DBLinks   []string
	InPathway []string
	RxnLocs   []string
	EnzRxns   []string
	Synonyms  []string
}

func ParseReactions(path string) []*Reaction {
	Reactions := []*Reaction{}

	dat, err := os.Open(path)
	check(err)
	er := new(Reaction)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "UNIQUE-ID") {
			er = new(Reaction)
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
		if strings.HasPrefix(scanner.Text(), "SOURCES") {
			er.Sources = append(er.Sources, strings.TrimPrefix(scanner.Text(), "SOURCES - "))
		}
		if strings.HasPrefix(scanner.Text(), "DBLINKS") {
			er.DBLinks = append(er.DBLinks, strings.TrimPrefix(scanner.Text(), "DBLINKS - "))
		}
		if strings.HasPrefix(scanner.Text(), "IN-PATHWAY") {
			er.InPathway = append(er.InPathway, strings.TrimPrefix(scanner.Text(), "IN-PATHWAY - "))
		}
		if strings.HasPrefix(scanner.Text(), "RXN-LOCATIONS") {
			er.RxnLocs = append(er.RxnLocs, strings.TrimPrefix(scanner.Text(), "RXN-LOCATIONS - "))
		}
		if strings.HasPrefix(scanner.Text(), "ENZYMATIC-REACTION") {
			er.EnzRxns = append(er.EnzRxns, strings.TrimPrefix(scanner.Text(), "ENZYMATIC-REACTION - "))
		}
		if strings.HasPrefix(scanner.Text(), "SYNONYMS") {
			er.Synonyms = append(er.Synonyms, strings.TrimPrefix(scanner.Text(), "SYNONYMS - "))
		}
		if scanner.Text() == "//" {
			Reactions = append(Reactions, er)
		}
	}
	return Reactions
}
