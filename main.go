package main

import (
	"fmt"
	"sync"
	"voucher-code-generator/vcgen"
)

func main() {

	// lolipop :P
	wg := sync.WaitGroup{}
	wg.Add(3)

	// normal
	go func(wg *sync.WaitGroup) {
		vc := vcgen.New(&vcgen.Generator{Count: 10, Pattern: "###-###-###", Charset: "123456789"})
		result, err := vc.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		wg.Done()
	}(&wg)

	// with prefix
	go func(wg *sync.WaitGroup) {
		vcPrefix := vcgen.New(&vcgen.Generator{Count: 10, Pattern: "######", Prefix: "WELC-"})
		result, err := vcPrefix.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		wg.Done()
	}(&wg)

	// with prefix + postfix
	go func(wg *sync.WaitGroup) {
		vcPrePostfix := vcgen.New(&vcgen.Generator{Count: 10, Pattern: "######", Prefix: "WELC-", Postfix: "-B"})
		result, err := vcPrePostfix.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		wg.Done()
	}(&wg)

	wg.Wait()

}
