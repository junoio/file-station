package aliyun

import (
	"errors"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Options struct {
	endpoint, accessKeyID, accessKeySecret string
}

func (o *Options) validate() error {
	if o.endpoint == "" || o.accessKeyID == "" || o.accessKeySecret == "" {
		return errors.New("endpoint accessKeyID accessKeySecret 错误")
	}
	return nil
}

type aliyunOss struct {
	client *oss.Client
}

//上传
func (a *aliyunOss) Upload(bucketName, objectKey, fileName string) error {
	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.PutObjectFromFile(objectKey, fileName)
	if err != nil {
		return err
	}
	//打印下载链接
	str, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载链接(有效期1小时)：%s\n", str)
	return nil
}

func NewOss(op *Options) (*aliyunOss, error) {
	err := op.validate()
	if err != nil {
		return nil, err
	}
	client, err := oss.New(op.endpoint, op.accessKeyID, op.accessKeySecret)
	return &aliyunOss{client}, err
}
