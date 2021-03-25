package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type GithubData struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	StargazersCount int    `json:"stargazers_count"`
	HtmlUrl         string `json:"html_url"`
	Forks           int    `json:"forks"`
	Language        string `json:"language"`
}

type ResData struct {
	Status string       `json:"status"`
	Data   []GithubData `json:"data"`
}

type GithubDatas []GithubData

func (g GithubDatas) Len() int { return len(g) }

func (g GithubDatas) Less(i, j int) bool {
	return g[i].StargazersCount > g[j].StargazersCount
}

func (g GithubDatas) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

//Github API
func Github(c *gin.Context) {

	var url string = "https://api.github.com/users/alpacabi/repos"

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return
	}

	var data GithubDatas

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}

	sort.Sort(data)

	data2 := data[:6]

	resData := ResData{
		Status: "ok",
		Data:   data2,
	}

	c.JSON(http.StatusOK, resData)
	return
}
