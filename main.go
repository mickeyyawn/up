package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func checkEndpointHTTPStatusCode(url string) {

	//Status     string // e.g. "200 OK"
	//StatusCode int    // e.g. 200
	//Proto      string // e.g. "HTTP/1.0"
	//ProtoMajor int    // e.g. 1
	//ProtoMinor int    // e.g. 0

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	logrus.Infof("Checking endpoint: %s  Status: %s Code: %s Protocol: %s",
		url, resp.Status, strconv.Itoa(resp.StatusCode), resp.Proto)

}

func main() {

	URLs := make([]string, 5)

	//
	// these are the endpoints we will be checking every 30 seconds
	//
	URLs[0] = "https://www.cnn.com"
	URLs[1] = "https://yawn.me"
	URLs[2] = "https://espn.go.com"
	URLs[3] = "https://www.gobyexample.com"
	URLs[4] = "https://www.medium.com"

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)
	logrus.Infof("Starting application: %s", "UP")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {

		<-ticker.C

		for i := 0; i < 5; i++ {
			go checkEndpointHTTPStatusCode(URLs[i])
		}

	}

}
