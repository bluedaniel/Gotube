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

// StopPointSearchResp contains stopPoint data
type StopPointSearchResp struct {
	Total   int
	Matches []struct {
		IcsID string
		Name  string
		ID    string
	}
}

// StopPointDataResp gets the station ID
type StopPointDataResp struct {
	Lines []struct {
		ID string
	}
}

// CmdStation runs `tube status`
func CmdStation(c *cli.Context) error {
	q := strings.Join(c.Args()[:], " ")
	query := &url.URL{Path: strings.Replace(q, "and", "&", -1)}

	stopPointSearch := utils.GetJSON("https://api.tfl.gov.uk/StopPoint/Search/" + query.String() + "?modes=tube")
	var arr1 StopPointSearchResp
	json.Unmarshal([]byte(stopPointSearch), &arr1)

	if arr1.Total == 0 {
		fmt.Println("No results found")
		os.Exit(2)
	}

	stopPointData := utils.GetJSON("https://api.tfl.gov.uk/StopPoint/" + arr1.Matches[0].ID)
	var arr2 StopPointDataResp
	json.Unmarshal([]byte(stopPointData), &arr2)

	fmt.Println(utils.GetTubeNames())
	//
	// fmt.Println("https://api.tfl.gov.uk/StopPoint/Search/" + query.String() + "?modes=tube")
	// fmt.Println(arr.Matches[0])

	// first := utils.GetJSON("https://tfl.gov.uk/Timetables/FirstLastServicesSummaryAjax?fromId=940GZZLUHAI&lines=victoria&firstNextDay=true")
	// last := utils.GetJSON("https://tfl.gov.uk/Timetables/FirstLastServicesSummaryAjax?fromId=940GZZLUHAI&lines=victoria&firstNextDay=false")

	// fmt.Printf(url.QueryEscape("https://api.tfl.gov.uk/StopPoint/Search/" + q + "?modes=tube"))
	return nil
}
