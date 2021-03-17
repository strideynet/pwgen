package main

import (
	"github.com/strideynet/pwgen"
	"log"
)

func main() {
	str, err := pwgen.Generate(pwgen.Number(12), pwgen.Length(18))
	if err != nil {
		log.Panicf("By golly, an error occured: %s", err)
	}

	log.Printf("The generated password is: %s", str)
}
