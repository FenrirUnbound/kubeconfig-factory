package driver_test

import (
	"os"
	"strings"
	"testing"

	"github.com/fenrirunbound/kubeconfig-factory/internal/driver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type envMock struct {
	mock.Mock
}

func (e *envMock) Set(name, value string) {
	e.Called(name, value)
}

func (e *envMock) Get(name string) string {
	args := e.Called(name)

	return args.String(0)
}

func TestNew(t *testing.T) {
	a := assert.New(t)

	env := new(envMock)
	d := driver.NewDriver(env)
	a.NotNil(d)
}

func TestLocateKubeconfig(t *testing.T) {
	a := assert.New(t)

	env := new(envMock)
	env.On("Get", "KUBECONFIG").Return("/tmp/whatever")

	d := driver.NewDriver(env)
	result := d.GetKubeconfig()

	a.Equal("/tmp/whatever", result)

	env.AssertExpectations(t)
}

func TestLocateFirstKubeconfig(t *testing.T) {
	a := assert.New(t)

	env := new(envMock)
	env.On("Get", "KUBECONFIG").Return("/tmp/onlythis:/tmp/notthis")

	d := driver.NewDriver(env)
	result := d.GetKubeconfig()

	a.Equal("/tmp/onlythis", result)
	env.AssertExpectations(t)
}

func TestDefaultKubeconfig(t *testing.T) {
	a := assert.New(t)

	env := new(envMock)
	env.On("Get", "KUBECONFIG").Return("")

	d := driver.NewDriver(env)
	result := d.GetKubeconfig()

	a.True(strings.HasSuffix(result, "kube/config"))
}

func TestGenerateConfig(t *testing.T) {
	a := assert.New(t)

	env := new(envMock)
	env.On("Get", "KUBECONFIG").Return("")

	d := driver.NewDriver(env)

	configLocation, err := d.GenerateConfig()
	a.Nil(err)

	tempInfo, err := os.Stat(configLocation)
	a.Nil(err)
	srcInfo, err := os.Stat(d.GetKubeconfig())
	a.Nil(err)

	// without a MD5 hash, assert some characteristics a copy would have in common
	a.Equal(srcInfo.Mode(), tempInfo.Mode())
	a.Equal(srcInfo.Size(), tempInfo.Size())

	env.AssertExpectations(t)
}
