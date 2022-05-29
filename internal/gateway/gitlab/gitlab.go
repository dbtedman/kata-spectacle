package gitlab

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type GitLab struct {
	Token string
	Url   string
}

type SearchQueryResult struct {
	Basename  string `json:"basename"`
	Data      string `json:"data"`
	Path      string `json:"path"`
	FileName  string `json:"filename"`
	Id        string `json:"id"`
	Ref       string `json:"ref"`
	StartLine int    `json:"startline"`
	ProjectId int    `json:"project_id"`
}

type ProjectResult struct {
	WebUrl            string `json:"web_url"`
	DefaultBranch     string `json:"default_branch"`
	NameWithNamespace string `json:"name_with_namespace"`
}

type InfoSpec struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type OpenAPISpec struct {
	OpenAPIVersion string   `yaml:"openapi"`
	Info           InfoSpec `yaml:"info"`
}

func (receiver GitLab) Search() ([]SearchQueryResult, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		receiver.Url+"/api/v4/search?scope=blobs&search=filename:openapi.yaml&per_page=5",
		nil,
	)

	if err != nil {
		return []SearchQueryResult{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Token)

	resp, err := client.Do(req)

	if err != nil {
		return []SearchQueryResult{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var results []SearchQueryResult

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return []SearchQueryResult{}, err
	}

	return results, nil
}

func (receiver GitLab) GetProject(projectId int) (ProjectResult, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(receiver.Url+"/api/v4/projects/%d", projectId),
		nil,
	)

	if err != nil {
		return ProjectResult{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Token)

	resp, err := client.Do(req)

	if err != nil {
		return ProjectResult{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var result ProjectResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ProjectResult{}, err
	}

	return result, nil
}

func (receiver GitLab) DownloadFileUrl(projectId int, filePath string, branchName string) string {
	return fmt.Sprintf(
		receiver.Url+"/api/v4/projects/%d/repository/files/%s/raw?ref=%s",
		projectId,
		url.QueryEscape(filePath),
		branchName,
	)
}

func (receiver GitLab) BrowseFileUrl(repositoryUrl string, branchName string, filePath string) string {
	return fmt.Sprintf(
		"%s/-/blob/%s/%s",
		repositoryUrl,
		branchName,
		filePath,
	)
}

func (receiver GitLab) GetSpec(specUrl string) (OpenAPISpec, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		specUrl,
		nil,
	)

	if err != nil {
		return OpenAPISpec{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Token)

	resp, err := client.Do(req)

	if err != nil {
		return OpenAPISpec{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var spec OpenAPISpec

	if err := yaml.Unmarshal(bodyBytes, &spec); err != nil {
		return OpenAPISpec{}, err
	}

	return spec, nil
}
