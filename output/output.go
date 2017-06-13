package output

import (
	"fmt"
	"strconv"
)

/*
	格式化输出异常URL，红色
 */
func ErrorOutput(url string, resCode int) {
	fmt.Println("\033[31;1m" + url + "\033[0m" + " " + "\033[31;1m" + strconv.Itoa(resCode) + "\033[0m")
}

/*
	输出正常URL,绿色
 */
func InfoOutput(url string, resCode int) {
	fmt.Println("\033[32m" + url + "\033[0m" + " " + "\033[32m" + strconv.Itoa(resCode) + "\033[0m")
}
