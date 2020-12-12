package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// define links slice with the type of string
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel with the data type of string
	c := make(chan string)

	// loop over the links slice
	// each link will be input of the function checkLink
	// each go-routine will take care for each time function
	// is called.
	for _, link := range links {
		go checkLink(link, c)
	}

	// continue run checklink
	// after finishing check the input link
	// Waiting 5s, the go routine will check that link again

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// check whether the link (url) is up or not
// using http.Get(link)
// Note that channel is one of the arguments of function
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		// Send result to channel
		c <- link
		return
	}
	fmt.Println(link + " is up!")
	// Send result to channel
	c <- link
}
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// MyStruct is an example structure for this program.
// type MyStruct struct {
// 	StructData int `json:"StructData"`
// }

// func main() {
// 	filename := "myFile.json"
// 	err := checkFile(filename)
// 	if err != nil {
// 		logrus.Error(err)
// 	}
// 	file, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		logrus.Error(err)
// 	}

// 	data := []MyStruct{}
// 	datas := 
// 	// Here the magic happens!
// 	json.Unmarshal(file, &data)
// 	for i := 0; i < 10; i++ {
// 		newStruct := &MyStruct{
// 			StructData: i,
// 		}

// 		data = append(data, *newStruct)
// 	}

// 	// Preparing the data to be marshalled and written.
// 	dataBytes, err := json.Marshal(data)
// 	if err != nil {
// 		logrus.Error(err)
// 	}

// 	err = ioutil.WriteFile(filename, dataBytes, 0644)
// 	if err != nil {
// 		logrus.Error(err)
// 	}
// }

// func checkFile(filename string) error {
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		_, err := os.Create(filename)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
