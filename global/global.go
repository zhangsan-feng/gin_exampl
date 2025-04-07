package global

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	_ "github.com/sirupsen/logrus"
	_ "go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var (
	DBConn    *gorm.DB
	DBErr     error
	RedisConn *redis.Client
	MinioConn *minio.Client
	MinioErr  error
)

const (
	PostgresqlAddress = "host=192.168.56.39 user=postgres password=root dbname='manager' port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	MysqlAddress      = "root:root@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=true&loc=Asia/Shanghai"
	RedisAddress      = "192.168.56.39:6379"
	MinioAddress      = "192.168.56.39:9000"
	MinioBucket       = "file-bucket"
	MinioAccessKey    = "5KVMp5GGqwBabT9P2lDq"
	MinioSecretKey    = "TwFsHoFziU58fOvdfOE3qfQvunEDf4cvEWWNg0cL"
)

func newPostgresql() {

	DBConn, DBErr = gorm.Open(postgres.Open(PostgresqlAddress), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if DBErr != nil {
		log.Fatalln("Failed to connect to database:", DBErr)
	}

	sqlDB, dbErr := DBConn.DB()
	if dbErr != nil {
		log.Fatalln("Failed to get underlying DB connection:", dbErr)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func newRedis() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:         RedisAddress, // Redis 服务器地址
		Password:     "",           // 密码，如果没有设置则为空
		DB:           0,            // 使用默认 DB
		DialTimeout:  3 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     100,
	})
	if _, connErr := RedisConn.Ping(context.Background()).Result(); connErr != nil {
		log.Fatalln("redis conn fail:", connErr)
	}

}

func initMinio() {
	MinioConn, MinioErr = minio.New(MinioAddress, &minio.Options{
		Creds: credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
	})
	if MinioErr != nil {
		log.Fatalln("minio client conn fail", MinioErr)
	}

	ctx := context.Background()
	exists, err := MinioConn.BucketExists(ctx, MinioBucket)
	if err != nil {
		log.Fatalln("check bucket fail:", err)
	}

	if !exists {
		err = MinioConn.MakeBucket(ctx, MinioBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln("create bucket fail:", err)
		}
		log.Println("create bucket success:", MinioBucket)
	}
}

func newGfOrm() {

}

func newLogger() {}

func New() {
	newPostgresql()
	newRedis()
	initMinio()
}
