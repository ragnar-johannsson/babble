package babble

import (
	"strings"
)

type Chain struct {
	model     map[string]map[string]int
	stateSize int
}

func (c *Chain) move(state string) string {
	return aggregateLinks(c.model[state]).Next()
}

func (c *Chain) initState() []string {
	state := []string{}

	for i := 0; i < c.stateSize; i++ {
		state = append(state, chainStart)
	}

	return state
}

func (c *Chain) Generate() string {
	sentence := []string{}
	state := strings.Join(c.initState(), " ")

	for {
		next := c.move(state)

		if next == chainEnd {
			break
		}

		sentence = append(sentence, next)
		state = strings.Join(append(strings.Split(state, " "), next)[1:], " ")
	}

	return strings.Join(sentence, " ")
}

func (c *Chain) build(corpus []([]string)) {
	c.model = map[string]map[string]int{}

	for k := range corpus {
		run := corpus[k]
		items := []string{}

		items = append(items, c.initState()...)
		items = append(items, run...)
		items = append(items, chainEnd)

		for i := 0; i < len(run)+1; i++ {
			state := strings.Join(items[i:i+c.stateSize], " ")
			follow := items[i+c.stateSize]

			if _, exists := c.model[state]; !exists {
				c.model[state] = map[string]int{}
			}

			if _, exists := c.model[state][follow]; !exists {
				c.model[state][follow] = 0
			}

			c.model[state][follow]++
		}

	}
}

func NewChain(corpus []([]string), stateSize int) Chain {
	chain := Chain{
		stateSize: stateSize,
	}

	chain.build(corpus)

	return chain

}
