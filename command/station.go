package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
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

	stopPointData := utils.FetchJSON(utils.StopPointURL(arr1.Matches[0].ID))
	var arr2 utils.StopPointDataResp
	json.Unmarshal([]byte(stopPointData), &arr2)

	tubesAtStation := []string{}
	for _, line := range arr2.Lines {
		if utils.StringInSlice(line.ID, utils.GetTubeNames()) {
			tubesAtStation = append(tubesAtStation, line.ID)
		}
	}

	matcher := func(n *html.Node) bool {
		return scrape.Attr(n, "class") == "first-last-train-item"
	}

	boldFormat := color.New(color.FgWhite).Add(color.Bold).SprintFunc()

	stationLength := utf8.RuneCountInString(arr1.Matches[0].Name)
	fmt.Printf("%s", boldFormat(strings.Repeat("=", stationLength+23)))
	fmt.Printf("\n%s %s %s", "| Last trains from ", boldFormat(arr1.Matches[0].Name), " |")
	fmt.Printf("\n%s\n", boldFormat(strings.Repeat("=", stationLength+23)))

	for _, line := range tubesAtStation {
		firstHTML := utils.FetchHTML(utils.StopPointDeadline(arr1.Matches[0].ID, line, false))
		articles := scrape.FindAll(firstHTML, matcher)
		fmt.Printf("%s\n", boldFormat(" "+line))
		for _, article := range articles {
			fmt.Println("  ➡ " + scrape.Text(article))
		}
		fmt.Println(" ")
	}
	return nil
}
