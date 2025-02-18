package goflow_client

import (
	"fmt"
	"os"
	"testing"
)

func Test_GetFile(t *testing.T) {
	client := Create(Option{
		Http:        "http",
		Ip:          "10.254.47.239",
		Port:        9999,
		ContextPath: "gofs",
		AppKey:      "brain",
		AppSecret:   "brain",
		Env:         "prod",
	})
	file := client.GetFileService().GetFile("582050d6-3bdd-4b1f-a64f-70965855b7ad")
	fmt.Printf("%+v\n", file)
}

func Test_GetFilesByBusinessId(t *testing.T) {
	client := Create(Option{
		Http:        "http",
		Ip:          "10.254.47.239",
		Port:        9999,
		ContextPath: "gofs",
		AppKey:      "brain",
		AppSecret:   "brain",
		Env:         "prod",
	})
	file := client.GetFileService().GetFilesByBusinessId("zwy0001", "case_voice")
	fmt.Printf("%+v\n", file)
}

func Test_GetFileBytesById(t *testing.T) {
	client := Create(Option{
		Http:        "http",
		Ip:          "10.254.47.239",
		Port:        9999,
		ContextPath: "gofs",
		AppKey:      "brain",
		AppSecret:   "brain",
		Env:         "prod",
	})
	data := client.GetFileService().GetFileBytesById("582050d6-3bdd-4b1f-a64f-70965855b7ad")
	os.WriteFile("C:\\Work\\test_word\\upload\\zwy.pdf", data, 0777)
}

func Test_GetFileBytesByStorePath(t *testing.T) {
	client := Create(Option{
		Http:        "http",
		Ip:          "10.254.47.239",
		Port:        9999,
		ContextPath: "gofs",
		AppKey:      "brain",
		AppSecret:   "brain",
		Env:         "prod",
	})
	data := client.GetFileService().GetFileBytesByStorePath("/temp/zwy_yf/加油站签名_20250218115431.docx")
	os.WriteFile("C:\\Work\\test_word\\upload\\加油站签名_20250218115431.docx", data, 0777)
}

func Test_UploadTempFile(t *testing.T) {
	client := Create(Option{
		Http:        "http",
		Ip:          "10.254.47.239",
		Port:        9999,
		ContextPath: "gofs",
		AppKey:      "brain",
		AppSecret:   "brain",
		Env:         "prod",
	})

	data, _ := os.ReadFile("C:\\Work\\test_word\\upload\\加油站签名.docx")
	file := client.GetFileService().UploadTempFile(data, "加油站签名.docx", "zwy_yf", true)
	fmt.Printf("%+v\n", file)
}
