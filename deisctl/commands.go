package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/grengojbo/deisctl/client"
	// "github.com/grengojbo/deisctl/utils"
)

var taskRestart = []string{"service", "platform", "unit"}
var cliUnit = []string{"unit"}

var DeisDefaulService = []string{"cache", "router", "database", "controller", "registry", "builder"}

var compRestartUnit = []string{"test_v1", "naxyi_v2", "putin_v3"}

var Commands = []cli.Command{
	commandList,
	commandScale,
	commandStart,
	commandRestart,
	commandStop,
	commandStatus,
	commandJournal,
	commandInstall,
	commandUninstall,
	commandConfig,
	commandUpdate,
	commandRefreshUnits,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "List service",
	Description: `
	deisctl list        - list deis services
	deisctl list unit   - list deis units
`,
	BashComplete: bashList,
	Action:       doList,
}

var commandScale = cli.Command{
	Name:  "scale",
	Usage: "[<service>=<num>]",
	Description: `
`,
	Action: doScale,
}

var commandStart = cli.Command{
	Name:  "start",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
`,
	Action: doStart,
}

var commandRestart = cli.Command{
	Name:  "restart",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
	Restaring:
	 - service
	 - platform
`,
	// BashComplete: taskRestart,
	Subcommands: []cli.Command{
		{
			Name:   "unit",
			Usage:  "restart only service",
			Action: doRestart,
		},
		{
			Name:   "platform",
			Usage:  "restart only service",
			Action: doRestart,
		},
		{
			Name:   "service",
			Usage:  "restart only service",
			Action: doRestart,
		},
	},
}

var commandStop = cli.Command{
	Name:  "stop",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
`,
	Subcommands: []cli.Command{
		// {
		// 	Name:   "unit",
		// 	Usage:  "restart only service",
		// 	Action: doRestart,
		// },
		{
			Name:   "platform",
			Usage:  "restart only service",
			Action: doRestart,
		},
		{
			Name:   "service",
			Usage:  "restart only service",
			Action: doRestart,
		},
	},
}

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
`,
	Action: doStatus,
}

var commandJournal = cli.Command{
	Name:  "journal",
	Usage: "[service <service name> | unit <unit name>]",
	Description: `
`,
	Action: doJournal,
}

var commandInstall = cli.Command{
	Name:  "install",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
	deisctl install service <cache | router | database | controller | registry | builder>
	deisctl install service platform
	deisctl install service unit <unit name>
`,
	Subcommands: []cli.Command{
		{
			Name:         "unit",
			Usage:        "install <unit name>",
			BashComplete: bashUnit,
			Action:       doInstall,
		},
		{
			Name:   "platform",
			Usage:  "install platform",
			Action: doInstall,
		},
		{
			Name:         "service",
			Usage:        "install services <service name>",
			BashComplete: bashService,
			Action:       doInstall,
		},
	},
	// Action: doInstall,
}

var commandUninstall = cli.Command{
	Name:  "uninstall",
	Usage: "[service <service name> | platform | unit <unit name>]",
	Description: `
`,
	Action: doUninstall,
}

var commandConfig = cli.Command{
	Name:  "config",
	Usage: "<component> <get|set> <args>",
	Description: `
`,
	Action: doConfig,
}

var commandUpdate = cli.Command{
	Name:  "update",
	Usage: "",
	Description: `
`,
	Action: doUpdate,
}

var commandRefreshUnits = cli.Command{
	Name:  "refresh-units",
	Usage: "",
	Description: `
`,
	Action: doRefreshUnits,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exit(err error, code int) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(code)
}

func bashList(c *cli.Context) {
	if len(c.Args()) > 0 {
		return
	}
	for _, t := range cliUnit {
		fmt.Println(t)
	}
}

func bashService(c *cli.Context) {
	if len(c.Args()) > 0 {
		return
	}
	for _, t := range DeisDefaulService {
		fmt.Println(t)
	}
}

func bashUnit(c *cli.Context) {
	if len(c.Args()) > 0 {
		return
	}
	for _, t := range compRestartUnit {
		fmt.Println(t)
	}
}

func doList(c *cli.Context) {
	// fmt.Println("fleet.Flags.Endpoint:", fleet.Flags.Endpoint)
	// fmt.Println("fleet.Flags.EtcdKeyPrefix:", fleet.Flags.EtcdKeyPrefix)
	// fmt.Println("fleet.Flags.EtcdKeyFile:", fleet.Flags.EtcdKeyFile)
	// fmt.Println("fleet.Flags.EtcdCertFile:", fleet.Flags.EtcdCertFile)
	// fmt.Println("fleet.Flags.EtcdCAFile:", fleet.Flags.EtcdCAFile)
	// fmt.Println("fleet.Flags.KnownHostsFile:", fleet.Flags.KnownHostsFile)
	// fmt.Println("fleet.Flags.StrictHostKeyChecking:", fleet.Flags.StrictHostKeyChecking)
	// fmt.Println("fleet.Flags.RequestTimeout:", float64(fleet.Flags.RequestTimeout))
	// fmt.Println("fleet tunnel:", fleet.Flags.Tunnel)

	f, err := client.NewClient("fleet")
	assert(err)
	if c.Args().Present() {
		// fmt.Println("List:", c.Args().First(), "...")
		err = f.List("unit")
		assert(err)
	} else {
		err = f.List("service")
		assert(err)
	}
}

func doScale(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Scale(targets)
	// assert(err)
}

func doStart(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Start(targets)
	// assert(err)
}

func doRestart(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Restart(targets)
	// assert(err)
}

func doStop(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Stop(targets)
	// assert(err)
}

func doStatus(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Status(targets)
	// assert(err)
}

func doJournal(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Journal(targets)
	// assert(err)
}

func doInstallService(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	target := []string{"service"}
	err = f.Install(targets)
	assert(err)
}

func doInstallUnit(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	target := []string{"unit"}
	err = f.Install(targets)
	assert(err)
}

func doInstall(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	target := []string{"platform"}
	err = f.Install(targets)
	assert(err)
}

func doUninstall(c *cli.Context) {
	// f, err := client.NewClient("fleet")
	// assert(err)
	// err = f.Uninstall(targets)
	// assert(err)
}

func doConfig(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	err = f.Config()
	assert(err)
}

func doUpdate(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	err = f.Update()
	assert(err)
}

func doRefreshUnits(c *cli.Context) {
	f, err := client.NewClient("fleet")
	assert(err)
	err = f.RefreshUnits()
	assert(err)
}
