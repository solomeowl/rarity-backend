package routers

import (
	"rarity-backend/app/controllers"
	raritycraftingone "rarity-backend/app/controllers/rarity-crafting-one"
	raritymarket "rarity-backend/app/controllers/rarity-market"
	"rarity-backend/utils/e"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) (int, int, int, interface{})

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		httpCode, msgCode, total, data := handler(c)
		controllers.Response(c, httpCode, msgCode, total, e.GetMsg(msgCode), data)

	}
}
func Init() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := r.Group("/backend")
	api.GET("/summoners", wrapper(raritymarket.GetAllSummoners))
	api.GET("/crafting_one", wrapper(raritycraftingone.GetCraftingByOwner))
	return r
}
