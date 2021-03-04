# REDIS 版签到

## 场景
1. 签到1天送1积分，连续签到2天送2积分，3天送3积分，3天以上均送3积分等。 

2. 如果连续签到中断，则重置计数，每月初重置计数。

3. 当月签到满3天领取奖励1，满5天领取奖励2，满7天领取奖励3……等等。

4. 显示用户某个月的签到次数和首次签到时间。

5. 在日历控件上展示用户每月签到情况，可以切换年月显示……等等。


## REDIS 相关命令
[文档](https://cloud.tencent.com/developer/section/1374167)

SETBIT

GETBIT

BITCOUNT

BITPOS

BITOP

BITFIELD