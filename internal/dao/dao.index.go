package dao

import (
	"github.com/go-redis/redis"
	"github.com/tommytan/garen/configs"
	"github.com/tommytan/garen/internal/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Dao struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func New() (dao *Dao) {
	connection, err := gorm.Open("postgres", configs.GetConfiguration().DSN)
	if err != nil {
		log.Fatal(err)
	}

	connection.AutoMigrate(&models.CarOrder{})
	connection.AutoMigrate(&models.User{})

	// GetInstance func return a db connection instance
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: configs.GetConfiguration().RedisPwd, // no password set
		PoolSize: 8,
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal("Redis init failed: ", err)
	}

	return &Dao{Db: connection,Redis:client}
}

// Close close the resource.
func (d *Dao) Close() {
	_ = d.Db.Close()
}
