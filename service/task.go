package service

import (
	"InternetPlus/model"
	"InternetPlus/pkg/e"
	"InternetPlus/serializer"
	logging "github.com/sirupsen/logrus"
)

//展示诗歌详情的服务
type ShowPoetryService struct {
}

//搜索诗歌的服务
type SearchPoetryService struct {
	Info string `form:"info" json:"info"`
}

func (service *ShowPoetryService) Show(id string) serializer.Response {
	var Poetry model.Poetry
	code := e.SUCCESS
	err := model.DB.First(&Poetry, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	Poetry.AddView() //增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildPoetry(Poetry),
		Msg:    e.GetMsg(code),
	}
}

func (service *SearchPoetryService) Search() serializer.Response {
	var Poetrys []model.Poetry
	code := e.SUCCESS
	err := model.DB.Model(&model.Poetry{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Find(&Poetrys).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildPoetrys(Poetrys),
	}
}
