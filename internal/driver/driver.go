package driver

import (
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	environment "github.com/fenrirunbound/kubeconfig-factory/internal/env"
)

const kubeConfigEnvVariable = "KUBECONFIG"

type Driver struct {
	env environment.Env
}

func (d *Driver) defaultConfig() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(u.HomeDir, ".kube/config"), nil
}

func (d *Driver) copyContents(source, target string) (copyErr error) {
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() {
		closeError := output.Close()
		if copyErr == nil {
			copyErr = closeError
		}
	}()

	if _, err = io.Copy(output, input); err != nil {
		return err
	}

	return output.Sync()
}

func (d *Driver) GetKubeconfig() string {
	configList := d.env.Get(kubeConfigEnvVariable)
	parts := strings.Split(configList, ":")
	result := parts[0]

	if len(result) == 0 {
		// todo: make this OS agnostic
		result, _ = d.defaultConfig()
		return result
	}

	return result
}

func (d *Driver) GenerateConfig() (string, error) {
	source := d.GetKubeconfig()
	target, err := ioutil.TempFile(os.TempDir(), "kubeconfig-")
	if err != nil {
		return "", err
	}

	if err = os.Link(source, target.Name()); err != nil {
		err = d.copyContents(source, target.Name())
	}

	if err != nil {
		return "", err
	}

	return target.Name(), nil
}

func NewDriver(e environment.Env) *Driver {
	return &Driver{
		env: e,
	}
}
