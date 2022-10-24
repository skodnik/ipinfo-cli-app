package main

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type richIpInfo struct {
	Input string `json:"input"`
	Data  ipInfo `json:"data"`
}

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
	Asn      asn     `json:"asn,omitempty"`
	Company  company `json:"company,omitempty"`
	Privacy  privacy `json:"privacy,omitempty"`
	Abuse    abuse   `json:"abuse,omitempty"`
	Domains  domains `json:"domains,omitempty"`
}

type asn struct {
	Asn    string `json:"asn,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Route  string `json:"route,omitempty"`
	Type   string `json:"type,omitempty"`
}

type privacy struct {
	Vpn     bool   `json:"vpn,omitempty"`
	Proxy   bool   `json:"proxy,omitempty"`
	Tor     bool   `json:"tor,omitempty"`
	Relay   bool   `json:"relay,omitempty"`
	Hosting bool   `json:"hosting,omitempty"`
	Service string `json:"service,omitempty"`
}

type abuse struct {
	Address string `json:"address,omitempty"`
	Country string `json:"country,omitempty"`
	Email   string `json:"email,omitempty"`
	Name    string `json:"name,omitempty"`
	Network string `json:"network,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type company struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
}

type domains struct {
	Total   int           `json:"total,omitempty"`
	Domains []interface{} `json:"domains,omitempty"`
}

var host = "https://ipinfo.io/"
var hostAlt = "https://ip.zxq.co/"
var hostSly = "https://ipinfo.io/widget/demo/"

func main() {
	app := &cli.App{
		Name:    "ipinfo",
		Usage:   "get ip information",
		Version: "v1.0.9",
		Action: func(cCtx *cli.Context) error {
			ip := cCtx.String("ip")
			token := cCtx.String("token")
			jsonb := cCtx.Bool("json")
			pretty := cCtx.Bool("pretty")
			sly := cCtx.Bool("sly")
			alt := cCtx.Bool("alt")
			printIpInfo(ip, token, jsonb, pretty, sly, alt)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "ip",
				Aliases: []string{"i"},
				Value:   "",
				Usage:   "ip to search",
			},
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Value:   "",
				Usage:   "access token",
			},
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Value:   false,
				Usage:   "result to json",
			},
			&cli.BoolFlag{
				Name:    "pretty",
				Aliases: []string{"p"},
				Value:   false,
				Usage:   "prettier json",
			},
			&cli.BoolFlag{
				Name:    "sly",
				Aliases: []string{"s"},
				Value:   false,
				Usage:   "rich info without token (only main host)",
			},
			&cli.BoolFlag{
				Name:    "alt",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "use an alternative host ip.zxq.co",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printIpInfo(ip string, token string, jsonb bool, pretty bool, sly bool, alt bool) {
	ipInfo := ipInfo{}
	if sly {
		if ip == "" {
			log.Fatalln("Are you sure you're a sly?")
		}
		ipInfo = convertToIpInfoSly(getBody(makeRequestSly(ip, token)))
	} else {
		ipInfo = convertToIpInfo(getBody(makeRequest(ip, token, alt)))
	}

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

func makeRequest(ip string, token string, alt bool) *http.Response {
	params := url.Values{}
	params.Add("token", token)
	if alt {
		host = hostAlt
	}
	response, err := http.Get(host + ip + "?" + params.Encode())
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func makeRequestSly(ip string, token string) *http.Response {
	params := url.Values{}
	params.Add("token", token)
	client := &http.Client{}
	request, err := http.NewRequest("GET", hostSly+ip, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Referer", "https://ipinfo.io/")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		request, err := httputil.DumpRequestOut(request, true)
		log.Printf("Request: %s\n", string(request))
		log.Fatalln(err)
	}

	return response
}

func getBody(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response: %s\n", string(body))
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Printf("Response [%d]: %s\n", resp.StatusCode, string(body))
		log.Fatalf("http error")
	}

	return body
}

func convertToIpInfo(body []byte) ipInfo {
	var ipInfo ipInfo
	err := json.Unmarshal(body, &ipInfo)
	if err != nil {
		log.Printf("Response: %s\n", string(body))
		log.Fatalln(err)
	}

	return ipInfo
}

func convertToIpInfoSly(body []byte) ipInfo {
	var richIpInfo richIpInfo
	err := json.Unmarshal(body, &richIpInfo)
	if err != nil {
		log.Printf("Response: %s\n", string(body))
		log.Fatalln(err)
	}

	return richIpInfo.Data
}
