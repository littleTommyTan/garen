package service

import (
	"github.com/tommytan/garen/internal/helpers"
	"log"
	"mime/multipart"
)

func (s *Service) BucketUpload(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	err := s.Dao.Bucket.PutObject(helpers.UUID()+"-"+fileHeader.Filename, file)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return "https://tommytan-oss.oss-cn-shanghai.aliyuncs.com/" + fileHeader.Filename, nil
}
