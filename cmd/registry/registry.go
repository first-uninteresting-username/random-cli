package registry

import (
	"fmt"
)

var actions []Action

func Register(a Action) error {
	if a.Name == "" {
		return fmt.Errorf("action with description %q has empty name", a.Description)
	}
	if a.Description == "" {
		return fmt.Errorf("action %q has empty description", a.Name)
	}
	if a.Run == nil {
		return fmt.Errorf("action %q has no action", a.Name)
	}

	for _, ac := range actions {
		if ac.Name == a.Name {
			return fmt.Errorf("there is an existing action with name %q", a.Name)
		}
	}

	actions = append(actions, a)

	return nil
}
