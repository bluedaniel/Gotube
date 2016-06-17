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

func GetTubeNames() [11]string {
	return [11]string{
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
