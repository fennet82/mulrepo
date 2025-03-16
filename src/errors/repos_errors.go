package errors

type ErrRepoNotFoundError struct {
	RepoName string
}

func (e *ErrRepoNotFoundError) Error() string {
    return "repo not found: " + e.RepoName
}