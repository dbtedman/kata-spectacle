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
	"strings"
)

type GitLab struct {
	Config ConfigReader
}

type ConfigReader interface {
	ReadToken() string
	ReadUrl() string
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
	Visibility        string `json:"visibility"`
}

type InfoSpec struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type OpenAPISpec struct {
	OpenAPIVersion string   `yaml:"openapi"`
	Info           InfoSpec `yaml:"info"`
}

type ProjectRepositoryBranch struct {
	Commit ProjectRepositoryBranchCommit `json:"commit"`
}

type ProjectRepositoryBranchCommit struct {
	Id string `json:"id"`
}

func (receiver GitLab) Search() ([]SearchQueryResult, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		receiver.Config.ReadUrl()+"/api/v4/search?scope=blobs&search=filename:openapi.yaml&per_page=10",
		nil,
	)

	log.Println(req.URL)

	if err != nil {
		return []SearchQueryResult{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Config.ReadToken())

	log.Println(req.URL)
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

	bodyString, bodyStringError := io.ReadAll(resp.Body)

	if bodyStringError != nil {
		log.Fatal(bodyStringError)
	}

	if resp.StatusCode != 200 {
		log.Fatal(string(bodyString))
	}

	var results []SearchQueryResult

	if err := json.NewDecoder(strings.NewReader(string(bodyString))).Decode(&results); err != nil {
		return []SearchQueryResult{}, err
	}

	return results, nil
}

func (receiver GitLab) GetProject(projectId int) (ProjectResult, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(receiver.Config.ReadUrl()+"/api/v4/projects/%d", projectId),
		nil,
	)

	if err != nil {
		return ProjectResult{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Config.ReadToken())

	log.Println(req.URL)
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

func (receiver GitLab) GetProjectRepositoryBranches(projectId int, branch string) (ProjectRepositoryBranch, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(receiver.Config.ReadUrl()+"/api/v4/projects/%d/repository/branches/%s", projectId, branch),
		nil,
	)

	if err != nil {
		return ProjectRepositoryBranch{}, err
	}

	req.Header.Add("PRIVATE-TOKEN", receiver.Config.ReadToken())

	log.Println(req.URL)
	resp, err := client.Do(req)

	if err != nil {
		return ProjectRepositoryBranch{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var result ProjectRepositoryBranch

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ProjectRepositoryBranch{}, err
	}

	return result, nil
}

func (receiver GitLab) DownloadFileUrl(projectId int, filePath string, ref string) string {
	return fmt.Sprintf(
		receiver.Config.ReadUrl()+"/api/v4/projects/%d/repository/files/%s/raw?ref=%s",
		projectId,
		url.QueryEscape(filePath),
		ref,
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

	req.Header.Add("PRIVATE-TOKEN", receiver.Config.ReadToken())

	log.Println(req.URL)
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
