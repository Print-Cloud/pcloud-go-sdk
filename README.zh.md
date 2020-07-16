- **中文** 
- [English](./README.md)

<h1 align="center">云打印服务 Cloud Print Service (CPS) SDK for Go</h1>

## 介绍

- 打印云服务 CPS(Cloud Print Service) Go SDK (API: V3.2)
- 这里向您介绍如何获取 [SDK] 并开始调用。

## 环境要求
- Go 环境版本必须不低于 1.12.x.

## 安装

使用 `go get` 下载安装 SDK

`go get -u github.com/print-cloud/pcloud-go-sdk`

## 用法

```go
func main() {
	// 必要步骤：
	// 实例化一个客户端对象，入参需要传入打印云地址，密钥 APIUrl，Key。
	client := sdk.NewClient(APIUrl,Key)
	// 必要步骤：
	// 实例化一个请求对象，入参需要传入打印机编号 和 所需打印内容
	request := api.NewPrintMsg(100000000, "打印云，你好！")
	// 必要步骤
	// 通过GET方式GetSend请求或者POST方式PostSend请求
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, _err := client.GetSend(request)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if _err != nil {
		panic(_err)
	}
	// 打印返回的json字符串
	fmt.Println(response)
}

```

所有的 API 在 [example](./example/) 目录下都有对应的使用示例。

服务API:

* [x] PrinterAdd（使用示例：[printer_add.go](./example/service/printer_add.go)） 
* [x] PrintMsg（使用示例： [print_msg.go](./example/service/print_msg.go)）
* [x] PrinterDel（使用示例： [printer_del.go](./example/service/printer_del.go)）
* [x] PrinterEdit（使用示例：[printer_edit.go](./example/service/printer_edit.go)）
* [x] QueryPrinterList（使用示例：[query_printer_list.go](./example/service/query_printer_list.go)）
* [x] QueryOrderState（使用示例：[query_order_state.go](./example/service/query_order_state.go)）
* [x] QueryOrderInfoByDate（使用示例：[query_order_info_by_date.go](./example/service/query_order_info_by_date.go)）
* [x] QueryPrinterStatus（使用示例：[query_printer_status.go](./example/service/query_printer_status.go)）
* [x] DelPrinterSqs（使用示例：[del_printer_sqs.go](./example/service/del_printer_sqs.go)）
