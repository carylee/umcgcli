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
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://api.umcgc.com/petitions/%s", petitionNumber)
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.Header.Set("Content-Type", "text/plain")
	response, _ := client.Do(request)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println(string(contents))
}
