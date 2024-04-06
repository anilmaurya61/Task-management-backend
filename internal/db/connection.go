package db

import (

)

const uri = ""
func ConnectDB() {
	serverAPI := options.serverAPI(options.serverAPIVersions1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
}