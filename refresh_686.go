package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	url := "http://blog.dota2.com/"
	needle := "6.86"
	for {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			s := string(contents[:])
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			if strings.Contains(s, needle) {
				fmt.Println("6.86 Patch Released")
				open.Run(url)
				break
			}
		}
	}

}
