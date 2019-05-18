package dao

import (
	"github.com/tommytan/garen/internal/helpers"
	"mime/multipart"
)

func (d *Dao) BucketUpload(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	err := d.Bucket.PutObject(helpers.UUID()+"-"+fileHeader.Filename, file)
	if err != nil {
		return "", err
	}
	return "https://tommytan-oss.oss-cn-shanghai.aliyuncs.com/" + fileHeader.Filename, nil
}
