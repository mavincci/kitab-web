package api

import "github.com/gin-gonic/gin"

var Routes map[string]gin.HandlerFunc

func init() {
	Routes = make(map[string]gin.HandlerFunc)
	Routes["/api"] = Index
	Routes["/api/login"] = UserLogin
	Routes["/api/register"] = UserRegister
}
