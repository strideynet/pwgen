// main provides a brief example of using the pwgen library
package main

import (
	"flag"
	"github.com/strideynet/pwgen"
	"log"
	"os"
)

var (
	length    = flag.Int("length", -1, "total number of characters in password")
	lowercase = flag.Int("lowercase", -1, "minimum number of lowercase characters in password")
	uppercase = flag.Int("uppercase", -1, "minimum number of uppercase characters in password")
	special   = flag.Int("special", -1, "minimum number of special characters in password")
	number    = flag.Int("number", -1, "minimum number of number characters in password")
)

func main() {
	flag.Parse()
	log.Printf("Generating a password!")

	var opts []pwgen.Option
	if *length != -1 {
		opts = append(opts, pwgen.Length(*length))
	}
	if *lowercase != -1 {
		opts = append(opts, pwgen.Lowercase(*lowercase))
	}
	if *uppercase != -1 {
		opts = append(opts, pwgen.Uppercase(*uppercase))
	}
	if *special != -1 {
		opts = append(opts, pwgen.Special(*special))
	}
	if *number != -1 {
		opts = append(opts, pwgen.Number(*number))
	}

	str, err := pwgen.Generate(opts...)
	if err != nil {
		log.Printf("By golly, an error occured: %s", err)
		os.Exit(1)
	}

	log.Printf("The generated password is: %s", str)
}
