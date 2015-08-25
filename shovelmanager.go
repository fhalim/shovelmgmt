package shovelmgmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	rh "github.com/michaelklishin/rabbit-hole"
)

// CreateShovel creates the specified cluster on the clusterInfo
// PUT /api/parameters/shovel/%2f/my-shovel
func (clusterInfo ClusterInfo) CreateShovel(shovelName string, shovelDefinition ShovelDefinition) (res *http.Response, err error) {
	log.Printf("Creating shovel %v for %v->%v", shovelName, shovelDefinition.SourceQueue, shovelDefinition.DestinationQueue)
	parm := ShovelParameter{Value: shovelDefinition}
	body, err := json.Marshal(parm)
	if err != nil {
		return nil, err
	}

	req, err := clusterInfo.newRequestWithBody("PUT", "parameters/shovel/"+url.QueryEscape(clusterInfo.Vhost)+"/"+url.QueryEscape(shovelName), body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AmqpURL gives the AMQP url for the cluster.
func (clusterInfo ClusterInfo) AmqpURL() string {
	return fmt.Sprintf("amqp://%v:%v/%v?heartbeat=10", clusterInfo.HostName, clusterInfo.AmqpPort, url.QueryEscape(clusterInfo.Vhost))
}

// AdminURL gives the HTTP AdminPort url for the cluster.
func (clusterInfo ClusterInfo) AdminURL() string {
	return fmt.Sprintf("http://%v:%v", clusterInfo.HostName, clusterInfo.AdminPort)
}

// CreateQueues creates all the specified queues on the given cluster
func (clusterInfo ClusterInfo) CreateQueues(queues []rh.QueueInfo) error {
	rmqc, err := rh.NewClient(clusterInfo.AdminURL(), clusterInfo.UserName, clusterInfo.Password)
	if err != nil {
		return err
	}
	for _, queue := range queues {
		log.Printf("Creating queue %v", queue.Name)
		_, err = rmqc.DeclareQueue(clusterInfo.Vhost, queue.Name, rh.QueueSettings{Durable: queue.Durable, AutoDelete: queue.AutoDelete, Arguments: queue.Arguments})
		if err != nil {
			return err
		}
	}
	return nil
}
func (clusterInfo ClusterInfo) newRequestWithBody(method string, path string, body []byte) (*http.Request, error) {
	s := clusterInfo.AdminURL() + "/api/" + path

	req, err := http.NewRequest(method, s, bytes.NewReader(body))
	req.SetBasicAuth(clusterInfo.UserName, clusterInfo.Password)
	// set Opaque to preserve percent-encoded path. MK.
	req.URL.Opaque = clusterInfo.AdminURL() + "/api/" + path

	req.Header.Add("Content-Type", "application/json")

	return req, err
}

func executeRequest(req *http.Request) (res *http.Response, err error) {
	var httpc *http.Client

	httpc = &http.Client{}
	res, err = httpc.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
