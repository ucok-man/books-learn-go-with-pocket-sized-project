package main

import "github.com/ucok-man/05-gordle/internal/gordle"

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("internal/corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(maxAttempts))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
