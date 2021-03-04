/**
 *@Description
 *@ClassName sign
 *@Date 2021/3/4 下午3:25
 *@Author ckhero
 */

package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sign/common/util"
	"time"
)

type Sign struct {
	ctx         context.Context
	userId      uint64
	redisClient *redis.Client
}

func NewSign(ctx context.Context, userId uint64, redisClient *redis.Client) *Sign {
	return &Sign{
		ctx:         ctx,
		userId:      userId,
		redisClient: redisClient,
	}
}

// 签到
func (s *Sign) DoSign(date time.Time) error {
	key := util.BuildSignKey(s.userId, date)
	s.redisClient.SetBit(s.ctx, key, int64(date.Day()), 1)
	return nil
}

// 获取签到次数
func (s *Sign) GetSignCount() (int64, error) {
	currTime := time.Now()
	key := util.BuildSignKey(s.userId, currTime)

	count, err := s.redisClient.BitCount(s.ctx, key, &redis.BitCount{
		Start: 0,
		End:   int64(currTime.Day()),
	}).Result()

	if err == redis.Nil {
		return 0, nil
	}

	return count, err
}

// 获取当月第一次签到的时间
func (s *Sign) GetFirstSignDate() (*time.Time, error) {
	currTime := time.Now()
	key := util.BuildSignKey(s.userId, currTime)
	signDay, err := s.redisClient.BitPos(s.ctx, key, 1).Result()

	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	firstSignDate := currTime.AddDate(0, 0, int(signDay)-currTime.Day())
	return &firstSignDate, err
}

//  获取持续签到的天数
func (s *Sign) GetContinuousSignCount(date time.Time) (int64, error) {
	key := util.BuildSignKey(s.userId, date)
	res, err := s.redisClient.BitField(s.ctx, key, "get", fmt.Sprintf("u%d", date.Day()), 1).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	var count int64
	for i := len(res) - 1; i >= 0; i-- {
		bitVal := res[i]
		for bitVal > 0  {
			if bitVal & 1 == 1 {
				count++
				bitVal = bitVal >> 1
			} else {
				return count, nil
			}
		}
	}
	// 计算签到天数
	return count, nil
}

// 获取该月的签到信息
func (s *Sign) GetSignInfo(date time.Time) (map[int]bool, error) {
	key := util.BuildSignKey(s.userId, date)
	day :=  date.Day()
	signInfo := map[int]bool{}
	res, err := s.redisClient.BitField(s.ctx, key, "get", fmt.Sprintf("u%d", date.Day()), 1).Result()
	if err == redis.Nil {
		return signInfo, nil
	}
	if err != nil {
		return nil, err
	}

	for i := len(res) - 1; i >= 0; i-- {
		bitVal := res[i]
		for bitVal > 0  {
			if bitVal & 1 == 1 {
				signInfo[day] = true
			} else {
				signInfo[day] = false
			}
			bitVal = bitVal >> 1
			day--
		}
	}
	for i := day; i > 0; i-- {
		signInfo[i] = false
	}
	// 计算签到天数
	return signInfo, nil
}
