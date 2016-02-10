package builder

import (
	"github.com/n0rad/go-erlog/logs"
	"os"
)

func (p *Pod) Clean() {
	logs.WithF(p.fields).Info("Cleaning")

	if err := os.RemoveAll(p.target + "/"); err != nil {
		panic("Cannot clean" + p.manifest.Name.String() + err.Error())
	}

	for _, e := range p.manifest.Pod.Apps {
		aci, err := NewAciWithManifest(p.path+"/"+e.Name, p.args, p.toAciManifest(e))
		if err != nil {
			logs.WithEF(err, p.fields).WithField("name", e.Name).Fatal("Cannot prepare aci")
		}
		aci.podName = &p.manifest.Name
		aci.Clean()
	}

}
