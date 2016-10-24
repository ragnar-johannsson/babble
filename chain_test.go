package babble

import (
	"strings"
	"testing"
)

func TestNewChain(t *testing.T) {
	sentence := strings.Split("This is not a drill!", " ")
	corpus := [][]string{sentence}

	chain := NewChain(corpus, DEFAULT_STATE_SIZE)

	if len(chain.model) != 6 {
		t.Fail()
	}
}

func TestInitState(t *testing.T) {
	sentence := strings.Split("This is not a drill!", " ")
	corpus := [][]string{sentence}

	chain := NewChain(corpus, 3)

	if _, ok := chain.model[strings.Join([]string{chainStart, chainStart, chainStart}, " ")]; !ok {
		t.Fail()
	}
}

func TestGenerate(t *testing.T) {
	sentence3 := strings.Split("This is not a drill!", " ")
	sentence2 := strings.Split("This is a test nonetheless.", " ")
	sentence1 := strings.Split("No matter how you look at it.", " ")
	sentence4 := strings.Split("But is this an excercise?", " ")
	corpus := [][]string{sentence1, sentence2, sentence3, sentence4}

	chain := NewChain(corpus, DEFAULT_STATE_SIZE)

	if chain.Generate() == "" {
		t.Fail()
	}

	chain = NewChain(corpus, 1)

	if chain.Generate() == "" {
		t.Fail()
	}

	chain = NewChain(corpus, 3)

	if chain.Generate() == "" {
		t.Fail()
	}
}
