package db

import (
	"fmt"
	"net/url"
)

type DbConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func (c DbConfig) toDatabaseUrl() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}
	return u.String()
}
