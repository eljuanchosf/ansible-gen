package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

var cliVersion string

func main() {

	app := cli.NewApp()
	app.Name = "ansible-gen"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Juan Pablo Genovese",
			Email: "juanpgenovese@gmail.com",
		},
	}
	app.Copyright = "(c) 2016 Juan Pablo Genovese"
	app.HelpName = "ansible-gen"
	app.Usage = "Generates and scaffolds Ansible projects and roles"

	var customRoles string
	var galaxyRoles string
	var projectGit bool

	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "Creates a new Ansible project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "c",
					Value:       "",
					Usage:       "Specify the custom roles for the project",
					Destination: &customRoles,
				},
				cli.StringFlag{
					Name:        "g",
					Value:       "",
					Usage:       "Specify the Ansible Galaxy roles for the project",
					Destination: &galaxyRoles,
				},
				cli.BoolFlag{
					Name:        "skip-git",
					Usage:       "Do not initialize a Git repository for the project",
					Destination: &projectGit,
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("Please provide a project name")
					cli.ShowSubcommandHelp(c)
					return nil
				}
				fmt.Printf("Create the project %s with custom roles %s and galaxy roles %s and %v git", c.Args().First(), customRoles, galaxyRoles, projectGit)
				return nil
			},
		},
		{
			Name:    "role",
			Aliases: []string{"r"},
			Usage:   "Creates a new Ansible role",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("Please provide a role name")
					cli.ShowSubcommandHelp(c)
					return nil
				}
				fmt.Println("Create the role: ", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
