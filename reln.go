package plantcyc

import "bufio"

func WriteEnzymePaths(w *bufio.Writer, e []*Enzyme, p []*Pathway) {
	for j := range p {
		for i := range e {
			for k := range e[i].Path {
				if e[i].Path[k] == p[j].ID {
					_, err := w.WriteString("PCYC:" + rmchars(e[i].ID, "*_-") + "|Chemicals|PCYC:" + rmchars(p[j].ID, "*_-") + "|is_part_of\n")
					check(err)
				}
			}
		}
	}
}

func WritePathGenes(w *bufio.Writer, p []*Pathway, g []*Gene) {
	for i := range p {
		// fmt.Println(p[i])
		for j := range p[i].GeneID {
			// fmt.Println("here")
			for k := range g {
				if p[i].GeneID[j] == g[k].ID {
					// fmt.Println("here")
					_, err := w.WriteString("PCYC:" + rmchars(p[i].ID, "*_-") + "|Pathways|PCYC:" + rmchars(g[k].ID, "*_-") + "|is_part_of\n")
					check(err)
				}
			}
		}
	}
}

func WriteProteinEnzrxns(w *bufio.Writer, pr []*Protein, er []*Enzrxn) {
	for i := range pr {
		for j := range pr[i].Catalyzes {
			for k := range er {
				if pr[i].Catalyzes[j] == er[k].ID {
					_, err := w.WriteString("PCYC:" + rmchars(pr[i].ID, "*_-") + "|Proteins|PCYC:" + rmchars(er[k].ID, "*_-") + "|catalyzes\n")
					check(err)
				}
			}
		}
	}
}
