package GoSystemD

import (
	"fmt"
	"os"

	Util "github.com/artziel/go-utilities"
)

const parametersFormat = `[Unit]
Description=%s
After=%s
StartLimitIntervalSec=%d

[Service]
Type=%s
Restart=%s
RestartSec=%d
User=%s
ExecStart=%s

[Install]
WantedBy=%s
`

type UnitSection struct {
	Description           string
	After                 string
	StartLimitIntervalSec int
}
type ServiceSection struct {
	Type       string
	Restart    string
	RestartSec int
	User       string
	ExecStart  string
}
type InstallSection struct {
	WantedBy string
}
type Parameters struct {
	Unit    UnitSection
	Service ServiceSection
	Install InstallSection
}

func (p *Parameters) toString() string {
	result := ""

	result = fmt.Sprintf(
		parametersFormat,
		p.Unit.Description,
		p.Unit.After,
		p.Unit.StartLimitIntervalSec,
		p.Service.Type,
		p.Service.Restart,
		p.Service.RestartSec,
		p.Service.User,
		p.Service.ExecStart,
		p.Install.WantedBy,
	)

	return result
}

type Service struct {
	Name       string
	Params     Parameters
	BinaryPath string
}

func (s *Service) Install() error {
	ok, err := Util.IsRoot()
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotRootUser
	}

	if s.isInstalled() {
		return ErrServiceIsInstalled
	}

	err = Util.SaveToNewTXTFile("/etc/systemd/system/"+s.Name+".service", s.Params.toString())
	if err != nil {
		return err
	}

	return err
}

func (s *Service) isInstalled() bool {
	file := "/etc/systemd/system/" + s.Name + ".service"
	if _, err := os.Stat(file); err != nil {
		return false
	}

	return true
}

func (s *Service) Uninstall() error {
	ok, err := Util.IsRoot()
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotRootUser
	}

	if !s.isInstalled() {
		return ErrServiceIsNotInstalled
	}

	file := "/etc/systemd/system/" + s.Name + ".service"

	return os.Remove(file)
}
