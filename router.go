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
		m, p, rule := methodToRoute(routeType.Method(i).Name)
		url := baseURL + "/" + p
		// fmt.Println(m, p)
		switch m {
		case "Get":
			if rule {
				router.GET(url, AuthMiddleware, contextAgent(method))
			} else {
				router.GET(url, contextAgent(method))
			}
		case "Post":

			if rule {
				router.POST(url, AuthMiddleware, contextAgent(method))
			} else {
				router.POST(url, contextAgent(method))
			}
		case "Put":

			if rule {
				router.PUT(url, AuthMiddleware, contextAgent(method))
			} else {
				router.PUT(url, contextAgent(method))
			}
		case "Patch":

			if rule {
				router.PATCH(url, AuthMiddleware, contextAgent(method))
			} else {
				router.PATCH(url, contextAgent(method))
			}
		case "Delete":

			if rule {
				router.DELETE(url, AuthMiddleware, contextAgent(method))
			} else {
				router.DELETE(url, contextAgent(method))
			}
		}

	}

}

func contextAgent(method reflect.Value) func(*gin.Context) {
	return func(c *gin.Context) {
		rv := reflect.ValueOf(c)
		method.Call([]reflect.Value{rv})
	}
}

func methodToRoute(methodName string) (method string, path string, rule bool) {
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

	checkRule := false
	if a[len(a)-1] == "Rule" {
		a = a[0 : len(a)-1]
		checkRule = true
	}
	// fmt.Println(a)
	return a[0], strings.ToLower(strings.Join(a[1:len(a)], "/")), checkRule
}
