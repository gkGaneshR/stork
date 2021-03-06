package schedops

import (
	"time"

	"github.com/libopenstorage/openstorage/api"
	"github.com/portworx/torpedo/drivers/node"
	"github.com/portworx/torpedo/drivers/volume"
	"github.com/portworx/torpedo/pkg/errors"
)

const (
	// DcosServiceName is the portworx service name in DC/OS
	DcosServiceName = "portworx.service"
)

type dcosSchedOps struct{}

func (d *dcosSchedOps) DisableOnNode(n node.Node, ndriver node.Driver) error {
	options := node.SystemctlOpts{
		ConnectionOpts: node.ConnectionOpts{
			Timeout:         2 * time.Minute,
			TimeBeforeRetry: 20 * time.Second,
		},
		Action: "stop",
	}
	return ndriver.Systemctl(n, DcosServiceName, options)
}

func (d *dcosSchedOps) ValidateOnNode(n node.Node) error {
	return &errors.ErrNotSupported{
		Type:      "Function",
		Operation: "ValidateOnNode",
	}
}

func (d *dcosSchedOps) EnableOnNode(n node.Node, ndriver node.Driver) error {
	options := node.SystemctlOpts{
		ConnectionOpts: node.ConnectionOpts{
			Timeout:         2 * time.Minute,
			TimeBeforeRetry: 20 * time.Second,
		},
		Action: "start",
	}
	return ndriver.Systemctl(n, DcosServiceName, options)
}

func (d *dcosSchedOps) ValidateAddLabels(replicaNodes []api.Node, vol *api.Volume) error {
	// We do not have labels in DC/OS currently
	return nil
}

func (d *dcosSchedOps) ValidateRemoveLabels(vol *volume.Volume) error {
	// We do not have labels in DC/OS currently
	return nil
}

func (d *dcosSchedOps) GetVolumeName(vol *volume.Volume) string {
	return vol.Name
}

func (d *dcosSchedOps) ValidateVolumeCleanup(n node.Driver) error {
	// TODO: Implement this
	return nil
}

func (d *dcosSchedOps) GetServiceEndpoint() (string, error) {
	// PX driver is accessed directly on agent nodes. There is no DC/OS level
	// service endpoint which can be used to redirect the calls to PX driver
	return "", nil
}

func (d *dcosSchedOps) UpgradePortworx(version string) error {
	// TOOD: Implement this method
	return nil
}

func init() {
	d := &dcosSchedOps{}
	Register("dcos", d)
}
