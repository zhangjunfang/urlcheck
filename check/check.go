package check

import (
	"github.com/ghdjl/urlcheck/output"
	"net/http"
	"net/url"
	"sync"
	"time"
	"net"
	"github.com/ghdjl/urlcheck/urllist"
)

/*
	检查url，并返回所有检查结果
*/

func CheckAll(urlname string, w *sync.WaitGroup) {
	IP, PORT := urllist.IpPort(urlname)

	timeout := make(chan int, 1)
	resChan := make(chan int, 1)

	// 设置超时时间
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- 1
	}()

	/*
		检测URL
	 */
	go func() {
		_, err := net.Dial("tcp", IP+":"+PORT)
		if err != nil {
			resChan <- 502
		}else {
			u, _ := url.Parse(urlname)
			q := u.Query()
			u.RawQuery = q.Encode()
			res, err := http.Get(u.String())
			if err != nil {
				resChan <- 404
			}
			resCode := res.StatusCode
			res.Body.Close()
			if err != nil {
				resChan <- 404
			}
			resChan <- resCode
		}

	}()
	// 设置超时时间，超过timeout设置的时间直接返回502
	for {
		select {
		case <-timeout:
			output.ErrorOutput(urlname, 502)
			w.Done()
			return
		case resCode := <-resChan:
			if resCode == 200 {
				output.InfoOutput(urlname, resCode)
			} else {
				output.ErrorOutput(urlname, resCode)
			}
			w.Done()
			return
		}
	}
}

/*
	检查url，只返回异常的URL
*/
func CheckFail(urlname string, w *sync.WaitGroup) {
	IP, PORT := urllist.IpPort(urlname)

	timeout := make(chan int, 1)
	resChan := make(chan int, 1)

	// 设置超时时间
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- 1
	}()

	/*
		检测URL
 	*/
	go func() {
		_, err := net.Dial("tcp", IP+":"+PORT)
		if err != nil {
			resChan <- 502
		}else {
			u, _ := url.Parse(urlname)
			q := u.Query()
			u.RawQuery = q.Encode()
			res, err := http.Get(u.String())
			if err != nil {
				resChan <- 404
			}
			resCode := res.StatusCode
			res.Body.Close()
			if err != nil {
				resChan <- 404
			}
			resChan <- resCode
		}

	}()
	// 设置超时时间，超过timeout设置的时间直接返回502
	for {
		select {
		case <-timeout:
			output.ErrorOutput(urlname, 502)
			w.Done()
			return
		case resCode := <-resChan:
			if resCode != 200 {
				output.ErrorOutput(urlname, resCode)
			}
			w.Done()
			return
		}
	}
}
