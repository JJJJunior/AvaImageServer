package app

import (
	"AvaImageServer/pkg/e"
	"AvaImageServer/pkg/logging"
	"AvaImageServer/pkg/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)
	t := time.Now()
	year := t.Year()   // type int
	month := t.Month() // type time.Month
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath() + strconv.Itoa(year) + strconv.Itoa(int(month))
		savePath := upload.GetImagePath() + strconv.Itoa(year) + strconv.Itoa(int(month))
		src := fullPath + "/" + imageName
		fmt.Println(imageName, fullPath, savePath, src)
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
			fmt.Println(e.GetMsg(code))
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				logging.Info("访问者IP:", GetRequestIP(c), "上传地址连接:", upload.GetImageFullUrl(imageName))
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + "/" + imageName
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
