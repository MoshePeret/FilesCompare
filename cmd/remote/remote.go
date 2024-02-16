package remote

import (
	"FilesCompare/pkg"
	"github.com/bramvdbogaerde/go-scp"
	auth2 "github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"log"
)

func Connection(ip string, auth pkg.Auth) {
	clientConfig, _ := auth2.PasswordKey(auth.Username, auth.Password, ssh.InsecureIgnoreHostKey())

	client := scp.NewClient(ip+":22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		log.Println("Couldn't establish a connection to the remote server ", err)
		return
	} else {
		log.Println("Connection has been established successfully", err)
	}
}
