package sandbox

import "github.com/zekroTJA/ranna/pkg/models"

type Sandbox interface {
	Run() (stdout, stderr string, err error)
	Kill() error
	Delete() error
}

type Provider interface {
	CreateSandbox(spec models.Spec) (Sandbox, error)
}
