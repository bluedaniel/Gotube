package utils

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

// Fetch is a fetch helper
func Fetch(url string) *http.Response {
	// fmt.Println(url)
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	return r
}

// FetchHTML returns a parsed HTML object
func FetchHTML(url string) *html.Node {
	r := Fetch(url)
	rb, _ := html.Parse(r.Body)
	defer r.Body.Close()
	return rb
}

// FetchJSON returns a map of type
func FetchJSON(url string) string {
	r := Fetch(url)
	rb, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return string(rb[:])
}

// StringInSlice checks if `x in arr`
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetTubeNames is a static list of available tube lines
func GetTubeNames() []string {
	return []string{
		"bakerloo",
		"central",
		"circle",
		"district",
		"hammersmith-city",
		"jubilee",
		"metropolitan",
		"northern",
		"piccadilly",
		"victoria",
		"waterloo-city",
	}
}
