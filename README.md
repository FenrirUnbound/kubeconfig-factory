# kubeconfig-factory

Generate a temporary Kubeconfig

Some workflows have you opening new tabs & shells at any given moment. Unfortunately, your kubeconfig
does not create a separate instance of itself when you try to diversify your work environment: your
context is shared across the different shell instances or you can't execute commands in parallel
because you don't want to interfere with your long-running kubectl operation.

This tool helps you generate a temporary copy of your Kubeconfig file so you can generate as many or
as few "unique" replicas.

## Usage

```
$ kubeconfig-factory
/tmp/kubeconfig-872134088

$ export KUBECONFIG=`kubeconfig-factory`
```
