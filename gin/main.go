package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "notes/gin/docs"
)

//https://www.lixueduan.com/posts/go/swagger/

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /api

// @Security apiKey
// @securityDefinitions.apiKey apiKey
// @in header
// @name apiKey

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", ShowAccount)
			accounts.POST("/post", HandlePost)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Security
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Success      200  {object}  Account
// @Failure      400  {object}  Account
// @Failure      404  {object}  Account
// @Failure      500  {object}  Account
// @Router       /v1/accounts/{id} [get]
func ShowAccount(ctx *gin.Context) {
	fmt.Println("apiKey:", ctx.GetHeader("apiKey"))
	id := ctx.Param("id")
	account := Account{
		AccountID: id,
	}
	ctx.JSON(http.StatusOK, account)
}

type Account struct {
	AccountID string `json:"accountID"   extensions:"x-nullable,x-abc=def,!x-omitempty"`
} // @name Account

type PostReq struct {
	ID   int64  `json:"ID" example:"1" minimum:"1" maximum:"20" validate:"required" binding:"required,gte=1,lte=20"`
	Name string `json:"Name" example:"wxxx"`
}

type PostResp struct {
	// ID description
	ID int64 `json:"ID" example:"5"`
	//user's name
	Name string `json:"Name" example:"xxxxxxx"`
} // @name PostResp

type GeneralResp struct {
	Code    int64       `json:"code" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// HandlePost godoc
// @Summary      Post
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param request body PostReq true "query params"
// @Success 200 {object} GeneralResp{data=PostResp} "desc"
// @Success 200 {object} GeneralResp{data=[]PostResp} "desc"
// @Success 201 {object} GeneralResp{data=[]PostResp} "desc"
// @Success 202 {object} GeneralResp{data=interface{}} "desc"
// @Success 203 {object} GeneralResp{data=[]nil} "desc"
// @Success 204 {object} GeneralResp{data1=string,data2=[]PostResp} "desc"
// @Failure      404  {object}  Account
// @Failure      500  {object}  Account
// @Router       /v1/accounts/post [post]
func HandlePost(ctx *gin.Context) {
	fmt.Println("apiKey:", ctx.GetHeader("apiKey"))
	req := PostReq{}
	err := ctx.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, GeneralResp{
			Code:    111,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	resp := GeneralResp{
		Code:    200,
		Message: "success",
		Data: &PostResp{
			ID:   req.ID,
			Name: req.Name,
		},
	}
	ctx.JSON(http.StatusOK, &resp)
}
