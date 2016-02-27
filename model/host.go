package model

type Host struct {
	HostId                string `json: "host_id,omitempty" gorethink:"host_id,omitempty"`
	BitType               string `json: "bit_type" gorethink: "bit_type"`
	EnvironmentId         string `json: "environment_id" gorethink: "environment_id"`
	OrganizationId        string `json: "organization_id" gorethink: "organization_id"`
	HostName              string `json: "host_name" gorethink:"host_name"`
	OSDistribution        string `json: "os_distribution" gorethink: "os_distribution"`
	OSDistributionRelease string `json: "os_distribution_release" gorethink: "os_distribution_release"`
	OSDistributionVersion string `json: "os_distribution_version" gorethink: "os_distribution_version"`
	OSFamily              string `json: "os_family" gorethink: "os_family"`
	SystemType            string `json: "system_type" gorethink: "system_type"`
}
