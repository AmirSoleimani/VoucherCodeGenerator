package main

import (
	"fmt"
	"sync"

	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	// normal
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		vc, _ := vcgen.NewWithOptions(
			vcgen.SetCount(10),
			vcgen.SetPattern("###-###-###"),
			vcgen.SetCharset("0123456789"),
		)
		result, err := vc.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}(&wg)

	// with prefix
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		vc, _ := vcgen.NewWithOptions(
			vcgen.SetCount(10),
			vcgen.SetPattern("######"),
			vcgen.SetPrefix("WELC-"),
		)
		result, err := vc.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}(&wg)

	// with prefix + suffix
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		vc, _ := vcgen.NewWithOptions(
			vcgen.SetCount(10),
			vcgen.SetPattern("######"),
			vcgen.SetPrefix("WELC-"),
			vcgen.SetSuffix("-B"),
		)
		result, err := vc.Run()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}(&wg)

	wg.Wait()
}
