package main

import (
	"fmt"
	"os"

	gonode "github.com/jgranstrom/gonodepkg"
	json "github.com/jgranstrom/go-simplejson"
)

var (
	f 	*os.File
)

func process(cmd *json.Json) (response *json.Json) {
	fmt.Println("Process called.")
	response, m, err := json.MakeMap()
	m["text"] = "No response"

	f, err = os.OpenFile("/home/user/temp2", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening ouput file.")
		m["text"] = "Error with output file"
		return 
	}
	defer f.Close()

	text := cmd.Get("text").MustString()

	switch text {
		case "Initialize":
			f.WriteString("Initialize command read.")
			f.WriteString("\n")
			m["text"] = "Initialize succesful"
/*
		addrModbus := cmd.Get("addrModbus").MustString()
		portModbus := cmd.Get("portModbus").MustString()
		configFile := cmd.Get("configFile").MustString()

		if modbus.LoggerEnable == true {
			log.Println("Loading config file...")
		}

		// Create smart meter with settings according to config file
		smartMeter := modbus.NewSmartMeter(*configFile)
		if smartMeter == nil {
			log.Println("Config file was not succefully loaded")
			return
		}

		if modbus.LoggerEnable == true {
			log.Println(smartMeter)
		}

		// Initialize and start modbus TCP server
		server := modbus.NewTCPServer(*portModbus, *addrModbus, smartMeter)
		if server == nil {
			log.Println("Server was not succesfully initialize")
			return
		}
		log.Println("Server starts.................")
		server.ServerStart()

		m["text"] = "Initiation of Modbus server successful"
		//TMP waiting feature
		//time.Sleep(10 * time.Second)
*/
		case "Input":
			f.WriteString("Input command read.")
			f.WriteString("\n")
			m["text"] = "Input succesful"
/*
		devID := cmd.Get("dev_id").MustString()
		regID := cmd.Get("reg_id").MustString()
		value := cmd.Get("value").MustString()

		// appID := cmd.Get("appID").MustString()

		if strings.HasPrefix(devID, "smart") {
			smartMeter.WriteValues(devID+"/"+regID, value)
		} else {
			log.Println("Invalid node id for received topic")
		}
		m["text"] = "Input succesfully added"
*/
		default:
			f.WriteString("\n")
			f.WriteString("Unrecognized input text.")
			f.WriteString("\n")
			m["text"] = "Unknown command"
	}
	return 
}

func main() {
	fmt.Println("Program started.")
	f.WriteString("\n")
	f.WriteString("Starting program.\n")
	gonode.Start(process)
	f.WriteString("\n")
	f.WriteString("Program closing.")
	f.WriteString("\n")
	fmt.Println("Program finished.")
}
