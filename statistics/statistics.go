package statistics

var Servers []Server

type Server struct {
	HostID     string `db:"hostID"`
	OwnerID    int    `db:"ownerID"`
	Visible    bool   `db:"visible"`
	AuthToken  string `db:"authtoken"`
	Hostname   string `db:"hostname"`
	Identifier string `db:"identifier"`
}
type Update struct {
	Id               int    `db:"id"`
	Frequency        int    `db:"frequency"`
	Date             int64  `db:"date"`
	HostID           string `db:"hostID"`
	Connections      string `db:"connections"`
	Cpucores         string `db:"cpucores"`
	Cpufreq          string `db:"cpufreq"`
	Cpuname          string `db:"cpuname"`
	Diskarray        string `db:"diskarray"`
	Disktotal        string `db:"disktotal"`
	Diskusage        string `db:"diskusage"`
	Filehandles      string `db:"filehandles"`
	Filehandleslimit string `db:"filehandleslimit"`
	Hostname         string `db:"hostname"`
	HostnameShort    string `db:"hostnameshort"`
	SSID             string `db:"ssid"`
	Ipv4             string `db:"ipv4"`
	Ipv4Public       string `db:"ipv4public"`
	Ipv6             string `db:"ipv6"`
	Load             string `db:"load"`
	Loadcpu          string `db:"loadcpu"`
	Loadio           string `db:"loadio"`
	Nic              string `db:"nic"`
	Osarch           string `db:"osarch"`
	Oskernel         string `db:"oskernel"`
	Osname           string `db:"osname"`
	Ping             string `db:"ping"`
	Packages         string `db:"Packages"`
	Processes        string `db:"processes"`
	Processesarray   string `db:"processesarray"`
	Ramtotal         string `db:"ramtotal"`
	Ramusage         string `db:"ramusage"`
	Rx               string `db:"rx"`
	Rxdiff           string `db:"rxdiff"`
	Sessions         string `db:"sessions"`
	Swaptotal        string `db:"swaptotal"`
	Swapusage        string `db:"swapusage"`
	Tx               string `db:"tx"`
	Txdiff           string `db:"txdiff"`
	Uptime           string `db:"uptime"`
	DockerInstalled  int    `db:"dockerinstalled"`
	DockerPS         string `db:"dockerps"`
}
