package backup

import (
	"gorm.io/gorm"
)

type BackupParam struct {
	CipherKey string
	AppDir    string
	Account   string
}

type Contact struct {
	gorm.Model
	Username            string `gorm:"primaryKey;size:255"`
	Nickname            string `gorm:"size:255"`
	Alias               string `gorm:"size:255"`
	UsernamePinyin      string `gorm:"size:255"`
	UsernamePinyinShort string `gorm:"size:128"`
	Remark              string `gorm:"size:255"`
	RemarkPinyin        string `gorm:"size:255"`
	RemarkPinyinShort   string `gorm:"size:128"`
	Sex                 int16
	HeadImageURL        string `gorm:"type:text"`
	HeadImageURLHd      string `gorm:"type:text"`
	BrandIconURL        string `gorm:"type:text"`
}

func (Contact) TableName() string {
	return "contact"
}
