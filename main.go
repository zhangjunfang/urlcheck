package main

import (
	"fmt"
	"github.com/ghdjl/urlcheck/check"
	"github.com/ghdjl/urlcheck/input"
	"github.com/ghdjl/urlcheck/urllist"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup
	// 获取用户输入参数
	filename, arg := input.HandleUserInput()
	//切割文件，生成URL
	s := urllist.GetLines(filename)
	u := urllist.GetUrls(s)

	u_length := len(u)
	if arg == "all" {
		for i := 0; i < u_length; i++ {
			waitgroup.Add(1)
			go check.CheckAll(u[i], &waitgroup)
		}
		waitgroup.Wait()
	}else if arg == "fail"{
		for i := 0; i < u_length; i++ {
			waitgroup.Add(1)
			go check.CheckFail(u[i], &waitgroup)
		}
		waitgroup.Wait()
	}

	fmt.Println("\033[33m检测完毕\033[0m")
}
