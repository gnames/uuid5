// Package uuid5 creates UUID v5 in DNS globalnames.org namespace
package uuid5

import (
	"strings"

	"github.com/satori/go.uuid"
)

var gnNamespace = uuid.NewV5(uuid.NamespaceDNS, "globalnames.org")

// UUID5 returns UUID version 5 generated from a string
func UUID5(input string) uuid.UUID {
	return uuid.NewV5(gnNamespace, normalizeName(input))
}

// UUID5s returns array of UUIDs v5 generated from an array of strings
func UUID5s(input []string) []uuid.UUID {
	uuids := make([]uuid.UUID, len(input))
	for i := range input {
		uuids[i] = UUID5(input[i])
	}
	return uuids
}

// String converts one UUID5 to its string representation
func String(uuid uuid.UUID) string {
	return uuid.String()
}

// Strings converts array of UUIDs to their string representation
func Strings(uuids []uuid.UUID) []string {
	uuidStrings := make([]string, len(uuids))
	for i := range uuids {
		uuidStrings[i] = uuids[i].String()
	}
	return uuidStrings
}

// PipeDelimited takes names separated by | and returns their UUID5
func PipeDelimited(pipeDelim string) []uuid.UUID {
	return delimited(pipeDelim, 0x007C) // = '|'
}

// NewLineDelimited takes names separated by \n (line feed) and returns
// their UUID5
func NewLineDelimited(newLineDelim string) []uuid.UUID {
	return delimited(newLineDelim, 0x000A) // = "\n"
}

func delimited(delimited string, delim rune) []uuid.UUID {
	names := splitNames(delimited, delim)
	uuids := make([]uuid.UUID, len(names))
	for i := range names {
		uuids[i] = UUID5(names[i])
	}
	return uuids
}

func splitNames(input string, delim rune) []string {
	f := func(ch rune) bool {
		return ch == delim
	}
	return strings.FieldsFunc(input, f)
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}
