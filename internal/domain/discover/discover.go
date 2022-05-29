package discover

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
)

type Discover struct {
	GitLab gitlab.GitLab
}

type DiscoverRequest struct {
}

type DiscoverResult struct {
	RepositoryName     string
	RepositorySpecLink string
	SpecTitle          string
	SpecDescription    string
}

func (receiver Discover) Execute(request DiscoverRequest) ([]DiscoverResult, error) {
	gitlabSearchResults, err := receiver.GitLab.Search()

	var results []DiscoverResult

	for _, searchResult := range gitlabSearchResults {
		project, err := receiver.GitLab.GetProject(searchResult.ProjectId)
		if err != nil {
			return []DiscoverResult{}, err
		}

		specFile := receiver.GitLab.DownloadFileUrl(
			searchResult.ProjectId,
			searchResult.Path,
			project.DefaultBranch,
		)

		spec, err := receiver.GitLab.GetSpec(specFile)
		if err != nil {
			return []DiscoverResult{}, err
		}

		results = append(results, DiscoverResult{
			RepositoryName: project.NameWithNamespace,
			RepositorySpecLink: receiver.GitLab.BrowseFileUrl(
				project.WebUrl,
				project.DefaultBranch,
				searchResult.FileName,
			),
			SpecTitle:       spec.Info.Title,
			SpecDescription: spec.Info.Description,
		})
	}

	return results, err
}
