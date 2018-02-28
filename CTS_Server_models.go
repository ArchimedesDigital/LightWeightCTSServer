// TODO: move other structs to here
package main

import (
	"strings"
)

type Text struct {
	RawURN             string
	MapCitationPassage map[string]string
}

type URN struct {
	rawURN       string
	WorkFileName string
	PassageStart string
	PassageEnd   string
}

// construct a URN struct from rawURN string
func NewURN(s string) *URN {
	// initialize workID as rawURN
	var workFileName, passageStart, passageEnd string

	// check and split by ':'
	workSeparator := ":"
	if strings.ContainsAny(s, workSeparator+"|") {
		splitedURN := strings.Split(s, workSeparator)
		workFileName = splitedURN[0] // trim the passage part off
		passageStart = splitedURN[1] // now passage query exists
	} else {
		workFileName = s
	}

	workFileName = strings.Join([]string{workFileName, "xml"}, ".") // append file extension TODO: support other file types

	// check and split by '-'
	passageSeparator := "-"
	if (len(s) > 0) && (strings.ContainsAny(passageStart, passageSeparator+"|")) {
		splitedPassage := strings.Split(passageStart, passageSeparator)
		passageStart = splitedPassage[0] // trim the end part off
		passageEnd = splitedPassage[1]   // now passage end exists
	}

	return &URN{
		rawURN:       s,
		WorkFileName: workFileName,
		PassageStart: passageStart,
		PassageEnd:   passageEnd,
	}
}
