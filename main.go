package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fenrirunbound/kubeconfig-factory/internal/driver"
	"github.com/spf13/cobra"
)

type myEnv struct{}

func (m *myEnv) Get(name string) string {
	return os.Getenv(name)
}

func (m *myEnv) Set(name, value string) {
	os.Setenv(name, value)
}

const longDescription = `kubeconfig-factory generates a temp copy of your Kubeconfig

It generates a temporary file that replicates what your current Kubeconfig is. It
determines your Kubeconfig by the first config defined in the $KUBECONFIG environment
variable, or the default config if undefined.
`

func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "kubeconfig-factory",
		Short: "Generate temporary kubeconfigs",
		Long:  longDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			d := driver.NewDriver(&myEnv{})

			tempConfig, err := d.GenerateConfig()

			fmt.Printf("%s\n", tempConfig)

			return err
		},
	}
}

func main() {
	cmd := createCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
