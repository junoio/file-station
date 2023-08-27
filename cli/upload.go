package cli

import (
	"errors"
	"file-station/store"
	"file-station/store/aliyun"
	"file-station/store/aws"
	"file-station/store/tencent"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name:     "accessKey",
		Prompt:   &survey.Input{Message: "Please input access-key:"},
		Validate: survey.Required,
	},
	{
		Name: "accessSecret",
		Prompt: &survey.Password{
			Message: "Please type your accessSecret",
		},
		Validate: survey.Required,
	},
}

var (
	ossProvider string
	// accessKey    string
	// accessSecret string
	bucketName string
	filePath   string
)

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "云中转站，上传文件",
	Long:    "云中转站，上传文件",
	Example: "upload filepath",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		answers := struct {
			AccessKey    string
			AccessSecret string
		}{}
		err = survey.Ask(qs, &answers)
		if err != nil {
			return
		}
		var uploader store.Uploader
		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewOss(&aliyun.Options{
				Endpoint:        "https://oss-cn-beijing.aliyuncs.com",
				AccessKeyID:     answers.AccessKey,
				AccessKeySecret: answers.AccessSecret,
			})
		case "tencent":
			uploader, err = tencent.NewOss(&tencent.Options{})
		case "aws":
			uploader, err = aws.NewOss(&aws.Options{})
		default:
			return errors.New("no support provider")
		}
		if err != nil {
			return
		}
		return uploader.Upload(bucketName, filePath, filePath)
	},
}

func init() {
	RootCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().StringVarP(&ossProvider, "provider", "p", "aliyun", "server for upload aliyun/tencent/aws")
	// uploadCmd.PersistentFlags().StringVarP(&accessKey, "access-key", "k", "", "access key for server")
	// uploadCmd.PersistentFlags().StringVarP(&accessSecret, "access-secret", "s", "", "access secret for server")
	uploadCmd.PersistentFlags().StringVarP(&bucketName, "bucket-name", "b", "", "bucket name for server")
	uploadCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "upload filepath")
}
