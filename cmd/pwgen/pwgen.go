// main provides a brief example of using the pwgen library
package main

import (
	"github.com/strideynet/pwgen"
	"log"
	"os"
)

func main() {
	log.Printf("Generating a password!")

	str, err := pwgen.Generate(pwgen.Number(12), pwgen.Length(2))
	if err != nil {
		log.Printf("By golly, an error occured: %s", err)
		os.Exit(1)
	}

	log.Printf("The generated password is: %s", str)
}
