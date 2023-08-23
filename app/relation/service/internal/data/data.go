package data

import (
	"Atreus/app/relation/service/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var ProviderSet = wire.NewSet(NewData, NewRelationRepo, NewUserRepo, NewMysqlConn, NewRedisConn)

// CacheClient relation 服务的 Redis 缓存客户端
type CacheClient struct {
	followRelation   *redis.Client // 用户关注关系缓存
	followedRelation *redis.Client // 用户被关注关系缓存
}

type Data struct {
	db    *gorm.DB
	cache *CacheClient
	log   *log.Helper
}

func NewData(db *gorm.DB, cache *CacheClient, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "data/comment"))
	// 关闭Redis连接
	cleanup := func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			_, err := cache.followedRelation.Ping(context.Background()).Result()
			if err != nil {
				return
			}
			if err = cache.followedRelation.Close(); err != nil {
				logHelper.Errorf("Redis connection closure failed, err: %w", err)
			}
		}()
		go func() {
			defer wg.Done()
			_, err := cache.followRelation.Ping(context.Background()).Result()
			if err != nil {
				return
			}
			if err = cache.followRelation.Close(); err != nil {
				logHelper.Errorf("Redis connection closure failed, err: %w", err)
			}
		}()
		wg.Wait()
		logHelper.Info("Successfully close the Redis connection")
	}

	data := &Data{
		db:    db.Model(&Followers{}),
		cache: cache,
		log:   logHelper,
	}
	return data, cleanup, nil
}

// NewMysqlConn mysql数据库连接
func NewMysqlConn(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Database connection failure, err : %v", err)
	}
	InitDB(db)
	log.Info("Database enabled successfully!")
	return db
}

// NewRedisConn Redis数据库连接
func NewRedisConn(c *conf.Data) (cache *CacheClient) {
	var wg sync.WaitGroup
	cache = &CacheClient{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		cache.followedRelation = redis.NewClient(&redis.Options{
			DB:           int(c.Redis.FollowedRelationDb),
			Addr:         c.Redis.Addr,
			Username:     c.Redis.Username,
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
			Password:     c.Redis.Password,
		})

		// ping Redis客户端，判断连接是否存在
		_, err := cache.followedRelation.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Redis database connection failure, err : %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		cache.followRelation = redis.NewClient(&redis.Options{
			DB:           int(c.Redis.FollowRelationDb),
			Addr:         c.Redis.Addr,
			Username:     c.Redis.Username,
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
			Password:     c.Redis.Password,
		})

		// ping Redis客户端，判断连接是否存在
		_, err := cache.followRelation.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Redis database connection failure, err : %v", err)
		}
	}()
	wg.Wait()
	log.Info("Cache enabled successfully!")
	return
}

// InitDB 创建followers数据表，并自动迁移
func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(&Followers{}); err != nil {
		log.Fatalf("Database initialization error, err : %v", err)
	}
}
