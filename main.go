package main

import (
	"flag"
	"log"
	"strings"

	"wechat-backup/internal/backup"
)

var param backup.BackupParam

func parseParam() {
	flag.StringVar(&param.CipherKey, "key", "", "Sqlite3 cipher key")
	flag.StringVar(&param.AppDir, "appdir", "", "Dir location of wechat app")
	flag.StringVar(&param.Account, "account", "", "Wechat account hash")
	flag.Parse()

	if param.CipherKey == "" {
		log.Fatal("key is required")
	}
	if strings.HasPrefix(param.CipherKey, "0x") {
		if len(param.CipherKey) != 66 {
			log.Fatal("key must be 32 bytes hex string")
		}
		param.CipherKey = param.CipherKey[2:]
	}
	// TODO: check key 32 bytes long
	// TODO: check account 16 bytes long
	if param.AppDir == "" {
		log.Fatal("appdir is required")
	}
}

func main() {
	parseParam()
	if err := backup.Backup(param); err != nil {
		log.Fatalf("fail to backup: %v", err)
	}
}
