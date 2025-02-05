package contracts

import "fmt"

type InputContract struct {
	Edges [][]int `json:"edges"`
	Start int     `json:"start"`
	End   int     `json:"end"`
}

func (c *InputContract) Validate() error {
	if len(c.Edges) == 0 {
		return fmt.Errorf("InputContract contains no edges")
	}

	maxNode := -1

	for _, edge := range c.Edges {
		if len(edge) != 2 {
			return fmt.Errorf("InputContract contains invalid edges")
		}
		if edge[0] > maxNode {
			maxNode = edge[0]
		}
		if edge[1] > maxNode {
			maxNode = edge[1]
		}
	}

	if c.Start > maxNode {
		return fmt.Errorf("InputContract does not contain start node")
	}
	if c.End > maxNode {
		return fmt.Errorf("InputContract does not contain end node")
	}

	return nil
}
