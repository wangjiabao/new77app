package biz

import (
	"context"
	"crypto/ecdsa"
	v1 "dhb/app/app/api"
	"encoding/base64"
	"fmt"
	sdk "github.com/BioforestChain/go-bfmeta-wallet-sdk"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID              int64
	Address         string
	Password        string
	Undo            int64
	AddressTwo      string
	PrivateKey      string
	AddressThree    string
	WordThree       string
	PrivateKeyThree string
	Last            uint64
	Amount          uint64
	AmountBiw       uint64
	Total           uint64
	IsDelete        int64
	Out             int64
	CreatedAt       time.Time
	Lock            int64
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	UseVip           int64
	HistoryRecommend int64
	TeamCsdBalance   int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	Total         int64
	CreatedAt     time.Time
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
	CreatedAt     time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID             int64
	UserId         int64
	BalanceUsdt    int64
	BalanceUsdt2   int64
	BalanceDhb     int64
	RecommendTotal int64
	AreaTotal      int64
	FourTotal      int64
	LocationTotal  int64
	BalanceC       int64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	CreatedAt       time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Num               int64
	Status            string
	Current           int64
	CurrentMax        int64
	CurrentMaxNew     int64
	StopLocationAgain int64
	OutRate           int64
	Count             int64
	StopCoin          int64
	Top               int64
	Usdt              int64
	Total             int64
	TotalTwo          int64
	TotalThree        int64
	Biw               int64
	TopNum            int64
	LastLevel         int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Amount    int64
	CoinType  string
	Balance   int64
	Type      string
	CreatedAt time.Time
}

type BalanceReward struct {
	ID        int64
	UserId    int64
	Status    int64
	Amount    int64
	SetDate   time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          int64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	CreatedAt        time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
}

type UserBalanceRepo interface {
	RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	CreateUserBalanceLock(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context) (*Reward, error)
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetRewardFourYes(ctx context.Context) (*Reward, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetLocationsToday(ctx context.Context) ([]*LocationNew, error)
	GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserSortRecommendReward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, relAmount int64, amount int64, amountFee int64, coinType string, address string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64, tmpRecommendUserIdsInt []int64) error
	WithdrawUsdt2(ctx context.Context, userId int64, amount int64) error
	Exchange(ctx context.Context, userId int64, amount int64, amountUsdtSubFee int64, amountUsdt int64, locationId int64) error
	Exchange2(ctx context.Context, userId int64, amount int64, amountUsdtSubFee int64, amountUsdt int64, locationId int64) error
	WithdrawUsdt3(ctx context.Context, userId int64, amount int64) error
	TranUsdt(ctx context.Context, userId int64, toUserId int64, amount int64, tmpRecommendUserIdsInt []int64, tmpRecommendUserIdsInt2 []int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	WithdrawC(ctx context.Context, userId int64, amount int64) error
	TranDhb(ctx context.Context, userId int64, toUserId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64, typeCoin string) ([]*Withdraw, error)
	GetWithdrawByUserId2(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetUserBalanceRecordByUserId(ctx context.Context, userId int64, typeCoin string, tran string) ([]*UserBalanceRecord, error)
	GetUserBalanceRecordsByUserId(ctx context.Context, userId int64) ([]*UserBalanceRecord, error)
	GetTradeByUserId(ctx context.Context, userId int64) ([]*Trade, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*UserSortRecommendReward, error)

	SetBalanceReward(ctx context.Context, userId int64, amount int64) error
	UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error
	GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*BalanceReward, error)

	GetUserBalanceLock(ctx context.Context, userId int64) (*UserBalance, error)
	Trade(ctx context.Context, userId int64, amount int64, amountB int64, amountRel int64, amountBRel int64, tmpRecommendUserIdsInt []int64, amount2 int64) error
}

type UserRecommendRepo interface {
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	UpdateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	CreateUserRecommendArea(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error)
	GetUserRecommendLowArea(ctx context.Context, code string) ([]*UserRecommendArea, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
	UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
}

type UserRepo interface {
	GetEthUserRecordListByUserId(ctx context.Context, userId int64) ([]*EthUserRecord, error)
	InRecordNew(ctx context.Context, userId int64, address string, amount int64, coinType string) error
	UpdateUserNewTwoNew(ctx context.Context, userId int64, amountUsdt uint64, amountBiw uint64, coinType string) error
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUserByUserIdsTwo(ctx context.Context, userIds []int64) (map[int64]*User, error)
	GetUsers(ctx context.Context, b *Pagination, address string) ([]*User, error, int64)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user          *User
		recommendUser *UserRecommend
		err           error
		userId        int64
		decodeBytes   []byte
	)

	user, err = uuc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			decodeBytes, err = base64.StdEncoding.DecodeString(code)
			code = string(decodeBytes)
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}
			if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}

			var (
				locationNew []*LocationNew
			)
			locationNew, err = uuc.locationRepo.GetLocationsByUserId(ctx, userId)
			if nil != err {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}

			if 0 == len(locationNew) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}

			// 查询推荐人的相关信息
			recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
			if err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
			}
		}

		// 创建私钥
		var (
			address    string
			privateKey string
		)
		address, privateKey, err = generateKey()
		if 0 >= len(address) || 0 >= len(privateKey) || err != nil {
			return nil, errors.New(500, "USER_ERROR", "生成地址错误")
		}

		u.PrivateKey = privateKey
		u.AddressTwo = address

		var (
			addressThree    string
			privateKeyThree string
		)
		u.WordThree = generateWord() // 生成助剂词
		if 20 >= len(u.WordThree) {
			return nil, errors.New(500, "USER_ERROR", "生成助记词错误")
		}

		privateKeyThree, addressThree = generateKeyBiw(u.WordThree)
		if 0 >= len(addressThree) || 0 >= len(privateKeyThree) || err != nil {
			return nil, errors.New(500, "USER_ERROR", "生成地址错误")
		}

		u.AddressThree = addressThree
		u.PrivateKeyThree = privateKeyThree

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uuc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			_, err = uuc.uiRepo.CreateUserInfo(ctx, user) // 创建用户信息
			if err != nil {
				return err
			}

			_, err = uuc.urRepo.CreateUserRecommend(ctx, user, recommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			_, err = uuc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func generateWord() string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 定义总组数和每组的字符数
	numGroups := 12
	groupSize := 3

	// 生成随机字符串
	var result []string
	for i := 0; i < numGroups; i++ {
		result = append(result, randString(groupSize))
	}

	// 将字符串数组用逗号连接
	finalString := strings.Join(result, ",")
	return finalString
}

func randString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var sdkClient = sdk.NewBCFWalletSDK()
var bCFSignUtil = sdkClient.NewBCFSignUtil("b")

func generateKeyBiw(word string) (string, string) {
	bCFSignUtil_CreateKeypair, _ := bCFSignUtil.CreateKeypair(word)
	got, _ := bCFSignUtil.GetAddressFromPublicKey(bCFSignUtil_CreateKeypair.PublicKey.StringBuffer, "b")
	return bCFSignUtil_CreateKeypair.SecretKey.Value, got
}

func generateKey() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", nil
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}

func (uuc *UserUseCase) UpdateUserRecommend(ctx context.Context, u *User, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	var (
		err                   error
		userId                int64
		recommendUser         *UserRecommend
		userRecommend         *UserRecommend
		locations             []*LocationNew
		myRecommendUser       *User
		myUserRecommendUserId int64
		Address               string
		decodeBytes           []byte
	)

	code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
	if "abf00dd52c08a9213f225827bc3fb100" != code {
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 1 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}
		if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 现有推荐人信息，判断推荐人是否改变
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, u.ID)
		if nil == userRecommend {
			return nil, err
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return nil, err
			}
		}

		if nil == myRecommendUser {
			return &v1.RecommendUpdateReply{InviteUserAddress: ""}, nil
		}

		if myRecommendUser.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		if u.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 我的占位信息
		locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, u.ID)
		if nil != err {
			return nil, err
		}
		if nil != locations && 0 < len(locations) {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 查询推荐人的相关信息
		recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 推荐人信息
		myRecommendUser, err = uuc.repo.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		// 更新
		_, err = uuc.urRepo.UpdateUserRecommend(ctx, u, recommendUser)
		if err != nil {
			return nil, err
		}
		Address = myRecommendUser.Address
	}

	return &v1.RecommendUpdateReply{InviteUserAddress: Address}, err
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	var (
		err                   error
		myUser                *User
		userRecommend         *UserRecommend
		myCode                string
		encodeString          string
		myUserRecommendUserId int64
		inviteUserAddress     string
		myRecommendUser       *User
		//userInfo      *UserInfo
		configs               []*Config
		locations             []*LocationNew
		userBalance           *UserBalance
		myLocations           []*v1.UserInfoReply_List
		bPrice                int64
		cPrice                int64
		bPriceBase            int64
		buyOne                int64
		buyTwo                int64
		buyThree              int64
		buyFour               int64
		buyFive               int64
		buySix                int64
		areaMin               int64
		areaMax               int64
		areaAll               int64
		locationUsdt          string
		locationCurrent       string
		locationCurrentMaxSub string
		locationBiw           int64
		total                 int64
		one                   int64
		two                   int64
		three                 int64
		four                  int64
		exchangeRate          int64
		exchangeRateC         int64
		lastLevel             int64 = -1
		areaOne               int64
		areaTwo               int64
		areaThree             int64
		areaFour              int64
		areaFive              int64
		configThree           string
		configFour            string
		status                = "stop"
		totalYesReward        int64
		buyLimit              int64
		withdrawMin           int64
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"b_price",
		"c_price",
		"exchange_rate_c",
		"exchange_rate",
		"b_price_base",
		"buy_one",
		"buy_two",
		"buy_three",
		"buy_four",
		"buy_five", "buy_six",
		"total",
		"one", "two", "three", "four",
		"area_one", "area_two", "area_three", "area_four", "area_five",
		"config_one", "config_two", "config_three", "config_four", "withdraw_amount_min", "buy_limit",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_limit" == vConfig.KeyName {
				buyLimit, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "c_price" == vConfig.KeyName {
				cPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "exchange_rate" == vConfig.KeyName {
				exchangeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "exchange_rate_c" == vConfig.KeyName {
				exchangeRateC, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_one" == vConfig.KeyName {
				buyOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_two" == vConfig.KeyName {
				buyTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_three" == vConfig.KeyName {
				buyThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_four" == vConfig.KeyName {
				buyFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_five" == vConfig.KeyName {
				buyFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_six" == vConfig.KeyName {
				buySix, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "total" == vConfig.KeyName {
				total, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "one" == vConfig.KeyName {
				one, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "two" == vConfig.KeyName {
				two, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "three" == vConfig.KeyName {
				three, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "four" == vConfig.KeyName {
				four, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_one" == vConfig.KeyName {
				areaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_two" == vConfig.KeyName {
				areaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_three" == vConfig.KeyName {
				areaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_four" == vConfig.KeyName {
				areaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_five" == vConfig.KeyName {
				areaFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "config_three" == vConfig.KeyName {
				configThree = vConfig.Value
			}
			if "config_four" == vConfig.KeyName {
				configFour = vConfig.Value
			}

		}
	}

	myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if 1 == myUser.IsDelete {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "用户已删除")
	}

	//userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}

	// 系统

	count6 := uuc.locationRepo.GetAllLocationsCount(ctx, 10000000)
	count1 := uuc.locationRepo.GetAllLocationsCount(ctx, 30000000)
	count2 := uuc.locationRepo.GetAllLocationsCount(ctx, 100000000)
	count3 := uuc.locationRepo.GetAllLocationsCount(ctx, 300000000)
	count4 := uuc.locationRepo.GetAllLocationsCount(ctx, 500000000)
	count5 := uuc.locationRepo.GetAllLocationsCount(ctx, 1000000000)

	// 入金
	locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}
	var (
		currentAmountBiw string
	)
	myLocations = make([]*v1.UserInfoReply_List, 0)
	if nil != locations && 0 < len(locations) {

		var tmpC int64
		if locations[0].Current <= locations[0].CurrentMax {
			tmpC = locations[0].CurrentMax - locations[0].Current
		}
		myLocations = append(myLocations, &v1.UserInfoReply_List{
			Current:              fmt.Sprintf("%.2f", float64(locations[0].Current)/float64(100000)),
			CurrentMaxSubCurrent: fmt.Sprintf("%.2f", float64(tmpC)/float64(100000)),
			Amount:               fmt.Sprintf("%.2f", float64(locations[0].CurrentMax)/float64(100000)/2.5),
		})

		for _, v := range locations {
			var tmp int64
			if v.Current <= v.CurrentMax {
				tmp = v.CurrentMax - v.Current
			}

			locationBiw += v.Biw

			if "running" == v.Status {
				status = "running"
				currentAmountBiw = fmt.Sprintf("%.2f", float64(tmp)/float64(100000))
				areaAll = v.Total + v.TotalThree + v.TotalTwo
				if v.TotalTwo >= v.Total && v.TotalTwo >= v.TotalThree {
					areaMax = v.TotalTwo
					areaMin = v.Total + v.TotalThree
				}
				if v.Total >= v.TotalTwo && v.Total >= v.TotalThree {
					areaMax = v.Total
					areaMin = v.TotalTwo + v.TotalThree
				}
				if v.TotalThree >= v.Total && v.TotalThree >= v.TotalTwo {
					areaMax = v.TotalThree
					areaMin = v.TotalTwo + v.Total
				}
				locationUsdt = fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000))

				locationCurrent = fmt.Sprintf("%.2f", float64(v.Current)/float64(100000))
				locationCurrentMaxSub = fmt.Sprintf("%.2f", float64(v.CurrentMax-v.Current)/float64(100000))
			}

			//if "stop" == v.Status {
			//	stopCount++
			//}

			//myLocations = append(myLocations, &v1.UserInfoReply_List{
			//	Current:              fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			//	CurrentMaxSubCurrent: fmt.Sprintf("%.2f", float64(tmp)/float64(100000)),
			//	Amount:               fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)/2.5),
			//})
			var tmpLastLevel int64
			// 1大区
			if v.Total >= v.TotalTwo && v.Total >= v.TotalThree {
				if areaOne <= v.TotalTwo+v.TotalThree {
					tmpLastLevel = 1
				}
				if areaTwo <= v.TotalTwo+v.TotalThree {
					tmpLastLevel = 2
				}
				if areaThree <= v.TotalTwo+v.TotalThree {
					tmpLastLevel = 3
				}
				if areaFour <= v.TotalTwo+v.TotalThree {
					tmpLastLevel = 4
				}
				if areaFive <= v.TotalTwo+v.TotalThree {
					tmpLastLevel = 5
				}
			} else if v.TotalTwo >= v.Total && v.TotalTwo >= v.TotalThree {
				if areaOne <= v.Total+v.TotalThree {
					tmpLastLevel = 1
				}
				if areaTwo <= v.Total+v.TotalThree {
					tmpLastLevel = 2
				}
				if areaThree <= v.Total+v.TotalThree {
					tmpLastLevel = 3
				}
				if areaFour <= v.Total+v.TotalThree {
					tmpLastLevel = 4
				}
				if areaFive <= v.Total+v.TotalThree {
					tmpLastLevel = 5
				}
			} else if v.TotalThree >= v.Total && v.TotalThree >= v.TotalTwo {
				if areaOne <= v.TotalTwo+v.Total {
					tmpLastLevel = 1
				}
				if areaTwo <= v.TotalTwo+v.Total {
					tmpLastLevel = 2
				}
				if areaThree <= v.TotalTwo+v.Total {
					tmpLastLevel = 3
				}
				if areaFour <= v.TotalTwo+v.Total {
					tmpLastLevel = 4
				}
				if areaFive <= v.TotalTwo+v.Total {
					tmpLastLevel = 5
				}
			}

			if tmpLastLevel > lastLevel {
				lastLevel = tmpLastLevel
			}

			if v.LastLevel > lastLevel {
				lastLevel = v.LastLevel
			}
		}
	}

	// 余额，收益总数
	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	if nil == userRecommend {
		return nil, err
	}

	myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	codeByte := []byte(myCode)
	encodeString = base64.StdEncoding.EncodeToString(codeByte)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
		myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
		if nil != err {
			return nil, err
		}
		inviteUserAddress = myRecommendUser.Address
		myCode = userRecommend.RecommendCode + myCode
	}

	var (
		myUserRecommend []*UserRecommend
		recommendTotal  int64
	)
	myUserRecommend, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
	if nil != err {
		return nil, err
	}
	myRecommendList := make([]*v1.UserInfoReply_ListRecommend, 0)
	if nil != myUserRecommend {
		for _, vMyUserRecommend := range myUserRecommend {
			var (
				tmpMyRecommendLocations []*LocationNew
			)
			tmpMyRecommendLocations, err = uuc.locationRepo.GetLocationsByUserId(ctx, vMyUserRecommend.UserId)
			if nil != err {
				return nil, err
			}

			if 0 < len(tmpMyRecommendLocations) {
				recommendTotal++
				var (
					myAllRecommendUser *User
				)
				myAllRecommendUser, err = uuc.repo.GetUserById(ctx, vMyUserRecommend.UserId)
				if nil != err {
					return nil, err
				}

				if nil == myAllRecommendUser {
					continue
				}

				myRecommendList = append(myRecommendList, &v1.UserInfoReply_ListRecommend{Address: myAllRecommendUser.Address})
			}
		}
	}

	// 提现
	var (
		withdraws      []*Withdraw
		withdrawAmount int64
		withdrawList   []*v1.UserInfoReply_ListWithdraw
	)

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	withdrawList = make([]*v1.UserInfoReply_ListWithdraw, 0)
	for _, v := range withdraws {
		if "usdt" == v.Type {
			withdrawAmount += v.RelAmount
		}

		withdrawList = append(withdrawList, &v1.UserInfoReply_ListWithdraw{
			Amount:   fmt.Sprintf("%.2f", float64(v.RelAmount)/float64(100000)),
			CreateAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}
	startDate := time.Now().AddDate(0, 0, -1)

	// 分红
	var (
		userRewards []*Reward
	)
	listReward := make([]*v1.UserInfoReply_ListReward, 0)
	userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	tmpUserIds := make(map[int64]int64, 0)
	for _, vUserRewards := range userRewards {
		if 0 >= vUserRewards.TypeRecordId {
			continue
		}
		tmpUserIds[vUserRewards.TypeRecordId] = vUserRewards.TypeRecordId
	}

	var (
		addressUsers map[int64]*User
	)
	if 0 < len(tmpUserIds) {
		tmpAddressUserIds := make([]int64, 0)
		for _, v := range tmpUserIds {
			tmpAddressUserIds = append(tmpAddressUserIds, v)
		}

		addressUsers, err = uuc.repo.GetUserByUserIdsTwo(ctx, tmpAddressUserIds)
	}

	tmpMyLocations := make([]*v1.UserInfoReply_List, 0)
	if nil != userRewards {
		for _, vUserReward := range userRewards {
			if "location" == vUserReward.Reason {
				if vUserReward.CreatedAt.After(startDate) {
					totalYesReward += vUserReward.Amount
				}
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)) + "BIW",
					Type:      1,
				})
			} else if "area" == vUserReward.Reason {
				if vUserReward.CreatedAt.After(startDate) {
					totalYesReward += vUserReward.Amount
				}
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)) + "BIW",
					Type:      4,
				})
			} else if "recommend" == vUserReward.Reason {
				if vUserReward.CreatedAt.After(startDate) {
					totalYesReward += vUserReward.Amount
				}

				addressTmp := ""
				if nil != addressUsers {
					if _, ok := addressUsers[vUserReward.TypeRecordId]; ok {
						addressTmp = addressUsers[vUserReward.TypeRecordId].Address
					}
				}
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)) + "BIW",
					Type:      2,
					Num:       vUserReward.ReasonLocationId,
					Address:   addressTmp,
				})
			} else if "recommend_location" == vUserReward.Reason {
				if vUserReward.CreatedAt.After(startDate) {
					totalYesReward += vUserReward.Amount
				}
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)),
					Type:      8,
				})
			} else if "four" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)),
					Type:      3,
				})
			} else if "exchange" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)) + "BIW",
					Type:      5,
				})
			} else if "exchange_2" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)) + "BIW",
					Type:      5,
				})
			} else if "withdraw" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)/float64(100000)),
					Type:      6,
				})
			} else if "buy" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    fmt.Sprintf("%.4f", float64(vUserReward.Amount)),
					Type:      7,
				})
				tmpMyLocations = append(tmpMyLocations, &v1.UserInfoReply_List{
					Current:              fmt.Sprintf("%.2f", float64(vUserReward.Amount)*2.5),
					CurrentMaxSubCurrent: "0.00",
					Amount:               fmt.Sprintf("%.2f", float64(vUserReward.Amount)),
				})
			} else {
				continue
			}
		}
	}

	for k, vTmpMyLocations := range tmpMyLocations {
		if 0 == k {
			continue
		}
		myLocations = append(myLocations, vTmpMyLocations)
	}

	// 充值
	var (
		userEth []*EthUserRecord
	)
	userEth, err = uuc.repo.GetEthUserRecordListByUserId(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}
	listUserEth := make([]*v1.UserInfoReply_ListEthRecord, 0)
	for _, vUserEth := range userEth {
		coinType := "BIW"
		if "USDT" == vUserEth.CoinType {
			coinType = vUserEth.CoinType
		}
		listUserEth = append(listUserEth, &v1.UserInfoReply_ListEthRecord{
			Amount:    uint64(vUserEth.AmountTwo),
			CoinType:  coinType,
			CreatedAt: vUserEth.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	// 全球
	var (
		day                    = -1
		userLocationsYes       []*LocationNew
		rewardLocationYes      int64
		totalRewardYes         int64
		fourUserRecommendTotal map[int64]int64
	)

	fourUserRecommendTotal = make(map[int64]int64, 0)
	userLocationsYes, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	for _, userLocationYes := range userLocationsYes {
		rewardLocationYes += userLocationYes.Usdt

		// 获取直推

		var (
			fourUserRecommend         *UserRecommend
			myFourUserRecommendUserId int64
			//myFourRecommendUser *User
		)
		fourUserRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userLocationYes.UserId)
		if nil == fourUserRecommend {
			continue
		}

		if "" != fourUserRecommend.RecommendCode {
			tmpFourRecommendUserIds := strings.Split(fourUserRecommend.RecommendCode, "D")
			if 2 <= len(tmpFourRecommendUserIds) {
				myFourUserRecommendUserId, _ = strconv.ParseInt(tmpFourRecommendUserIds[len(tmpFourRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			//myFourRecommendUser, err = uuc.repo.GetUserById(ctx, myFourUserRecommendUserId)
			//if nil != err {
			//	return nil, err
			//}

			if _, ok := fourUserRecommendTotal[myFourUserRecommendUserId]; ok {
				fourUserRecommendTotal[myFourUserRecommendUserId] += userLocationYes.Usdt
			} else {
				fourUserRecommendTotal[myFourUserRecommendUserId] = userLocationYes.Usdt
			}
		}
	}

	// 前四名
	type KeyValuePair struct {
		Key   int64
		Value int64
	}
	var keyValuePairs []KeyValuePair
	for key, value := range fourUserRecommendTotal {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// 按值排序切片
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// 全网前天
	var (
		rewardFourYes *Reward
		yesAmount     int64
	)
	totalRewardYes = rewardLocationYes / 100 * total
	rewardFourYes, err = uuc.ubRepo.GetRewardFourYes(ctx) // 推荐人奖励
	if nil == err && nil != rewardFourYes {
		totalRewardYes += rewardFourYes.Amount
		yesAmount = rewardFourYes.Amount
	}
	totalReward := totalRewardYes / 100 * 70

	fourList := make([]*v1.UserInfoReply_ListFour, 0)

	// 获取前四项
	var topFour []KeyValuePair
	if 4 <= len(keyValuePairs) {
		topFour = keyValuePairs[:4]
	} else {
		topFour = keyValuePairs[:len(keyValuePairs)]
	}
	for k, vTopFour := range topFour {
		var (
			fourUser *User
		)
		fourUser, err = uuc.repo.GetUserById(ctx, vTopFour.Key)
		if nil != err {
			return nil, err
		}

		if nil == fourUser {
			continue
		}

		var (
			tmpMyRecommendAmount int64
		)
		if 0 == k {
			tmpMyRecommendAmount = totalReward / 100 * one
		} else if 1 == k {
			tmpMyRecommendAmount = totalReward / 100 * two
		} else if 2 == k {
			tmpMyRecommendAmount = totalReward / 100 * three
		} else if 3 == k {
			tmpMyRecommendAmount = totalReward / 100 * four
		}

		var address1 string
		if 20 <= len(fourUser.Address) {
			address1 = fourUser.Address[:6] + "..." + fourUser.Address[len(fourUser.Address)-4:]
		}
		fourList = append(fourList, &v1.UserInfoReply_ListFour{
			Location: address1,
			Amount:   fmt.Sprintf("%.2f", float64(vTopFour.Value)/float64(100000)),
			Reward:   fmt.Sprintf("%.2f", float64(tmpMyRecommendAmount)/float64(100000)),
		})
	}

	return &v1.UserInfoReply{
		Status:                status,
		BiwPrice:              float64(bPrice) / float64(bPriceBase),
		PriceC:                float64(cPrice) / float64(bPriceBase),
		BalanceC:              fmt.Sprintf("%.2f", float64(userBalance.BalanceC)/float64(100000)) + "ISPAY",
		Price:                 float64(bPrice) / float64(bPriceBase),
		ExchangeRate:          float64(exchangeRate) / 1000,
		ExchangeRateC:         float64(exchangeRateC) / 1000,
		BalanceBiw:            fmt.Sprintf("%.2f", float64(userBalance.BalanceDhb)/float64(100000)) + "BIW",
		BalanceUsdt:           fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt)/float64(100000)),
		BiwDaily:              "",
		BuyNumTwo:             count2,
		BuyNumThree:           count3,
		BuyNumFour:            count4,
		BuyNumFive:            count5,
		BuyNumOne:             count1,
		BuyNumSix:             count6,
		SellNumOne:            float64(buyOne-count1) / float64(buyOne) * 100,
		SellNumTwo:            float64(buyTwo-count2) / float64(buyTwo) * 100,
		SellNumThree:          float64(buyThree-count3) / float64(buyThree) * 100,
		SellNumFour:           float64(buyFour-count4) / float64(buyFour) * 100,
		SellNumFive:           float64(buyFive-count5) / float64(buyFive) * 100,
		SellNumSix:            float64(buySix-count6) / float64(buySix) * 100,
		DailyRate:             0,
		BiwDailySpeed:         0,
		CurrentAmountBiw:      currentAmountBiw,
		RecommendNum:          int64(len(myUserRecommend)),
		Time:                  time.Now().Unix(),
		LocationList:          myLocations,
		WithdrawList:          withdrawList,
		InviteUserAddress:     inviteUserAddress,
		InviteUrl:             encodeString,
		Count:                 myUser.Out,
		LocationReward:        fmt.Sprintf("%.2f", float64(userBalance.LocationTotal)/float64(100000)) + "BIW",
		RecommendReward:       fmt.Sprintf("%.2f", float64(userBalance.RecommendTotal)/float64(100000)) + "BIW",
		FourReward:            fmt.Sprintf("%.2f", float64(userBalance.FourTotal)/float64(100000)),
		AreaReward:            fmt.Sprintf("%.2f", float64(userBalance.AreaTotal)/float64(100000)) + "BIW",
		FourRewardPool:        fmt.Sprintf("%.2f", float64(rewardLocationYes/100*total)/float64(100000)),
		FourRewardPoolYes:     fmt.Sprintf("%.2f", float64(yesAmount)/float64(100000)),
		Four:                  fourList,
		AreaMax:               areaMax,
		AreaMin:               areaMin,
		AreaAll:               areaAll,
		RecommendTotal:        recommendTotal,
		LocationUsdt:          locationUsdt,
		LocationCurrentMaxSub: locationCurrentMaxSub,
		LocationCurrentSub:    locationCurrent,
		WithdrawTotal:         "",
		LocationUsdtAll:       "",
		ListReward:            listReward,
		ListRecommend:         myRecommendList,
		LastLevel:             lastLevel,
		ConfigFour:            configFour,
		ConfigOne:             fmt.Sprintf("%.2f", float64(locationBiw)/float64(100000)),
		ConfigThree:           configThree,
		ConfigTwo:             fmt.Sprintf("%.2f", float64(totalYesReward)/float64(100000)),
		WithdrawMin:           withdrawMin,
		BuyLimit:              buyLimit,
		AmountBiw:             myUser.AmountBiw,
		AmountUsdt:            myUser.Amount,
		Address:               myUser.AddressTwo,
		AddressBiw:            myUser.AddressThree,
		ListEth:               listUserEth,
	}, nil
}

func (uuc *UserUseCase) UserArea(ctx context.Context, req *v1.UserAreaRequest, user *User) (*v1.UserAreaReply, error) {
	var (
		err             error
		locationId      = req.LocationId
		Locations       []*LocationNew
		LocationRunning *LocationNew
	)

	res := make([]*v1.UserAreaReply_List, 0)
	if 0 >= locationId {
		Locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, user.ID)
		if nil != err {
			return nil, err
		}
		for _, vLocations := range Locations {
			if "running" == vLocations.Status {
				LocationRunning = vLocations
				break
			}
		}

		if nil == LocationRunning {
			return &v1.UserAreaReply{Area: res}, nil
		}

		locationId = LocationRunning.ID
	}

	var (
		myLowLocations []*LocationNew
	)

	myLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, locationId)
	if nil != err {
		return nil, err
	}

	for _, vMyLowLocations := range myLowLocations {
		var (
			userLow           *User
			tmpMyLowLocations []*LocationNew
		)

		userLow, err = uuc.repo.GetUserById(ctx, vMyLowLocations.UserId)
		if nil != err {
			continue
		}

		tmpMyLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, vMyLowLocations.ID)
		if nil != err {
			return nil, err
		}

		var address1 string
		if 20 <= len(userLow.Address) {
			address1 = userLow.Address[:6] + "..." + userLow.Address[len(userLow.Address)-4:]
		}

		res = append(res, &v1.UserAreaReply_List{
			Address:    address1,
			LocationId: vMyLowLocations.ID,
			CountLow:   int64(len(tmpMyLowLocations)),
		})
	}

	return &v1.UserAreaReply{Area: res}, nil
}

func (uuc *UserUseCase) UserInfoOld(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	//var (
	//	myUser                  *User
	//	userInfo                *UserInfo
	//	locations               []*LocationNew
	//	myLastStopLocations     []*LocationNew
	//	userBalance             *UserBalance
	//	userRecommend           *UserRecommend
	//	userRecommends          []*UserRecommend
	//	userRewards             []*Reward
	//	userRewardTotal         int64
	//	userRewardWithdrawTotal int64
	//	encodeString            string
	//	recommendTeamNum        int64
	//	myCode                  string
	//	amount                  = "0"
	//	amount2                 = "0"
	//	configs                 []*Config
	//	myLastLocationCurrent   int64
	//	withdrawAmount          int64
	//	withdrawAmount2         int64
	//	myUserRecommendUserId   int64
	//	inviteUserAddress       string
	//	myRecommendUser         *User
	//	withdrawRate            int64
	//	myLocations             []*v1.UserInfoReply_List
	//	myLocations2            []*v1.UserInfoReply_List22
	//	allRewardList           []*v1.UserInfoReply_List9
	//	timeAgain               int64
	//	err                     error
	//)
	//
	//// 配置
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"time_again",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "time_again" == vConfig.KeyName {
	//			timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}
	//
	//myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	//if nil != locations && 0 < len(locations) {
	//	for _, v := range locations {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//		if "running" == v.Status {
	//			amount = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations = append(myLocations, &v1.UserInfoReply_List{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			Num:       v.Num,
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//// 冻结
	//myLastStopLocations, err = uuc.locationRepo.GetMyStopLocationsLast(ctx, myUser.ID)
	//now := time.Now().UTC()
	//tmpNow := now.Add(8 * time.Hour)
	//if nil != myLastStopLocations {
	//	for _, vMyLastStopLocations := range myLastStopLocations {
	//		if tmpNow.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
	//			myLastLocationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax
	//		}
	//	}
	//}
	//
	//var (
	//	locations2 []*LocationNew
	//)
	//locations2, err = uuc.locationRepo.GetLocationsByUserId2(ctx, myUser.ID)
	//if nil != locations2 && 0 < len(locations2) {
	//	for _, v := range locations2 {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//
	//		if "running" == v.Status {
	//			amount2 = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations2 = append(myLocations2, &v1.UserInfoReply_List22{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	//if nil == userRecommend {
	//	return nil, err
	//}
	//
	//myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	//codeByte := []byte(myCode)
	//encodeString = base64.StdEncoding.EncodeToString(codeByte)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
	//	if 2 <= len(tmpRecommendUserIds) {
	//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
	//	}
	//	myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
	//	if nil != err {
	//		return nil, err
	//	}
	//	inviteUserAddress = myRecommendUser.Address
	//	myCode = userRecommend.RecommendCode + myCode
	//}
	//
	//// 团队
	//var (
	//	teamUserIds        []int64
	//	teamUsers          map[int64]*User
	//	teamUserInfos      map[int64]*UserInfo
	//	teamUserAddresses  []*v1.UserInfoReply_List7
	//	recommendAddresses []*v1.UserInfoReply_List11
	//	teamLocations      map[int64][]*Location
	//	recommendUserIds   map[int64]int64
	//	userBalanceMap     map[int64]*UserBalance
	//)
	//recommendUserIds = make(map[int64]int64, 0)
	//userRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
	//if nil != userRecommends {
	//	for _, vUserRecommends := range userRecommends {
	//		if myCode == vUserRecommends.RecommendCode {
	//			recommendUserIds[vUserRecommends.UserId] = vUserRecommends.UserId
	//		}
	//		teamUserIds = append(teamUserIds, vUserRecommends.UserId)
	//	}
	//
	//	// 用户信息
	//	recommendTeamNum = int64(len(userRecommends))
	//	teamUsers, _ = uuc.repo.GetUserByUserIds(ctx, teamUserIds...)
	//	teamUserInfos, _ = uuc.uiRepo.GetUserInfoByUserIds(ctx, teamUserIds...)
	//	teamLocations, _ = uuc.locationRepo.GetLocationMapByIds(ctx, teamUserIds...)
	//	userBalanceMap, _ = uuc.ubRepo.GetUserBalanceByUserIds(ctx, teamUserIds...)
	//	if nil != teamUsers {
	//		for _, vTeamUsers := range teamUsers {
	//			var locationAmount int64
	//			if _, ok := teamUserInfos[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := userBalanceMap[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := teamLocations[vTeamUsers.ID]; ok {
	//
	//				for _, vTeamLocations := range teamLocations[vTeamUsers.ID] {
	//					locationAmount += vTeamLocations.Usdt
	//				}
	//			}
	//			if _, ok := recommendUserIds[vTeamUsers.ID]; ok {
	//				recommendAddresses = append(recommendAddresses, &v1.UserInfoReply_List11{
	//					Address: vTeamUsers.Address,
	//					Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//					Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//				})
	//			}
	//
	//			teamUserAddresses = append(teamUserAddresses, &v1.UserInfoReply_List7{
	//				Address: vTeamUsers.Address,
	//				Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//				Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//			})
	//		}
	//	}
	//}
	//
	//var (
	//	totalDeposit      int64
	//	userBalanceRecord []*UserBalanceRecord
	//	depositList       []*v1.UserInfoReply_List13
	//)
	//
	//userBalanceRecord, _ = uuc.ubRepo.GetUserBalanceRecordsByUserId(ctx, myUser.ID)
	//for _, vUserBalanceRecord := range userBalanceRecord {
	//	totalDeposit += vUserBalanceRecord.Amount
	//	depositList = append(depositList, &v1.UserInfoReply_List13{
	//		CreatedAt: vUserBalanceRecord.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//		Amount:    fmt.Sprintf("%.4f", float64(vUserBalanceRecord.Amount)/float64(100000)),
	//	})
	//}
	//
	//// 累计奖励
	//var (
	//	locationRewardList            []*v1.UserInfoReply_List2
	//	recommendRewardList           []*v1.UserInfoReply_List5
	//	locationTotal                 int64
	//	yesterdayLocationTotal        int64
	//	recommendRewardTotal          int64
	//	recommendRewardDhbTotal       int64
	//	yesterdayRecommendRewardTotal int64
	//)
	//
	//var startDate time.Time
	//var endDate time.Time
	//if 16 <= now.Hour() {
	//	startDate = now.AddDate(0, 0, -1)
	//	endDate = startDate.AddDate(0, 0, 1)
	//} else {
	//	startDate = now.AddDate(0, 0, -2)
	//	endDate = startDate.AddDate(0, 0, 1)
	//}
	//yesterdayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	//yesterdayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)
	//
	//fmt.Println(now, yesterdayStart, yesterdayEnd)
	//userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	//if nil != userRewards {
	//	for _, vUserReward := range userRewards {
	//		if "location" == vUserReward.Reason {
	//			locationTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayLocationTotal += vUserReward.Amount
	//			}
	//			locationRewardList = append(locationRewardList, &v1.UserInfoReply_List2{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "recommend" == vUserReward.Reason {
	//
	//			recommendRewardList = append(recommendRewardList, &v1.UserInfoReply_List5{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			recommendRewardTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayRecommendRewardTotal += vUserReward.Amount
	//			}
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "reward_withdraw" == vUserReward.Reason {
	//			userRewardTotal += vUserReward.Amount
	//			userRewardWithdrawTotal += vUserReward.Amount
	//		} else if "recommend_token" == vUserReward.Reason {
	//			recommendRewardDhbTotal += vUserReward.Amount
	//		}
	//	}
	//}
	//
	//var (
	//	withdraws []*Withdraw
	//)
	//withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//for _, v := range withdraws {
	//	if "usdt" == v.Type {
	//		withdrawAmount += v.RelAmount
	//	}
	//	if "usdt_2" == v.Type {
	//		withdrawAmount2 += v.RelAmount
	//	}
	//}
	//
	//return &v1.UserInfoReply{
	//	Address:                 myUser.Address,
	//	Level:                   userInfo.Vip,
	//	Amount:                  amount,
	//	Amount2:                 amount2,
	//	LocationList2:           myLocations2,
	//	AmountB:                 fmt.Sprintf("%.2f", float64(myLastLocationCurrent)/float64(100000)),
	//	DepositList:             depositList,
	//	BalanceUsdt:             fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt)/float64(100000)),
	//	BalanceUsdt2:            fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt2)/float64(100000)),
	//	BalanceDhb:              fmt.Sprintf("%.2f", float64(userBalance.BalanceDhb)/float64(100000)),
	//	InviteUrl:               encodeString,
	//	RecommendNum:            userInfo.HistoryRecommend,
	//	RecommendTeamNum:        recommendTeamNum,
	//	Total:                   fmt.Sprintf("%.2f", float64(userRewardTotal)/float64(100000)),
	//	RewardWithdraw:          fmt.Sprintf("%.2f", float64(userRewardWithdrawTotal)/float64(100000)),
	//	WithdrawAmount:          fmt.Sprintf("%.2f", float64(withdrawAmount)/float64(100000)),
	//	WithdrawAmount2:         fmt.Sprintf("%.2f", float64(withdrawAmount2)/float64(100000)),
	//	LocationTotal:           fmt.Sprintf("%.2f", float64(locationTotal)/float64(100000)),
	//	Account:                 "0xAfC39F3592A1024144D1ba6DC256397F4DbfbE2f",
	//	LocationList:            myLocations,
	//	TeamAddressList:         teamUserAddresses,
	//	AllRewardList:           allRewardList,
	//	InviteUserAddress:       inviteUserAddress,
	//	RecommendAddressList:    recommendAddresses,
	//	WithdrawRate:            withdrawRate,
	//	RecommendRewardTotal:    fmt.Sprintf("%.2f", float64(recommendRewardTotal)/float64(100000)),
	//	RecommendRewardDhbTotal: fmt.Sprintf("%.2f", float64(recommendRewardDhbTotal)/float64(100000)),
	//	TotalDeposit:            fmt.Sprintf("%.2f", float64(totalDeposit)/float64(100000)),
	//	WithdrawAll:             fmt.Sprintf("%.2f", float64(withdrawAmount+withdrawAmount2)/float64(100000)),
	//}, nil
	return nil, nil

}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {

	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {

	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) TranList(ctx context.Context, user *User, reqTypeCoin string, reqTran string) (*v1.TranListReply, error) {

	var (
		userBalanceRecord []*UserBalanceRecord
		typeCoin          = "usdt"
		tran              = "tran"
		err               error
	)

	res := &v1.TranListReply{
		Tran: make([]*v1.TranListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	if "tran_to" == reqTran {
		tran = reqTran
	}

	userBalanceRecord, err = uuc.ubRepo.GetUserBalanceRecordByUserId(ctx, user.ID, typeCoin, tran)
	if nil != err {
		return res, err
	}

	for _, v := range userBalanceRecord {
		res.Tran = append(res.Tran, &v1.TranListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User, reqTypeCoin string) (*v1.WithdrawListReply, error) {

	var (
		withdraws []*Withdraw
		typeCoin  = "usdt"
		err       error
	)

	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId(ctx, user.ID, typeCoin)
	if nil != err {
		return res, err
	}

	for _, v := range withdraws {
		res.Withdraw = append(res.Withdraw, &v1.WithdrawListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
			Status:    v.Status,
			Type:      v.Type,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) TradeList(ctx context.Context, user *User) (*v1.TradeListReply, error) {

	var (
		trades []*Trade
		err    error
	)

	res := &v1.TradeListReply{
		Trade: make([]*v1.TradeListReply_List, 0),
	}

	trades, err = uuc.ubRepo.GetTradeByUserId(ctx, user.ID)
	if nil != err {
		return res, err
	}

	for _, v := range trades {
		res.Trade = append(res.Trade, &v1.TradeListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(v.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(v.AmountHbs)/float64(100000)),
			Status:    v.Status,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) Buy(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	var (
		amount     = req.SendBody.Amount
		configs    []*Config
		bPrice     int64
		bPriceBase int64
		coinType   string
		err        error
	)

	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"b_price",
		"b_price_base",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price" == vConfig.KeyName {
				bPrice, err = strconv.ParseInt(vConfig.Value, 10, 64)
				if nil != err || bPrice < 1 {
					return &v1.BuyReply{
						Status: "币价错误",
					}, nil
				}
			}

			if "b_price_base" == vConfig.KeyName {
				bPriceBase, err = strconv.ParseInt(vConfig.Value, 10, 64)
				if nil != err || bPriceBase < 1 {
					return &v1.BuyReply{
						Status: "币价错误",
					}, nil
				}
			}
		}
	}

	var (
		tmpValue int64
		strValue string
	)
	if 100 == amount {
		tmpValue = int64(10000000)
		strValue = "100000000000000000000"
	} else if 1000 == amount {
		tmpValue = int64(100000000)
		strValue = "1000000000000000000000"
	} else if 3000 == amount {
		tmpValue = int64(300000000)
		strValue = "3000000000000000000000"
	} else if 5000 == amount {
		tmpValue = int64(500000000)
		strValue = "5000000000000000000000"
	} else if 10000 == amount {
		tmpValue = int64(1000000000)
		strValue = "10000000000000000000000"
	} else if 300 == amount {
		tmpValue = int64(30000000)
		strValue = "300000000000000000000"
	} else if 15000 == amount {
		tmpValue = int64(1500000000)
		strValue = "15000000000000000000000"
	} else if 30000 == amount {
		tmpValue = int64(3000000000)
		strValue = "30000000000000000000000"
	} else {
		return &v1.BuyReply{
			Status: "金额错误",
		}, nil
	}

	var (
		amountUsdt uint64
		amountBiw  uint64
	)
	amountUsdt = amount // 半个
	//amountSub := amount * 30 / 100
	//amountSub := 0
	if 1 == req.SendBody.Type {
		if amountUsdt > user.Amount {
			return &v1.BuyReply{
				Status: "usdt余额不足",
			}, nil
		}
		coinType = "USDT"

		//amountBiw = amountSub * uint64(bPriceBase) / uint64(bPrice)
		//if 0 >= amountBiw {
		//	return &v1.BuyReply{
		//		Status: "所需biw为0，错误",
		//	}, nil
		//}
		//
		//if amountBiw > user.AmountBiw {
		//	return &v1.BuyReply{
		//		Status: "biw余额不足",
		//	}, nil
		//}
		//} else if 2 == req.SendBody.Type {
		//	amountBiw := amount * uint64(bPriceBase) / uint64(bPrice)
		//	if 0 >= amountBiw {
		//		return &v1.BuyReply{
		//			Status: "所需biw为0，错误",
		//		}, nil
		//	}
		//
		//	if amountBiw > user.AmountBiw {
		//		return &v1.BuyReply{
		//			Status: "biw余额不足",
		//		}, nil
		//	}
		//	coinType = "DHB"
		//	tmpAmount = amountBiw
		//} else {
	} else {
		return &v1.BuyReply{
			Status: "类型错误",
		}, nil
	}

	notExistDepositResult := make([]*EthUserRecord, 0)
	notExistDepositResult = append(notExistDepositResult, &EthUserRecord{ // 两种币的记录
		UserId:    user.ID,
		Status:    "success",
		Type:      "deposit",
		Amount:    strValue,
		RelAmount: tmpValue,
		CoinType:  coinType,
		Last:      0,
	})

	var (
		res  bool
		code int64
	)
	res, code, err = uuc.EthUserRecordHandle(ctx, amount, amountUsdt, amountBiw, coinType, notExistDepositResult...)
	if nil != err {
		fmt.Println(err)
	}

	if !res {
		if 1 == code {
			return &v1.BuyReply{
				Status: "已认购",
			}, nil
		} else {
			fmt.Println(user.ID, "buy 错误", err)
			return &v1.BuyReply{
				Status: "已认购",
			}, nil
		}
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) EthUserRecordHandle(ctx context.Context, amount uint64, amountUsdt uint64, amountBiw uint64, coinType string, ethUserRecord ...*EthUserRecord) (bool, int64, error) {

	var (
		err          error
		configs      []*Config
		buyOne       int64
		buyTwo       int64
		buyThree     int64
		buyFour      int64
		buyFive      int64
		buySix       int64
		buySeven     int64
		buyEight     int64
		areaOne      int64
		areaTwo      int64
		areaThree    int64
		areaFour     int64
		areaFive     int64
		recommendOne int64
		recommendTwo int64
		bPrice       int64
		bPriceBase   int64
		feeRate      int64
		//recommendRate1 int64
		//recommendRate2 int64
		//recommendRate3 int64
		//recommendRate4 int64
		//recommendRate5 int64
		//recommendRate6 int64
		//recommendRate7 int64
		//recommendRate8 int64
		//recommendBase  = int64(100)
	)
	// 配置
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"area_one", "area_two", "area_three", "area_four", "area_five", "recommend_new_one", "recommend_new_two", "exchange_rate",
		"buy_one", "buy_two", "buy_six", "buy_three", "buy_four", "buy_five", "b_price", "b_price_base", "recommend_rate_1", "recommend_rate_2", "recommend_rate_3", "recommend_rate_4", "recommend_rate_5", "recommend_rate_6", "recommend_rate_7", "recommend_rate_8")
	if nil != configs {
		for _, vConfig := range configs {
			if "buy_one" == vConfig.KeyName {
				buyOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_two" == vConfig.KeyName {
				buyTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_three" == vConfig.KeyName {
				buyThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_four" == vConfig.KeyName {
				buyFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_five" == vConfig.KeyName {
				buyFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "buy_six" == vConfig.KeyName {
				buySix, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "buy_seven" == vConfig.KeyName {
				buySeven, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_eight" == vConfig.KeyName {
				buyEight, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "area_one" == vConfig.KeyName {
				areaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_two" == vConfig.KeyName {
				areaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_three" == vConfig.KeyName {
				areaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_four" == vConfig.KeyName {
				areaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_five" == vConfig.KeyName {
				areaFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "recommend_new_one" == vConfig.KeyName {
				recommendOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "recommend_new_two" == vConfig.KeyName {
				recommendTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, err = strconv.ParseInt(vConfig.Value, 10, 64)
				if nil != err {
					fmt.Println(err, "b_price err")
					return false, 0, nil
				}
			}
			if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			//		if "b_price_base" == vConfig.KeyName {
			//			bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_1" == vConfig.KeyName {
			//			recommendRate1, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_2" == vConfig.KeyName {
			//			recommendRate2, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_3" == vConfig.KeyName {
			//			recommendRate3, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_4" == vConfig.KeyName {
			//			recommendRate4, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_5" == vConfig.KeyName {
			//			recommendRate5, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_6" == vConfig.KeyName {
			//			recommendRate6, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_7" == vConfig.KeyName {
			//			recommendRate7, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_8" == vConfig.KeyName {
			//			recommendRate8, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
		}
	}

	for _, v := range ethUserRecord {
		var (
			user           *User
			allLocations   []*LocationNew
			myLocations    []*LocationNew
			myLastLocation *LocationNew
			lastLevel      int64
			err            error
		)
		user, err = uuc.repo.GetUserById(ctx, v.UserId)
		if nil != err || nil == user {
			fmt.Println(err, "用户不存在，buy")
			continue
		}

		// 获取当前用户的占位信息，已经有运行中的跳过
		allLocations, err = uuc.locationRepo.GetAllLocationsNew(ctx, v.RelAmount*25/10) // 同倍率的
		if nil == allLocations {                                                        // 查询异常跳过本次循环
			fmt.Println(err, "错误投资", v)
			continue
		}
		if nil != err {
			fmt.Println(err, "123")
			continue
		}

		if 10000000 == v.RelAmount && int64(len(allLocations)) < buySix {

		} else if 30000000 == v.RelAmount && int64(len(allLocations)) < buyOne {

		} else if 100000000 == v.RelAmount && int64(len(allLocations)) < buyTwo {

		} else if 300000000 == v.RelAmount && int64(len(allLocations)) < buyThree {

		} else if 500000000 == v.RelAmount && int64(len(allLocations)) < buyFour {

		} else if 1000000000 == v.RelAmount && int64(len(allLocations)) < buyFive {

		} else if 1500000000 == v.RelAmount && int64(len(allLocations)) < buySeven {

		} else if 3000000000 == v.RelAmount && int64(len(allLocations)) < buyEight {

		} else {
			fmt.Println(v, "1234")
			continue
		}

		// 获取当前用户的占位信息，已经有运行中的跳过
		myLocations, err = uuc.locationRepo.GetLocationsNewByUserId(ctx, v.UserId)
		if nil == myLocations { // 查询异常跳过本次循环
			fmt.Println(err, "错误投资2", v)
			continue
		}
		if nil != err {
			fmt.Println(err, "12")
			continue
		}

		if 0 < len(myLocations) {
			var (
				stop bool
			)

			for _, vMyLocations := range myLocations {
				if "stop" != vMyLocations.Status {
					stop = true
					fmt.Println(err, "已投资", v)
					break
				}

				var tmpLastLevel int64
				// 1大区
				if vMyLocations.Total >= vMyLocations.TotalTwo && vMyLocations.Total >= vMyLocations.TotalThree {
					if areaOne <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 5
					}
				} else if vMyLocations.TotalTwo >= vMyLocations.Total && vMyLocations.TotalTwo >= vMyLocations.TotalThree {
					if areaOne <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 5
					}
				} else if vMyLocations.TotalThree >= vMyLocations.Total && vMyLocations.TotalThree >= vMyLocations.TotalTwo {
					if areaOne <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 5
					}
				}

				if tmpLastLevel > lastLevel {
					lastLevel = tmpLastLevel
				}

				if vMyLocations.LastLevel > lastLevel {
					lastLevel = vMyLocations.LastLevel
				}
			}

			if stop {
				return false, 1, nil
			}

			myLastLocation = myLocations[0] // 最新一个
		}

		// 推荐人
		var (
			userRecommend         *UserRecommend
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, v.UserId)
		if nil != err {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
		}

		// 直推人投资
		var (
			myRecommmendLocation *LocationNew
		)
		if 0 < myUserRecommendUserId {
			myRecommmendLocation, err = uuc.locationRepo.GetMyLocationLastRunning(ctx, myUserRecommendUserId)
			if nil != err {
				continue
			}
		}

		// 顺位
		var (
			lastLocation     *LocationNew
			isOriginLocation bool
		)

		// 有直推人占位且第一次入金，挂在直推人名下，按位查找
		if nil != myRecommmendLocation && nil == myLastLocation {
			fmt.Println("认购用户：", v.UserId)
			var (
				selectLocation *LocationNew // 选中的
			)
			if 3 <= myRecommmendLocation.Count {
				fmt.Println("推荐人下面数量", myRecommmendLocation.Count)
				var tmpSopFor bool

				tmpIds := make([]int64, 0)
				tmpIds = append(tmpIds, myRecommmendLocation.ID)

				i := 0
				for i < len(tmpIds) { // 小于3个人
					// 查找
					vTmpId := tmpIds[i]
					fmt.Println("Processing:", vTmpId)

					//
					i++

					var (
						topLocations []*LocationNew
					)
					topLocations, err = uuc.locationRepo.GetLocationsByTopTwo(ctx, vTmpId)
					if nil != err {
						break
					}

					// 没数据没数据, 正常最少三个
					if 0 >= len(topLocations) {
						tmpSopFor = true
						break
					}

					for _, vTopLocations := range topLocations {
						if 3 > vTopLocations.Count {
							selectLocation = vTopLocations
							break
						}
						tmpIds = append(tmpIds, vTopLocations.ID)
					}

					if nil != selectLocation {
						break
					}
				}

				if tmpSopFor {
					continue
				}
			} else {
				selectLocation = myRecommmendLocation
			}

			lastLocation = selectLocation
		} else if nil != myLastLocation { // 2复投，原点
			fmt.Println("复投,认购用户：", v.UserId)
			lastLocation, err = uuc.locationRepo.GetLocationById(ctx, myLastLocation.Top)
			if nil != err {
				fmt.Println("查找错误，投资", myLastLocation, v)
				continue
			}
			isOriginLocation = true
		} else if nil == myRecommmendLocation { // 直推无位置或一号用户无直推人，顺位补齐
			fmt.Println("无推荐人,认购用户：", v.UserId)
			var (
				firstLocation  *LocationNew
				tmpSopFor      bool
				selectLocation *LocationNew // 选中的
			)

			firstLocation, err = uuc.locationRepo.GetLocationFirst(ctx)
			if nil != err {
				continue
			}
			if nil != firstLocation {
				if 3 <= firstLocation.Count {
					tmpIds := make([]int64, 0)
					tmpIds = append(tmpIds, firstLocation.ID)

					i := 0
					for i < len(tmpIds) { // 小于3个人
						// 查找
						vTmpId := tmpIds[i]
						fmt.Println("Processing2:", vTmpId)

						//
						i++

						// 查找
						var (
							topLocations []*LocationNew
						)
						topLocations, err = uuc.locationRepo.GetLocationsByTopTwo(ctx, vTmpId)
						if nil != err {
							break
						}

						// 没数据, 正常最少三个
						if 0 >= len(topLocations) {
							tmpSopFor = true
							break
						}

						for _, vTopLocations := range topLocations {
							if 3 > vTopLocations.Count {
								selectLocation = vTopLocations
								break
							}
							tmpIds = append(tmpIds, vTopLocations.ID)
						}

						if nil != selectLocation {
							break
						}
					}

					if tmpSopFor {
						continue
					}
				} else {
					selectLocation = firstLocation
				}
			}

			lastLocation = selectLocation
		} else {
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

			// 推荐人
			if 0 < len(tmpRecommendUserIds) {
				lastKey := len(tmpRecommendUserIds) - 1 // 有直推len比>=2 ,key是0则是空格，1是直推，键位最后一个人
				if 1 <= lastKey {
					for i := 0; i <= 1; i++ { // 两代
						// 有占位信息，推荐人推荐人的上一代
						if lastKey-i <= 0 {
							break
						}

						tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
						if 0 >= tmpMyTopUserRecommendUserId {
							break
						}

						var tmpUser *User
						tmpUser, err = uuc.repo.GetUserById(ctx, tmpMyTopUserRecommendUserId)
						if nil == tmpUser {
							continue
						}

						if 1 == tmpUser.Lock {
							continue
						}

						var myUserRecommendUserLocationsLast []*LocationNew
						myUserRecommendUserLocationsLast, err = uuc.locationRepo.GetLocationsNewByUserId(ctx, tmpMyTopUserRecommendUserId)
						if nil != myUserRecommendUserLocationsLast {

							var tmpMyTopUserRecommendUserLocationLast *LocationNew
							if 1 <= len(myUserRecommendUserLocationsLast) {
								for _, vMyUserRecommendUserLocationLast := range myUserRecommendUserLocationsLast {
									if "running" == vMyUserRecommendUserLocationLast.Status {
										tmpMyTopUserRecommendUserLocationLast = vMyUserRecommendUserLocationLast
										break
									}
								}

								if nil == tmpMyTopUserRecommendUserLocationLast { // 无位
									continue
								}

								tmpMinUsdt := tmpMyTopUserRecommendUserLocationLast.Usdt
								if v.RelAmount < tmpMinUsdt {
									tmpMinUsdt = v.RelAmount
								}

								var tmpMyRecommendAmount int64
								if 0 == i { // 当前用户被此人直推
									tmpMyRecommendAmount = tmpMinUsdt / 1000 * recommendOne
								} else if 1 == i {
									tmpMyRecommendAmount = tmpMinUsdt / 1000 * recommendTwo
								} else {
									continue
								}

								if 0 < tmpMyRecommendAmount { // 扣除推荐人分红
									bAmount := tmpMyRecommendAmount * bPriceBase / bPrice
									tmpStatus := tmpMyTopUserRecommendUserLocationLast.Status
									tmpStopDate := time.Now().UTC().Add(8 * time.Hour)
									// 过了
									if tmpMyTopUserRecommendUserLocationLast.Current+tmpMyRecommendAmount >= tmpMyTopUserRecommendUserLocationLast.CurrentMax { // 占位分红人分满停止
										tmpStatus = "stop"
										tmpStopDate = time.Now().UTC().Add(8 * time.Hour)

										tmpMyRecommendAmount = tmpMyTopUserRecommendUserLocationLast.CurrentMax - tmpMyTopUserRecommendUserLocationLast.Current
										bAmount = tmpMyRecommendAmount * bPriceBase / bPrice
									}

									if 0 < tmpMyRecommendAmount && 0 < bAmount {
										var tmpMaxNew int64
										if tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew < tmpMyTopUserRecommendUserLocationLast.CurrentMax {
											tmpMaxNew = tmpMyTopUserRecommendUserLocationLast.CurrentMax - tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew
										}

										if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
											err = uuc.locationRepo.UpdateLocationNewNew(ctx, tmpMyTopUserRecommendUserLocationLast.ID, tmpMyTopUserRecommendUserLocationLast.UserId, tmpStatus, tmpMyRecommendAmount, tmpMaxNew, bAmount, tmpStopDate, tmpMyTopUserRecommendUserLocationLast.CurrentMax) // 分红占位数据修改
											if nil != err {
												return err
											}

											_, err = uuc.ubRepo.RecommendLocationRewardBiw(ctx, tmpMyTopUserRecommendUserId, bAmount, int64(i+1), tmpStatus, tmpMaxNew, feeRate) // 推荐人奖励
											if nil != err {
												return err
											}

											// 业绩减掉
											if "stop" == tmpStatus {
												tmpTop := tmpMyTopUserRecommendUserLocationLast.Top
												tmpTopNum := tmpMyTopUserRecommendUserLocationLast.TopNum
												for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
													err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, tmpMyTopUserRecommendUserLocationLast.Usdt/100000)
													if nil != err {
														return err
													}

													var (
														currentLocation *LocationNew
													)
													currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
													if nil != err {
														return err
													}

													if nil != currentLocation && 0 < currentLocation.Top {
														tmpTop = currentLocation.Top
														tmpTopNum = currentLocation.TopNum
														continue
													}

													break
												}
											}

											return nil
										}); nil != err {
											fmt.Println("err reward daily recommend", err, myUserRecommendUserLocationsLast)
											continue
										}
									}
								}
							}
						}

					}

				}

			}

			var (
				tmpTop int64
				tmpNum int64
			)

			// 顺位
			if nil != lastLocation {
				if isOriginLocation && nil != myLastLocation {
					fmt.Println("复投", v.UserId)
					err = uuc.locationRepo.UpdateLocationNewCountTwo(ctx, myLastLocation.Top, myLastLocation.TopNum, v.RelAmount/100000)
					if nil != err {
						return err
					}
				} else {
					fmt.Println("新位置", v.UserId, lastLocation.ID, lastLocation.Count+1)
					err = uuc.locationRepo.UpdateLocationNewCount(ctx, lastLocation.ID, lastLocation.Count+1, v.RelAmount/100000)
					if nil != err {
						return err
					}
					tmpTop = lastLocation.ID
					tmpNum = lastLocation.Count + 1
				}

				var (
					currentTop    = lastLocation.Top
					currentTopNum = lastLocation.TopNum
				)
				// 大小区业绩
				for j := 0; j < 10000 && 0 < currentTop && 0 < currentTopNum; j++ {
					err = uuc.locationRepo.UpdateLocationNewTotal(ctx, currentTop, currentTopNum, v.RelAmount/100000)
					if nil != err {
						return err
					}

					var (
						currentLocation *LocationNew
					)
					currentLocation, err = uuc.locationRepo.GetLocationById(ctx, currentTop)
					if nil != err {
						return err
					}

					if nil != currentLocation && 0 < currentLocation.Top && 0 < currentLocation.TopNum {
						currentTop = currentLocation.Top
						currentTopNum = currentLocation.TopNum
						continue
					}

					break
				}
			}

			if isOriginLocation && nil != myLastLocation {
				_, err = uuc.locationRepo.UpdateLocationNew(ctx, myLastLocation.ID,
					v.UserId, v.RelAmount*25/10, v.RelAmount, int64(amount), user.Address, coinType)
			} else {
				_, err = uuc.locationRepo.CreateLocationNew(ctx, &LocationNew{ // 占位
					UserId:     v.UserId,
					Status:     "running",
					Current:    0,
					CurrentMax: v.RelAmount * 25 / 10, // 2.5倍率
					Num:        1,
					Top:        tmpTop,
					TopNum:     tmpNum,
					LastLevel:  lastLevel,
				}, v.RelAmount, int64(amount), user.Address, coinType)
			}

			if nil != err {
				return err
			}

			if 0 < myUserRecommendUserId {
				err = uuc.urRepo.UpdateUserRecommendTotal(ctx, myUserRecommendUserId, v.RelAmount/100000)
				if nil != err {
					return err
				}
			}

			err = uuc.repo.UpdateUserNewTwoNew(ctx, v.UserId, amountUsdt, amountBiw, coinType)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "错误投资3", v)
			return false, 0, err
		}
	}

	return true, 0, nil
}

// Exchange Exchange.
func (uuc *UserUseCase) Exchange(ctx context.Context, req *v1.ExchangeRequest, user *User) (*v1.ExchangeReply, error) {
	var (
		//u           *User
		err         error
		userBalance *UserBalance
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	if userBalance.BalanceDhb < amount {
		amount = userBalance.BalanceDhb
	}

	if 100000 > amount {
		return &v1.ExchangeReply{
			Status: "fail",
		}, nil
	}

	// 配置
	var (
		configs       []*Config
		exchangeRate  int64
		exchangeRateC int64
		bPrice        int64
		cPrice        int64
		bPriceBase    int64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"exchange_rate",
		"exchange_rate_c",
		"b_price",
		"b_price_base",
		"c_price",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "exchange_rate" == vConfig.KeyName {
				exchangeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "exchange_rate_c" == vConfig.KeyName {
				exchangeRateC, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "c_price" == vConfig.KeyName {
				cPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	var (
		amountUsdtSubFee int64
		amountUsdt       int64
	)
	if "ispay" == req.SendBody.Type {
		amountUsdt = int64(float64(amount) / float64(bPriceBase) * float64(cPrice))
		amountUsdtSubFee = amountUsdt - amountUsdt*exchangeRateC/1000
		//if amountUsdt <= 0 {
		//	return &v1.ExchangeReply{
		//		Status: "fail price 2",
		//	}, nil
		//}
		return &v1.ExchangeReply{
			Status: "close",
		}, nil
	} else {
		amountUsdt = int64(float64(amount) / float64(bPriceBase) * float64(bPrice))
		amountUsdtSubFee = amountUsdt - amountUsdt*exchangeRate/1000
		if amountUsdt <= 0 {
			return &v1.ExchangeReply{
				Status: "fail price",
			}, nil
		}
	}

	//var (
	//	locations       []*LocationNew
	//	runningLocation *LocationNew
	//)
	//
	//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//if 0 >= len(locations) {
	//	return &v1.ExchangeReply{
	//		Status: "fail location",
	//	}, nil
	//}

	//runningLocation = locations[0]
	//if "running" != runningLocation.Status {
	//	return &v1.ExchangeReply{
	//		Status: "fail location",
	//	}, nil
	//}
	//
	//if runningLocation.CurrentMax < runningLocation.CurrentMaxNew+amountUsdt {
	//	return &v1.ExchangeReply{
	//		Status: "fail location max",
	//	}, nil
	//}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "ispay" == req.SendBody.Type {
			err = uuc.ubRepo.Exchange2(ctx, user.ID, amount, amountUsdtSubFee, amountUsdt, 0) // 提现
			if nil != err {
				return err
			}
		} else {
			err = uuc.ubRepo.Exchange(ctx, user.ID, amount, amountUsdtSubFee, amountUsdt, 0) // 提现
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.ExchangeReply{
		Status: "ok",
	}, nil

}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User, password string) (*v1.WithdrawReply, error) {
	var (
		//u           *User
		err         error
		userBalance *UserBalance
	)

	if "1" == req.SendBody.Type {
		req.SendBody.Type = "usdt"
	} else if "2" == req.SendBody.Type {
		req.SendBody.Type = "dhb"
	} else if "3" == req.SendBody.Type {
		req.SendBody.Type = "ispay"
	}
	//u, _ = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}

	//if "" == u.Password || 6 > len(u.Password) {
	//	return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	//}
	//
	//if u.Password != user.Password {
	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//}

	//if password != u.Password {
	//	return nil, errors.New(500, "密码错误", "密码错误")
	//}

	if "usdt" != req.SendBody.Type && "dhb" != req.SendBody.Type && "ispay" != req.SendBody.Type {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	//if "dhb" == req.SendBody.Type {
	//	if userBalance.BalanceDhb < amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//
	//	if 10000000 > amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//}

	// 配置
	var (
		configs         []*Config
		withdrawMax     int64
		withdrawMin     int64
		withdrawMaxBiw  int64
		withdrawMinBiw  int64
		withdrawMaxC    int64
		withdrawMinC    int64
		withdrawOpen    string
		withdrawOpenBiw string
		withdrawOpenC   string
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_amount_max",
		"withdraw_amount_max_biw",
		"withdraw_amount_min_biw",
		"withdraw_amount_min",
		"withdraw_amount_min_c",
		"withdraw_amount_max_c",
		"withdraw_open",
		"withdraw_open_biw",
		"withdraw_open_c",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_max" == vConfig.KeyName {
				withdrawMax, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}

			if "withdraw_amount_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}

			if "withdraw_amount_max_biw" == vConfig.KeyName {
				withdrawMaxBiw, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}

			if "withdraw_amount_min_biw" == vConfig.KeyName {
				withdrawMinBiw, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}
			if "withdraw_amount_max_c" == vConfig.KeyName {
				withdrawMaxC, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}

			if "withdraw_amount_min_c" == vConfig.KeyName {
				withdrawMinC, _ = strconv.ParseInt(vConfig.Value+"00000", 10, 64)
			}
			if "withdraw_open" == vConfig.KeyName {
				withdrawOpen = vConfig.Value
			}
			if "withdraw_open_biw" == vConfig.KeyName {
				withdrawOpenBiw = vConfig.Value
			}
			if "withdraw_open_c" == vConfig.KeyName {
				withdrawOpenC = vConfig.Value
			}
		}
	}
	//
	//var (
	//	locations       []*LocationNew
	//	runningLocation *LocationNew
	//)
	//
	//amountUsdt := amount / bPriceBase * bPrice
	if "usdt" == req.SendBody.Type {
		if "1" != withdrawOpen {
			return &v1.WithdrawReply{
				Status: "ok",
			}, nil
		}

		if 35 >= len(req.SendBody.Address) || 45 < len(req.SendBody.Address) {
			return &v1.WithdrawReply{
				Status: "地址长度不正确",
			}, nil
		}

		if userBalance.BalanceUsdt < amount {
			amount = userBalance.BalanceUsdt
		}

		if withdrawMax < amount {
			return &v1.WithdrawReply{
				Status: "fail max",
			}, nil
		}

		if withdrawMin > amount {
			return &v1.WithdrawReply{
				Status: "fail min",
			}, nil
		}
	} else if "dhb" == req.SendBody.Type {
		if "1" != withdrawOpenBiw {
			return &v1.WithdrawReply{
				Status: "ok",
			}, nil
		}

		if userBalance.BalanceDhb < amount {
			amount = userBalance.BalanceDhb
		}

		if withdrawMaxBiw < amount {
			return &v1.WithdrawReply{
				Status: "fail max",
			}, nil
		}

		if withdrawMinBiw > amount {
			return &v1.WithdrawReply{
				Status: "fail min",
			}, nil
		}

		//if amountUsdt <= 0 {
		//	return &v1.WithdrawReply{
		//		Status: "fail price",
		//	}, nil
		//}
		//
		//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, user.ID)
		//if nil != err {
		//	return nil, err
		//}
		//
		//if 0 >= len(locations) {
		//	return &v1.WithdrawReply{
		//		Status: "fail location",
		//	}, nil
		//}
		//
		//runningLocation = locations[0]
		//if "running" != runningLocation.Status {
		//	return &v1.WithdrawReply{
		//		Status: "fail location",
		//	}, nil
		//}
		//
		//if runningLocation.CurrentMax < runningLocation.CurrentMaxNew+amountUsdt {
		//	return &v1.WithdrawReply{
		//		Status: "fail location max",
		//	}, nil
		//}
	} else if "ispay" == req.SendBody.Type {
		if "1" != withdrawOpenC {
			return &v1.WithdrawReply{
				Status: "ok",
			}, nil
		}

		if 35 >= len(req.SendBody.Address) || 45 < len(req.SendBody.Address) {
			return &v1.WithdrawReply{
				Status: "地址长度不正确",
			}, nil
		}

		if userBalance.BalanceC < amount {
			amount = userBalance.BalanceC
		}

		if withdrawMaxC < amount {
			return &v1.WithdrawReply{
				Status: "fail max",
			}, nil
		}

		if withdrawMinC > amount {
			return &v1.WithdrawReply{
				Status: "fail min",
			}, nil
		}
	}

	//if "usdt_2" == req.SendBody.Type {
	//	if userBalance.BalanceUsdt2 < amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//
	//	if 1000000 > amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//}

	//var userRecommend *UserRecommend
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	//if nil == userRecommend {
	//	return &v1.WithdrawReply{
	//		Status: "信息错误",
	//	}, nil
	//}
	//
	//var (
	//	tmpRecommendUserIds    []string
	//	tmpRecommendUserIdsInt []int64
	//)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	//}
	//lastKey := len(tmpRecommendUserIds) - 1
	//if 1 <= lastKey {
	//	for i := 0; i <= lastKey; i++ {
	//		// 有占位信息，推荐人推荐人的上一代
	//		if lastKey-i <= 0 {
	//			break
	//		}
	//
	//		tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
	//		tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
	//	}
	//}

	// 配置
	//var (
	//	configs      []*Config
	//	withdrawRate int64
	//)
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"withdraw_rate",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "withdraw_rate" == vConfig.KeyName {
	//			withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "usdt" == req.SendBody.Type {
			err = uuc.ubRepo.WithdrawUsdt2(ctx, user.ID, amount) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, amount, 0, req.SendBody.Type, req.SendBody.Address)
			if nil != err {
				return err
			}
		} else if "dhb" == req.SendBody.Type {
			err = uuc.ubRepo.WithdrawDhb(ctx, user.ID, amount) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, amount, 0, req.SendBody.Type, "")
			if nil != err {
				return err
			}
		} else if "ispay" == req.SendBody.Type {
			err = uuc.ubRepo.WithdrawC(ctx, user.ID, amount) // 提现
			if nil != err {
				return err
			}
			_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, amount, 0, req.SendBody.Type, req.SendBody.Address)
			if nil != err {
				return err
			}
		}
		//else if "usdt_2" == req.SendBody.Type {
		//	err = uuc.ubRepo.WithdrawUsdt3(ctx, user.ID, amount) // 提现
		//	if nil != err {
		//		return err
		//	}
		//	_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, amount-amount*withdrawRate/100, amount*withdrawRate/100, req.SendBody.Type)
		//	if nil != err {
		//		return err
		//	}
		//
		//}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Tran(ctx context.Context, req *v1.TranRequest, user *User, password string) (*v1.TranReply, error) {
	var (
		err         error
		userBalance *UserBalance
		toUser      *User
		u           *User
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	if "" == req.SendBody.Address {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	toUser, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if nil != err {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	if user.ID == toUser.ID {
		return &v1.TranReply{
			Status: "不能给自己转账",
		}, nil
	}

	if "dhb" != req.SendBody.Type && "usdt" != req.SendBody.Type {
		return &v1.TranReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	if "dhb" == req.SendBody.Type {
		if userBalance.BalanceDhb < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 10000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	if "usdt" == req.SendBody.Type {
		if userBalance.BalanceUsdt < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 1000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	var (
		userRecommend  *UserRecommend
		userRecommend2 *UserRecommend
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds          []string
		tmpRecommendUserIdsInt       []int64
		toUserTmpRecommendUserIds    []string
		toUserTmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	if 1 < len(tmpRecommendUserIds) {
		lastKey := len(tmpRecommendUserIds) - 1
		if 1 <= lastKey {
			for i := 0; i <= lastKey; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey-i <= 0 {
					break
				}

				tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
				tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
			}
		}
	}

	userRecommend2, err = uuc.urRepo.GetUserRecommendByUserId(ctx, toUser.ID)
	if nil == userRecommend2 {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}
	if "" != userRecommend2.RecommendCode {
		toUserTmpRecommendUserIds = strings.Split(userRecommend2.RecommendCode, "D")
	}

	if 1 < len(toUserTmpRecommendUserIds) {
		lastKey2 := len(toUserTmpRecommendUserIds) - 1
		if 1 <= lastKey2 {
			for i := 0; i <= lastKey2; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey2-i <= 0 {
					break
				}

				toUserTmpMyTopUserRecommendUserId, _ := strconv.ParseInt(toUserTmpRecommendUserIds[lastKey2-i], 10, 64) // 最后一位是直推人
				toUserTmpRecommendUserIdsInt = append(toUserTmpRecommendUserIdsInt, toUserTmpMyTopUserRecommendUserId)
			}
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "usdt" == req.SendBody.Type {
			err = uuc.ubRepo.TranUsdt(ctx, user.ID, toUser.ID, amount, tmpRecommendUserIdsInt, toUserTmpRecommendUserIdsInt) // 提现
			if nil != err {
				return err
			}
		} else if "dhb" == req.SendBody.Type {
			err = uuc.ubRepo.TranDhb(ctx, user.ID, toUser.ID, amount) // 提现
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.TranReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Trade(ctx context.Context, req *v1.WithdrawRequest, user *User, amount int64, amountB int64, amount2 int64, password string) (*v1.WithdrawReply, error) {
	var (
		u                   *User
		userBalance         *UserBalance
		userBalance2        *UserBalance
		configs             []*Config
		userRecommend       *UserRecommend
		withdrawRate        int64
		withdrawDestroyRate int64
		err                 error
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "withdraw_rate",
		"withdraw_destroy_rate",
	)

	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_destroy_rate" == vConfig.KeyName {
				withdrawDestroyRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	userBalance, err = uuc.ubRepo.GetUserBalanceLock(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	userBalance2, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.WithdrawReply{
			Status: "csd锁定部分的余额不足",
		}, nil
	}

	if userBalance2.BalanceDhb < amountB {
		return &v1.WithdrawReply{
			Status: "hbs锁定部分的余额不足",
		}, nil
	}

	// 推荐人
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.WithdrawReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds    []string
		tmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}
	lastKey := len(tmpRecommendUserIds) - 1
	if 1 <= lastKey {
		for i := 0; i <= lastKey; i++ {
			// 有占位信息，推荐人推荐人的上一代
			if lastKey-i <= 0 {
				break
			}

			tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
			tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.Trade(ctx, user.ID, amount, amountB, amount-amount/100*(withdrawRate+withdrawDestroyRate), amountB-amountB/100*(withdrawRate+withdrawDestroyRate), tmpRecommendUserIdsInt, amount2) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest, user *User) (*v1.SetBalanceRewardReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.SetBalanceReward(ctx, user.ID, amount) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.SetBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest, user *User) (*v1.DeleteBalanceRewardReply, error) {
	var (
		err            error
		balanceRewards []*BalanceReward
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != err {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	var totalBalanceRewardAmount int64
	for _, vBalanceReward := range balanceRewards {
		totalBalanceRewardAmount += vBalanceReward.Amount
	}

	if totalBalanceRewardAmount < amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	for _, vBalanceReward := range balanceRewards {
		tmpAmount := int64(0)
		Status := int64(1)

		if amount-vBalanceReward.Amount < 0 {
			tmpAmount = amount
		} else {
			tmpAmount = vBalanceReward.Amount
			Status = 2
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UpdateBalanceReward(ctx, user.ID, vBalanceReward.ID, tmpAmount, Status) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
		amount -= tmpAmount

		if amount <= 0 {
			break
		}
	}

	return &v1.DeleteBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) GetUserByUserId(ctx context.Context, userId int64) (*User, error) {
	return uuc.repo.GetUserById(ctx, userId)
}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}
	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	res := &v1.AdminConfigUpdateReply{}
	return res, nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedList(ctx context.Context) ([]*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewarded(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	return &v1.AdminAllReply{}, nil
}

func (uuc *UserUseCase) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return &v1.AdminWithdrawReply{}, nil
}
