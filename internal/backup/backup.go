package backup

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
)

func Backup(param BackupParam) error {
	baseDir, err := DirCheck(param)
	if err != nil {
		return err
	}
	contactDir := filepath.Join(baseDir, CONTACT_DIR)
	if err = ProcessContact(contactDir, param.CipherKey); err != nil {
		return err
	}
	return nil
}

func ProcessContact(path, key string) error {
	files, _, err := listDirWithPattern(path, `.*\.db$`)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("fail to get contact db file")
	}
	db, err := OpenDB(filepath.Join(path, files[0]), key)
	if err != nil {
		return err
	}
	results := make([]map[string]any, 0)
	if err := db.Raw("SELECT * FROM WCContact").Scan(&results).Error; err != nil {
		return err
	}
	fmt.Printf("%+v", results)

	return nil
}

func DirCheck(param BackupParam) (accountBaseDir string, err error) {
	appDir, err := expandPath(param.AppDir)
	if err != nil {
		log.Fatalf("fail to expand path: %v", err)
	}
	param.AppDir = appDir
	_, dirs, err := listDir(param.AppDir)
	if err != nil {
		return
	}
	// only one wechat version app
	if len(dirs) != 1 {
		err = errors.New("fail to get wechat version dir")
		return
	}
	versionDir := filepath.Join(appDir, dirs[0])
	_, dirs, err = listDir(versionDir)
	if err != nil {
		return
	}

	var accountDirs []string
	for _, dir := range dirs {
		if len(dir) == 32 {
			accountDirs = append(accountDirs, dir)
		}
	}
	if len(accountDirs) == 0 {
		err = errors.New("no wechat account found")
		return
	}
	if len(accountDirs) > 1 && param.Account == "" {
		log.Printf("More than one account found, please use --account param to choose the account hash. Accounts: %v", accountDirs)
		err = errors.New("too many wechat account found")
		return
	}
	if param.Account == "" {
		param.Account = accountDirs[0]
	}
	accountBaseDir = filepath.Join(versionDir, param.Account)
	dirExists, err := dirFileExists(accountBaseDir)
	if err != nil {
		return
	}
	if !dirExists {
		err = errors.New("fail to find wechat account dir")
		return
	}
	return
}
