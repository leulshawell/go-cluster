package group

import (
	// "github.com/geohot/minikeyvalue"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Config struct {
	Interface          *net.Interface
	Port               int
	Key                string
	MasterSelectionRle int
}

type Group struct {
	volume string
	Name   string
	Config *Config
}

func getKey(key string) string {
	return "localhost:9000"
}

func GetGroupConfig(key string) (*Config, error) {
	//get the config string from minikeyvalue. omes in interfacename:port format
	configStr := getKey(key)

	inter_port := strings.Split(configStr, ":")
	interface_, err_ := net.InterfaceByName(inter_port[0])
	port, err := strconv.Atoi(inter_port[1])
	if err != nil || err_ != nil {
		panic(fmt.Sprintf("Invalid condfig string %s", configStr))
	}

	return &Config{
		Key:       key,
		Port:      port,
		Interface: interface_,
	}, nil

}

func (g *Group) Up() {
	//start minikeyvalue instance for group data

	//schedule

	//start start work

}

// Join the cluster
func (g *Group) Accept() error {
	return nil
}

func NewGroup(name string, interfaceName string, port int) (*Group, error) {
	inter, err := net.InterfaceByName(interfaceName)

	if err != nil {
		return nil, err
	}

	key := uuid.New().String()
	return &Group{
		Config: &Config{
			Interface: inter,
			Port:      port,
			Key:       key,
		},
		Name:   name,
		volume: key + "/" + name,
	}, nil

}

func Start() {
	//status check for nodes

	//schedule jobs from queue on free nodes

}
