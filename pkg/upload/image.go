package upload

import (
	"AvaImageServer/pkg/file"
	"AvaImageServer/pkg/logging"
	"AvaImageServer/pkg/util"
	"AvaImageServer/setting"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

/*
GetImageFullUrl：获取图片完整访问 URL
GetImageName：获取图片名称
GetImagePath：获取图片路径
GetImageFullPath：获取图片完整路径
CheckImageExt：检查图片后缀
CheckImageSize：检查图片大小
CheckImage：检查图片
*/

// GetImageFullUrl：获取图片完整访问 URL
func GetImageFullUrl(name string) string {
	t := time.Now()
	year := t.Year()   // type int
	month := t.Month() // type time.Month
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + strconv.Itoa(year) + strconv.Itoa(int(month)) + "/" + name
}

// GetImageName：获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath：获取图片路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath：获取图片完整路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt：检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize：检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage：检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
