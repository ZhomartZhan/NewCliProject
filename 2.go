package main

import  (
	"errors"
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "App for calculating and playing mp3 file"
	app.Description = "Just set three arguments for calculating and put mp3 to the directory"
	app.Usage = "Calculate inputted arguments and play the mp3 file"
	app.Commands = []cli.Command{
		{
			Name:  "calc",
			Usage: "Calculate inputted arguments",
			Action: func(c *cli.Context) error {
				z := c.Args().Get(0)
				a := c.Args().Get(1)
				b := c.Args().Get(2)
				if c.NArg() == 3 {
					total := 0
					num1, err := strconv.Atoi(a)
					if err != nil {
						return err
					}
					num2, err := strconv.Atoi(b)
					if err != nil {
						return err
					}
					if z == "+" {
						total = num1 + num2
					} else if z == "-" {
						total = num1 - num2
					} else {
						return errors.New("first argument should be + or -")
					}
					fmt.Println(total)
				} else {
					return errors.New("count of arguments should be only 3")
				}
				return nil
			},
		},
		{
			Name:    "player",
			Usage:   "play the mp3 file",
			Action: func(c *cli.Context) error {
				f, err := os.Open("1.mp3")
				if err != nil {
					log.Fatal(err)
				}
				st, format, err := mp3.Decode(f)
				if err != nil {
					log.Fatal(err)
				}
				defer st.Close()
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				speaker.Play(st)
				select {}
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
