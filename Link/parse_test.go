package Link

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"io/ioutil"
	"testing"
)

func TestHTMLParser(t *testing.T) {
	tt := []struct{
		name string
		file string
		expected []link
	}{
		{
			name: "checks a link struct is returned,",
			file: "../test-files/ex1.html",
			expected:[]link{
				{
					Href:"/other-page",
					Text:"A link to another page",
				},
			},
		},
		{
			name: "checks a link struct is returned,",
			file: "../test-files/ex2.html",
			expected:[]link{
				{
					Href:"https://www.twitter.com/joncalhoun",
					Text:"Check me out on twitter",
				},
				{
					Href:"https://github.com/gophercises",
					Text:"Gophercises is on Github!",
				},
			},
		},
		{
			name: "checks a link struct is returned for test case 3 ",
			file: "../test-files/ex3.html",
			expected:[]link{
				{
					Href:"#",
					Text:"Login",
				},
				{
					Href:"/lost",
					Text:"Lost? Need help?",
				},
				{
					Href:"https://twitter.com/marcusolsson",
					Text:"@marcusolsson",
				},
			},
		},
		{
			name: "checks a link struct is returned,",
			file: "../test-files/ex4.html",
			expected:[]link{
				{
					Href:"/dog-cat",
					Text:"dog cat",
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T){
			xb, err := ioutil.ReadFile(tc.file)
			errCheck(err)

			r := bytes.NewReader(xb)

			xl, err := Parse(r)
			errCheck(err)
			fmt.Printf("\nexpected\n actual\n")
			assert.Equal(t, tc.expected, xl)

		})
	}


}

func TestGrabText( t *testing.T) {
	testData := &html.Node{
		FirstChild: &html.Node{
			NextSibling: &html.Node{
				NextSibling: &html.Node{
					NextSibling: nil,
					Type:        html.TextNode,
					Data:        "!",
				},
				FirstChild: &html.Node{
					NextSibling: nil,
					Type:        html.TextNode,
					Data:        "Github",
				},
				Type: html.ElementNode,
				Data: "strong",
			},
		Type: html.TextNode,
		Data: "Gophercise is on ",
	},
	Type:html.ElementNode,
	}


	result := parseText(testData)

	assert.Equal(t,  "Gophercise is on Github!",result)
}

