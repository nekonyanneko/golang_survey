package batch

import (
	"encoding/json"
  "fmt"
	"log"
  "io/ioutil"
	"../error"
	"../ssh_client"
)

const (
	CONFIG_PATH = "YOUR_PATH/src/tools/batch/config.json"
)

type Config struct {
	Ssh_user   Ssh_user
	Server     Server
	Batch_list []Batch
}

type Ssh_user struct {
	User    string `json:"user"`
	Rsapath string `json:"keypath"`
	Rsapass string `json:"rsapass"`
}

type Server struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Batch struct {
	Batch_name string `json:"batch_name"`
	Command    string `json:"command"`
}

func (config *Config) Config_load(){
	file, err := ioutil.ReadFile(CONFIG_PATH)
	error.Ce(err, "Didn't open Config File!!")
	error.Ce(json.Unmarshal(file, &config), "Didn't parse in Config File!!")

log.Println("**** loaded config file ****")
	fmt.Printf("ssh user is :%s\n", config.Ssh_user.User)
	fmt.Printf("ssh rsapath is :%s\n", config.Ssh_user.Rsapath)
	fmt.Printf("ssh rsapass is :%s\n", config.Ssh_user.Rsapass)
	fmt.Printf("server ip is :%s\n", config.Server.Ip)
	fmt.Printf("server port is :%s\n", config.Server.Port)
	for b, v := range config.Batch_list {
		fmt.Printf("Batch %d\n", b)
		fmt.Printf("  batch_name is %s\n", v.Batch_name)
		fmt.Printf("  command is %s\n", v.Command)
	}
}

func (config *Config) Batch_exec(){
	log.Printf("**** batch start *****")
	for _, v:= range config.Batch_list{
		ssh_client.Ssh_execution(config.Server.Ip, config.Server.Port,
		config.Ssh_user.User, config.Ssh_user.Rsapath, config.Ssh_user.Rsapass,
		v.Command, "["+v.Batch_name+"]:")
	}
}
