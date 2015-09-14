package main

import (
	"fmt"
	"github.com/bgentry/speakeasy"
	"github.com/jfroche/vCloudOSLicenses"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	username = kingpin.Flag("username", "Name of user.").Short('u').Required().String()
	password = kingpin.Flag("password", "Password of user.").Short('p').String()
	host     = kingpin.Flag("host", "VCloud host.").Short('h').Required().String()
	vdc      = kingpin.Flag("vdc", "VCloud vdc").Short('c').Required().String()
)

func main() {
	kingpin.Parse()
	password := *password
	if password == "" {
		password, _ = speakeasy.Ask("Password: ")
	}
	url := fmt.Sprint("https://", *host)
	session := &vcloudoslicenses.VCloudSession{
		Host:     url,
		Username: *username,
		Password: password,
		Context:  *vdc,
	}
	session.Login()
	var VMs []*vcloudoslicenses.VAppVm
	VMs, _ = session.FindVMs(30, 1)
	for _, vm := range VMs {
		fmt.Printf("%s: %s \n", vm.Name, vm.NetworkConnectionSection.NetworkConnections[0].IpAddress)
		fmt.Printf("  metadata: %s\n", vm.Metadata)
	}
}
