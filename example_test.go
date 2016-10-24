package babble_test

import (
	"fmt"
	"io/ioutil"

	"github.com/ragnar-johannsson/babble"
)

func ExampleTextModel_MakeSentence() {
	b, _ := ioutil.ReadFile("./texts/pg5827.txt")
	t := babble.NewTextModel(string(b), babble.DEFAULT_STATE_SIZE)

	fmt.Println(t.MakeSentence())
}
