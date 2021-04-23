package export

import "gin-blog/pkg/setting"

func GetExcelFullURL(name string) string {
	return setting.AppSetting.PrefixURL + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
