# pwgen

[![Go Reference](https://pkg.go.dev/badge/github.com/strideynet/pwgen.svg)](https://pkg.go.dev/github.com/strideynet/pwgen)

Go password generation library (codetest).

## CLI usage

You can even use pwgen as a cli tool! All options can be configured via commandline opts.

```shell
go run ./cmd/pwgen # Generates standard 8 character password
go run ./cmd/pwgen -h # Outputs the help text detailing the possible parameters e.g
go run ./cmd/pwgen -length 32
```

## API Usage

pwgen exposes a simple functional option API. This means you can pass no options in, or as many as you want, to configure pwgen to behave the way you are expecting.

Simplest version:
```go
str, err := pwgen.Generate() // Generates a standard password with the default settings.

// You need to check the error!
if err != nil {
	log.Printf("Oh no!") // You'll want some nicer error handling here
}

log.Printf("Your password is %s", str)

pwgen.Generate(pwgen.WithLength(32), pwgen.WithLowercaseCount(12)) // Creates a password of length 32, with at least 12 lowercase characters
```

Checkout the docs at https://pkg.go.dev/github.com/strideynet/pwgen

## Notes

### Testing

I tend to use ``testify/assert`` for most projects, but have stuck with standard library tooling here.
