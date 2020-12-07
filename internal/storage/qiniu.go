package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	qiniustorage "github.com/qiniu/api.v7/v7/storage"
)

type _QiniuConfig struct {
	Domain string
	Bucket string
	Access string
	Secret string
}

var (
	qiniuConfig *_QiniuConfig
)

func InitQiniu(domain, bucket, access, secret string) {
	qiniuConfig = &_QiniuConfig{
		Domain: domain,
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
	return uploader.Put(context.Background(), nil, token, name, reader, reader.Size(), &qiniustorage.PutExtra{})
}

func DowloadFromQiniu(name string) ([]byte, error) {
	mac := qbox.NewMac(qiniuConfig.Access, qiniuConfig.Secret)
	deadline := time.Now().Add(5 * time.Minute).Unix()
	url := qiniustorage.MakePrivateURL(mac, qiniuConfig.Domain, name, deadline)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
