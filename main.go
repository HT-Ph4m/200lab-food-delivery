package main

import (
	"200lab-project-1/component/appctx"
	"200lab-project-1/middleware"
	"200lab-project-1/module/restaurant/transport/ginrestaurant"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	const MY_CONN = "food_delivery:12345678@tcp(localhost:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv(MY_CONN)
	dsn := MY_CONN
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	r := gin.Default()

	appctx := appctx.NewAppContext(db)

	r.Use(middleware.Recover(appctx))

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants")

	// Post //restaurants
	restaurant.POST("", ginrestaurant.CreateRestaurant(appctx))

	// Get list //restaurants
	restaurant.GET("", ginrestaurant.ListRestaurant(appctx))

	// Update restaurant
	restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appctx))

	// Delete restaurant
	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appctx))

	// Get one restaurant
	restaurant.GET("/:id", ginrestaurant.FindDataWithCondition(appctx))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// Update
	// addr := "Mỗ Lao - Hà Đông - Hà Nội"
	// updateRestaurant := RestaurantUpdate{Addr: &addr}
	// if err := db.Where("id = ?", 2).Updates(&updateRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(updateRestaurant)
}
