package task_invoker

import (
	"fmt"

	"github.com/spf13/cobra"
)

type InvokerFN func() error
type InvokerList map[string]InvokerFN

type Invoker struct {
	name        string
	invokerList InvokerList
	Cmd         *cobra.Command
	Args        []string
}

func NewInvoker(
	name string,
	cmd *cobra.Command,
	args []string,
) *Invoker {
	return &Invoker{
		name:        name,
		Cmd:         cmd,
		Args:        args,
		invokerList: InvokerList{},
	}
}

func (i *Invoker) Add(name string, fn InvokerFN) {
	i.invokerList[name] = fn
}

func (i *Invoker) Run() error {
	invoker, found := i.invokerList[i.name]

	if !found {
		return fmt.Errorf("task name %s is not registered", i.name)
	}

	return invoker()
}
