package main

import (
	"github.com/blablacar/dgr/dgr/common"
	"github.com/n0rad/go-erlog/errs"
	"github.com/n0rad/go-erlog/logs"
)

func (p *Pod) Push() error {
	logs.WithF(p.fields).Info("Pushing")

	if err := p.CleanAndBuild(); err != nil {
		return err
	}

	for _, e := range p.manifest.Pod.Apps {
		aci, err := p.toPodAci(e)
		if err != nil {
			return err
		}

		if err := aci.Push(); err != nil {
			return err
		}
	}

	if Home.Config.Push.Type == "maven" && p.manifest.Name.DomainName() == "aci.blbl.cr" {
		// TODO this definitely need to be removed
		if err := common.ExecCmd("curl", "-i",
			"-F", "r=releases",
			"-F", "hasPom=false",
			"-F", "e=pod",
			"-F", "g=com.blablacar.aci.linux.amd64",
			"-F", "p=pod",
			"-F", "v="+p.manifest.Name.Version(),
			"-F", "a="+p.manifest.Name.ShortName(),
			"-F", "file=@"+p.target+"/pod-manifest.json",
			"-u", Home.Config.Push.Username+":"+Home.Config.Push.Password,
			Home.Config.Push.Url+"/service/local/artifact/maven/content"); err != nil {

			return errs.WithEF(err, p.fields, "Failed to push pod")
		}
	}

	return nil
}
