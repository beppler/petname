// http://manpages.ubuntu.com/manpages/bionic/man1/petname.1.html
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"github.com/dustinkirkland/golang-petname"
)

var (
	words = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	flag.Parse()
	fmt.Println(petname.Generate(*words, *separator))
}
