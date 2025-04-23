package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Carousel struct {
	// ID
	ID int `json:"id"`
	// 標題
	Title string `json:"title"`
	// 排序
	Order int `json:"order"`
	// 图片URL地址
	ImageUrl string `json:"imageUrl"`
	// 图片URL地址
	RedirectUrl string `json:"redirectUrl"`
	// 是否启用
	IsEnabled bool `json:"isEnabled"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 修改时间
	UpdateTime string `json:"updateTime"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
}

func main() {
	r := gin.Default()
	var Carousels []Carousel
	Carousels = append(Carousels, Carousel{
		ID:       1,
		Title:    "測試",
		Order:    1,
		ImageUrl: "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
	})

	Carousels = append(Carousels, Carousel{
		ID:       1,
		Title:    "測試",
		Order:    1,
		ImageUrl: "@/assets/2.png",
	})
	Carousels = append(Carousels, Carousel{
		ID:       1,
		Title:    "測試",
		Order:    1,
		ImageUrl: "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
	})

	r.GET("/v1/carousels/getAllCarousels", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": Carousels,
		})
		// fmt.Println(Carousels)
	})

	r.GET("/v1/basicInformation/getAllBasicInformation", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": 100,
		})
	})

	r.GET("/v1/certificates/getAllCertificates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": 100,
		})
	})

	r.Run("127.0.0.1:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
