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

	i := cloud_scaffolder.Impl{}
	i.InitModelDb()
	i.InitSchema()

	cloud_scaffolder.PrepareVagrantControl()
	cloud_scaffolder.GenerateVagrantFile()

	//a := cloud_scaffolder.GenerateVagrantModel()
	//a.Vms[0].DeleteAlltVm(i)
	//s := a.Vms[0].CreateVm(i)
	//log.Println(s)

	v := cloud_scaffolder.GetVm(i)
	v.ShowVm()
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
