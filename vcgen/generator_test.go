package vcgen_test

import (
	"errors"
	"reflect"
	"testing"
	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
)

func TestGenerator(t *testing.T) {

	tables := []struct {
		value *vcgen.Generator
		want  error
	}{
		{value: &vcgen.Generator{}, want: nil},
		{value: &vcgen.Generator{Pattern: "##", Charset: "AB", Count: 1000}, want: errors.New("")},
		{value: &vcgen.Generator{Pattern: "##", Charset: "AB", Count: 2}, want: nil},
		{value: &vcgen.Generator{Pattern: "###-###-###", Charset: "123456789ABCDabfg", Count: 50}, want: nil},
	}

	for _, b := range tables {
		OkStep := vcgen.New(b.value)
		_, err := OkStep.Run()
		if reflect.TypeOf(err) != reflect.TypeOf(b.want) {
			t.Error(err)
		}
	}

}
