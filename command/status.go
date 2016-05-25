package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

type Message struct {
	Name, modeName []string
}

func getJSON(url string) *json.Decoder {
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body)
}

func CmdStatus(c *cli.Context) error {
	jsonData := getJSON("https://api.tfl.gov.uk/Line/Mode/tube")

	fmt.Println(jsonData)
	return nil
}
