package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}


func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	a := authenticationInfo{
		username: "Arijit Roy",
		password: "123213123213",
	}

	fmt.Println(a.getBasicAuth())
}