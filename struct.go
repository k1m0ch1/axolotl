package main

type Target struct {
	Url string `yaml:"url"`
	TechStack string `yaml:"tech-stack"`
	OpenPort string `yaml:"open-port"`
	HostIP string `yaml:"host-ip"`
}

type Host struct{
	Target Target `yaml:"target"`
}