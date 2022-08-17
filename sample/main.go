package main

import (
	"fmt"

	SystemD "github.com/artziel/go-systemd"
)

func main() {

	exec, _ := SystemD.GetExecutablePath()
	srv := SystemD.NewService(
		"sample-daemon",
		SystemD.Parameters{
			Unit: SystemD.UnitSection{
				Description:           "Sample Daemon",
				After:                 "mysqld.service",
				StartLimitIntervalSec: 0,
			},
			Service: SystemD.ServiceSection{
				Type:       "simple",
				Restart:    "always",
				RestartSec: 1,
				User:       "centos",
				ExecStart:  exec,
			},
			Install: SystemD.InstallSection{
				WantedBy: "multi-user.target",
			},
		},
	)

	err := srv.Install()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

}
