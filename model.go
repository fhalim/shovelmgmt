package shovelmgmt

type ShovelDefinition struct {
	SourceUri              string `json:"src-uri,omitempty"`
	SourceQueue            string `json:"src-queue,omitempty"`
	SourceExchange         string `json:"src-exchange,omitempty"`
	SourceExchangeKey      string `json:"src-exchange-key,omitempty"`
	DestinationUri         string `json:"dest-uri,omitempty"`
	DestinationQueue       string `json:"dest-queue,omitempty"`
	DestinationExchange    string `json:"dest-exchange,omitempty"`
	DestinationExchangeKey string `json:"dest-exchange-key,omitempty"`
}

type ShovelParameter struct {
	Value ShovelDefinition `json:"value"`
}

type ClusterInfo struct {
	HostName  string
	AdminPort int
	AmqpPort  int
	UserName  string
	Password  string
	Vhost     string
}
