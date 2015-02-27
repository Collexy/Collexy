package main

import (
	"flag" // https://gobyexample.com/command-line-flags
	"net"
	"log"
	"io/ioutil"
	"net/http"
    "collexy/core/api"
    "collexy/core/application"
    "encoding/json"
    "os"
    "collexy/globals"
    coreglobals "collexy/core/globals"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

func init () {
    
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
        if err1 = jsonParser.Decode(&globals.Conf); err1 != nil {
            log.Println("Error parsing config file")
            //printError("parsing config file", err1.Error())
        }
        // log.Println(globals.Conf.DbName)
        // log.Println(globals.Conf.DbUser)
        // log.Println(globals.Conf.DbPassword)
        // log.Println(globals.Conf.DbHost)
        // log.Println(globals.Conf.SslMode)
        coreglobals.Db = coreglobals.SetupDB()

    }

    
}
func main() {

	// After all flags are defined, call flag.parse() to parse the command line into the defined flags. 
	flag.Parse()
    
    api.Main()

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