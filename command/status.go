package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/urfave/cli"
)

// LineStatuses contains line data
type status struct {
	StatusSeverity            int
	StatusSeverityDescription string
	Reason                    string
}

// Tube contains ftl data
type Tube struct {
	Name         string
	LineStatuses []status
}

func getJSON(url string) string {
	r, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	rb, err := ioutil.ReadAll(r.Body)
	return string(rb[:])
}

func pickEmoji(v int) string {
	switch v {
	case 10:
		return ":thumbsup:"
	case 9:
		return ":ok_hand:"
	}
	return ":shit:"
}

func serviceStatus(tube Tube) int {
	return tube.LineStatuses[0].StatusSeverity
}

// CmdStatus runs `tube status`
func CmdStatus(c *cli.Context) error {
	s := getJSON("https://api.tfl.gov.uk/line/mode/tube/status")
	tubeTextFormat := color.New(color.FgWhite).Add(color.Bold).SprintFunc()

	var arr []Tube
	json.Unmarshal([]byte(s), &arr)

	for _, e := range arr {
		fmt.Printf("%s %s\n", emoji.Sprint(pickEmoji(serviceStatus(e))),
			tubeTextFormat(e.Name))

		if serviceStatus(e) < 10 {
			re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(e.Name) + " line:")
			for _, statuses := range e.LineStatuses {
				fmt.Printf("  %s %s\n", emoji.Sprint(":exclamation:"),
					strings.Trim(re.ReplaceAllString(statuses.Reason, ""), " "))
			}
		}
	}
	return nil
}
