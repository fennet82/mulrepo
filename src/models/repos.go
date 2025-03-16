package models

import (
	"fmt"
	logger "log/slog"
	"mulrepo/custom_errors"
	"slices"

	"github.com/go-playground/validator"
)

type Repos struct {
	Repos []Repo `validate:"are_names_unique"`
}

// validators
func ValidateNamesUniqueness(fl validator.FieldLevel) bool {
	repos, ok := fl.Field().Interface().([]Repo)
	if !ok {
		return false
	}

	seen := make(map[string]struct{})
	for _, repo := range repos {
		if _, exists := seen[repo.Name]; exists {
			err := &custom_errors.ErrDuplicateName{DuplicateName: repo.Name, ConfigFilePath: GetConfig().ConfigFilePath}
			logger.Error(err.Error())

			return false
		}
		seen[repo.Name] = struct{}{}
	}
	return true
}

func Validate(repos Repos) bool {
	validate := validator.New()

	err := validate.RegisterValidation("are_names_unique", ValidateNamesUniqueness)
	if err != nil {
		logger.Error(fmt.Sprintf("Error registering custom validation: %s", err.Error()))
	}

	validationErr := validate.Struct(repos)

	return validationErr == nil
}

// struct functions
func (repos *Repos) ListRepos() error {
	if repos == nil {
		return custom_errors.ErrNilPointerReferenced
	}
	for index, repo := range repos.Repos {
		fmt.Printf("%d - %+v", index, repo)
	}

	return nil
}

func (repos *Repos) DeleteRepoByName(Name string) error {
	originalLength := len(repos.Repos)

	repos.Repos = slices.DeleteFunc(repos.Repos, func(repo Repo) bool {
		return repo.Name == Name
	})

	if len(repos.Repos) == originalLength {
		return &custom_errors.ErrRepoNotFoundError{RepoName: Name}
	}

	return nil
}

func (repos *Repos) UpdateRepoByName(Name string, updatedRepo Repo) error {
	repo, err := repos.GetRepoByName(Name)
	if err != nil {
		return &custom_errors.ErrRepoNotFoundError{RepoName: Name}
	}
	*repo = updatedRepo

	return nil
}

func (repos *Repos) AddRepo(NewRepo Repo) error {
	repos.Repos = append(repos.Repos, NewRepo)

	return nil
}

func (repos *Repos) GetRepoByName(Name string) (*Repo, error) {
	for _, repo := range repos.Repos {
		if repo.Name == Name {
			return &repo, nil
		}
	}

	return nil, &custom_errors.ErrRepoNotFoundError{RepoName: Name}
}
