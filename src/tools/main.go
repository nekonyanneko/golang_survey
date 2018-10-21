package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"./args_parser"
	"./batch"
	"./error"
	"./ssh_client"
	"github.com/robfig/cron"
)

const (
	MYAPPLICATION_NAME = "[BATCH or ADHOC]"
	ADHOC_SUB_COMMAND  = "adhoc"
	BATCH_SUB_COMMAND  = "batch"
	CRON_TIME          = "* * * * * *"
)

func main() {
	log.Println("****" + MYAPPLICATION_NAME + "****")
	subcommand := os.Args[1]

	switch {
	case subcommand == ADHOC_SUB_COMMAND:
		flags := flag.NewFlagSet(ADHOC_SUB_COMMAND, flag.ExitOnError)
		args := args_parser.Parser_Var{}
		args.Parser_call(flags, os.Args[2:])

		ssh_client.Ssh_execution(
			args.Ip_address, args.Port, args.Server_user,
			args.Ssh_keypath, args.Rsa_Password, args.Command,
			"["+ADHOC_SUB_COMMAND+"]:")

	case subcommand == BATCH_SUB_COMMAND:
		batch_config := batch.Config{}
		batch_config.Config_load()

		c := cron.New()
		err := c.AddFunc(
			CRON_TIME,
			batch_config.Batch_exec)
		error.Ce(err, "Cron is Failed")
		c.Start()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println("input is", scanner.Text())

		c.Stop()

	default:
		fmt.Println("sub command please....")
	}
}
