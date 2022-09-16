package main

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
)

type ipInfo struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

var ipData ipInfo

func main() {
	app := &cli.App{
		Name:    "ipinfo",
		Usage:   "get ip information",
		Version: "1.0.0",
		Action: func(cCtx *cli.Context) error {
			ip := cCtx.String("ip")
			getIpInfo(ip)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ip",
				Value: "",
				Usage: "ip to search",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getIpInfo(ip string) {
	resp, err := http.Get("https://ipinfo.io/" + ip)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &ipData)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(color.Ize(color.Green, "\n"+ipData.Ip+" - "+ipData.Org))
	fmt.Println(color.Ize(color.Green, ipData.Country+", "+ipData.Region+", "+ipData.City))
}
