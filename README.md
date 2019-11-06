# validator

Golang library for chaining data validations on values.

This allows you to "chain" any number of validation logic and have them apply to your value in one call.

Features:
  * Flexibility to define validators dynamically and at run-time.
  * Ability to "save" the chain and pass it around where needed.
  * Re-usability of your "saved" chain.
  * Late-binding of the value to validate; the chain is not tied to a particular value.

## Usage

### Installation

Get the package.
```sh
go get github.com/jmong/validator
```

Import it into your Golang code.
```go
import "github.com/jmong/validator"
```

### How To Use

Here is a simple example. 

In the following code block, it creates a validation chain that checks that the text is at most 30 characters, all alphabetical characters are upper-case, and contains the characters "AM".<br> 
All conditions must be true in order for the text to be validated.

```go
text := "I AM SPARTACUS!"
isvalidated := validator.BuildStrChain().IsMaxLen(30).IsUpper().IsContains("AM").ValidateStr(text)
if isvalidated == true {
    // do something
}
```

Here is an example of dynamically building your chain at run-time.
```go
chain := validator.BuildStrChain()
if checkMaxLen == true {
    chain = chain.IsMaxLen(30)
}
if checkUpper == true {
    chain = chain.IsUpper()
}
// ... and so on ...

if chain.ValidateStr(text) == true {
    // do something
}
```

You can also build the chain and delay triggering the validation till later.
For example, 

```go
chains := validator.BuildStrChain().IsMaxLen(5).IsUpper()
for _, text := range []string{"ABC", "abc", "aBc", "ABCDEFGH"} {
    if chains.ValidateStr(text) == true {
        // do something
    }
}
```

This gives you flexibility to "save" a chain and pass it around where as needed.
For example,
```go
choiceA := validator.BuildIntChain().IsInRange(1, 5)
choiceB := validator.BuildIntChain().IsInRange(6, 10)

choices := []struct{
    value   int
    checks  validator.IntChainer
}{
    {value: 2, checks: choiceA},
    {value: 8, checks: choiceB},
}
```

## Coding

You first need to call one of the `Build*Chain()` functions to initialize an empty validation chain. There are various `Build*Chain()` functions that correspond to the value type you are validating.<br>
For example, `BuildStrChain()` creates a chain for validating string values.
```go
// Initialize an empty chain for validating string values.
chain := validator.BuildStrChain()
```

Once a chain is created, you attach any number of validators (applicable for that value type) to it.<br>
For example, if you want to validate that your string value is all upper-case and has at most 30 characters, you add `IsUpper()` and `IsMaxLen()` methods to `BuildStrChain()`.
```go
// Initialize and register these validations to the chain.
chain := validator.BuildStrChain().IsUpper().IsMaxLen(30)
```

Finally when you are satisfied that you created the final chain that contains all of the validators you want, you can run the actual validation on your value with one of the `Validate*()` methods.

Similarly there are various `Validate*()` methods that correspond to the value type you are validating. For example, `ValidateStr()` runs the validations on a string value. It returns true if the value passed _all_ validations; otherwise, it returns false.
```go
isvalidated := validator.BuildStrChain().IsUpper().IsMaxLen(30).ValidateStr("I AM SPARTACUS!")
if isvalidated == true {
    // do something
}
```

## License

See [LICENSE](LICENSE)

## Author

Joe Mong
* Github - [https://github.com/jmong](https://github.com/jmong)
