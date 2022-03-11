package utils

import(
	"os"
	"log"
	"io/ioutil"
	"path/filepath"

	yaml "github.com/goccy/go-yaml"
)

func (h *HostIdentity) Load(filename string) *HostIdentity {
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

func (h *UserConfig) Load(filename string) *UserConfig {
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

func (h *Finding) Load(filename string) *Finding {
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