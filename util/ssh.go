package util

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// Connect to a service
func Connect(node *Node) (*ssh.Session, error) {
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(node.Password))
	addr := fmt.Sprintf("%s:%d", node.Host, 22)

	clientConfig := &ssh.ClientConfig{
		User:    node.User,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	var (
		client  *ssh.Client
		session *ssh.Session
		err     error
	)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
