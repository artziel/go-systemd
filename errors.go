package GoSystemD

import "errors"

var ErrUnableToGetUser error = errors.New("unable to get current user")
var ErrNotRootUser error = errors.New("current user is not root")
var ErrServiceIsNotInstalled error = errors.New("systemd service entry not found")
var ErrServiceIsInstalled error = errors.New("systemd service entry exist")
