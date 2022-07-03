package vcgen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

var (
	// charset types
	numbers      = "0123456789"
	alphabetic   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric = numbers + alphabetic

	// defaults
	minCount  uint16 = 1
	minLength uint16 = 6
)

const patternChar = "#"

var (
	ErrNotFeasible       = errors.New("Not feasible to generate requested number of codes")
	ErrPatternIsNotMatch = errors.New("Pattern is not match with the length value")
)

// initialize random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

type Generator struct {
	// Length of the code
	Length uint16 `json:"length"`

	// Count of the codes
	Count uint16 `json:"count"`

	// Charset to use
	Charset string `json:"charset"`

	// Prefix of the code
	Prefix string `json:"prefix"`

	// Suffix of the code
	Suffix string `json:"suffix"`

	// Pattern of the code
	Pattern string `json:"pattern"`

	// Suffix of the code
	//
	// Deprecated: use Suffix instead
	Postfix string `json:"postfix"`
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
		Charset: alphanumeric,
		Pattern: repeatStr(minLength, patternChar),
	}
}

// New generator config
//
// Deprecated: use NewWithOptions or Default instead
func New(i *Generator) *Generator {

	// check functione entry args
	if i == nil {
		rps := repeatStr(minLength, "#")

		// --
		i = &Generator{
			Length:  minLength,
			Count:   minCount,
			Charset: alphanumeric,
			Pattern: rps,
		}
		goto AfterChecker
	}

	// check pattern
	if i.Pattern == "" {
		rps := repeatStr(i.Length, "#")
		i.Pattern = rps
	}

	// check length
	if i.Length == 0 {
		i.Length = uint16(strings.Count(i.Pattern, "#"))
	}

	// check count
	if i.Count == 0 {
		i.Count = minCount
	}

	// check charset
	if i.Charset == "" {
		i.Charset = alphanumeric
	}

AfterChecker:

	// return
	return i
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

	// TODO: Remove it when we deprecate Postfix
	if g.Postfix != "" {
		suffix = g.Postfix
	}

	return g.Prefix + strings.Join(pts, "") + suffix
}
