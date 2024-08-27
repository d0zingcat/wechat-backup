package backup

import (
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(usr.HomeDir, path[1:]), nil
	}
	return path, nil
}

// list dir with pattern match
func listDirWithPattern(dir string, pattern string) ([]string, []string, error) {
	var files []string
	var dirs []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}

	// pattern is a regular expression
	for _, entry := range entries {
		matched, err := regexp.MatchString(pattern, entry.Name())
		if err != nil {
			return nil, nil, err
		}
		if !matched {
			continue
		}
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return files, dirs, nil
}

func listDir(dir string) ([]string, []string, error) {
	var files []string
	var dirs []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		} else {
			dirs = append(dirs, entry.Name())
		}
	}

	return files, dirs, nil
}

func dirFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
