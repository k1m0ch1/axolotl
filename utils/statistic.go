package utils

type Stats struct {
	VulnType []VulnType
}

type VulnType struct {
	Type       string
	ListOfVuln []string
}
