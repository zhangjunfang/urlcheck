package urllist

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	获取配置文件中所有的行
 */
func GetLines(filename string) []string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(f), "\n")
	return s
}

/*
	检查有效行并返回
 */
func GetUrls(s []string) []string {
	var urls []string
	for i:=0;i<len(s);i++  {
		url := s[i]
		if len(url) >=7 {
			if (url != "") && (string(url[0]) != "#") && (string(url[0:7]) == "http://"){
				urls=append(urls,s[i])
			}else {
				fmt.Println("Invalid url: " + url)
				continue
			}
		}else if url == ""{
			continue
		} else {
			fmt.Println("Invalid url: " + url)
			continue
		}

	}
	return urls
}

func IpPort(url string) (ip string, port string) {
	var IP string
	var PORT string
	HOSTNAME := strings.Split(url, "/")[2]

	if len(strings.Split(HOSTNAME, ":")) == 2 {
		IP = strings.Split(HOSTNAME, ":")[0]
		PORT = strings.Split(HOSTNAME, ":")[1]
	} else if len(strings.Split(HOSTNAME, ":")) == 1 {
		IP = strings.Split(HOSTNAME, ":")[0]
		PORT = "80"
	}
	return IP, PORT
}
