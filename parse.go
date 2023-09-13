package conntrack

import (
	"fmt"
	"regexp"
	"strings"
)

type Row struct {
	isrc, isport, idst, idport string
	esrc, esport, edst, edport string
}

func (r Row) String() string {
	return fmt.Sprintf("%s:%s-%s:%s<->%s:%s-%s:%s", r.isrc, r.isport, r.idst, r.idport, r.esrc, r.esport, r.edst, r.edport)
}

func Parse(input string) []Row {
	var out = []Row{}
	var pattern = "src=([0-9a-f:\\.]+)\\s+dst=([0-9a-f:\\.]+)\\s+sport=([0-9]+)\\s+dport=([0-9]+)\\s+src=([0-9a-f:\\.]+)\\s+dst=([0-9a-f:\\.]+)\\s+sport=([0-9]+)\\s+dport=([0-9]+)\\s+"
	var match = regexp.MustCompile(pattern)
	for _, line := range strings.SplitN(input, "\n", -1) {
		var matches = match.FindStringSubmatch(line)
		if len(matches) < 9 {
			continue
		}
		var row = Row{
			isrc:   matches[1],
			idst:   matches[2],
			isport: matches[3],
			idport: matches[4],
			esrc:   matches[5],
			edst:   matches[6],
			esport: matches[7],
			edport: matches[8],
		}
		out = append(out, row)
	}
	return out
}
