package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client
var (
	Api             = "https://oss-cn-beijing.aliyuncs.com"
	AccessKeyId     = ""
	AccessKeySecret = ""
	Bucket          = ""
	Path            = ""
)

func init() {
	var err error
	client, err = oss.New(Api, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Fprintln(os.Stderr, "连接失败："+err.Error())
		os.Exit(1)
	}
}

func uploadFile(path string) (string, error) {
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		return "", err
	}
	err = bucket.PutObjectFromFile(path, path)
	if err != nil {
		return "", err
	}
	//打印下载链接
	return bucket.SignURL(path, oss.HTTPGet, 60*60)

}

func validParams() error {
	if Api == "" || AccessKeyId == "" || AccessKeySecret == "" || Bucket == "" {
		return errors.New("参数Api,AccessKeyId,AccessKeySecret,Bucket未设置")
	}

	if Path == "" {
		flag.Usage()
		os.Exit(1)
	}
	return nil
}

func loadParam() {
	//覆盖usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage : command -f filepath \nOptions:")
		flag.PrintDefaults()
	}
	flag.StringVar(&Path, "f", "", "-f filepath")
	flag.Parse()
}

func main() {
	loadParam()
	err := validParams()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err.Error())
		os.Exit(1)
	}
	s, err := uploadFile(Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "上传失败：", err.Error())
		os.Exit(1)
	}
	fmt.Printf("文件上传成功！\n 下载链接：%s\n下载链接有效时间1小时。", s)

}
