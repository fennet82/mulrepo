package models

import (
	"encoding/json"
	"fmt"
	"mulrepo/custom_errors"
	"os"
	"sync"
)

type Config struct {
	GlobalBasicGitAuth *BasicGitAuth
	ConfigFilePath     string
	ExportFilePath     string
	ReposInstance      *Repos
}

var (
	configInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{
			ConfigFilePath: "",
			ExportFilePath: ".",
			ReposInstance:  &Repos{},
		}
	})

	return configInstance
}

func (config *Config) MarshalReposToJSON(path string) error {
	if config.ReposInstance == nil {
		return fmt.Errorf("ReposInstance is nil")
	}

	data, err := json.MarshalIndent(config.ReposInstance, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal repos: %w", err)
	}

	if err := os.WriteFile(config.ExportFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (config *Config) LoadConfigFromJson(JsonFilePath string) (*Repos, error) {
	repos := &Repos{}

	bytes, err := os.ReadFile(JsonFilePath)
	if err != nil {
		return repos, err
	}

	err = json.Unmarshal(bytes, repos)
	if err != nil {
		return repos, err
	}

	return repos, nil
}

func (config *Config) PrintConfigTemplate() error {
	if config == nil {
		return custom_errors.ErrNilPointerReferenced
	}

	fmt.Printf(`
	{
		"repos": {
			{
				"name": <str>,                                    -   name of the repository, doesnt need to be the actual name of the repository.
				"path": <str>,                                      -   path is basically the path of the repository (it's better if youll use the full path of the repository)
				"include": <bool>,                               -  include is boolean value that indicates if the current repo will be included in the iteration of repos
				"critical_branches": <str_list>,             -   critical_branches is a mention of all the branches that needs protection (asking before commiting and asking before pushing)
				"auto_push": <bool>                           -  auto_push will automatically push the staged files (if the current branch in critical_branches auto_push will have no affect) to the repo without asking the user for permission
			},
			...
		}
	}
	`)

	return nil
}

func (config *Config) GetIncludedRepos() (*Repos, error) {
	IncludedRepos := make(map[*Repo]struct{})

	for _, repo := range config.ReposInstance.Repos {
		if repo.Include {
			IncludedRepos[&repo] = struct{}{}
		}
	}

	includedReposList := &Repos{}
	for repo := range IncludedRepos {
		includedReposList.Repos = append(includedReposList.Repos, *repo)
	}

	return includedReposList, nil
}
