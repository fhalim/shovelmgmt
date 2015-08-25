package main

import (
	"flag"
	"log"

	"github.com/fhalim/shovelmgmt"
)

func main() {
	clusterInfo := readFlags()
	autoShovels, err := clusterInfo.ListAutoShovels()
	if err != nil {
		log.Fatal(err)
	}

	for _, shovel := range autoShovels {
		log.Println(shovel.Name)
		clusterInfo.DeleteShovel(shovel)
	}
}

func readFlags() shovelmgmt.ClusterInfo {
	host := flag.String("host", "localhost", "")
	adminport := flag.Int("adminport", 15674, "")
	username := flag.String("username", "guest", "")
	password := flag.String("password", "guest", "")
	vhost := flag.String("vhost", "/", "")

	flag.Parse()

	clusterInfo := shovelmgmt.ClusterInfo{
		HostName:  *host,
		UserName:  *username,
		Password:  *password,
		AdminPort: *adminport,
		Vhost:     *vhost,
	}

	return clusterInfo
}
