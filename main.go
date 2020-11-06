package parse

import (
	"bytes"
	"github.com/Haydn0230/parse/Link"
	"io/ioutil"
	"log"
)

func main() {
	htmlDoc, err := ioutil.ReadFile("./test-files/ex2.html")

	r := bytes.NewReader(htmlDoc)

	links, err := Link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(links)
}