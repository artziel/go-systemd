package GoSystemD

import "os/exec"

func systemdExists() bool {
	_, err := exec.LookPath("systemctl")
	return err == nil
}

func validateParams(pms *Parameters) error {
	if pms.Service.Type == "" {
		pms.Service.Type = "simple"
	}
	if pms.Service.Restart == "" {
		pms.Service.Restart = "always"
	}
	if pms.Service.RestartSec == 0 {
		pms.Service.RestartSec = 1
	}
	if pms.Service.User == "" {
		pms.Service.User = "root"
	}
	if pms.Install.WantedBy == "" {
		pms.Install.WantedBy = "multi-user.target"
	}

	var err error
	if pms.Service.ExecStart == "" {
		err = ErrServiceServiceExecStartIsRequired
	}
	if pms.Unit.Description == "" {
		err = ErrServiceUnitDescriptionIsRequired
	}
	if pms.Name == "" {
		err = ErrServiceNameIsRequired
	}

	return err
}

func NewService(prms Parameters) (Service, error) {
	var err error
	srv := Service{
		Params: prms,
	}
	err = validateParams(&prms)

	if systemdExists() {
		err = ErrSystemCtlCommandNotFound
	}

	return srv, err
}
