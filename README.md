# Voucher Code Generator

Generate unique, random, and hard to guess coupon / voucher codes. Use cases: promo codes, loyalty coupons, gift vouchers, in-app purchases, referral links

### Instalation

Just use go get.
```go
    go get github.com/AmirSoleimani/VoucherCodeGenerator/vcgen
```

### Sample

```go

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
```

#### Prefix and Postfix

You can optionally surround each generated code with a prefix and/or postfix.

#### Pattern

Codes may follow a specified pattern. Use hash (`#`) as a placeholder for random characters. 

#### Infeasible configs

There exist some configs that are not feasible. For example it's not possible to generate 1000 codes if you want
your codes to be 2 characters long and consisting only of numbers. Voucher code generator detects such cases and
throws an error `"Not possible to generate requested number of codes."`.

#### Config reference

| attribute        | default value  | description                                                                     |
|------------------|:--------------:|---------------------------------------------------------------------------------|
| `length`         | `6`            | Number of characters in a generated code (excluding prefix and postfix)         |
| `count`          | `1`            | Number of codes generated.                                                      |
| `charset`        | `alphanumeric` | Characters that can appear in the code.                                         |
| `prefix`         | `""`           | A text appended before the code.                                                |
| `postfix`        | `""`           | A text appended after the code.                                                 |
| `pattern`        | `"######"`   | A pattern for codes where hashes (`#`) will be replaced with random characters. |


### License

Code released under the [MIT license](LICENSE).