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
