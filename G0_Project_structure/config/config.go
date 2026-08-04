package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configarations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
    
    err := godotenv.Load()
	if err!= nil{
		fmt.Println("failde to load the env varibale ")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VBersion is req ")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Servise name is Requeired ")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64) //func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)

	if err != nil {
		fmt.Println("port number needed")
		os.Exit(1)
	}

	configarations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configarations
}
