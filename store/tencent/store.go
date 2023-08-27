package tencent

type Options struct {
	// endpoint, accessKeyID, accessKeySecret string
}

func (o *Options) validate() error {
	return nil
}

type TencentOss struct {
}

//上传
func (a *TencentOss) Upload(bucketName, objectKey, fileName string) error {
	return nil
}

func NewOss(op *Options) (*TencentOss, error) {
	err := op.validate()
	if err != nil {
		return nil, err
	}
	return &TencentOss{}, err
}
