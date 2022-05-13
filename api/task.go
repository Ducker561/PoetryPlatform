package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"

	"InternetPlus/service"
)

func ShowPoetry(c *gin.Context) {
	showPoetryService := service.ShowPoetryService{}
	res := showPoetryService.Show(c.Param("id"))
	c.JSON(200, res)
}

func SearchPoetry(c *gin.Context) {
	searchPoetryService := service.SearchPoetryService{}
	if err := c.ShouldBind(&searchPoetryService); err == nil {
		res := searchPoetryService.Search()
		fmt.Println(res)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
