package v1

import (
	"net/http"

	"github.com/borentaylor05/streamline/db"
	"github.com/borentaylor05/streamline/server/api"
	"github.com/gin-gonic/gin"
)

type Retailer struct {
	router *gin.Engine
}

type RetailerLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RetailerCreate struct {
	Name     string
	Password string
	Phone    string
	Email    string
	Username string
	Address  string
	City     string
	State    string
	Zip      string
}

func Retailers(r *gin.Engine) {
	retailer := NewRetailer(r)
	group := retailer.router.Group("/api/v1/retailers")

	group.GET("", retailer.GetAll)
	group.GET("/:id", retailer.Get)
	group.POST("", retailer.Create)
	group.PATCH("", retailer.Update)
	group.DELETE("/:id", retailer.Destroy)

	group.POST("/login", retailer.RetailerLogin)
}

func (r *Retailer) GetAll(c *gin.Context) {

}

func (r *Retailer) Get(c *gin.Context) {

}

func (r *Retailer) Create(c *gin.Context) {
	var json RetailerCreate
	if c.BindJSON(&json) == nil {
		contactInfo := &db.ContactInfo{
			PhoneNumber: json.Phone,
			Email:       json.Email,
			Address:     json.Address,
			City:        json.City,
			State:       json.State,
			Zip:         json.Zip,
			Type:        "retailer",
		}

		if err := contactInfo.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			retailer := db.Retailer{
				Name:        json.Name,
				Password:    json.Password,
				Username:    json.Username,
				ContactInfo: contactInfo,
			}
			if err := retailer.Save(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Retailer created successfully"})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
	}
}

func (r *Retailer) Update(c *gin.Context) {

}

func (r *Retailer) Destroy(c *gin.Context) {

}

func (r *Retailer) RetailerLogin(c *gin.Context) {
	var json RetailerLogin
	if c.BindJSON(&json) == nil {
		retailer := &db.Retailer{
			Username: json.Username,
			Password: json.Password,
		}
		if authedRetailer, err := retailer.Login(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			token, _ := api.GenerateToken(json.Username, "retailer")
			refreshToken := db.RandomString()
			hashedRefreshToken, _ := db.HashPassword(refreshToken)
			authedRetailer.UpdateRefreshToken(refreshToken)
			c.JSON(http.StatusOK, gin.H{
				"message":      "Login Successful",
				"accessToken":  token,
				"refreshToken": hashedRefreshToken,
			})
		}
	}
}

func NewRetailer(r *gin.Engine) *Retailer {
	return &Retailer{
		router: r,
	}
}
