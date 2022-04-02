package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"unicode"

	yaml "github.com/goccy/go-yaml"
)

type App struct {
	Version string
}

type Pandoc []string

func (h *HostIdentity) Load(filename string) *HostIdentity {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, h)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return h
}

func (h *UserConfig) Load(filename string) *UserConfig {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		log.Printf("[?] Warning! File User Config not found\n")
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\n[?] Do you want me to create the `config.yml` file ? (y/N): ")
		text, _, _ := reader.ReadRune()
		resultText := fmt.Sprintf("%c", unicode.ToLower(text))
		if resultText == "n" {
			log.Println("[+] Well ok")
		} else if resultText == "y" {
			var uC UserConfig
			var name string
			fmt.Println("[!] Generate the template config")
			fmt.Printf("\n[+] Tell me your nick/name: ")
			fmt.Scanln(&name)
			err := uC.GenerateConfig(name)
			if err != nil {
				fmt.Println(err)
			}

			uC.Load("./config.yml")

			dirs := []string{
				uC.DirConfig.HostsIdentityDir,
				uC.DirConfig.VulnDir,
				uC.DirConfig.ToolsReports,
				uC.DirConfig.PocDir,
				uC.DirConfig.OutputReportsDir,
				uC.DirConfig.TemplatesReportDir,
			}

			for _, v := range dirs {
				path := fmt.Sprintf("./%s", v)
				_, err := CheckDirAndCreate(path)
				if err != nil {
					fmt.Println("[X] Weird error occured ", err)
				}
			}
		}
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, h)
	if err != nil {
		log.Println("Error to get value from config.yml")
		log.Fatalf("Unmarshal: %v", err)
	}
	return h
}

func (h *Finding) Load(filename string) *Finding {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, h)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return h
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func CheckDirAndCreate(paths string) (bool, error) {
	if _, err := os.Stat(paths); os.IsNotExist(err) {
		fmt.Printf("\n[?] The dir %s is not yet created, I will create this", paths)
		err := os.Mkdir(paths, 0644)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
