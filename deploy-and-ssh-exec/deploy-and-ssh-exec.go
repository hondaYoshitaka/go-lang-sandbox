package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"bytes"
	"fmt"
	"time"
)

func main() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("root"),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", "0.0.0.0:32772", config)
	if err != nil {
		log.Fatalf("[FAILED] %s", err)
	}
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("[FAILED] %s", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("gooo version"); err != nil {
		log.Fatalf("[FAILED] %s", err.Error())
	}
	fmt.Println(b.String())
}
