package main

import (
	"flag"
	"log"

	"github.com/fhalim/shovelmgmt"

	rh "github.com/michaelklishin/rabbit-hole"
)

func main() {
	upstreamClusterInfo, downstreamClusterInfo := readFlags()

	log.Println("Getting list of queues from upstream cluster...")
	upstreamQueues, err := getQueues(upstreamClusterInfo)

	if err != nil {
		log.Panic(err)
	}
	log.Println("Creating queues in downstream cluster...")
	err = downstreamClusterInfo.CreateQueues(upstreamQueues)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Creating shovels in downstream cluster...")
	for _, queue := range upstreamQueues {
		shovelDefinition := shovelmgmt.ShovelDefinition{SourceUri: upstreamClusterInfo.AmqpURL(), SourceQueue: queue.Name, DestinationUri: downstreamClusterInfo.AmqpURL(), DestinationQueue: queue.Name}
		_, err := downstreamClusterInfo.CreateAutoShovel(shovelDefinition)
		if err != nil {
			log.Panic(err)
		}
	}

}
func readFlags() (shovelmgmt.ClusterInfo, shovelmgmt.ClusterInfo) {
	upstreamhost := flag.String("upstreamhost", "localhost", "")
	upstreamadminport := flag.Int("upstreamadminport", 15673, "")
	upstreamamqpport := flag.Int("upstreamamqpport", 5673, "")
	upstreamusername := flag.String("upstreamusername", "guest", "")
	upstreampassword := flag.String("upstreampassword", "guest", "")
	upstreamvhost := flag.String("upstreamvhost", "/", "")

	downstreamhost := flag.String("downstreamhost", "localhost", "")
	downstreamadminport := flag.Int("downstreamadminport", 15674, "")
	downstreamamqpport := flag.Int("downstreamamqpport", 5674, "")
	downstreamusername := flag.String("downstreamusername", "guest", "")
	downstreampassword := flag.String("downstreampassword", "guest", "")
	downstreamvhost := flag.String("downstreamvhost", "/", "")

	flag.Parse()

	upstreamClusterInfo := shovelmgmt.ClusterInfo{
		HostName:  *upstreamhost,
		UserName:  *upstreamusername,
		Password:  *upstreampassword,
		AdminPort: *upstreamadminport,
		AmqpPort:  *upstreamamqpport,
		Vhost:     *upstreamvhost,
	}

	downstreamClusterInfo := shovelmgmt.ClusterInfo{
		HostName:  *downstreamhost,
		UserName:  *downstreamusername,
		Password:  *downstreampassword,
		AdminPort: *downstreamadminport,
		AmqpPort:  *downstreamamqpport,
		Vhost:     *downstreamvhost,
	}

	return upstreamClusterInfo, downstreamClusterInfo
}
func getQueues(clusterInfo shovelmgmt.ClusterInfo) ([]rh.QueueInfo, error) {
	rmqc, err := rh.NewClient(clusterInfo.AdminURL(), clusterInfo.UserName, clusterInfo.Password)
	if err != nil {
		return nil, err
	}
	queues, err := rmqc.ListQueuesIn(clusterInfo.Vhost)
	if err != nil {
		return nil, err
	}
	return queues, err
}
