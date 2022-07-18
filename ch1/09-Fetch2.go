//对08进行了改进，使用io.Copy(dis, src) 替换ioutil.ReadAll, 从而避免申请一个缓冲区来存储
//并且使用resp.Status变量获得HTTP协议的状态码
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		//在命令行运行时使用重定向 > 将结果输出到文件中
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() //关闭流，防止资源泄漏
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Status)
	}
}
