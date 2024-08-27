package backup

import (
	"errors"
	"fmt"

	"github.com/zhaobingss/gorm-sqlcipher"
	"gorm.io/gorm"
)

// PRESET sqlite3 encryption schema
// PRAGMA cipher_page_size = 1024;
// PRAGMA kdf_iter = 64000;
// PRAGMA cipher_kdf_algorithm = PBKDF2_HMAC_SHA1;
// PRAGMA cipher_hmac_algorithm = HMAC_SHA1;

// PRESET sqlite4 encryption schema
// PRAGMA cipher_page_size = 4096;
// PRAGMA kdf_iter = 256000;
// PRAGMA cipher_kdf_algorithm = PBKDF2_HMAC_SHA512;
// PRAGMA cipher_hmac_algorithm = HMAC_SHA512;

func OpenDB(path, key string) (*gorm.DB, error) {
	existed, err := dirFileExists(path)
	if err != nil {
		return nil, err
	}
	if !existed {
		return nil, errors.New("no such db file")
	}

	path = fmt.Sprintf("%s?_pragma_key=x'%s'&_pragma_cipher_page_size=1024&_pragma_cipher_page_size=1024&_pragma_kdf_iter=64000&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA1&_pragma_cipher_hmac_algorithm=HMAC_SHA1", path, key)
	println(path)
	db, err := gorm.Open(sqlcipher.Open(path), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
