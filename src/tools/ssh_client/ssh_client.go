package ssh_client

import (
	"io/ioutil"

	"../error"
	"golang.org/x/crypto/ssh"
)

const (
	TCP = "tcp"
)

const (
	RSA_KEY_ERROR           = "Rsa Key didn't loaded!!"
	RSA_PASS_ERROR          = "Rsa Key didn't Parse!! or One more input your password"
	SSH_DIAL_ERROR          = "Didn't established SSH Conneciton"
	SESSION_ESTABLISH_ERROR = "Didn't established New Session"
	BATCH_EXEC_ERROR        = "Didn't execute batch"
)

func Ssh_execution(server_ip, server_port, server_user,
	ssh_keypath, rsa_keypass, exec_command, exec_mode string) {
	ip := server_ip
	port := server_port
	user := server_user

	buf, err := ioutil.ReadFile(ssh_keypath)
	error.Ce(err, exec_mode+RSA_KEY_ERROR)

	signer, err := ssh.ParsePrivateKeyWithPassphrase(buf, []byte(rsa_keypass))
	error.Ce(err, exec_mode+RSA_PASS_ERROR)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial(TCP, ip+":"+port, sshConfig)
	error.Ce(err, exec_mode+SSH_DIAL_ERROR)

	defer conn.Close()

	session, err := conn.NewSession()
	error.Ce(err, exec_mode+SESSION_ESTABLISH_ERROR)
	defer session.Close()

	err = session.Run(exec_command)
	error.Ce(err, exec_mode+BATCH_EXEC_ERROR)
}
