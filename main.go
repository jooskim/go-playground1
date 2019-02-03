package main

import (
	"github.com/jooskim/playground1/sampleModule"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	var language string

	app := cli.NewApp()
	app.Name = "Test App"
	app.Version = "1.0.0"
	app.Usage = "Testing my test app"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
			Destination: &language,
		},
		cli.StringFlag{
			Name: "home",
			Usage: "your home directory",
			EnvVar: "HOME",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "complete",
			Aliases: []string{"c"},
			Usage: "complete a task on the list",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "base, b",
					Usage: "base folder",
				},
			},
			Action: func (c *cli.Context) error {
				log.Println("Your home is : " + c.Parent().String("home"))
				log.Println("base folder is : " + c.String("base"))
				channelSample()
				sampleModule.TestRun()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func channelSample() {
	c1 := make(chan int)
	c2 := make(chan int)

	for i := 0; i < 10; i++ {
		go func(val int) {
			time.Sleep(time.Duration(val * 500) * time.Millisecond)
			c1 <- val
		}(i)
	}

	go func() {
		c2 <- 8
	}()

	for {
		select {
		case val := <-c1:
			log.Printf("from c1 : %d\n", val)
		case val := <-c2:
			log.Printf("from c2 : %d\n", val)
		case <-time.After(10 * time.Second):
			log.Printf("timeout\n")
			return

		}
	}

}
