package openshift

import (
	"fmt"
)

// ErrOpenShiftFailedToStart is thrown when the OpenShift server failed to start
func ErrOpenShiftFailedToStart(container string) error {
	return fmt.Errorf("Could not start OpenShift container %q", container)
}

// ErrTimedOutWaitingForStart is thrown when the OpenShift server can't be pinged after reasonable
// amount of time.
func ErrTimedOutWaitingForStart(container string) error {
	return fmt.Errorf("Could not start OpenShift container %q", container)
}

type errPortsNotAvailable struct {
	ports []int
}

func (e *errPortsNotAvailable) Error() string {
	return fmt.Sprintf("ports in use: %v", e.ports)
}

func ErrPortsNotAvailable(ports []int) error {
	return &errPortsNotAvailable{
		ports: ports,
	}
}

func IsPortsNotAvailableErr(err error) bool {
	_, ok := err.(*errPortsNotAvailable)
	return ok
}

func UnavailablePorts(err error) []int {
	e, ok := err.(*errPortsNotAvailable)
	if !ok {
		return []int{}
	}
	return e.ports
}
