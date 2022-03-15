package utils

type UserConfig struct{
	ProjectOwner string `yaml:"project-owner"`
	Group string `yaml:"group"`
	Teams []Team `yaml:"teams"`
	Email string `yaml:"email"`
	DirConfig DirConfig `yaml:"dir-config"`
}

type Team struct{
	TeamName string `yaml:"team-name"`
	Members []Member `yaml:"members"`
	Desc string `yaml:"desc"`
}

type Member struct{
	MemberName string `yaml:"name"`
	Role string `yaml:"role"`
	Email string `yaml:"email"`
}

type DirConfig struct{
	HostsIdentityDir string `yaml:"hosts-identity-dir"`
	VulnDir string `yaml:"vuln-dir"`
	ToolsReports string `yaml:"tools-reports"`
	PocDir string `yaml:"poc-dir"`
	OutputReportsDir string `yaml:"output-reports"`
	TemplatesReportDir string `yaml:"template-report-dir"`
}