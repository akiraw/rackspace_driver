package executor

// // StorageExecutor is the part of a storage driver that is downloaded at
// // runtime by the libStorage client.
// type StorageExecutor interface {
// 	Driver
// 	StorageExecutorFunctions
// }

// // StorageExecutorFunctions is the collection of functions that are required of
// // a StorageExecutor.
// type StorageExecutorFunctions interface {
// 	// InstanceID returns the local system's InstanceID.
// 	InstanceID(
// 		ctx Context,
// 		opts Store) (*InstanceID, error)

// 	// NextDevice returns the next available device.
// 	NextDevice(
// 		ctx Context,
// 		opts Store) (string, error)

// 	// LocalDevices returns a map of the system's local devices.
// 	LocalDevices(
// 		ctx Context,
// 		opts *LocalDevicesOpts) (*LocalDevices, error)
// }



import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/akutz/gofig"
	"github.com/akutz/goof"

	"github.com/emccode/libstorage/api/registry"
	"github.com/emccode/libstorage/api/types"
	"github.com/emccode/libstorage/drivers/storage/rackpace"
)

// driver is the storage executor for the VFS storage driver.
type driver struct{
	config        	gofig.Config
}

func init() {
	registry.RegisterStorageExecutor(rackspace.Name, newdriver)
}

func newDriver() types.StorageExecutor {
	return &driver{}
}

func (d *driver) Init(ctx types.Context, config gofig.Config) error {
	d.config = config
	return nil
}

func (d *driver) Name() string {
	return rackspace.Name
}


// InstanceID returns the aws instance configuration
func (d *driver) InstanceID(
	ctx types.Context,
	opts types.Store) (*types.InstanceID, error) {
	cmd := exec.Command("/usr/bin/xenstore-read",  "name")
	cmd.Env = d.config.EnvVars()
	cmdOut, err := cmd.Output()

	if err != nil {
		return "",
			goof.WithFields(eff(goof.Fields{
				"moduleName": d.r.Context,
				"cmd.Path":   cmd.Path,
				"cmd.Args":   cmd.Args,
				"cmd.Out":    cmdOut,
			}), "error getting instance id")
	}

	instanceID := strings.Replace(string(cmdOut), "\n", "", -1)

	validInstanceID := regexp.MustCompile(`^instance-`)
	valid := validInstanceID.MatchString(instanceID)
	if !valid {
		return "", goof.WithFields(eff(goof.Fields{
			"instanceId": instanceID}), "error matching instance id")
	}
	instanceID = strings.Replace(instanceID, "instance-", "", 1)

	iid := &types.InstanceID{Driver: rackspace.Name}
	if err := iid.MarshalMetadata(instanceID); err != nil {
		return nil, err
	}
	return iid, nil
}

func (d *driver) NextDevice(
	ctx types.Context,
	opts types.Store) (string, error) {
	return "", types.ErrNotImplemented
}

func (d *driver) LocalDevices(
	ctx types.Context,
	opts *types.LocalDevicesOpts) (*types.LocalDevices, error) {

	out, err := exec.Command(
		"df", "--output=source,target").Output()
	if err != nil {
		return nil, goof.WithError("error running df", err)
	}

	input := string(out)

	re, _ := regexp.Compile(`^/dev/xvd([a-z])`)
	localDevices := make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	var prev string
	matched := false
	for scanner.Scan() {
		temp := scanner.Text()
		if matched {
			localDevices[prev] = temp
		}
		matched = re.MatchString(temp)
		prev = temp
	}

	return &types.LocalDevices{
		Driver:    rackspace.Name,
		DeviceMap: localDevices,
	}, nil
}

///////////////////////////////////////////////////////////////////////
//// HELPER FUNCTIONS FOR RACKSPACE EXECUTOR FROM THIS POINT ON   /////
///////////////////////////////////////////////////////////////////////

