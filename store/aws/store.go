package aws

type Options struct {
	// endpoint, accessKeyID, accessKeySecret string
}

func (o *Options) validate() error {
	return nil
}

type AwsOss struct {
}

//上传
func (a *AwsOss) Upload(bucketName, objectKey, fileName string) error {
	return nil
}

func NewOss(op *Options) (*AwsOss, error) {
	err := op.validate()
	if err != nil {
		return nil, err
	}
	return &AwsOss{}, err
}
