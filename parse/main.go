package main

import (
	"bufio"
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
var e []*plantcyc.Enzyme
var c []*plantcyc.Compound
var pr []*plantcyc.Protein
var er []*plantcyc.Enzrxn
var r []*plantcyc.Reaction

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	location := "plantcyc/tier1-tier2-flatfiles/"

	geneNode, err := os.Create("pCycGeneNodeOut.csv")
	check(err)
	pathNode, err := os.Create("pCycPathNodeOut.csv")
	check(err)
	// protNode, err := os.Create("pCycProtNodeOut.csv")
	// check(err)
	reln, err := os.Create("pCycRelnOut.csv")
	check(err)

	defer geneNode.Close()
	defer pathNode.Close()
	defer reln.Close()

	wGeneNode := bufio.NewWriter(geneNode)
	wPathNode := bufio.NewWriter(pathNode)
	// wProtNode := bufio.NewWriter(protNode)
	// wReln := bufio.NewWriter(reln)

	// Write headers
	_, err = geneNode.WriteString("GeneID:ID|Synonyms:String[]|Description|Source|:Label\n")
	check(err)
	_, err = pathNode.WriteString("Source_ID:ID|Name|Source|:LABEL\n")
	check(err)

	// Iterate through files, parse different node types and write to files
	err = filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "genes.col") {
			g = plantcyc.ParseGenes(path)
			err = plantcyc.WriteGenes(path, wGeneNode, g)
			check(err)
		} else if strings.HasSuffix(path, "pathways.col") {
			p = plantcyc.ParsePathways(path)
			err = plantcyc.WritePathways(path, wPathNode, p)
			check(err)
		} else if strings.HasSuffix(path, "enzymes.col") {
			e = plantcyc.ParseEnzymes(path)
			err = plantcyc.WriteEnzymes(path, wPathNode, e)
			check(err)
		} else if strings.HasSuffix(path, "compounds.dat") {
			c = plantcyc.ParseCompounds(path)
			// err = plantcyc.WriteCompounds(path, wPathNode, c)
			check(err)
		} else if strings.HasSuffix(path, "proteins.dat") {
			pr = plantcyc.ParseProteins(path)
			// err = plantcyc.WriteProteins(path, wProtNode, pr)
			check(err)
		} else if strings.HasSuffix(path, "enzrxns.dat") {
			er = plantcyc.ParseEnzrxns(path)
			// err = plantcyc.WriteEnzrxns(path, wPathNode, er)
			check(err)
		} else if strings.HasSuffix(path, "reactions.dat") {
			r = plantcyc.ParseReactions(path)
			// err = plantcyc.WriteReactions(path, wPathNode, er)
			check(err)
		}
		return nil
	})
	check(err)

	// Flush to ensure all buffered operations have been applied
	err = wGeneNode.Flush()
	check(err)
	err = wPathNode.Flush()
	check(err)
}
