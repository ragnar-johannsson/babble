package babble

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type link struct {
	score int
	value string
}

type links []link

func (n links) Next() string {
	rnd := rand.Intn(n.SumScore()) + 1

	for i := range n {
		if rnd <= n[i].score {
			return n[i].value
		}
	}

	return ""
}

func (n links) SumScore() int {
	return n[len(n)-1].score
}

func aggregateLinks(c map[string]int) links {
	acc := links{}
	total := 0

	for key := range c {
		total += c[key]

		acc = append(acc, link{
			score: total,
			value: key,
		})
	}

	return acc
}
