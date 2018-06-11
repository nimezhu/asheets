package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/nimezhu/asheets"
	"github.com/nimezhu/data"
	"github.com/urfave/cli"
	"golang.org/x/oauth2/google"
	sheets "google.golang.org/api/sheets/v4"
)

func readGSheet(sheetId string, dir string) {

}
func readCol(id string, srv *sheets.Service, sheetid string, col string) []string {
	readRange := id + "!" + col + "1:" + col
	resp, err := srv.Spreadsheets.Values.Get(sheetid, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	fmt.Println("Reading ", id)
	if len(resp.Values) > 0 {
		r := make([]string, len(resp.Values))
		for i, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			r[i] = row[0].(string)
		}
		return r
	} else {
		return []string{}
	}
}

const VERSION = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "sheetstools"
	app.Usage = "auto fill google sheets "
	app.EnableBashCompletion = true //TODO
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:   "fillExprDis",
			Usage:  "extract encode experiment description and fill in google sheet",
			Action: CmdFillExpr,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input gsheet id",
				},
				cli.StringFlag{
					Name:  "title,t",
					Usage: "sheet title",
					Value: "Sheet1",
				},
			},
		},
	}

	app.Run(os.Args)
}
func CmdFillExpr(c *cli.Context) {
	home := os.Getenv("HOME")
	dir := path.Join(home, ".cnb") //TODO
	ctx := context.Background()
	title := c.String("title")
	sheetid := c.String("input")
	b, err := data.Asset("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	gA := asheets.NewGAgent(dir)
	client := gA.GetClient(ctx, config)
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}
	value := readCol(title, srv, sheetid, "I")
	for i, v := range value {
		if i == 0 {
			continue
		} else {
			url := "https://www.encodeproject.org/experiments/" + v + "/?format=json"
			res, err := http.Get(url)
			if err != nil {
				panic(err.Error())
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err.Error())
			}
			v := make(map[string]interface{})
			json.Unmarshal(body, &v)

			if k, ok := v["description"]; ok {
				fmt.Println(k)
			}

		}
	}
}
