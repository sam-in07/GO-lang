package main

import (
	"fmt"

	"example.com/g0_project_structure/config"
)

func main() {
	cnf :=   config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)
	//cmd.Server()

}
