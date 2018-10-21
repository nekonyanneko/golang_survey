package args_parser

import "flag"

type Parser_Var struct {
	Ip_address   string
	Port         string
	Ssh_keypath  string
	Rsa_Password string
	Server_user  string
	Command      string
}

const (
	IP_ADDRESS_HELP_MESSAGE          = "server ip address"
	SSH_PORT_HELP_MESSAGE            = "ssh port number"
	SSH_KEY_PATH_HELP_MESSAGE        = "your rsa key path"
	RSA_KEY_PASS_HELP_MESSAGE        = "your rsa key password"
	SERVER_USER_HELP_MESSAGE         = "server's user"
	SERVER_EXEC_COMMAND_HELP_MESSAGE = "remote code execution in server"
)

func (args *Parser_Var) Parser_call(flags *flag.FlagSet, option []string) {
	flags.StringVar(&args.Ip_address, "ip", "203.0.113.10", IP_ADDRESS_HELP_MESSAGE)
	flags.StringVar(&args.Port, "port", "22", SSH_PORT_HELP_MESSAGE)
	flags.StringVar(&args.Ssh_keypath, "key_path", "~/.ssh/id_rsa", SSH_KEY_PATH_HELP_MESSAGE)
	flags.StringVar(&args.Rsa_Password, "keypass", "password", RSA_KEY_PASS_HELP_MESSAGE)
	flags.StringVar(&args.Server_user, "user", "user name", SERVER_USER_HELP_MESSAGE)
	flags.StringVar(&args.Command, "command", "echo", SERVER_EXEC_COMMAND_HELP_MESSAGE)
	flags.Parse(option)
}
