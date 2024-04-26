package vcgen

import (
	"errors"
	"strings"
)

var (
	ErrNotFeasible       = errors.New("Not feasible to generate requested number of codes")
	ErrPatternIsNotMatch = errors.New("Pattern is not match with the length value")
)

const (
	// Charset types
	CharsetNumbers      = "0123456789"
	CharsetAlphabetic   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetAlphanumeric = CharsetNumbers + CharsetAlphabetic
)

const (
	// Default minimums
	minCount  uint16 = 1
	minLength uint16 = 6

	patternChar = "#"
)

type Generator struct {
	// Length of the code to generate
	Length uint16 `json:"length"`

	// Count of the codes
	// How many codes to generate
	Count uint16 `json:"count"`

	// Charset to use
	// `CharsetNumbers`, `CharsetAlphabetic`, and `CharsetAlphanumeric`
	// are already defined and you can use them.
	Charset string `json:"charset"`

	// Prefix of the code
	Prefix string `json:"prefix"`

	// Suffix of the code
	Suffix string `json:"suffix"`

	// Pattern of the code
	// # is the placeholder for the charset
	Pattern string `json:"pattern"`
}

// Creates a new generator with options
func NewWithOptions(opts ...Option) (*Generator, error) {
	g := Default()
	if err := setOptions(opts...)(g); err != nil {
		return nil, err
	}

	return g, nil
}

// Creates a new generator with default values
func Default() *Generator {
	return &Generator{
		Length:  minLength,
		Count:   minCount,
		Charset: CharsetAlphanumeric,
		Pattern: repeatStr(minLength, patternChar),
	}
}

// Generates a list of codes
func (g *Generator) Run() ([]string, error) {
	if !isFeasible(g.Charset, g.Pattern, patternChar, g.Count) {
		return nil, ErrNotFeasible
	}

	result := make([]string, g.Count)

	var i uint16
	for i = 0; i < g.Count; i++ {
		result[i] = g.one()
	}

	return result, nil
}

// one generates one code
func (g *Generator) one() string {
	pts := strings.Split(g.Pattern, "")
	for i, v := range pts {
		if v == patternChar {
			pts[i] = randomChar([]byte(g.Charset))
		}
	}

	suffix := g.Suffix

	return g.Prefix + strings.Join(pts, "") + suffix
}
