package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/trobbi11/plantcyc"
)

// PlantCyc Data Parser
// By: Tyler Robbins
// Desciption:

var g []*plantcyc.Gene
var p []*plantcyc.Pathway

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteGenes(path string, w *bufio.Writer) error {
	// Parse gene nodes
	g = plantcyc.ParseGenes(path)
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

func main() {
	location := "/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/"

	node, err := os.Create("pCycNodeOut.csv")
	check(err)
	reln, err := os.Create("pCycRelnOut.csv")
	check(err)

	defer node.Close()
	defer reln.Close()

	wNode := bufio.NewWriter(node)
	// wReln := bufio.NewWriter(reln)

	// fmt.Println(reflect.TypeOf(wNode))

	// Write header
	_, err = wNode.WriteString("GeneID:ID|Synonyms:String[]|Description|Source|:Label\n")
	check(err)

	// File iterating
	err = filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "genes.col") {
			fmt.Println(path)
			err = WriteGenes(path, wNode)
			check(err)
		}

		err = wNode.Flush()
		check(err)
		return nil
	})
	check(err)

	// p = plantcyc.ParsePathways("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/pathways.col", g)
	// for i := range p {
	// 	fmt.Println(p[i])
	// }
	//
	// e := enzymes.ParseEnzymes("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/enzymes.col")
	//
	// for i := range e {
	// 	fmt.Println(e[i])
	// }
	//

	// er := enzrxns.ParseEnzrxns("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/enzrxns.dat")
	//
	// for i := range er {
	// 	fmt.Println(er[i])
	// }
	// pr := proteins.ParseProteins("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/proteins.dat")
	//
	// for i := range pr {
	// 	fmt.Println(pr[i])
	// }
	// c := classes.ParseClasses("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/classes.dat")
	// for i := range c {
	// 	fmt.Println(c[i])
	// }
	// cp := compounds.ParseCompounds("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/compounds.dat")
	//
	// for i := range cp {
	// 	fmt.Println(cp[i])
	// }
	// r := plantcyc.ParseReactions("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/reactions.dat")
	//
	// for i := range r {
	// 	fmt.Println(r[i])
	// }

	// Flush to ensure all buffered operations have been applied
	err = wNode.Flush()
	check(err)
}
