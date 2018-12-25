package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ListOSTools() {
	fmt.Println()
	fmt.Println("**************** env ****************")
	for key, val := range os.Environ() {
		fmt.Println(string(key)+`:       `, string(val))
	}
	fmt.Println("session name:       ", os.ExpandEnv(`${SESSIONNAME}`))
	fmt.Println("process id:         ", os.Getpid())
	process, err := os.FindProcess(os.Getpid())
	check(err)
	out, err := json.Marshal(process)
	check(err)
	fmt.Printf("current process:     %v", string(out))
	info, err := os.Stat("./settings.yml")
	check(err)
	infoStruct, err := json.Marshal(info)
	check(err)
	fmt.Printf("\nos Signal:           %v", string(infoStruct))
	fmt.Println()
}
