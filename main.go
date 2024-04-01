package main

import (
	"200lab-project-1/component/appctx"
	"200lab-project-1/component/uploadprovider"
	"200lab-project-1/middleware"
	"200lab-project-1/module/restaurant/transport/ginrestaurant"
	"200lab-project-1/module/upload/uploadtransport/ginupload"
	"200lab-project-1/module/user/usertransport/ginuser"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("MY_CONN_STRING")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	r := gin.Default()

	appctx := appctx.NewAppContext(db, s3Provider, secretKey)

	r.Use(middleware.Recover(appctx))

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Static("static", "./static")

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

	// Upload
	v1.POST("/upload", ginupload.Upload(appctx))

	v1.POST("/register", ginuser.Register(appctx))

	v1.POST("/authenticate", ginuser.Login(appctx))

	v1.GET("/profile", middleware.RequireAuth(appctx), ginuser.GetProfile(appctx))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// Update
	// addr := "Mỗ Lao - Hà Đông - Hà Nội"
	// updateRestaurant := RestaurantUpdate{Addr: &addr}
	// if err := db.Where("id = ?", 2).Updates(&updateRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(updateRestaurant)
}
