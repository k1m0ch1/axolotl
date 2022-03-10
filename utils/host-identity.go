package utils

type HostIdentity struct{
	ID string `yaml:"id"`
	Info HostInfo `yaml:"info"`
	Recons []Recon `yaml:"recons"`
	VulnScans []VulnScan `yaml:"vuln-scans"`
}

type Image struct {
	Path string `json:"path"`
	Caption string `json:"caption"`
}

type Recon struct {
	Tools string `yaml:"tools"`
	Type string `yaml:"type"`
	Report string `yaml:"report"`
	FileType string `yaml:"file-type"` // STDOUT, YAML, XML, JSON
	Desc string `yaml:"desc"`
	Images []Image `yaml:"images"`
}

type VulnScan struct {
	Tools string `yaml:"tools"`
	Type string `yaml:"type"`
	Report string `yaml:"report"`
	FileType string `yaml:"file-type"` // STDOUT, YAML, XML, JSON
	Desc string `yaml:"desc"`
	Images []Image `yaml:"images"`
}

type Other struct{
	Source string `yaml:"source"`
	Desc string `yaml:"desc"`
	URL string `yaml:"url"`
}

type HostInfo struct {
	URL string `yaml:"url"`
	TechStacks string `yaml:"tech-stacks"`
	Tag string `yaml:"tag"`
	OpenPorts string `yaml:"open-ports"`
	FilteredPort string `yaml:"filtered-port"`
	HostIP string `yaml:"host-ip"`
	Country string `yaml:"country"`
	City string `yaml:"city"`
	Organization string `yaml:"organization"`
	Others []Other `yaml:"others"`
	Desc string `yaml:"desc"`
}