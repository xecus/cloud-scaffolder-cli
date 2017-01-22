package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	cloud_scaffolder "github.com/xecus/cloud-scaffolder"
	"log"
	"os"
)

func scaffold(c *cli.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("added task: ", c.Args().First())

	// Init DB
	i := cloud_scaffolder.Impl{}
	i.InitModelDb()
	i.InitSchema()

	// Prepare
	cloud_scaffolder.PrepareVagrantControl()
	cloud_scaffolder.GenerateVagrantFile()

	cloud_scaffolder.DeleteAlltVm(i)

	a := cloud_scaffolder.GenerateVagrantModel()

	for _, vm := range a.Vms {
		vm.CreateVm(i)
	}

	cloud_scaffolder.Serve(&i)

}

func serve(c *cli.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("added task: ", c.Args().First())
}

func main() {
	app := cli.NewApp()
	app.Name = "CloudScaffolder"
	app.Usage = "This is test"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "scaffold",
			Aliases: []string{"sc"},
			Usage:   "scaffold cloud",
			Action: func(c *cli.Context) error {
				scaffold(c)
				return nil
			},
		},
		{
			Name:    "serve",
			Aliases: []string{"sv"},
			Usage:   "start server",
			Action: func(c *cli.Context) error {
				serve(c)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
