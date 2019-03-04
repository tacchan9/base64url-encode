package main

import (
	"encoding/base64"
	"os"
        "log"
	"fmt"
	"github.com/urfave/cli"
)

const version = "0.1.0"

func main() {

	var file string

	app := cli.NewApp()
	app.Name = "base64url-encode"
	app.Version = version
	app.Usage = "base64url-encode"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "file path",
			Destination: &file,
		},
	}

	app.Action = func(c *cli.Context) error {
		if file == "" {
			return fmt.Errorf("please set file path")
		}

		file, _ := os.Open(file)
		defer file.Close()
	
		fi, _ := file.Stat()
		size := fi.Size()
	
		data := make([]byte, size)
		file.Read(data)
	
		uEnc := base64.URLEncoding.EncodeToString(data)
		fmt.Println(uEnc)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
