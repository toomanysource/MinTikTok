package data

import (
	"Atreus/app/comment/service/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var ProviderSet = wire.NewSet(NewData, NewCommentRepo, NewUserRepo, NewMysqlConn, NewRedisConn)

// CacheClient comment服务的缓存客户端
type CacheClient struct {
	commentNumber *redis.Client
	commentList   *redis.Client
}

type Data struct {
	db    *gorm.DB
	cache *CacheClient
	log   *log.Helper
}

func NewData(db *gorm.DB, cacheClient *CacheClient, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "data/comment"))
	// 并发关闭所有数据库连接，后期根据Redis与Mysql是否数据同步修改
	cleanup := func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqlDB, err := db.DB()
			// 如果err不为空，则连接池中没有连接
			if err != nil {
				return
			}
			if err = sqlDB.Close(); err != nil {
				logHelper.Errorf("Mysql connection closure failed, err: %w", err)
				return
			}
			logHelper.Info("Successfully close the Mysql connection")
		}()
		wg.Add(2)
		go func() {
			defer wg.Done()
			_, err := cacheClient.commentNumber.Ping(context.Background()).Result()
			if err != nil {
				return
			}
			if err = cacheClient.commentNumber.Close(); err != nil {
				logHelper.Errorf("Redis connection closure failed, err: %w", err)
			}
			logHelper.Info("Successfully close the Redis connection")
		}()
		go func() {
			defer wg.Done()
			_, err := cacheClient.commentList.Ping(context.Background()).Result()
			if err != nil {
				return
			}
			if err = cacheClient.commentList.Close(); err != nil {
				logHelper.Errorf("Redis connection closure failed, err: %w", err)
			}
			logHelper.Info("Successfully close the Redis connection")
		}()
		wg.Wait()
	}

	data := &Data{
		db:    db,
		cache: cacheClient,
		log:   logHelper,
	}
	return data, cleanup, nil
}

// NewMysqlConn mysql数据库连接
func NewMysqlConn(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn))
	if err != nil {
		log.Fatalf("Database connection failure, err : %w", err)
	}
	InitDB(db)
	log.Info("Database enabled successfully!")
	return db
}

// NewRedisConn Redis数据库连接, 并发开启连接提高速率
func NewRedisConn(c *conf.Data) (cacheClient *CacheClient) {
	cacheClient = &CacheClient{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		client := redis.NewClient(&redis.Options{
			DB:           int(c.Redis.CommentNumberDb),
			Addr:         c.Redis.Addr,
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
			Password:     c.Redis.Password,
		})

		// ping Redis客户端，判断连接是否存在
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Redis database connection failure, err : %w", err)
		}
		cacheClient.commentNumber = client
		log.Info("CommentNumberCache enabled successfully!")
	}()
	go func() {
		defer wg.Done()
		client := redis.NewClient(&redis.Options{
			DB:           int(c.Redis.CommentListDb),
			Addr:         c.Redis.Addr,
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
			Password:     c.Redis.Password,
		})

		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Redis database connection failure, err : %w", err)
		}
		cacheClient.commentList = client
		log.Info("CommentListCache enabled successfully!")
	}()
	wg.Wait()
	return
}

// InitDB 创建User数据表，并自动迁移
func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(&Comment{}); err != nil {
		log.Fatalf("Database initialization error, err : %w", err)
	}
}
