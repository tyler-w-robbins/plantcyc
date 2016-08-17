package plantcyc

import (
	"bufio"
	"os"
	"strings"
)

type Class struct {
	ID       string
	Types    []string
	Comment  []string
	Synonyms []string
	Name     string
}

func ParseClasses(path string) []*Class {
	Classes := []*Class{}

	dat, err := os.Open(path)
	check(err)
	c := new(Class)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "UNIQUE-ID") {
			c = new(Class)
			c.ID = strings.TrimPrefix(scanner.Text(), "UNIQUE-ID - ")
		}
		if strings.HasPrefix(scanner.Text(), "TYPES") {
			c.Types = append(c.Types, strings.TrimPrefix(scanner.Text(), "TYPES - "))
		}
		if strings.HasPrefix(scanner.Text(), "COMMENT") {
			c.Comment = append(c.Comment, strings.TrimPrefix(scanner.Text(), "COMMENT - "))
		}
		if strings.HasPrefix(scanner.Text(), "/") && !strings.HasPrefix(scanner.Text(), "//") {
			c.Comment = append(c.Comment, strings.TrimPrefix(scanner.Text(), "/"))
		}
		if strings.HasPrefix(scanner.Text(), "SYNONYMS") {
			c.Synonyms = append(c.Synonyms, strings.TrimPrefix(scanner.Text(), "SYNONYMS - "))
		}
		if strings.HasPrefix(scanner.Text(), "COMMON-NAME") {
			c.Name = strings.TrimPrefix(scanner.Text(), "COMMON-NAME - ")
		}
		if scanner.Text() == "//" {
			Classes = append(Classes, c)
		}
	}
	return Classes
}
