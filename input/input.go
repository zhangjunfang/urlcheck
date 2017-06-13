package input

import (
	"os"
	"fmt"
)

/*
	检测用户输入的参数是否合法
 */
func HandleUserInput() (filename string,arg string) {
	args_lenth := len(os.Args)
	if args_lenth != 3 {
		fmt.Println("usage: \n\turlcheck FILENAME [all|fail]")
		os.Exit(0)
	} else if e,_ :=IsExists(os.Args[1]);!e{
		fmt.Println("file \"" + os.Args[1] + "\" is not exist!")
		os.Exit(0)
	}else if os.Args[2] != "all" && os.Args[2] != "fail"{
		fmt.Println("usage: \n\turlcheck FILENAME [all|fail]")
		os.Exit(0)
	}
	return os.Args[1],os.Args[2]
}

/*
	判断用户输入的文件是否存在
 */
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}