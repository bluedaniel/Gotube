package utils

import (
	"io/ioutil"
	"net/http"
)

// GetJSON is a fetch helper
func GetJSON(url string) string {
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	rb, err := ioutil.ReadAll(r.Body)
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
