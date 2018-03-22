package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
)

func main() {

	// init
	fmt.Println("Incrementing...")

	// key setting
	key := ""
	if os.Getenv("INCREMENT_KEY") != "" {
		key = os.Getenv("INCREMENT_KEY")
	} else {
		panic("Error: no INCREMENT_KEY environment variable")
	}

	// add file
	addFile := "/local/add"
	if len(os.Args) > 1 {
		addFile = os.Args[1]
	}

	// increment setting
	add := 1
	if fileExists(addFile) {
		file, err := ioutil.ReadFile(addFile)
		if err != nil {
			panic("Error: cannot read file")
		}
		add, err = strconv.Atoi(strings.TrimSpace(string(file[:])))
		if err != nil {
			panic("Error: add file not a number")
		}
	}

	// consul client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// key value client
	kv := client.KV()

	// lookup value
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}

	// convert value to int
	i, err := strconv.Atoi(string(pair.Value[:]))
	if err != nil {
		panic(err)
	}

	// increment int
	i = i + add

	// set incremented value
	p := &api.KVPair{Key: key, Value: []byte(strconv.Itoa(i))}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
