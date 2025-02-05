package services

import (
	"LLd-Test/contracts"
	"LLd-Test/domains"
)

type PathService struct {
	allPaths [][]int
}

func NewPathService() *PathService {
	return &PathService{}
}

func (ps *PathService) dfs(curr int, dest int, adjList [][]int, vis []bool, path []int) {
	path = append(path, curr)

	if curr == dest {
		p := path
		ps.allPaths = append(ps.allPaths, p)
		path = path[:len(path)-1]
		return
	}

	vis[curr] = true
	for _, child := range adjList[curr] {
		if !vis[child] {
			ps.dfs(child, dest, adjList, vis, path)
		}
	}
	vis[curr] = false
}

func (ps *PathService) GetAllPathsFromGraph(request contracts.InputContract) (domains.Response, error) {
	edgesList := request.Edges
	maxNodeValue := -1

	for _, edge := range edgesList {
		if edge[0] > maxNodeValue {
			maxNodeValue = edge[0]
		}
		if edge[1] > maxNodeValue {
			maxNodeValue = edge[1]
		}
	}

	adjList := make([][]int, maxNodeValue+1)
	for _, edge := range edgesList {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
	}

	visited := make([]bool, maxNodeValue+1)
	path := make([]int, 0)
	ps.dfs(request.Start, request.End, adjList, visited, path)

	response := domains.Response{
		Paths: ps.allPaths,
	}

	return response, nil
}
