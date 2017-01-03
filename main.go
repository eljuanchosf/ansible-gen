package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eljuanchosf/ansible-gen/ansibleGen"
	"github.com/urfave/cli"
)

var cliVersion string

func main() {

	app := cli.NewApp()
	app.Name = "ansible-gen"
	app.Version = cliVersion
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
	var skipGit bool
	var dryRun bool

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "dry-run, d",
			Usage:       "just print results, do not modify the filesystem",
			Destination: &dryRun,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "Creates a new Ansible project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "c",
					Value:       "",
					Usage:       "A comma separated list of the custom roles for the project",
					Destination: &customRoles,
				},
				cli.StringFlag{
					Name:        "g",
					Value:       "",
					Usage:       "A comma separated list of the Ansible Galaxy roles for the project",
					Destination: &galaxyRoles,
				},
				cli.BoolFlag{
					Name:        "skip-git",
					Usage:       "Do not initialize a Git repository for the project",
					Destination: &skipGit,
				},
				cli.StringFlag{
					Name:        "t",
					Value:       "",
					Usage:       "A template name (get the list of templates with the 'template' command)",
					Destination: &galaxyRoles,
				},
			},
			Action: func(c *cli.Context) error {
				projectName := c.Args().First()
				if projectName == "" {
					fmt.Println("Please provide a project name")
					cli.ShowSubcommandHelp(c)
					return nil
				}
				ansibleProject := *ansibleGen.NewAnsibleProject(projectName, customRoles, galaxyRoles)
				ansibleProject.Save(dryRun)
				if !skipGit {
					ansibleProject.InitGit(dryRun)
				}
				return nil
			},
		},
		{
			Name:    "role",
			Aliases: []string{"r"},
			Usage:   "Creates a new Ansible role",
			Action: func(c *cli.Context) error {
				roleName := c.Args().First()
				if c.Args().First() == "" {
					fmt.Println("Please provide a role name")
					cli.ShowSubcommandHelp(c)
					return nil
				}
				ansibleRole := *ansibleGen.NewAnsibleRole(roleName)
				ansibleRole.Save(dryRun)
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "Manage templates",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "list the available project templates",
					Action: func(c *cli.Context) error {
						fmt.Println("Listing templates")
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "update-cache",
							Value:       "",
							Usage:       "Force the project templates cache update",
							Destination: &customRoles,
						},
					},
				},
				{
					Name:  "create",
					Usage: "create a new project template",
					Action: func(c *cli.Context) error {
						fmt.Println("Created a new template")
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
