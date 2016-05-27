package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
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

	colors := map[string]color.Attribute{
		"Bakerloo":           color.FgHiBlack,
		"Central":            color.FgRed,
		"Circle":             color.FgHiYellow,
		"District":           color.FgGreen,
		"Hammersmith & City": color.FgHiYellow,
		"Jubilee":            color.FgHiCyan,
		"Metropolitan":       color.FgHiYellow,
		"Northern":           color.FgWhite,
		"Piccadilly":         color.FgHiBlue,
		"Victoria":           color.FgBlue,
		"Waterloo & City":    color.FgHiCyan,
	}

	for _, e := range arr {
		colFn := color.New(colors[e.Name]).Add(color.Bold).SprintFunc()
		fmt.Printf("%s %s (%s)\n", emoji.Sprint(":rage:"), colFn(e.Name), e.ModeName)
	}
	return nil
}
