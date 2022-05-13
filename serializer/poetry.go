package serializer

import (
	"InternetPlus/model"
	"fmt"
)

const (
	timeLayout = "2006-01-02 15:04:05" //必须用这个 有点逆天
)

type Poetry struct {
	ID		  uint `json:"ID" example:"1"`
	Writer	  string `json:"writer" example:"李白"`
	Title     string `json:"title" example:"静夜思"`
	Info 	  string `json:"info" example:"诗人"`
	Content   string `json:"content" example:"窗前明月光"`
}

func BuildPoetry(item model.Poetry) Poetry {
	fmt.Println(item)
	return Poetry{
		ID: 	 item.ID,
		Writer:  item.Writer,
		Title:   item.Title,
		Info:    item.Info,
		Content: item.Content,
	}
}

func BuildPoetrys(items []model.Poetry) (Poetrys []Poetry) {
	for _, item := range items {
		Poetry := BuildPoetry(item)
		Poetrys = append(Poetrys, Poetry)
	}
	return Poetrys
}
