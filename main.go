package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"os"
	"server/components/appctx"
	"server/middleware"
	"server/module/restaurant/transport/ginrestaurant"
	ginrestaurantlike "server/module/restaurantlike/transport/ginrestauraurantlike"
	"server/module/user/transport/ginuser"
	"server/skio"

	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connection to database")
	}

	db.Debug()

	secret := os.Getenv("SYSTEM_SECRET")

	appContext := appctx.NewAppContext(db, secret)

	r := gin.Default()
	r.StaticFile("/demo/", "./socket.html")
	r.Use(middleware.Recover(appContext))
	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/login", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequiredAuthen(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuthen(appContext))
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))
	restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))
	restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.POST("/:id/like", ginrestaurantlike.LikeRestaurant(appContext))
	restaurants.DELETE("/:id/unlike", ginrestaurantlike.UnlikeRestaurant(appContext))
	restaurants.GET("/:id/list-like", ginrestaurantlike.ListUser(appContext))

	rtEngine := skio.NewEngine()
	appContext.SetRealTimeEngine(rtEngine)

	_ = rtEngine.Run(appContext, r)

	je, err := jg.NewExporter(jg.Options{
		AgentEndpoint: "localhost:6831",
		Process:       jg.Process{ServiceName: "food-delivery"},
	})

	if err != nil {
		log.Panicln(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler((1))})

	http.ListenAndServe(
		":"+os.Getenv("PORT"),
		&ochttp.Handler{
			Handler: r,
		},
	)

	// r.Run(":" + os.Getenv("PORT"))
}
