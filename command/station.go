package command

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/bluedaniel/gotube/utils"
	"github.com/urfave/cli"
)

// CmdStation runs `tube status`
func CmdStation(c *cli.Context) error {
	q := strings.Join(c.Args()[:], " ")
	query := &url.URL{Path: strings.Replace(q, "and", "&", -1)}

	stopPointSearch := utils.GetJSON(utils.StopPointSearchURL(query.String()))
	var arr1 utils.StopPointSearchResp
	json.Unmarshal([]byte(stopPointSearch), &arr1)

	if arr1.Total == 0 {
		fmt.Println("No results found")
		os.Exit(2)
	}

	stopPointData := utils.GetJSON(utils.StopPointURL(arr1.Matches[0].ID))
	var arr2 utils.StopPointDataResp
	json.Unmarshal([]byte(stopPointData), &arr2)

	tubesAtStation := []string{}
	for _, line := range arr2.Lines {
		if utils.StringInSlice(line.ID, utils.GetTubeNames()) {
			tubesAtStation = append(tubesAtStation, line.ID)
		}
	}

	for i, line := range tubesAtStation {
		if i == 0 {
			first := utils.GetJSON(utils.StopPointDeadline(arr1.Matches[0].ID, line, true))
			fmt.Println(first)
		}
	}

	// first := utils.GetJSON("https://tfl.gov.uk/Timetables/FirstLastServicesSummaryAjax?fromId=940GZZLUHAI&lines=victoria&firstNextDay=true")
	// last := utils.GetJSON("https://tfl.gov.uk/Timetables/FirstLastServicesSummaryAjax?fromId=940GZZLUHAI&lines=victoria&firstNextDay=false")

	// fmt.Printf(url.QueryEscape("https://api.tfl.gov.uk/StopPoint/Search/" + q + "?modes=tube"))
	return nil
}
