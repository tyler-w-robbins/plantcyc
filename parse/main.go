package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/trobbi11/plantcyc"
)

// PlantCyc Data Parser
// By: Tyler Robbins
// Desciption:
func printFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if strings.HasSuffix(path, "pathways.dat") {
		fmt.Println(path)
	}
	return nil
}
func main() {
	// g := plantcyc.ParseGenes("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/genes.col")
	//
	// for i := range g {
	// }
	//
	// e := enzymes.ParseEnzymes("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/enzymes.col")
	//
	// for i := range e {
	// 	fmt.Println(e[i])
	// }
	//
	// p := pathways.ParsePathways("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/pathways.col")
	// for i := range p {
	// 	fmt.Println(p[i])
	// }
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
	r := plantcyc.ParseReactions("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/10403s_rastcyc/reactions.dat")

	for i := range r {
		fmt.Println(r[i])
	}
	err := filepath.Walk("/Users/trobbi11/plantcyc/tier1-tier2-flatfiles/", printFile)
	if err != nil {
		log.Fatal(err)
	}
}
