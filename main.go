package main

import (
	"collexy/core/application"
	"encoding/json"
	"flag" // https://gobyexample.com/command-line-flags
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	//"collexy/globals"
	coreglobals "collexy/core/globals"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

func init() {

	if _, err := os.Stat("./config/config.json"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("Config file does not exist")
		} else {
			// other error
		}
	} else {

		configFile, err1 := os.Open("./config/config.json")
		defer configFile.Close()
		if err1 != nil {
			log.Println("Error opening config file")
			//printError("opening config file", err1.Error())
		}

		jsonParser := json.NewDecoder(configFile)
		if err1 = jsonParser.Decode(&coreglobals.Conf); err1 != nil {
			log.Println("Error parsing config file")
			//printError("parsing config file", err1.Error())
		}
		// log.Println(coreglobals.Conf.DbName)
		// log.Println(coreglobals.Conf.DbUser)
		// log.Println(coreglobals.Conf.DbPassword)
		// log.Println(coreglobals.Conf.DbHost)
		// log.Println(coreglobals.Conf.SslMode)
		coreglobals.Db = coreglobals.SetupDB()

	}

}
func main() {

	// After all flags are defined, call flag.parse() to parse the command line into the defined flags.
	flag.Parse()

	application.Main()

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
