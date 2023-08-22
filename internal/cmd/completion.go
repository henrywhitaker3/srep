package cmd

import "github.com/henrywhitaker3/srep/internal/metadata"

func ScenarioCompletion() []string {
	out := []string{}
	s, err := metadata.Get()

	if err != nil {
		return out
	}

	for _, scenario := range *s {
		out = append(out, scenario.Name)
	}

	return out
}
