package cmd

import (
	"fmt"
	"os"
	"strings"

	jh "github.com/jdrivas/sponde/jupyterhub"
	t "github.com/jdrivas/sponde/term"
	"github.com/juju/ansiterm"
)

// Info is a proxy for the jupyterhub Info
type Info jh.Info

// List displays the Info object.
func (i Info) List() {
	info := jh.Info(i)
	lines := [][2]string{
		{t.Title("JupyterHub"), t.Text(getCurrentConnection().HubURL)},
		{t.Title("JupyterHub Version:"), t.Text(info.Version)},
		{t.Title("JupyterHub System Executable:"), t.Text(info.SysExecutable)},
		{t.Title("Authenticator Class:"), t.Text(info.Authenticator.Class)},
		{t.Title("Authenticator Version:"), t.Text(info.Authenticator.Version)},
		{t.Title("Spawner Class:"), t.Text(info.Spawner.Class)},
		{t.Title("Spawner Version:"), t.Text(info.Spawner.Version)},
	}
	w := ansiterm.NewTabWriter(os.Stdout, 4, 4, 3, ' ', 0)
	for _, l := range lines {
		fmt.Fprintf(w, "%s\t%s\n", l[0], l[1])
	}

	python := strings.Split(info.Python, "\n")
	if len(python) > 0 {
		fmt.Fprintf(w, "Python:\t%s\n", python[0])
		for _, l := range python[1:] {
			fmt.Fprintf(w, "\t%s\n", l)
		}
	}
	w.Flush()
}
