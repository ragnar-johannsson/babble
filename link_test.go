package babble

import (
	"testing"
)

func TestAggregateLinks(t *testing.T) {
	data := map[string]int{
		"one": 1,
		"two": 2,
		"six": 6,
	}
	tested := aggregateLinks(data)

	if tested.SumScore() != 9 {
		t.Fail()
	}

	data = map[string]int{
		"1": 100,
		"2": 0,
		"3": 0,
	}
	tested = aggregateLinks(data)

	if tested.Next() != "1" {
		t.Fail()
	}

	data = map[string]int{
		"1": 0,
		"2": 100,
		"3": 0,
	}
	tested = aggregateLinks(data)

	if tested.Next() != "2" {
		t.Fail()
	}

	data = map[string]int{
		"1": 0,
		"2": 0,
		"3": 100,
	}
	tested = aggregateLinks(data)

	if tested.Next() != "3" {
		t.Fail()
	}
}

func TestLinksNext(t *testing.T) {

	data := map[string]int{
		"1": 5,
		"2": 1,
		"3": 3,
		"4": 1,
	}
	testSet := aggregateLinks(data)

	count := map[string]int{
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
	}

	for i := 0; i < 10000; i++ {
		count[testSet.Next()]++
	}

	switch {
	case !(4850 < count["1"] && count["1"] < 5150):
		t.Fatal(count["1"])
	case !(850 < count["2"] && count["2"] < 1150):
		t.Fatal(count["2"])
	case !(2850 < count["3"] && count["3"] < 3150):
		t.Fatal(count["3"])
	case !(850 < count["4"] && count["4"] < 1150):
		t.Fatal(count["4"])
	}
}
