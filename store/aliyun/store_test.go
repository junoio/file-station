package aliyun

import (
	"file-station/store"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Api             = os.Getenv("API")
	AccessKeyId     = os.Getenv("ACCESSKEY")
	AccessKeySecret = os.Getenv("ACCESSSEC")
	Bucket          = os.Getenv("BUCKET")
	client          *aliyunOss
)

func TestValidateError(t *testing.T) {
	_, e := NewOss(&Options{Api, AccessKeyId, ""})
	a := assert.New(t)
	a.Error(e, "endpoint accessKeyID accessKeySecret 错误")
}

func TestUpload(t *testing.T) {
	a := assert.New(t)
	client, err := NewOss(&Options{Api, AccessKeyId, AccessKeySecret})
	a.NoError(err)
	var uploader store.Uploader = client
	err = uploader.Upload(Bucket, "testfile", "/Users/fenghengjun/Downloads/123_070215.docx")
	a.NoError(err)
}

// func TestUploadError(t *testing.T) {
// 	a := assert.New(t)
// 	var uploader store.Uploader = client
// 	err := uploader.Upload(Bucket, "testfile", "123")
// 	a.Error(err, "open 123: no such file or directory")
// }
