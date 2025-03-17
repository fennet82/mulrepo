package custom_errors

import "fmt"

type BaseGitError struct {
	RepoName  string
	RepoPath  string
	Traceback error
}

func NewGitErrorCtor (repoName string, repoPath string, traceback error, errPointer any) ()

type ErrRepoNotFoundError struct {
	BaseError BaseGitError
}

func (e *ErrRepoNotFoundError) Error() string {
	return fmt.Sprintf("repo: %s, in path: %s not found.\nderived from:\n%s\n",
		e.BaseError.RepoName,
		e.BaseError.RepoPath,
		e.BaseError.Traceback.Error(),
	)
}

type ErrRepoCantBeOpened struct {
	BaseError BaseGitError
}

func (e *ErrRepoCantBeOpened) Error() string {
	return fmt.Sprintf("Repo - InitRepo: repo: %s, in path: %s can be opened.\nderived from:\n%s\n",
		e.BaseError.RepoName,
		e.BaseError.RepoPath,
		e.BaseError.Traceback.Error(),
	)
}

type ErrGettingRepoWorkTree struct {
	BaseError BaseGitError
}

func (e *ErrGettingRepoWorkTree) Error() string {
	return fmt.Sprintf("Repo - InitRepo: repo: %s,  in path: %s can't get repo's worktree.\nderived from:\n%s\n",
		e.BaseError.RepoName,
		e.BaseError.RepoPath,
		e.BaseError.Traceback.Error(),
	)
}

type ErrGettingRepoStatus struct {
	BaseError BaseGitError
}

func (e *ErrGettingRepoStatus) Error() string {
	return fmt.Sprintf("Repo - InitRepo: repo: %s,  in path: %s can't get repo's status.\nderived from:\n%s\n",
		e.BaseError.RepoName,
		e.BaseError.RepoPath,
		e.BaseError.Traceback.Error(),
	)
}
