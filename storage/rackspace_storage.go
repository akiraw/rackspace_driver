// /*
// StorageDriver is a libStorage driver used by the routes to implement the
// backend functionality.
// Functions that inspect a resource or send an operation to a resource should
// always return ErrResourceNotFound if the acted upon resource cannot be found.
// */
// type StorageDriver interface {
// 	Driver

// 	// NextDeviceInfo returns the information about the driver's next available
// 	// device workflow.
// 	NextDeviceInfo(
// 		ctx Context) (*NextDeviceInfo, error)

// 	// Type returns the type of storage the driver provides.
// 	Type(
// 		ctx Context) (StorageType, error)

// 	// InstanceInspect returns an instance.
// 	InstanceInspect(
// 		ctx Context,
// 		opts Store) (*Instance, error)

// 	// Volumes returns all volumes or a filtered list of volumes.
// 	Volumes(
// 		ctx Context,
// 		opts *VolumesOpts) ([]*Volume, error)

// 	// VolumeInspect inspects a single volume.
// 	VolumeInspect(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeInspectOpts) (*Volume, error)

// 	// VolumeCreate creates a new volume.
// 	VolumeCreate(
// 		ctx Context,
// 		name string,
// 		opts *VolumeCreateOpts) (*Volume, error)

// 	// VolumeCreateFromSnapshot creates a new volume from an existing snapshot.
// 	VolumeCreateFromSnapshot(
// 		ctx Context,
// 		snapshotID,
// 		volumeName string,
// 		opts *VolumeCreateOpts) (*Volume, error)

// 	// VolumeCopy copies an existing volume.
// 	VolumeCopy(
// 		ctx Context,
// 		volumeID,
// 		volumeName string,
// 		opts Store) (*Volume, error)

// 	// VolumeSnapshot snapshots a volume.
// 	VolumeSnapshot(
// 		ctx Context,
// 		volumeID,
// 		snapshotName string,
// 		opts Store) (*Snapshot, error)

// 	// VolumeRemove removes a volume.
// 	VolumeRemove(
// 		ctx Context,
// 		volumeID string,
// 		opts Store) error

// 	// VolumeAttach attaches a volume and provides a token clients can use
// 	// to validate that device has appeared locally.
// 	VolumeAttach(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeAttachOpts) (*Volume, string, error)

// 	// VolumeDetach detaches a volume.
// 	VolumeDetach(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeDetachOpts) (*Volume, error)

// 	// Snapshots returns all volumes or a filtered list of snapshots.
// 	Snapshots(
// 		ctx Context,
// 		opts Store) ([]*Snapshot, error)

// 	// SnapshotInspect inspects a single snapshot.
// 	SnapshotInspect(
// 		ctx Context,
// 		snapshotID string,
// 		opts Store) (*Snapshot, error)

// 	// SnapshotCopy copies an existing snapshot.
// 	SnapshotCopy(
// 		ctx Context,
// 		snapshotID,
// 		snapshotName,
// 		destinationID string,
// 		opts Store) (*Snapshot, error)

// 	// SnapshotRemove removes a snapshot.
// 	SnapshotRemove(
// 		ctx Context,
// 		snapshotID string,
// 		opts Store) error
// }


package storage

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/akutz/gofig"
	"github.com/akutz/goof"

	"github.com/emccode/libstorage/api/context"
	"github.com/emccode/libstorage/api/registry"
	"github.com/emccode/libstorage/api/types"
	"github.com/emccode/libstorage/drivers/storage/rackspace"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/blockstorage/v1/snapshots"
	"github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

onst (
	providerName = "Rackspace"
	minSize      = 75 //rackspace is 75
)

type driver struct {
	provider           	*gophercloud.ProviderClient
	client             	*gophercloud.ServiceClient
	clientBlockStorage 	*gophercloud.ServiceClient
	region             	string
	instanceID         	string
	config				gofig.Config
	//sync.Mutex
}

func init() {
	registry.RegisterStorageDriver(rackspace.Name, newDriver)
}

func newDriver() types.StorageDriver {
	return &driver{}
}

func (d *driver) Name() string {
	return rackspace.Name
}

func (d *driver) Init(context types.Context, config gofig.Config) error {
	d.config = config
	// fields := eff(map[string]interface{}{
	// 	"endpoint": d.endpoint(),
	// 	"insecure": d.insecure(),
	// 	"useCerts": d.useCerts(),
	// })

	log.WithFields(fields).Debug("starting scaleio driver")

	var err error

	//IMPLEMENT THIS

	log.WithFields(fields).Info("storage driver initialized")

	return nil
}

func (d *driver) Type(ctx types.Context) (types.StorageType, error) {
	return types.Block, nil
}

// 	// NextDeviceInfo returns the information about the driver's next available
// 	// device workflow.
// 	NextDeviceInfo(
// 		ctx Context) (*NextDeviceInfo, error)
//  Currently not implemented in ScaleIO
func (d *driver) NextDeviceInfo(
	ctx types.Context) (*types.NextDeviceInfo, error) {
	return nil, nil
}

// 	// InstanceInspect returns an instance.
// 	InstanceInspect(
// 		ctx Context,
// 		opts Store) (*Instance, error)
func (d *driver) InstanceInspect(
	ctx types.Context,
	opts types.Store) (*types.Instance, error) {

	//IMPLEMENT THIS

	iid := context.MustInstanceID(ctx)
	return nil, nil
}

// 	// Volumes returns all volumes or a filtered list of volumes.
// 	Volumes(
// 		ctx Context,
// 		opts *VolumesOpts) ([]*Volume, error)
func (d *driver) Volumes(
	ctx types.Context,
	opts *types.VolumesOpts) ([]*types.Volume, error) {

	//IMPLEMENT THIS

	return nil, nil
}

// 	// VolumeInspect inspects a single volume.
// 	VolumeInspect(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeInspectOpts) (*Volume, error)
func (d *driver) VolumeInspect(
	ctx types.Context,
	volumeID string,
	opts *types.VolumeInspectOpts) (*types.Volume, error) {

	if volumeID == "" {
		return nil, goof.New("no volumeID specified")
	}

	//IMPLEMENT THIS

	return nil, nil
}

// 	// VolumeCreate creates a new volume.
// 	VolumeCreate(
// 		ctx Context,
// 		name string,
// 		opts *VolumeCreateOpts) (*Volume, error)
func (d *driver) VolumeCreate(ctx types.Context, volumeName string,
	opts *types.VolumeCreateOpts) (*types.Volume, error) {

	//IMPLEMENT THIS

	return nil, nil
}

// 	// VolumeCreateFromSnapshot creates a new volume from an existing snapshot.
// 	VolumeCreateFromSnapshot(
// 		ctx Context,
// 		snapshotID,
// 		volumeName string,
// 		opts *VolumeCreateOpts) (*Volume, error)
func (d *driver) VolumeCreateFromSnapshot(
	ctx types.Context,
	snapshotID, volumeName string,
	opts *types.VolumeCreateOpts) (*types.Volume, error) {

	// notUsed bool,volumeName, volumeID, snapshotID, volumeType string,
	// IOPS, size int64, availabilityZone string) (*types.VolumeResp, error)
	if volumeName == "" {
		return nil, goof.New("no volume name specified")
	}

	//IMPLEMENT THIS
	
	return nil, nil
}

// 	// VolumeCopy copies an existing volume.
// 	VolumeCopy(
// 		ctx Context,
// 		volumeID,
// 		volumeName string,
// 		opts Store) (*Volume, error)
// Currently not implemented in ScaleIO
func (d *driver) VolumeCopy(
	ctx types.Context,
	volumeID, volumeName string,
	opts types.Store) (*types.Volume, error) {
	return nil, nil
}

// 	// VolumeSnapshot snapshots a volume.
// 	VolumeSnapshot(
// 		ctx Context,
// 		volumeID,
// 		snapshotName string,
// 		opts Store) (*Snapshot, error)
// Currently not implemented in ScaleIO
func (d *driver) VolumeSnapshot(
	ctx types.Context,
	volumeID, snapshotName string,
	opts types.Store) (*types.Snapshot, error) {
	return nil, nil
}

// 	// VolumeRemove removes a volume.
// 	VolumeRemove(
// 		ctx Context,
// 		volumeID string,
// 		opts Store) error
func (d *driver) VolumeRemove(
	ctx types.Context,
	volumeID string,
	opts types.Store) error {

	//IMPLEMENT THIS

	log.WithFields(fields).Debug("removed volume")
	return nil
}

// 	// VolumeAttach attaches a volume and provides a token clients can use
// 	// to validate that device has appeared locally.
// 	VolumeAttach(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeAttachOpts) (*Volume, string, error)
func (d *driver) VolumeAttach(
	ctx types.Context,
	volumeID string,
	opts *types.VolumeAttachOpts) (*types.Volume, string, error) {

	//IMPLEMENT THIS
	
	return nil, nil, nil
}

// 	// VolumeDetach detaches a volume.
// 	VolumeDetach(
// 		ctx Context,
// 		volumeID string,
// 		opts *VolumeDetachOpts) (*Volume, error)
func (d *driver) VolumeDetach(
	ctx types.Context,
	volumeID string,
	opts *types.VolumeDetachOpts) (*types.Volume, error) {

	//IMPLEMENT THIS

	return nil, nil
}

//  // Not a part of storage interface
// Not implemented in ScaleIO
func (d *driver) VolumeDetachAll(
	ctx types.Context,
	volumeID string,
	opts types.Store) error {
	return nil
}

// 	// Snapshots returns all volumes or a filtered list of snapshots.
// 	Snapshots(
// 		ctx Context,
// 		opts Store) ([]*Snapshot, error)
// Not implemented in ScaleIO
func (d *driver) Snapshots(
	ctx types.Context,
	opts types.Store) ([]*types.Snapshot, error) {
	return nil, nil
}

// 	// SnapshotInspect inspects a single snapshot.
// 	SnapshotInspect(
// 		ctx Context,
// 		snapshotID string,
// 		opts Store) (*Snapshot, error)
// Not implemented in ScaleIO
func (d *driver) SnapshotInspect(
	ctx types.Context,
	snapshotID string,
	opts types.Store) (*types.Snapshot, error) {
	return nil, nil
}

// 	// SnapshotCopy copies an existing snapshot.
// 	SnapshotCopy(
// 		ctx Context,
// 		snapshotID,
// 		snapshotName,
// 		destinationID string,
// 		opts Store) (*Snapshot, error)
// Not implemented in ScaleIO
func (d *driver) SnapshotCopy(
	ctx types.Context,
	snapshotID, snapshotName, destinationID string,
	opts types.Store) (*types.Snapshot, error) {
	return nil, nil
}

// 	// SnapshotRemove removes a snapshot.
// 	SnapshotRemove(
// 		ctx Context,
// 		snapshotID string,
// 		opts Store) error
// }
// Not implemented in ScaleIO
func (d *driver) SnapshotRemove(
	ctx types.Context,
	snapshotID string,
	opts types.Store) error {
	return nil
}


///////////////////////////////////////////////////////////////////////
///// HELPER FUNCTIONS FOR RACKSPACE DRIVER FROM THIS POINT ON ////////
///////////////////////////////////////////////////////////////////////
