package data

import (
	"Atreus/app/user/service/internal/biz"
	"Atreus/app/user/service/internal/server"
	"Atreus/pkg/gorms"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var userTableName = "users"

type User struct {
	*biz.User
	gorm.DeletedAt
}

func (User) TableName() string {
	return userTableName
}

type RelationRepo interface {
	IsFollow(ctx context.Context, userId uint32, toUserId []uint32) ([]bool, error)
}

type userRepo struct {
	data         *Data
	relationRepo RelationRepo
	log          *log.Helper
}

func (*userRepo) initDB(conn *gorms.GormConn) {
	db := conn.Session(context.Background())
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Database %s initialization error, err : %s", userTableName, err.Error())
	}
}

// NewUserRepo .
func NewUserRepo(data *Data, relationRepo server.RelationConn, logger log.Logger) biz.UserRepo {
	repo := &userRepo{
		data:         data,
		relationRepo: NewRelationRepo(relationRepo),
		log:          log.NewHelper(logger),
	}
	conn := data.db
	repo.initDB(conn)
	return repo
}

// Save .
func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{
		User: user,
	}
	session := r.data.db.Session(ctx)
	err := session.Save(&u).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindById .
func (r *userRepo) FindById(ctx context.Context, id uint32) (*biz.User, error) {
	u := new(User)
	session := r.data.db.Session(ctx)
	err := session.Model(&User{}).
		Where("id = ?", id).Limit(1).
		Take(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.User{}, nil
	}
	if err != nil {
		return nil, err
	}

	return u.User, nil
}

// FindByIds .
func (r *userRepo) FindByIds(ctx context.Context, userId uint32, ids []uint32) ([]*biz.User, error) {
	us := make([]*User, len(ids))
	session := r.data.db.Session(ctx)
	err := session.Model(&User{}).
		Where("id IN ?", ids).
		Find(&us).Error
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, nil
	}
	// gorm使用IN查询，如果是根据主键，将会略重复数据。若非主键，则不会忽略
	resultMap := make(map[uint32]*biz.User, len(ids))
	for _, u := range us {
		resultMap[u.Id] = u.User
	}
	result := make([]*biz.User, len(ids))
	for i, id := range ids {
		result[i] = resultMap[id]
	}

	// 登陆用户才需要判断是否关注
	if userId != 0 {
		isFollow, err := r.relationRepo.IsFollow(ctx, userId, ids)
		if err != nil {
			return nil, err
		}
		for i, u := range us {
			u.IsFollow = isFollow[i]
			result[i] = u.User
		}
		return result, nil
	}
	for i, u := range us {
		u.IsFollow = false
		result[i] = u.User
	}
	return result, nil
}

// FindByUsername .
func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	u := new(User)
	session := r.data.db.Session(ctx)
	err := session.Model(&User{}).
		Where("username = ?", username).Limit(1).
		Take(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &biz.User{}, nil
	}
	if err != nil {
		return nil, err
	}

	return u.User, nil
}

// UpdateFollow .
func (r *userRepo) UpdateFollow(ctx context.Context, id uint32, followChange int32) error {
	user, err := r.FindById(ctx, id)
	if err != nil {
		return err
	}
	modifiedValue := calculateValidUint32(user.FollowCount, followChange)
	session := r.data.db.Session(ctx)
	sqlStmt := "update " + userTableName + " set follow_count = ? where id = ?"
	return session.Exec(sqlStmt, modifiedValue, id).Error
}

// UpdateFollower .
func (r *userRepo) UpdateFollower(ctx context.Context, id uint32, followerChange int32) error {
	user, err := r.FindById(ctx, id)
	if err != nil {
		return err
	}
	modifiedValue := calculateValidUint32(user.FollowerCount, followerChange)
	session := r.data.db.Session(ctx)
	sqlStmt := "update " + userTableName + " set follower_count = ? where id = ?"
	return session.Exec(sqlStmt, modifiedValue, id).Error
}

// UpdateFavorited .
func (r *userRepo) UpdateFavorited(ctx context.Context, id uint32, favoritedChange int32) error {
	user, err := r.FindById(ctx, id)
	if err != nil {
		return err
	}
	modifiedValue := calculateValidUint32(user.TotalFavorited, favoritedChange)
	session := r.data.db.Session(ctx)
	sqlStmt := "update " + userTableName + " set total_favorited = ? where id = ?"
	return session.Exec(sqlStmt, modifiedValue, id).Error
}

// UpdateWork .
func (r *userRepo) UpdateWork(ctx context.Context, id uint32, workChange int32) error {
	user, err := r.FindById(ctx, id)
	if err != nil {
		return err
	}
	modifiedValue := calculateValidUint32(user.WorkCount, workChange)
	session := r.data.db.Session(ctx)
	sqlStmt := "update " + userTableName + " set work_count = ? where id = ?"
	return session.Exec(sqlStmt, modifiedValue, id).Error
}

// UpdateFavorite .
func (r *userRepo) UpdateFavorite(ctx context.Context, id uint32, favoriteChange int32) error {
	user, err := r.FindById(ctx, id)
	if err != nil {
		return err
	}
	modifiedValue := calculateValidUint32(user.FavoriteCount, favoriteChange)
	session := r.data.db.Session(ctx)
	sqlStmt := "update " + userTableName + " set favorite_count = ? where id = ?"
	return session.Exec(sqlStmt, modifiedValue, id).Error
}

func calculateValidUint32(src uint32, mod int32) uint32 {
	if mod < 0 {
		mod = -mod
		if src < uint32(mod) {
			return 0
		}
		return src - uint32(mod)
	}
	return src + uint32(mod)
}
