package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Cluster struct {
	config        Config `yaml:"cluster"`
	cluserDatadir string `yaml:"cluster_data_dir"`
}

type Config struct {
	Name      string `yaml:"name"`
	Port      string `yaml:"port"`
	Interface string `yaml:"interface"`
}

func ReadConig(path string) {
	configString := make([]byte, 64)

	var cfg Config

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	file.Read(configString)
	// configString = []byte{
	// 	`name: cluster
	// 	port: 8080
	// 	interface: wlp2s0`,

	// }
	// fmt.Println(string(configString))

	err = yaml.Unmarshal([]byte(configString), &cfg)
	fmt.Printf("%s %s", string(configString), err.Error())

}

func main() {
	ReadConig("./go-cluster.yaml")
	return
}
