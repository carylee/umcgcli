package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "umcgcli"
	app.Usage = "UMC General Conference 2016 Command Line Tool!"
	app.Action = func(c *cli.Context) error {
		getPetitionText(c.Args()[0])
		return nil
	}

	app.Run(os.Args)
}

func getPetitionText(petitionNumber string) {
	response, err := http.Get(fmt.Sprintf("http://umcgc.herokuapp.com/api/petitions/%s", petitionNumber))
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println(string(contents))
}
