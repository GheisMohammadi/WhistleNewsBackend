package config

const (
	// MongoURI is full url of mongo db 
	MongoURI = "127.0.0.1:27017"
	
	// DataBaseName for storing docs
	DataBaseName = "WhistleNews"

	// DatabaseUserName for connect to db
	DatabaseUserName = "mongo"

	// DatabasePassword is passowrd for db user 
	DatabasePassword = "1234567"

	// NSQURL is full address of NSQ server
	NSQURL = "127.0.0.1:4150"

	// APIPrefix is prefix at the beginning of routes like http://server_ip:port/{prefix}/... 
	APIPrefix = "/counter/v1"

	// ServerPort is port number of API server
	ServerPort = "3085"
)