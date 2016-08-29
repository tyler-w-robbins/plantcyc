package plantcyc

import "bufio"

func WriteEnzymePaths(w *bufio.Writer, e []*Enzyme, p []*Pathway) {
	for i := range e {
		for j := range p {
			for k := range e[i].Path {
				if e[i].Path[k] == p[j].ID {
					_, err := w.WriteString("PCYC:" + rmchars(e[i].ID, "-") + "|Chemicals|PCYC:" + rmchars(p[j].ID, "-") + "|is_part_of\n")
					check(err)
				}
			}
		}
	}
}
