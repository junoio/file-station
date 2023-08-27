package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

type ProgressListener struct {
	bar *progressbar.ProgressBar //进度条
}

func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		p.bar = progressbar.DefaultBytes(
			event.TotalBytes,
			"uploading",
		)
	case oss.TransferDataEvent:
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Println("上传完成")
	case oss.TransferFailedEvent:
		fmt.Println("上传失败")
	}
}
