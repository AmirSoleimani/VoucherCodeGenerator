package vcgen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

//Charset types
var (
	numbers      = "0123456789"
	alphabetic   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// defaults
	minCount  uint16 = 1
	minLength uint16 = 6
)

//initialize package base config
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Generator request struct
type Generator struct {
	Length  uint16 `json:"length"` // ;) api...
	Count   uint16 `json:"count"`
	Charset string `json:"charset"`
	Prefix  string `json:"prefix"`
	Postfix string `json:"postfix"`
	Pattern string `json:"pattern"`
}

//New generator config
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

//Run voucher code generator
func (g *Generator) Run() (*[]string, error) {

	if !isFeasible(g.Charset, g.Pattern, g.Count) {
		return nil, errors.New("Not possible to generate requested number of codes")
	}

	result := make([]string, g.Count)

	var i uint16
	for i = 0; i < g.Count; i++ {
		result[i] = g.one()
	}

	return &result, nil
}

// generate one vouchers code
func (g *Generator) one() string {

	pts := strings.Split(g.Pattern, "")
	for i, v := range pts {
		if v == "#" {
			pts[i] = randomChar([]byte(g.Charset))
		}
	}

	return g.Prefix + strings.Join(pts, "") + g.Postfix

}
