package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

func MainAction(c *cli.Context) error {
	a := c.Args().Get(0)
	b := c.Args().Get(1)
	d := c.Args().Get(2)
	if c.NArg() == 3 {
		if a == "+" {
			num1, err := strconv.Atoi(b)
			if err != nil {
				return err
			}
			num2, err := strconv.Atoi(d)
			if err != nil {
				return err
			}
			fmt.Println(num1 + num2)
		} else if a == "-" {
			num1, err := strconv.Atoi(b)
			if err != nil {
				return err
			}
			num2, err := strconv.Atoi(d)
			if err != nil {
				return err
			}
			fmt.Println(num1 - num2)
		} else {
			return errors.New("symbol must be '+' or '-' ")
		}

	} else {
		return errors.New("count of arguments should be only 3")
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	app.Action = MainAction
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
