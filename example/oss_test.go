package example_test

import (
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

var (
	Api             = os.Getenv("API")
	AccessKeyId     = os.Getenv("ACCESSKEY")
	AccessKeySecret = os.Getenv("ACCESSSEC")
	Bucket          = os.Getenv("BUCKET")
)

func init() {
	var err error
	client, err = oss.New(Api, AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}
}

func TestListBucket(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Fatal(err)
	}
	for _, bucket := range lsRes.Buckets {
		t.Log(bucket)
	}
}

func TestPutObject(t *testing.T) {
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		t.Fatal(err)
	}
	err = bucket.PutObjectFromFile("testfile1", "/Users/fenghengjun/Downloads/123_070215.docx")
	if err != nil {
		t.Fatal(err)
	}
}

// func TestPutObject(t *testing.T) {
// 	bucket, err := client.Bucket("f-devclould-station")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = bucket.PutObjectFromFile("testfile1", "/Users/fenghengjun/Downloads/123_070215.docx")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
