package gonylon

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Router x
func Router(router *gin.Engine, baseURL string, ct interface{}) {

	routeValue := reflect.ValueOf(ct)
	routeType := routeValue.Type()

	for i := 0; i < routeType.NumMethod(); i++ {
		method := routeValue.Method(i)
		m, p := methodToRoute(routeType.Method(i).Name)
		url := baseURL + "/" + p
		// fmt.Println(m, p)
		switch m {
		case "Get":
			router.GET(url, contextAgent(method))
		case "Post":
			router.POST(url, contextAgent(method))
		case "Put":
			router.PUT(url, contextAgent(method))
		case "Patch":
			router.PATCH(url, contextAgent(method))
		case "Delete":
			router.DELETE(url, contextAgent(method))
		}

		// router.GET("/x", contextAgent(method))

	}

}

func contextAgent(method reflect.Value) func(*gin.Context) {
	return func(c *gin.Context) {
		rv := reflect.ValueOf(c)
		method.Call([]reflect.Value{rv})
	}
}

func methodToRoute(methodName string) (method string, path string) {
	var camel = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(methodName, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	//fmt.Println(a)
	return a[0], strings.ToLower(strings.Join(a[1:len(a)], "/"))
}
