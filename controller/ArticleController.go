package controller

import (
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Like(c *gin.Context) {
	var articleId string
	//获取点赞的文章id
	articleId = c.Param("articleId")
	//判断是否点赞
	var articleLikeKey string
	articleLikeKey = "article_user_like:1"
	articleCount := "article_like_count"
	member := service.IsMember(articleLikeKey, articleId)
	fmt.Println("M", member)
	message := ""
	if member == true {
		service.SRemove(articleLikeKey, articleId)
		service.HDecr(articleCount, articleId)
		message = "取消点赞"
		c.String(http.StatusOK, message)
	} else {
		// 未点赞则增加文章id
		service.SAdd(articleLikeKey, articleId)
		service.HIncr(articleCount, articleId)
		message = "点赞成功"
		c.String(http.StatusOK, message)
	}
}
