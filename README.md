# Voucher Code Generator

Generate unique, random, and hard to guess coupon / voucher codes. Use cases: promo codes, loyalty coupons, gift vouchers, in-app purchases, referral links

### Instalation

Just use go get.
```go
    go get -u github.com/AmirSoleimani/VoucherCodeGenerator/vcgen
```

### Tested in the following Golang releases
All releases from Go1.13.x to Go1.18.x.

### How to use

```go

func main() {

	// lolipop :P
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
```

#### Options
```sh
SetLength(length uint16)
SetCount(count uint16)
SetCharset(charset string)
SetPrefix(prefix string)
SetSuffix(suffix string)
SetPattern(pattern string)
```

#### Prefix and Suffix

You can optionally surround each generated code with a prefix and/or suffix.

#### Pattern

Codes may follow a specified pattern. Use hash (`#`) as a placeholder for random characters. 

#### Infeasible configs

There exist some configs that are not feasible. For example it's not possible to generate 1000 codes if you want
your codes to be 2 characters long and consisting only of numbers. Voucher code generator detects such cases and
throws an error `"Not possible to generate requested number of codes."`.

#### Config reference

| attribute        | default value  | description                                                                     |
|------------------|:--------------:|---------------------------------------------------------------------------------|
| `length`         | `6`            | Number of characters in a generated code (excluding prefix and suffix)         |
| `count`          | `1`            | Number of codes generated.                                                      |
| `charset`        | `alphanumeric` | Characters that can appear in the code.                                         |
| `prefix`         | `""`           | A text appended before the code.                                                |
| `suffix`        | `""`           | A text appended after the code.                                                 |
| `pattern`        | `"######"`   | A pattern for codes where hashes (`#`) will be replaced with random characters. |


### License

Code released under the [MIT license](LICENSE).