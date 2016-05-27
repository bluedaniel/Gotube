package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// Tube contains ftl data
type Tube struct {
	Name, ModeName string
}

func getJSON(url string) []byte {
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	rb, err := ioutil.ReadAll(r.Body)
	return rb
}

// CmdStatus runs `tube status`
func CmdStatus(c *cli.Context) error {
	jsonByte := getJSON("https://api.tfl.gov.uk/Line/Mode/tube")

	var arr []Tube
	json.Unmarshal(jsonByte, &arr)

	col := color.New(color.FgCyan).Add(color.Underline)
	col.Println("Prints cyan text with an underline.")

	for _, e := range arr {
		fmt.Printf("%v line is of type %v \n", e.Name, e.ModeName)
	}
	return nil
}
