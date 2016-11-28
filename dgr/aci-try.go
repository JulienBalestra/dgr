package main

import "github.com/blablacar/dgr/dgr/common"

func (aci *Aci) CleanAndTry() error {
	aci.Clean()
	return aci.RunBuilderCommand(common.CommandTry)
}
