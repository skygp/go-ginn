package qiniuyun

import (
	"context"
	"ginn/config"
	errMsg "ginn/utils/code"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)

func UploadFile(file multipart.File, fileSize int64) (string, errMsg.CodeType) {
	qiNiuInfo := config.Gconfig.QiNiu
	putPolicy := storage.PutPolicy{
		Scope: qiNiuInfo.Bucket,
	}
	mac := qbox.NewMac(qiNiuInfo.AccessKey, qiNiuInfo.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errMsg.CodeError
	}
	url := qiNiuInfo.ImageUrl + ret.Key
	return url, errMsg.CodeSuccess
}
