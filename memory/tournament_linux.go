//go:build linux
// +build linux

package memory

import (
	"errors"

	"github.com/lekluge/gosumemory/mem"
)

func resolveTourneyClients(procs []mem.Process) ([]mem.Process, error) {
	return nil, errors.New("Not implemented!")
}

func getTourneyGameplayData(proc mem.Process, iterator int) {
	return
}

func getTourneyIPC() error {
	return errors.New("Not implemented!")
}
