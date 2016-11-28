package main

import (
	"github.com/blablacar/dgr/dgr/common"
)

func (aci *Aci) Init() error {
	defer giveBackUserRights(aci.path)

	aci.Clean()
	err := aci.RunBuilderCommand(common.CommandInit)
	aci.Clean() // TODO this is a ack because init will leave an empty aci in target. should better be processed by stage1
	return err
}
