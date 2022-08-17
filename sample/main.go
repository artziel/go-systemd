package main

import (
	"fmt"
	"os"
	"path/filepath"

	SystemD "github.com/artziel/go-systemd"
)

func getExecutablePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return dir + "/" + filepath.Base(os.Args[0]), err
}

func main() {

	exec, _ := getExecutablePath()
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
