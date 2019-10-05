package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func main() {

	//init logging
	logFile, err := os.OpenFile("./sf_lhm_api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
		log.Fatal("could not open log-file sf_lhm_api.log, exiting.")
		os.Exit(1)
	}

	logger = log.New(logFile, "sf_lhm_api:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("sf_lhm_api mainline started.")

	endPoint := ":8080"
	if len(os.Args) == 2 {
		endPoint = ":" + os.Args[1]
	}

	router := NewRouter()
	fmt.Printf("serving on %s\n", endPoint[1:])
	logger.Fatal(http.ListenAndServe(endPoint, router)) //non-tls
	//logger.Fatal(http.ListenAndServeTLS(endPoint, "cert.pem", "key.pem", router)) //tls mode
}
