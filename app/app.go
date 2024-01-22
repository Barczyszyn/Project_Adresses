package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Generate will return the ready-to-run command line application.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Web Address App"
	app.Usage = "Search for Information About Web Addresses"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Web IP Search",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "ns",
			Usage:  "Web Server Name Search",
			Flags:  flags,
			Action: searchNs,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchNs(c *cli.Context) {
	host := c.String("host")

	names, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, name := range names {
		fmt.Println(name.Host)
	}
}
