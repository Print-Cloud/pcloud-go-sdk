- **English**
- [中文](./README.zh.md) 

<h1 align="center">Cloud Print Service (CPS) SDK for Go</h1>

## Introduction

- Cloud Print Service (CPS) Go SDK (API: V3.2)

## Requirements

- installing a Go environment which is new than 1.10.x.

## Install

`go get -u github.com/print-cloud/pcloud-go-sdk`

## Usage

```go
func main() {
	// Necessary steps:
	// Instantiating a client object requires passing in the print cloud address, Key APIUrl, Key.
	client := sdk.NewClient(APIUrl,Key)
	// Necessary steps:
	// Instantiate a request object by passing in the printer number and what you want to print.
	request := api.NewPrintMsg(550517385, "Hello,Cloud Print.")
	// Necessary steps:
	// GET GetSend or POST PostSend
	// To invoke the interface you want to access through the Client object, you pass in the request object
	response, _err := client.GetSend(request)
	// Non-sdk exception, direct failure.Additional processing can be added to the actual code.
	if _err != nil {
		panic(_err)
	}
	// Prints the returned JSON string
	fmt.Println(response)
}

```

All API have corresponding usage examples in the [Example](./example/) directory.

Service API:

* [x] PrinterAdd（Examples: [printer_add.go](./example/service/printer_add.go)） 
* [x] PrintMsg（Examples: [print_msg.go](./example/service/print_msg.go)）
* [x] PrinterDel（Examples: [printer_del.go](./example/service/printer_del.go)）
* [x] PrinterEdit（Examples: [printer_add.go](./example/service/printer_edit.go)）
* [x] QueryPrinterList（Examples: [printer_add.go](./example/service/query_printer_list.go)）
* [x] QueryOrderState（Examples: [printer_add.go](./example/service/query_order_state.go)）
* [x] QueryOrderInfoByDate（Examples: [printer_add.go](./example/service/query_order_info_by_date.go)）
* [x] QueryPrinterStatus（Examples: [printer_add.go](./example/service/query_printer_status.go)）
* [x] DelPrinterSqs（Examples: [printer_add.go](./example/service/del_printer_sqs.go)）
