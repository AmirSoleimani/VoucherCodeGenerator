package vcgen

import (
	"errors"
)

var (
	ErrInvalidCount   = errors.New("invalid count, It should be greater than 0")
	ErrInvalidCharset = errors.New("invalid charset, charset length should be greater than 0")
	ErrInvalidPattern = errors.New("invalid pattern, pattern cannot be empty")
)

type Option func(*Generator) error

// SetLength sets the length of the code
func SetLength(length uint16) Option {
	return func(g *Generator) error {
		if length == 0 {
			length = numberOfChar(g.Pattern, patternChar)
		}
		g.Length = length
		return nil
	}
}

// SetCount sets the count of the code
func SetCount(count uint16) Option {
	return func(g *Generator) error {
		if count == 0 {
			return ErrInvalidCount
		}
		g.Count = count
		return nil
	}
}

// SetCharset sets the charset of the code
func SetCharset(charset string) Option {
	return func(g *Generator) error {
		if len(charset) == 0 {
			return ErrInvalidCharset
		}
		g.Charset = charset
		return nil
	}
}

// SetPrefix sets the prefix of the code
func SetPrefix(prefix string) Option {
	return func(g *Generator) error {
		g.Prefix = prefix
		return nil
	}
}

// SetSuffix sets the suffix of the code
func SetSuffix(suffix string) Option {
	return func(g *Generator) error {
		g.Suffix = suffix
		return nil
	}
}

// SetPattern sets the pattern of the code
func SetPattern(pattern string) Option {
	return func(g *Generator) error {
		if pattern == "" {
			return ErrInvalidPattern
		}

		numPatternChar := numberOfChar(pattern, patternChar)
		if g.Length == 0 || g.Length != numPatternChar {
			g.Length = numPatternChar
		}

		g.Pattern = pattern
		return nil
	}
}

func setOptions(opts ...Option) Option {
	return func(g *Generator) error {
		for _, opt := range opts {
			if err := opt(g); err != nil {
				return err
			}
		}
		return nil
	}
}
