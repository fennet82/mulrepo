package models

import (
	"errors"
	"fmt"
	"os"
	"bufio"
    "fmt"
    "os"
    "strings"

	logger "log/slog"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name            string   `json:"name"`
	Path            string   `json:"path"`
	Include         bool     `json:"include"`
	CrticalBranches []string `json:"crtical_branches"`
	AutoPush        bool     `json:"auto_push"`
}

func (repo *Repo) GetGitStatus() (error) {

}

func (repo *Repo) GitCommit(CommitMsg string) (error) {

}

func (repo *Repo) GitPush() (error) {
	git_repo, err := git.PlainOpen(repo.Path)
	if err != nil {
		fmt.Errorf("GitPush(): error occured in function on repo: %s, %w", repo.Name, err)
	}

	if !repo.AutoPush {
		var 
	}
}

func (repo *Repo) GitFetch() (error) {

}
