package controller

import (
	"github.com/openshift/gcp-project-operator/pkg/controller/projectclaim"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, projectclaim.Add)
}
