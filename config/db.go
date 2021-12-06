package config

import (
	"net"
	"net/url"
	"os"
	"strings"
)

// DBConnectionInfo is a struct of db info
type DBConnectionInfo struct {
	Name string
	User string
	Pass string
	HOST string
	PORT string
}

// GetDBConnectionInfo returns DBConnectionInfo
func GetDBConnectionInfo() DBConnectionInfo {
	dbURL := os.Getenv("DATABASE_URL")

	u, err := url.Parse(dbURL)
	if err != nil {
		panic(err)
	}

	p, _ := u.User.Password()
	host, port, _ := net.SplitHostPort(u.Host)
	dbName := strings.Replace(u.Path, "/", "", -1)

	return DBConnectionInfo{
		Name: dbName,
		User: u.User.Username(),
		Pass: p,
		HOST: host,
		PORT: port,
	}
}
