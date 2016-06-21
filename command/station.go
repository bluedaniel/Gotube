package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	"golang.org/x/net/html"

	"github.com/bluedaniel/gotube/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/yhat/scrape"
)

// CmdStation runs `tube status`
func CmdStation(c *cli.Context) error {
	q := strings.Join(c.Args()[:], " ")
	query := &url.URL{Path: strings.Replace(q, "and", "&", -1)}

	var arr1 utils.StopPointSearchResp
	stopPointSearch := utils.FetchJSON(utils.StopPointSearchURL(query.String()))
	json.Unmarshal([]byte(stopPointSearch), &arr1)

	if arr1.Total == 0 {
		fmt.Println("No results found")
		os.Exit(2)
	}

	stationID := arr1.Matches[0].ID

	stopPointData := utils.FetchJSON(utils.StopPointURL(stationID))
	var arr2 utils.StopPointDataResp
	json.Unmarshal([]byte(stopPointData), &arr2)

	tubesAtStation := []string{}
	for _, line := range arr2.Lines {
		if utils.StringInSlice(line.ID, utils.GetTubeNames()) {
			tubesAtStation = append(tubesAtStation, line.ID)
		}
	}

	fmt.Printf("\n%s %s", "Last trains from", utils.BoldFormat(arr1.Matches[0].Name))
	fmt.Printf("\n%s\n", utils.BoldFormat(strings.Repeat("=", utf8.RuneCountInString(arr1.Matches[0].Name)+17)))

	var lines = map[string]*html.Node{}

	var wg sync.WaitGroup
	for _, line := range tubesAtStation {
		wg.Add(1)
		go func(stationID string, line string) {
			defer wg.Done()
			lines[line] = utils.FetchHTML(utils.StopPointDeadline(stationID, line, false))
		}(stationID, line)
	}
	wg.Wait()

	for key, value := range lines {
		articles := scrape.FindAll(value, func(n *html.Node) bool {
			return scrape.Attr(n, "class") == "first-last-train-item"
		})
		fmt.Printf("%s\n", utils.BoldFormat(strings.Title(strings.Replace(key, "-", " & ", -1))))
		for _, article := range articles {
			fmt.Println(color.GreenString("➡ ") + scrape.Text(article))
		}
	}
	return nil
}
