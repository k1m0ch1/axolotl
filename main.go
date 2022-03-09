package main

import(
	"log"
	"fmt"
	"io/ioutil"

	yaml "github.com/goccy/go-yaml"
)

func main(){
	files, err := ioutil.ReadDir("hosts")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		var h Host
		fileName := fmt.Sprintf("./hosts/%s", f.Name())
		h.LoadFile(fileName)
		fmt.Println("Host: ", h.Target.Url)
		fmt.Println("TechStack: ", h.Target.TechStack)
	}
}

func (h *Host) LoadFile(filename string) *Host {
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, h);
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return h
}