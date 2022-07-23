package repo_manager

import "errors"

var NotGitRepoErorr = errors.New("Current directoru is a not a git repository.")
var NoGitReposInTheDirErorr = errors.New("There are no git repositories in this directory.")
