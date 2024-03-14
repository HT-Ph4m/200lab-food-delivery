package main

import (
	"200lab-project-1/component/appctx"
	"200lab-project-1/middleware"
	"200lab-project-1/module/restaurant/transport/ginrestaurant"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// "food_delivery:12345678@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MY_CONN")
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

	// Post //restaurants
	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants")

	restaurant.POST("", ginrestaurant.CreateRestaurant(appctx))

	// restaurant.GET("/:id", func(c *gin.Context) {
	// 	id, err := strconv.Atoi(c.Param("id"))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	var data Restaurant

	// 	db.Where("id = ?", id).First(&data)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })

	restaurant.GET("", ginrestaurant.ListRestaurant(appctx))

	// restaurant.PATCH("/:id", func(c *gin.Context) {
	// 	id, err := strconv.Atoi(c.Param("id"))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	var data RestaurantUpdate

	// 	if err := c.ShouldBind(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})

	// 		return
	// 	}

	// 	db.Where("id = ?", id).Updates(&data)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appctx))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// create data
	// newRestaurants := Restaurant{Name: "BBQ", Addr: "Mỗ Lao - Hà Đông - Hà Nội - Việt Nam"}
	// if err := db.Create(&newRestaurants).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println("New id:", newRestaurants.Id)

	// Find
	// var myRestaurants Restaurant
	// if err := db.Where("id = ?", 2).First(&myRestaurants).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(&myRestaurants)

	// Update
	// addr := "Mỗ Lao - Hà Đông - Hà Nội"
	// updateRestaurant := RestaurantUpdate{Addr: &addr}
	// if err := db.Where("id = ?", 2).Updates(&updateRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(updateRestaurant)

	// Delete
	// if err := db.Table("restaurants").Where("id = ?", 3).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }
}
