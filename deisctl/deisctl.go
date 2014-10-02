package main

import (
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/grengojbo/deisctl/backend/fleet"
)

func main() {
	app := cli.NewApp()
	app.Name = "deisctl"
	app.Version = Version
	app.Usage = "Deis Control Utility"
	// app.Author = "Oleg Dolya"
	// app.Email = "oleg.dolya@gmail.com"
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "path to configuration file",
		},
		cli.StringFlag{
			Name:  "tunnel, t",
			Usage: "establish an SSH tunnel for communication with fleet and etcd",
		},
		cli.StringFlag{
			Name:  "endpoint, e",
			Value: "http://127.0.0.1:4001",
			Usage: "etcd endpoint for fleet",
		},
		cli.StringFlag{
			Name:  "etcd-key-prefix",
			Value: "/_coreos.com/fleet/",
			Usage: "keyspace for fleet data in etcd",
		},
		cli.StringFlag{
			Name:  "etcd-keyfile",
			Usage: "etcd key file authentication",
		},
		cli.StringFlag{
			Name:  "etcd-certfile",
			Usage: "etcd cert file authentication",
		},
		cli.StringFlag{
			Name:  "etcd-cafile",
			Usage: "etcd CA file authentication",
		},
		cli.StringFlag{
			Name:  "known-hosts-file",
			Value: "~/.ssh/known_hosts",
			Usage: "file used to store remote machine fingerprints",
		},
		cli.BoolTFlag{
			Name:  "strict-host-key-checking",
			Usage: "verify SSH host keys [default: true]",
		},
		cli.StringFlag{
			Name:  "request-timeout",
			Value: "3.0",
			Usage: "amount of time to allow a single request before considering it failed.",
		},
	}
	app.Before = func(c *cli.Context) error {
		fleet.Flags.Endpoint = c.GlobalString("endpoint")
		fleet.Flags.EtcdKeyPrefix = c.GlobalString("etcd-key-prefix")
		fleet.Flags.EtcdKeyFile = c.GlobalString("etcd-keyfile")
		fleet.Flags.EtcdCertFile = c.GlobalString("etcd-certfile")
		fleet.Flags.EtcdCAFile = c.GlobalString("etcd-cafile")
		fleet.Flags.KnownHostsFile = c.GlobalString("known-hosts-file")
		fleet.Flags.StrictHostKeyChecking = c.GlobalBool("strict-host-key-checking")
		timeout, _ := strconv.ParseFloat(c.GlobalString("request-timeout"), 64)
		fleet.Flags.RequestTimeout = timeout
		tunnel := c.GlobalString("tunnel")
		if tunnel != "" {
			// fmt.Println("Set tunnel ...")
			fleet.Flags.Tunnel = tunnel
		} else {
			fleet.Flags.Tunnel = os.Getenv("DEISCTL_TUNNEL")
		}
		return nil
	}

	// 	deisctlMotd := utils.DeisIfy("Deis Control Utility")
	// usage := `
	// Usage:
	//   deisctl <command> [<target>...] [options]

	// Commands:
	//   deisctl install [<service> | platform]
	//   deisctl uninstall [<service> | platform]
	//   deisctl start [<service> | platform]
	//   deisctl stop [<service> | platform]
	//   deisctl restart [<service> | platform]
	//   deisctl journal <service>
	//   deisctl config <component> <get|set> <args>
	//   deisctl update
	//   deisctl refresh-units

	// Example Commands:
	//   deisctl install platform
	//   deisctl uninstall builder@1
	//   deisctl scale router=2
	//   deisctl start router@2
	//   deisctl stop router builder
	//   deisctl status controller
	//   deisctl journal controller

	app.Run(os.Args)
}
