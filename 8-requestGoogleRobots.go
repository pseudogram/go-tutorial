package requestGoogleRobots

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// open a resource
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("There was a failure getting robots")
	}
	defer res.Body.Close() // close a resouce at the same time as opening using defer (so next line runs)
	// careful if you do this inside a loop, as you could open 1000000 resources open simultaneously

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
