// challenges/packages/begin/proverbs.go
package main

// import the proverbs package
import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() (proverb string) {
	proverb = proverbs.Random().Saying
	return
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
