package common

import (
	"fmt"
	"io/ioutil"
	"ssh_manage/config"
)

func WriteCaptchaToFile(captcha, phone string) {
	path := config.Config.CaptchaDir.Path
	filePath := path + "/" + phone + ".txt"
	content := []byte(captcha)
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("write file error,error is:", err)
	}

}
