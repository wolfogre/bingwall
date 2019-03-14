package storage

import (
	"bytes"
	"context"

	"github.com/qiniu/api.v7/auth/qbox"
	qiniustorage "github.com/qiniu/api.v7/storage"
)

type _QiniuConfig struct {
	Bucket string
	Access string
	Secret string
}

var (
	qiniuConfig *_QiniuConfig
)

func InitQiniu(bucket, access, secret string) {
	qiniuConfig = &_QiniuConfig{
		Bucket: bucket,
		Access: access,
		Secret: secret,
	}
}

func UploadToQiniu(name string, content []byte) error {
	policy := &qiniustorage.PutPolicy{
		Scope: qiniuConfig.Bucket,
	}
	token := policy.UploadToken(qbox.NewMac(qiniuConfig.Access, qiniuConfig.Secret))
	uploader := qiniustorage.NewFormUploader(&qiniustorage.Config{
		Zone:          &qiniustorage.ZoneHuadong, // TODO: need to support specified zone
		UseHTTPS:      false,
		UseCdnDomains: false,
	})
	reader := bytes.NewReader(content)
	if err := uploader.Put(context.Background(), nil, token, name, reader, reader.Size(), &qiniustorage.PutExtra{}); err != nil {
		return err
	}
}