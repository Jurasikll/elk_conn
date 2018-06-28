package rest_db_conn

import (
	"database/sql"
	"fmt"
	"io"
	"net"

	tomb "gopkg.in/tomb.v2"
	//	"os"

	//	re "temp_change_encode/restya_entities"

	"golang.org/x/crypto/ssh"
	//	"golang.org/x/crypto/ssh/agent"

	_ "github.com/lib/pq"
)

type Rest_conn struct {
	*sql.DB
	tunn *SSHtunnel
}

func Init(remote_tunn_user string, remote_tunn_pwd string, local *Endpoint, server *Endpoint, remote *Endpoint, host string, port int, user string, pwd string, db_name string) *Rest_conn {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, db_name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sshConfig := &ssh.ClientConfig{
		User: remote_tunn_user,
		Auth: []ssh.AuthMethod{
			ssh.Password(remote_tunn_pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	tunn := &SSHtunnel{
		Config: sshConfig,
		Local:  local,
		Server: server,
		Remote: remote,
	}
	res := &Rest_conn{db, tunn}
	res.tunn.t.Go(res.tunn.Start)
	//	go res.tunn.Start()
	return res
}

type Endpoint struct {
	Host string
	Port int
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

type SSHtunnel struct {
	Local  *Endpoint
	Server *Endpoint
	Remote *Endpoint
	t      tomb.Tomb

	Config *ssh.ClientConfig
}

func (tunnel *SSHtunnel) Start() error {
	i := 1
	//	p := 1000
	listener, err := net.Listen("tcp", tunnel.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		fmt.Printf("%d\r\n", i)
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		tunnel.forward(conn)
		select {
		case <-tunnel.t.Dying():
			fmt.Println("End goroutine")
			return nil
		}

		i = i + 1

		fmt.Println("ASFDSGFWEOPKFW")
	}

	return nil
}

func (tunnel *SSHtunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

func (rc *Rest_conn) Close_rc() error {
	fmt.Println("close")
	rc.Close()
	rc.tunn.t.Kill(nil)
	return rc.tunn.t.Wait()
}
