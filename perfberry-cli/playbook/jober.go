package playbook

import (
	"github.com/perfberry/cli/api"
	"github.com/perfberry/cli/ui"
)

type Jober interface {
	Run(ac *api.Client, uc *ui.Client) error
}
