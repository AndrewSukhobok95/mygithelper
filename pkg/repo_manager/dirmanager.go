package repo_manager

import (
	"os"
	"io/ioutil"
	"path/filepath"
)

type DirManager struct {
	baseDir   string
	baseName  string
    repoDirs  []string
	repoNames []string
    otherDirs []string
}

func NewDirManager(baseDir string) (*DirManager, error) {
	if baseDir[len(baseDir)-1] != '/' {
        baseDir += "/"
    }

	dirManager := &DirManager{
		baseDir : baseDir, 
		baseName : filepath.Base(baseDir),
	}

	files, err := ioutil.ReadDir(baseDir)
    if err != nil {
		return nil, err
	}

	for _, fileInfo := range files {
        if !fileInfo.IsDir() {
            continue
        }

		path := baseDir + fileInfo.Name()
		if isRepo, _ := IsGitRepo(path); isRepo {
			dirManager.repoDirs = append(dirManager.repoDirs, path)
			dirManager.repoNames = append(dirManager.repoNames, fileInfo.Name())
		} else {
			dirManager.repoDirs = append(dirManager.otherDirs, path)
		}
    }
	return dirManager, nil
}

func (m *DirManager) RunMghFunc(
	mghFunc func (...string) (string, error),
	multiRun bool,
	mghFuncArgs ...string) (map[string]string, error) {
	output := make(map[string]string)
	if multiRun {
		if len(m.repoDirs) == 0 {
			return output, NoGitReposInTheDirErorr
		}
		for rDirIndex, rDir := range m.repoDirs {
			os.Chdir(rDir)
			out, err := mghFunc(mghFuncArgs...)
			output[m.repoNames[rDirIndex]] = out
			if err != nil {
				return output, err
			}
		}
	} else {
		if isRepo, _ := IsGitRepo(m.baseDir); !isRepo {
			return output, NotGitRepoErorr
		}
		out, err := mghFunc(mghFuncArgs...)
		output[m.baseName] = out
		if err != nil {
			return output, err
		}
	}
	return output, nil
}


