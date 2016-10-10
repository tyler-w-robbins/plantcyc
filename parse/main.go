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
	chemNode, err := os.Create("pCycChemNodeOut.csv")
	check(err)
	protNode, err := os.Create("pCycProtNodeOut.csv")
	check(err)
	reln, err := os.Create("pCycRelnOut.csv")
	check(err)

	defer geneNode.Close()
	defer pathNode.Close()
	defer chemNode.Close()
	defer protNode.Close()
	defer reln.Close()

	wGeneNode := bufio.NewWriter(geneNode)
	wPathNode := bufio.NewWriter(pathNode)
	wChemNode := bufio.NewWriter(chemNode)
	wProtNode := bufio.NewWriter(protNode)
	wReln := bufio.NewWriter(reln)

	// Write headers
	_, err = wGeneNode.WriteString("GeneID:ID|Synonyms:String[]|Description|Source|:LABEL\n")
	check(err)
	_, err = wPathNode.WriteString("Source_ID:ID|Name|Source|:LABEL\n")
	check(err)
	_, err = wChemNode.WriteString("Source_ID:ID|Name|Source|Definition|Synonyms:string[]|:LABEL\n")
	check(err)
	_, err = wProtNode.WriteString("Source_ID:ID|Name|Source|Function|Diseases|Synonyms:string[]|KEGG_Pathway|Wiki_Pathway|:LABEL\n")
	check(err)
	_, err = wReln.WriteString(":START_ID|Source|:END_ID|:TYPE\n")
	check(err)

	// Iterate through files, parse different node types and write to files
	err = filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "genes.col") {
			g = plantcyc.ParseGenes(path, g)
			plantcyc.WriteGenes(wGeneNode, g)
		} else if strings.HasSuffix(path, "pathways.col") {
			p = plantcyc.ParsePathways(path, p)
			plantcyc.WritePathways(wPathNode, p)
		} else if strings.HasSuffix(path, "enzymes.col") {
			e = plantcyc.ParseEnzymes(path, e)
			plantcyc.WriteEnzymes(wPathNode, e)
		} else if strings.HasSuffix(path, "compounds.dat") {
			c = plantcyc.ParseCompounds(path, c)
			plantcyc.WriteCompounds(wChemNode, c)
		} else if strings.HasSuffix(path, "proteins.dat") {
			pr = plantcyc.ParseProteins(path, pr)
			plantcyc.WriteProteins(wProtNode, pr)
		} else if strings.HasSuffix(path, "enzrxns.dat") {
			er = plantcyc.ParseEnzrxns(path, er)
			plantcyc.WriteEnzrxns(wPathNode, er)
		} else if strings.HasSuffix(path, "reactions.dat") {
			r = plantcyc.ParseReactions(path, r)
			plantcyc.WriteReactions(wPathNode, r)
		}
		return nil
	})
	//
	// plantcyc.WriteEnzymePaths(wReln, e, p)
	// plantcyc.WritePathGenes(wReln, p, g)
	// plantcyc.WriteProteinEnzrxns(wReln, pr, er)
	// plantcyc.WriteCompoundChebi(wReln, c)

	// Flush to ensure all buffered operations have been applied
	err = wGeneNode.Flush()
	check(err)
	err = wPathNode.Flush()
	check(err)
	err = wChemNode.Flush()
	check(err)
	err = wProtNode.Flush()
	check(err)
	err = wReln.Flush()
	check(err)
}
