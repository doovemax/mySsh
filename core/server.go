package core

import (
	"errors"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Server struct {
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Method   string `json:"method"`
	Key      string `json:"key"`
}

func parseAuthMethods(config string, server *Server) ([]ssh.AuthMethod, error) {
	sshs := []ssh.AuthMethod{}

	switch server.Method {
	case "password":
		sshs = append(sshs, ssh.Password(server.Password))
	case "pem":
		method, err := pemKey(config, server)
		if err != nil {
			return nil, err
		}
		sshs = append(sshs, method)
	default:
		return nil, errors.New("无效的密码方式： " + server.Method)
	}
	return sshs, nil

}

func pemKey(config string, server *Server) (ssh.AuthMethod, error) {
	if !filepath.IsAbs(server.Key) {
		server.Key = config + "/" + server.Key
	}
	//fmt.Println(server.Key)
	pemBytes, err := ioutil.ReadFile(server.Key)
	if err != nil {
		return nil, err
	}

	var signer ssh.Signer
	if server.Password == "" {
		signer, err = ssh.ParsePrivateKey(pemBytes)
	} else {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(server.Password))
	}
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signer), nil
}

func (server *Server) Connection(conf string) {
	auths, err := parseAuthMethods(conf, server)
	if err != nil {
		Printer.Errorln("auth error: ", err)
	}
	config := &ssh.ClientConfig{
		User: server.User,
		Auth: auths,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: time.Second * 60,
	}
	addr := server.Ip + ":" + strconv.Itoa(server.Port)
	//client, err := ssh.Dial("tcp", addr, config)
	//if err != nil {
	//	Printer.Errorln("建立连接出错： ", err)
	//	return
	//}

	tcpConn, err := net.DialTimeout("tcp", addr, config.Timeout)
	if err != nil {
		Printer.Errorln(err)
	}
	err = tcpConn.(*net.TCPConn).SetKeepAlivePeriod(time.Second * 10)
	if err != nil {
		Printer.Errorln(err)
	}
	err= tcpConn.(*net.TCPConn).SetKeepAlive(true)
    if err != nil {
        Printer.Errorln(err)
    }

	clientConn, chans, reqs, err := ssh.NewClientConn(tcpConn, addr, config)

	client := ssh.NewClient(clientConn, chans, reqs)

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		Printer.Errorln("创建Session出错： ", err)
		return
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		Printer.Errorln("创建文件描述符出错： ", err)
		return
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		Printer.Errorln("获取窗口高度出错： ", err)
		return
	}
	defer terminal.Restore(fd, oldState)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		Printer.Errorln("创建终端出错： ", err)
		return
	}

	err = session.Shell()
	if err != nil {
		Printer.Errorln("执行Shell出错： ", err)
		return
	}

	//定时发送空串,防止 ssh 自动断开
	ticker := time.NewTicker(time.Second * 10)
	go func(ticker *time.Ticker, session *ssh.Session) {
		for _ = range ticker.C {
			session.Run(" ")
		}
	}(ticker, session)
	err = session.Wait()
	if err != nil {
		Printer.Errorln("执行Wait出错： ", err)
		return
	}
	//os.Exit(0)
}
