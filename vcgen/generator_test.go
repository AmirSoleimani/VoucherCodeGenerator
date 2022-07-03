package vcgen_test

import (
	"testing"

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
)

func TestGenerator(t *testing.T) {
	tables := []struct {
		options     []vcgen.Option
		want        error
		outputCount int
	}{
		{
			options:     nil,
			want:        nil,
			outputCount: 1,
		},
		{
			options: []vcgen.Option{
				vcgen.SetCount(0),
			},
			want: vcgen.ErrInvalidCount,
		},
		{
			options: []vcgen.Option{
				vcgen.SetPattern("##"),
				vcgen.SetCharset("AB"),
				vcgen.SetCount(1000),
			},
			want: vcgen.ErrNotFeasible,
		},
		{
			options: []vcgen.Option{
				vcgen.SetPattern("##"),
				vcgen.SetCharset("AB"),
				vcgen.SetCount(2),
			},
			want:        nil,
			outputCount: 2,
		},
		{
			options: []vcgen.Option{
				vcgen.SetPattern("###-###-###"),
				vcgen.SetCharset("123456789ABCDabfg"),
				vcgen.SetCount(50),
			},
			want:        nil,
			outputCount: 50,
		},
	}

	for _, b := range tables {
		g, err := vcgen.NewWithOptions(b.options...)
		if err != nil {
			if err != b.want {
				t.Error(err)
			}
			continue
		}

		codes, err := g.Run()
		if err != nil {
			if err != b.want {
				t.Error(err)
			}
			continue
		}

		if len(codes) != b.outputCount {
			t.Errorf("expected %d codes, got %d", b.outputCount, len(codes))
		}
	}
}
