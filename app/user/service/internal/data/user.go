package data

import (
	"Atreus/app/user/service/internal/biz"
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

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (*userRepo) initDB(conn *gorms.GormConn) {
	db := conn.Session(context.Background())
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Database %s initialization error, err : %s", userTableName, err.Error())
	}
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	repo := &userRepo{
		data: data,
		log:  log.NewHelper(logger),
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
func (r *userRepo) FindByIds(ctx context.Context, ids []uint32) ([]*biz.User, error) {
	var us []*User
	session := r.data.db.Session(ctx)
	err := session.Model(&User{}).
		Where("id IN ?", ids).
		Find(&us).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*biz.User{}, nil
	}
	if err != nil {
		return nil, err
	}

	bizus := make([]*biz.User, len(us))
	for i := range bizus {
		bizus[i] = us[i].User
	}
	return bizus, nil
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

// UpdateInfo not implement yet
func (r *userRepo) UpdateInfo(ctx context.Context, info *biz.UserInfo) error {
	return errors.New("not implement")
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
