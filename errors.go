package GoSystemD

import "errors"

var ErrUnableToGetUser error = errors.New("unable to get current user")
var ErrNotRootUser error = errors.New("current user is not root")
var ErrServiceIsNotInstalled error = errors.New("systemd service entry not found")
var ErrServiceIsInstalled error = errors.New("systemd service entry exist")
var ErrSystemCtlCommandNotFound error = errors.New("systemctl command not found")

var ErrServiceUnitDescriptionIsRequired error = errors.New("parameter Unit.Description is required")
var ErrServiceServiceExecStartIsRequired error = errors.New("parameter Service.ExecStart is required")
var ErrServiceNameIsRequired error = errors.New("parameter name is required")
