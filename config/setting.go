package config

import (
	"gopkg.in/ini.v1"
	"pdfUtil/lib"
)

var (
	PORT string
	SAVEFILE string
	SENDFILE string
	COVER  string
	)

// 引入时自动加载init函数
func init() {
	file, err := ini.Load("config.ini")

	if err != nil {
		lib.CheckErr(err, "[x] ini loading error!")
	}
	LoadService(file)
}

func LoadService(file *ini.File) {
	// 选择分区，MustString是默认值
	PORT = file.Section("service").Key("PORT").MustString("8001")
	SAVEFILE = file.Section("service").Key("SAVEFILE").MustString("./data/in/")
	SENDFILE = file.Section("service").Key("SENDFILE").MustString("./data/out/watermark/pdf/")
	COVER = file.Section("service").Key("COVER").MustString("./data/cover.pdf")
}