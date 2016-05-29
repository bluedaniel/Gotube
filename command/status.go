package command

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bluedaniel/gotube/utils"
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

func serviceStatus(tube Tube) int { return tube.LineStatuses[0].StatusSeverity }

func formatTubeReason(name, reason string) string {
	name = strings.Replace(name, "&", "and", -1)
	re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(name) + " line:")
	return strings.Trim(re.ReplaceAllString(reason, ""), " ")
}

func pickEmoji(v int) string {
	switch v {
	case 20:
		return ":no_entry_sign:"
	case 10:
		return ":thumbsup:"
	case 9:
		return ":ok_hand:"
	}
	return ":shit:"
}

type byStatus []Tube

func (a byStatus) Len() int      { return len(a) }
func (a byStatus) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byStatus) Less(i, j int) bool {
	return a[i].LineStatuses[0].StatusSeverity > a[j].LineStatuses[0].StatusSeverity
}

// CmdStatus runs `tube status`
func CmdStatus(c *cli.Context) error {
	s := utils.GetJSON("https://api.tfl.gov.uk/line/mode/tube/status")
	tubeTextFormat := color.New(color.FgWhite).Add(color.Bold).SprintFunc()

	var arr []Tube
	json.Unmarshal([]byte(s), &arr)
	sort.Sort(byStatus(arr))

	for i, e := range arr {
		if i > 0 && serviceStatus(arr[i-1]) != serviceStatus(e) {
			fmt.Println("----------------------------")
		}
		fmt.Printf("%s %s\n", emoji.Sprint(pickEmoji(serviceStatus(e))),
			tubeTextFormat(e.Name))

		if serviceStatus(e) != 10 {
			for _, statuses := range e.LineStatuses {
				fmt.Printf("  %s %s\n", emoji.Sprint(":exclamation:"),
					formatTubeReason(e.Name, statuses.Reason))
			}
		}
	}
	return nil
}
