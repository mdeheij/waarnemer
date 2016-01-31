package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mdeheij/monitoring/statistics"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

func statsInit(r *gin.Engine) {

	statsGroup := r.Group("/stats")
	{
		statsGroup.GET("/server/update", func(c *gin.Context) {
			c.String(400, "You should post!")
		})
		statsGroup.POST("/server/update", statsServerUpdate)
	}
	statsServerGroup := r.Group("/stats", AuthRequired())
	{
		statsServerGroup.GET("/", statsPage)
		statsServerGroup.GET("/server/list", getServers)
		statsServerGroup.POST("/server/new", newServer)
		statsServerGroup.GET("/server/view/:hostID", statsViewByHostID)
	}
}

func statsPage(c *gin.Context) {
	c.HTML(200, "stats.tmpl", gin.H{
		"title":    "Serverstat",
		"subtitle": "Easy server statistics monitoring",
	})
}
func newServer(c *gin.Context) {
	u1 := uuid.NewV4()
	c.String(200, u1.String())
}

func getServers(c *gin.Context) {

	//clear []Server cache
	statistics.Servers = []statistics.Server{}
	ownerID := 1

	_, err := dbmap.Select(&statistics.Servers, "SELECT server.hostID, server.ownerID, server.visible, server.identifier, IFNULL(serverupdate.hostname,server.hostID) AS hostname FROM server LEFT JOIN serverupdate ON server.hostID=serverupdate.hostID WHERE ownerID = ? GROUP BY server.hostID ORDER BY server.identifier DESC", ownerID)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, statistics.Servers)
}

func getIdentifier(HostID string) string {
	for _, item := range statistics.Servers {
		//Only when it matches
		if item.HostID == HostID {
			return item.Identifier
		}

	}
	return "unknown-identifier"
}

func statsViewByHostID(c *gin.Context) {
	hostID := c.Param("hostID")

	recentUpdates := []statistics.Update{}

	_, err := dbmap.Select(&recentUpdates, "SELECT * FROM serverupdate WHERE hostID = ? ORDER BY `serverupdate`.`id` DESC LIMIT 25", hostID)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"identifier": getIdentifier(hostID),
		"hostID":     hostID,
		"obj":        recentUpdates,
	})
}

func statsServerUpdate(c *gin.Context) {
	postedUpdate := statistics.Update{}
	postedUpdate.HostID = base64Decode(c.PostForm("hostID"))

	var datetime = time.Now()

	postedUpdate.Date = datetime.Unix()

	postedUpdate.Frequency, _ = strconv.Atoi(c.PostForm("frequency"))
	postedUpdate.Connections = base64Decode(c.PostForm("connections"))
	postedUpdate.Cpucores = base64Decode(c.PostForm("cpucores"))
	postedUpdate.Cpufreq = base64Decode(c.PostForm("cpufreq"))
	postedUpdate.Cpuname = base64Decode(c.PostForm("cpuname"))
	postedUpdate.DockerInstalled, _ = strconv.Atoi(c.PostForm("dockerinstalled"))
	postedUpdate.DockerPS = base64Decode(c.PostForm("dockerps"))
	postedUpdate.Diskarray = base64Decode(c.PostForm("diskarray"))
	postedUpdate.Disktotal = base64Decode(c.PostForm("disktotal"))
	postedUpdate.Diskusage = base64Decode(c.PostForm("diskusage"))
	postedUpdate.Filehandles = base64Decode(c.PostForm("filehandles"))
	postedUpdate.Filehandleslimit = base64Decode(c.PostForm("filehandleslimit"))
	postedUpdate.Hostname = base64Decode(c.PostForm("hostname"))
	postedUpdate.HostnameShort = base64Decode(c.PostForm("hostnameshort"))
	postedUpdate.SSID = base64Decode(c.PostForm("ssid"))
	postedUpdate.Ipv4 = base64Decode(c.PostForm("ipv4"))
	postedUpdate.Ipv4Public = base64Decode(c.PostForm("ipv4public"))
	postedUpdate.Ipv6 = base64Decode(c.PostForm("ipv6"))
	postedUpdate.Load = base64Decode(c.PostForm("load"))
	postedUpdate.Loadcpu = base64Decode(c.PostForm("loadcpu"))
	postedUpdate.Loadio = base64Decode(c.PostForm("loadio"))
	postedUpdate.Nic = base64Decode(c.PostForm("nic"))
	postedUpdate.Osarch = base64Decode(c.PostForm("osarch"))
	postedUpdate.Oskernel = base64Decode(c.PostForm("oskernel"))
	postedUpdate.Osname = base64Decode(c.PostForm("osname"))
	postedUpdate.Ping = base64Decode(c.PostForm("ping"))
	postedUpdate.Packages = base64Decode(c.PostForm("packages"))
	postedUpdate.Processes = base64Decode(c.PostForm("processes"))
	postedUpdate.Processesarray = base64Decode(c.PostForm("processesarray"))
	postedUpdate.Ramtotal = base64Decode(c.PostForm("ramtotal"))
	postedUpdate.Ramusage = base64Decode(c.PostForm("ramusage"))
	postedUpdate.Rx = base64Decode(c.PostForm("rx"))
	postedUpdate.Rxdiff = base64Decode(c.PostForm("rxdiff"))
	postedUpdate.Sessions = base64Decode(c.PostForm("sessions"))
	postedUpdate.Swaptotal = base64Decode(c.PostForm("swaptotal"))
	postedUpdate.Swapusage = base64Decode(c.PostForm("swapusage"))
	postedUpdate.Tx = base64Decode(c.PostForm("tx"))
	postedUpdate.Txdiff = base64Decode(c.PostForm("txdiff"))
	postedUpdate.Uptime = base64Decode(c.PostForm("uptime"))

	err := dbmap.Insert(&postedUpdate)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"status": "received",
		"update": postedUpdate,
	})
}
