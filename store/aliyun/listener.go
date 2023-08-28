package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

type ProgressListener struct {
	bar *progressbar.ProgressBar //进度条
}

func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		// p.bar = progressbar.DefaultBytes(
		// 	event.TotalBytes,
		// 	"uploading",
		// )

		p.bar = progressbar.NewOptions(int(event.TotalBytes),
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //修复windows 进度条
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(50),
			progressbar.OptionSetDescription("uploading"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))

	case oss.TransferDataEvent:
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Println("上传完成")
	case oss.TransferFailedEvent:
		fmt.Println("上传失败")
	}
}
