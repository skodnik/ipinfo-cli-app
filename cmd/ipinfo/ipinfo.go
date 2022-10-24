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
	Ip       string  `json:"ip"`
	Hostname string  `json:"hostname"`
	Anycast  bool    `json:"anycast,omitempty"`
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Loc      string  `json:"loc"`
	Org      string  `json:"org"`
	Postal   string  `json:"postal"`
	Timezone string  `json:"timezone"`
	Readme   string  `json:"readme"`
	Asn      Asn     `json:"asn,omitempty"`
	Company  Company `json:"company,omitempty"`
	Privacy  Privacy `json:"privacy,omitempty"`
	Abuse    Abuse   `json:"abuse,omitempty"`
	Domains  Domains `json:"domains,omitempty"`
}

type Asn struct {
	Asn    string `json:"asn,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Route  string `json:"route,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Privacy struct {
	Vpn     bool   `json:"vpn,omitempty"`
	Proxy   bool   `json:"proxy,omitempty"`
	Tor     bool   `json:"tor,omitempty"`
	Relay   bool   `json:"relay,omitempty"`
	Hosting bool   `json:"hosting,omitempty"`
	Service string `json:"service,omitempty"`
}

type Abuse struct {
	Address string `json:"address,omitempty"`
	Country string `json:"country,omitempty"`
	Email   string `json:"email,omitempty"`
	Name    string `json:"name,omitempty"`
	Network string `json:"network,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type Company struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Domains struct {
	Total   int           `json:"total,omitempty"`
	Domains []interface{} `json:"domains,omitempty"`
}

const host = "https://ipinfo.io/"

func main() {
	app := &cli.App{
		Name:    "ipinfo",
		Usage:   "get ip information",
		Version: "v1.0.5",
		Action: func(cCtx *cli.Context) error {
			ip := cCtx.String("ip")
			token := cCtx.String("token")
			jsonb := cCtx.Bool("json")
			pretty := cCtx.Bool("pretty")
			printIpInfo(ip, token, jsonb, pretty)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ip",
				Value: "",
				Usage: "ip to search",
			},
			&cli.StringFlag{
				Name:  "token",
				Value: "",
				Usage: "access token",
			},
			&cli.BoolFlag{
				Name:  "json",
				Value: false,
				Usage: "result to json",
			},
			&cli.BoolFlag{
				Name:  "pretty",
				Value: false,
				Usage: "prettier json",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printIpInfo(ip string, token string, jsonb bool, pretty bool) {
	ipInfo := convertToIpInfo(getBody(makeRequest(ip, token)))

	if ipInfo.Ip == "" {
		log.Fatalln("Incorrect input data, token perhaps?")
	}

	if jsonb {
		if pretty {
			marshal, err := json.MarshalIndent(ipInfo, "", "    ")
			if err != nil {
				return
			}
			fmt.Println(string(marshal))

			return
		}

		marshal, err := json.Marshal(ipInfo)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))

		return
	}

	fmt.Println(color.Ize(color.Green, "\n"+ipInfo.Ip+" - "+ipInfo.Org))
	fmt.Println(color.Ize(color.Green, ipInfo.Country+", "+ipInfo.Region+", "+ipInfo.City))
}

func makeRequest(ip string, token string) *http.Response {
	resp, err := http.Get(host + ip + "?token=" + token)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func getBody(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func convertToIpInfo(body []byte) ipInfo {
	var ipInfo ipInfo
	err := json.Unmarshal(body, &ipInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return ipInfo
}
