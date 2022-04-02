package utils

type ReportVar struct {
	Host  HostIdentity
	Vulns []Finding
	User  UserConfig
}
