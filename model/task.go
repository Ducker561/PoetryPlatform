package model

import (
	"InternetPlus/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

//诗歌模型
type Poetry struct {
	gorm.Model
	Writer	  string `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Info 	  string `gorm:"type:longtext"`
	Content   string `gorm:"type:longtext"`
}

func (p *Poetry) View() uint64 {
	//增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView
func (p *Poetry) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(p.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(p.ID))) //增加排行点击数
}
