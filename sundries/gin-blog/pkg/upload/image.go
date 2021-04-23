// Package upload
// @file: image.go
// @date: 2020/12/31
package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"gin-blog/pkg/file"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
)

// GetImageFullUrl get the full access path of image
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetImagePath() + name
}

// GetImageName get the image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImageSize check image file size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
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
