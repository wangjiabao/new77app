package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type EthUserRecord struct {
	ID        int64
	UserId    int64
	Hash      string
	Status    string
	Type      string
	Amount    string
	AmountTwo int64
	CoinType  string
	RelAmount int64
	Last      uint64
	CreatedAt time.Time
}

type Location struct {
	ID           int64
	UserId       int64
	Status       string
	CurrentLevel int64
	Usdt         int64
	Current      int64
	CurrentMax   int64
	Row          int64
	Col          int64
	StopDate     time.Time
	CreatedAt    time.Time
}

type GlobalLock struct {
	ID     int64
	Status int64
}

type RecordUseCase struct {
	ethUserRecordRepo             EthUserRecordRepo
	userRecommendRepo             UserRecommendRepo
	configRepo                    ConfigRepo
	locationRepo                  LocationRepo
	userBalanceRepo               UserBalanceRepo
	userInfoRepo                  UserInfoRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type EthUserRecordRepo interface {
	GetEthUserRecordListByHash(ctx context.Context, hash ...string) (map[string]*EthUserRecord, error)
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
}

type LocationRepo interface {
	GetLocationById(ctx context.Context, id int64) (*LocationNew, error)
	UpdateLocationNewTotalSub(ctx context.Context, id int64, count int64, total int64) error
	UpdateLocationNewNew(ctx context.Context, id int64, userId int64, status string, current int64, amountB int64, biw int64, stopDate time.Time, usdt int64) error
	GetMyLocationLastRunning(ctx context.Context, userId int64) (*LocationNew, error)
	GetLocationsNewByUserId(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetAllLocationsNew(ctx context.Context, currentMax int64) ([]*LocationNew, error)
	CreateLocation(ctx context.Context, rel *Location) (*Location, error)
	GetLocationLast(ctx context.Context) (*Location, error)
	GetLocationDaily(ctx context.Context) ([]*Location, error)
	GetMyLocationLast(ctx context.Context, userId int64) (*Location, error)
	GetMyStopLocationLast(ctx context.Context, userId int64) (*Location, error)
	GetMyLocationRunningLast(ctx context.Context, userId int64) (*Location, error)
	GetLocationsByUserId(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetLocationsByTop(ctx context.Context, top int64) ([]*LocationNew, error)
	GetLocationsByTopTwo(ctx context.Context, top int64) ([]*LocationNew, error)
	GetLocationFirst(ctx context.Context) (*LocationNew, error)
	GetLocationsByUserId2(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetAllLocationsCount(ctx context.Context, usdt int64) int64
	GetRewardLocationByRowOrCol(ctx context.Context, row int64, col int64, locationRowConfig int64) ([]*Location, error)
	GetRewardLocationByIds(ctx context.Context, ids ...int64) (map[int64]*Location, error)
	GetLocationMapByIds(ctx context.Context, userIds ...int64) (map[int64][]*Location, error)
	GetLocationByIds(ctx context.Context, userIds ...int64) ([]*Location, error)
	UpdateLocation(ctx context.Context, id int64, status string, current int64, stopDate time.Time) error
	GetLocations(ctx context.Context, b *Pagination, userId int64) ([]*Location, error, int64)
	UpdateLocationRowAndCol(ctx context.Context, id int64) error
	GetLocationsStopNotUpdate(ctx context.Context) ([]*Location, error)
	LockGlobalLocation(ctx context.Context) (bool, error)
	UnLockGlobalLocation(ctx context.Context) (bool, error)
	LockGlobalWithdraw(ctx context.Context) (bool, error)
	UnLockGlobalWithdraw(ctx context.Context) (bool, error)
	GetLockGlobalLocation(ctx context.Context) (*GlobalLock, error)

	GetMyStopLocationsLast(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetLocationDailyYesterday(ctx context.Context, day int) ([]*LocationNew, error)
	UpdateLocationNewCount(ctx context.Context, id int64, count int64, total int64) error
	UpdateLocationNewCountTwo(ctx context.Context, id int64, count int64, total int64) error
	UpdateLocationNewTotal(ctx context.Context, id int64, count int64, total int64) error
	CreateLocationNew(ctx context.Context, rel *LocationNew, amount int64, amountB int64, address string, coinType string) (*LocationNew, error)
	UpdateLocationNew(ctx context.Context, id int64, userId int64, currentMax int64, amount int64, amountB int64, address string, coinType string) (*LocationNew, error)
}

func NewRecordUseCase(
	ethUserRecordRepo EthUserRecordRepo,
	locationRepo LocationRepo,
	userBalanceRepo UserBalanceRepo,
	userRecommendRepo UserRecommendRepo,
	userInfoRepo UserInfoRepo,
	configRepo ConfigRepo,
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo,
	tx Transaction,
	logger log.Logger) *RecordUseCase {
	return &RecordUseCase{
		ethUserRecordRepo:             ethUserRecordRepo,
		locationRepo:                  locationRepo,
		configRepo:                    configRepo,
		userRecommendRepo:             userRecommendRepo,
		userBalanceRepo:               userBalanceRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		userInfoRepo:                  userInfoRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger),
	}
}

func (ruc *RecordUseCase) GetEthUserRecordByTxHash(ctx context.Context, txHash ...string) (map[string]*EthUserRecord, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordListByHash(ctx, txHash...)
}

func (ruc *RecordUseCase) EthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	return true, nil
}

func (ruc *RecordUseCase) LockEthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	var (
		lock bool
	)
	// todo 全局锁
	for i := 0; i < 3; i++ {
		lock, _ = ruc.locationRepo.LockGlobalLocation(ctx)
		if lock {
			return true, nil
		}
		time.Sleep(5 * time.Second)
	}

	return false, nil
}

func (ruc *RecordUseCase) UnLockEthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	return ruc.locationRepo.UnLockGlobalLocation(ctx)
}
