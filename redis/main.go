/**
 *@Description
 *@ClassName main
 *@Date 2021/3/4 下午3:58
 *@Author ckhero
 */

package main

import (
	"context"
	"fmt"
	"sign/common/db/redis"
	"sign/redis/service"
	"time"
)

func init() {
	redis.ConnectRedis()
}
func main() {
	currTime := time.Now()
	signSrv := service.NewSign(context.Background(), uint64(111), redis.GlobalRedisClient)
	_ = signSrv.DoSign(time.Now())
	_ = signSrv.DoSign(time.Now().AddDate(0, 0, -1))
	_ = signSrv.DoSign(time.Now().AddDate(0, 0, -2))

	// 第一次签到
	firstSignDate, _ := signSrv.GetFirstSignDate()
	fmt.Println(fmt.Sprintf("第一次签到 : %s", firstSignDate.String()))
	// 连续签到
	cCount, _ := signSrv.GetContinuousSignCount(currTime)
	fmt.Println(fmt.Sprintf("连续签到天数 : %d", cCount))
	// 当月签到情况
	signInfo, _ := signSrv.GetSignInfo(currTime)
	for i := 1; i <= currTime.Day(); i++ {
		signText := "NO"
		if signInfo[i] {
			signText = "YES"
		}
		fmt.Println(fmt.Sprintf("%d-%d-%d : %s", currTime.Year(), currTime.Month(), i, signText))
	}

}
