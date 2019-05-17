package dao

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tommytan/garen/configs"
	"github.com/tommytan/garen/internal/models"
	"log"
)

type Dao struct {
	Db     *gorm.DB
	Redis  *redis.Client
	Bucket *oss.Bucket
}

func New() (dao *Dao) {
	connection, err := gorm.Open("postgres", configs.GetConfiguration().DSN)
	if err != nil {
		log.Print(err)
	}

	connection.AutoMigrate(&models.CarOrder{})
	connection.AutoMigrate(&models.User{})

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: configs.GetConfiguration().RedisPwd,
		PoolSize: 8,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Print("Redis init failed: ", err)
	}
	ossClient, _ := oss.New("http://oss-cn-shanghai.aliyuncs.com/",
		configs.GetConfiguration().OssAccessKey,
		configs.GetConfiguration().OssAccessKeySecret)
	bucket, err := ossClient.Bucket("tommytan-oss")
	if err != nil {
		log.Print("Oss init failed: ", err)
	}
	return &Dao{Db: connection, Redis: redisClient, Bucket: bucket}
}

// Close close the resource.
func (d *Dao) Close() {
	_ = d.Db.Close()
}
