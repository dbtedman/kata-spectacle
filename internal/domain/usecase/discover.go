package usecase

import (
	"github.com/dbtedman/kata-spectacle/internal/gateway/gitlab"
	"log"
)

type Discover struct {
	GitLab gitlab.GitLab
}

type Request struct {
}

type Result struct {
	RepositoryName     string
	RepositorySpecLink string
	SpecTitle          string
	SpecDescription    string
	CommitRef          string
}

func (receiver Discover) Execute(request Request) ([]Result, error) {
	// TODO: Load existing spec list

	gitlabSearchResults, err := receiver.GitLab.Search()

	var results []Result

	for _, searchResult := range gitlabSearchResults {
		project, err := receiver.GitLab.GetProject(searchResult.ProjectId)
		if err != nil {
			log.Println(err)
			continue
		}

		if project.Visibility == "private" {
			log.Printf("%s is private", project.WebUrl)
			continue
		}

		projectRepositoryBranches, _ := receiver.GitLab.GetProjectRepositoryBranches(searchResult.ProjectId, project.DefaultBranch)

		// If projectRepositoryBranches.Commit.Id has not changed, we don't need to re-fetch the spec

		specFile := receiver.GitLab.DownloadFileUrl(
			searchResult.ProjectId,
			searchResult.Path,
			projectRepositoryBranches.Commit.Id,
		)

		spec, err := receiver.GitLab.GetSpec(specFile)
		if err != nil {
			log.Println(err)
			continue
		}

		results = append(results, Result{
			RepositoryName: project.NameWithNamespace,
			RepositorySpecLink: receiver.GitLab.BrowseFileUrl(
				project.WebUrl,
				project.DefaultBranch,
				searchResult.FileName,
			),
			SpecTitle:       spec.Info.Title,
			SpecDescription: spec.Info.Description,
			CommitRef:       projectRepositoryBranches.Commit.Id,
		})
	}

	// TODO: Save updated spec list

	return results, err
}
