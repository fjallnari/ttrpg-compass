package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var DATA [30]string = [30]string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "Mauris", "faucibus", "lectus", "eget", "cursus", "tempus", "ligula", "orci", "mattis", "massa", "nec", "eleifend", "lorem", "ipsum", "congue", "erat", "Pellentesque", "suscipit", "semper", "sapien", "sed", "luctus"}

func LoremWords(quantity int) string {

	lorem := ""

	for i := 0; i < quantity; i++ {
		lorem += DATA[rand.Intn(30)] + " "
	}

	return lorem
}

func Sentence(words int) string {
	text := LoremWords(words)
	return strings.TrimSpace(strings.ToUpper(text[0:1])+text[1:]) + "."
}

func createMockStruct() *TTRPGSystem {
	system := TTRPGSystem{Title: "lorem"}

	return &system
}

func genMockData() {
	system := createMockStruct()

	f, err := os.Create(fmt.Sprintf("%s.toml", system.Title))

	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
	}
	if err := toml.NewEncoder(f).Encode(system); err != nil {
		// failed to encode
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		// failed to close the file
		log.Fatal(err)

	}

}
