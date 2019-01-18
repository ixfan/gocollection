package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getAwardUserWeight(users map[string]int64) (name string) {
	type awardUser struct {
		name   string
		offset int64
		count  int64
	}

	userSli := make([]*awardUser, 0,len(users))
	var sumCount int64 = 0
	for n, c := range users {
		a := awardUser{
			name:   n,
			offset: sumCount,
			count:  c,
		}
		//整理所有用户的count数据为数轴
		userSli = append(userSli, &a)
		sumCount += c
	}

	awardIndex := rand.Int63n(sumCount)
	for _, u := range userSli {
		//判断获奖index落在那个用户区间内
		if u.offset+u.count>awardIndex {
			name = u.name
			return
		}
	}
	return
}

func main() {

	var users map[string]int64 = map[string]int64{
		"a": 1000,
		"b": 2000,
		"c": 3000,
		"d": 2500,
		"f": 1499,
		"g":1,
	}
	rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i <= 100000; i++ {
		awardName := getAwardUserWeight(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 1
		}
	}
	for n,c := range awardCount {
		fmt.Printf("%v:%v \n",n,c)
	}
}