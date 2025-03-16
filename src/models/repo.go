package models

import (
	"fmt"
	"mulrepo/custom_errors"

	logger "log/slog"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name             string   `json:"name"`
	Path             string   `json:"path"`
	Include          bool     `json:"include"`
	CriticalBranches []string `json:"critical_branches"`
	AutoPush         bool     `json:"auto_push"`
	WorkTree         *git.Worktree
}

func (repo *Repo) InitRepo() error {
	if repo == nil {
		return &custom_errors.ErrNilPointerReferenced{}
	}

	r, err := git.PlainOpen(repo.Path)
	if err != nil {
		err := &custom_errors.ErrRepoCantBeOpened{RepoName: repo.Name, RepoPath: repo.Path, Traceback: err}
		logger.Error(err.Error())

		return err
	}

	worktree, err := r.Worktree()
	if err != nil {
		err := &custom_errors.ErrGettingRepoWorkTree{RepoName: repo.Name, Traceback: err}
		logger.Error(err.Error())

		return err
	}

	repo.WorkTree = worktree
	logger.Info(fmt.Sprintf("got repo's: %s worktree successfully", repo.Name))

	return nil
}

func (repo *Repo) GetGitStatus() (string, bool, error) {
	status, err := repo.WorkTree.Status()
	if err != nil {
		err := &custom_errors.ErrGettingRepoStatus{RepoName: repo.Name, Traceback: err}
		logger.Error(err.Error())

		return "", true, err
	}

	return status.String(), status.IsClean(), nil
}

func (repo *Repo) GitCommit(CommitMsg string) error {

}

func (repo *Repo) GitPush() error {

}
