## Babble [![GoDoc](https://godoc.org/github.com/ragnar-johannsson/babble?status.png)](https://godoc.org/github.com/ragnar-johannsson/babble)

Package babble implements a finite Markov chain sentence generator, inspired by Jeremy Singer-Vine's excellent [markovify](https://github.com/jsvine/markovify) for Python. Eric Bower's [sentences](https://github.com/neurosnap/sentences) package is used for sentence extraction.


### Usage

```go
package main

import (
        "fmt"
        "io/ioutil"

        "github.com/ragnar-johannsson/babble"
)

func main() {
        // The Problems of Philosophy by Bertrand Russell
        b, _ := ioutil.ReadFile("./texts/pg5827.txt")
        t := babble.NewTextModel(string(b), babble.DEFAULT_STATE_SIZE)

        fmt.Println(t.MakeShortSentence(140))
        // Our nature is the characteristic of the universe, we have discovered so far.
}
```

### License

BSD 2-Clause. See the LICENSE file for details.
