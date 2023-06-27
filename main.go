package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	NewRegistry string   `yaml:"newRegistry"`
	Images      []string `yaml:"images"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("please provide a config file path")
		os.Exit(2)
	}

	var config Config
	configFile, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Printf("couldn't read config file: %v\n", err)
		os.Exit(2)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("couldn't unmarshal config file: %v\n", err)
		os.Exit(2)
	}

	newRegistry := config.NewRegistry

	for _, image := range config.Images {
		imageParts := strings.SplitN(image, ":", 2)
		imageName := imageParts[0]
		imageTag := "latest"
		if len(imageParts) > 1 {
			imageTag = imageParts[1]
		}

		newImage := fmt.Sprintf("%s/%s:%s", newRegistry, imageName, imageTag)

		err := execCommand("docker", "pull", image)
		if err != nil {
			fmt.Printf("\n-----------> Error pulling %s: %v\n", image, err)
			continue
		}

		err = execCommand("docker", "tag", image, newImage)
		if err != nil {
			fmt.Printf("\n-----------> Error tagging %s: %v\n", newImage, err)
			continue
		}

		err = execCommand("docker", "push", newImage)
		if err != nil {
			fmt.Printf("--------Error pushing %s: %v\n", newImage, err)
			continue
		}

		fmt.Printf("-----------> Pushed %s 🎉\n\n", newImage)
	}
}

func execCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
