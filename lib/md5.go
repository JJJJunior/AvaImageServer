package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

const password = "avavin"
const salt = "asldjd23qie2bdnjkbnaf!@#$@$"

// md5WithSalt generates a salted MD5 hash of the input string.
func MD5WithSalt() string {
	hasher := md5.New()
	// 获取当前时间戳（以秒为单位）
	timestampInSeconds := time.Now().Unix()
	// 转换为分钟
	timestampInMinutes := strconv.FormatInt(timestampInSeconds/60, 10)
	hasher.Write([]byte(password + "$" + timestampInMinutes + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}
