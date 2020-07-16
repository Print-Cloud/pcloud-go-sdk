package main

import (
	"fmt"
	sdk "github.com/print-cloud/pcloud-go-sdk"
	"github.com/print-cloud/pcloud-go-sdk/api"
)

func main() {
	// 必要步骤：
	// 实例化一个客户端对象，入参需要传入打印云地址，密钥 APIUrl，Key。
	client := sdk.NewClient(APIUrl, Key)
	// 必要步骤：
	// 实例化一个请求对象，入参需要传入打印机编号
	request := api.NewQueryPrinterStatus(100000000)
	// 必要步骤
	// 通过GET 方式GetSend请求 或者 PostSend请求
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, _err := client.GetSend(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if _err != nil {
		panic(_err)
	}
	// 打印返回的json字符串
	fmt.Println(response)
}
