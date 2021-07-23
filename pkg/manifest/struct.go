package manifest

import "github.com/stevenle/topsort"

type Package struct {
	Name         string   `yaml:"name"`
	Dependencies []string `yaml:"dependencies"`
}

type Manifest struct {
	Name     string    `yaml:"name"`
	Packages []Package `yaml:"packages"`
}

func (m *Manifest) Graph() (*topsort.Graph, error) {
	// Initialize the graph.
	graph := topsort.NewGraph()

	for _, p := range m.Packages {
		for _, d := range p.Dependencies {
			err := graph.AddEdge(p.Name, d)
			if err != nil {
				return nil, err
			}
		}
	}

	return graph, nil
}

func (m *Manifest) Dependencies(name string) ([]string, error) {
	graph, err := m.Graph()
	if err != nil {
		return nil, err
	}

	return graph.TopSort(name)
}

func (m *Manifest) TopLevelPackages() ([]string, error) {
	topLevel := []string{}
	// for each package
	for _, p := range m.Packages {
		isDependency := false
		for _, p2 := range m.Packages {
			if p.Name == p2.Name {
				// skip
			}
			// if current package is used as a dependency elsewhere, do not add
			if arrayContains(p2.Dependencies, p.Name) {
				isDependency = true
			}
		}

		if !isDependency {
			topLevel = append(topLevel, p.Name)
		}
	}
	return topLevel, nil
}

func arrayContains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
