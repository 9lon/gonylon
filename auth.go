package gonylon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware x
func AuthMiddleware(c *gin.Context) {
	//fmt.Println("AuthMiddleware log")
	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "StatusUnauthorized"})
	} else {
		response, _ := http.Get("http://localhost:3000/api/auth/check/token?token=" + c.Request.Header["Authorization"][0])
		contents, _ := ioutil.ReadAll(response.Body)

		var objmap gin.H
		json.Unmarshal(contents, &objmap)
		c.Set("info", objmap)

		c.Next()
	}
	//fmt.Println(len(c.Request.Header["Authorization"]))

}
