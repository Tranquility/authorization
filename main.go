package main

import (
    "os"
    "github.com/codegangsta/cli"
)

func main() {
        app := cli.NewApp()
        app.Name = "authorization"
        app.Usage = "Manager ssh authorizations"
        app.Version = "0.1.0"
        app.Commands = []cli.Command {
                {
                        Name:      "add",
                        ShortName: "a",
                        Usage:     "Add a user",
                        Action: func(c *cli.Context) {
                                Add(c.Args().First())
                        },
                },
                {
                        Name:      "remove",
                        ShortName: "r",
                        Usage:     "Remove a user",
                        Action: func(c *cli.Context) {
                                Remove(c.Args().First())
                        },
                },
        }
        app.Run(os.Args)
}
