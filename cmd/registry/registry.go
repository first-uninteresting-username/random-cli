package registry

import (
	"fmt"
	"math/rand/v2"
	"os"
	"text/tabwriter"
)

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

func GetAction(name string, allowDestructive bool) (Action, error) {
	if len(actions) == 0 {
		return Action{}, fmt.Errorf("action list is empty")
	}

	if name == "" {
		return getRandomAction(allowDestructive)
	}

	for _, a := range actions {
		if a.Name == name {
			return a, nil
		}
	}

	return Action{}, fmt.Errorf("no action named %q", name)
}

func getRandomAction(allowDestructive bool) (Action, error) {
	if len(actions) == 0 {
		return Action{}, fmt.Errorf("action list is empty")
	}

	if allowDestructive == true {
		index := rand.IntN(len(actions))
		return actions[index], nil
	}

	a := Action{
		IsDestructive: true,
	}
	for a.IsDestructive {
		a = actions[rand.IntN(len(actions))]
	}

	return a, nil
}

func ListActions() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	sorted := sortActions()

	fmt.Fprint(w, "action\tdescription\tis destructive\n")

	for _, action := range sorted {
		fmt.Fprintf(w, "%s\t%s\t%t\n", action.Name, action.Description, action.IsDestructive)
	}

	err := w.Flush()
	if err != nil {
		return err
	}
	return nil
}
