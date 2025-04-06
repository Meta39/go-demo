package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ResourceHandler 定义资源操作接口
type ResourceHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// RegisterResource 通用资源注册函数，功能有限。企业级开发推荐使用第三方开源库：GORM、sqlx等
func RegisterResource(path string, handler ResourceHandler) {
	// 注册 /{path} 路由
	http.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("请求方法：%s，请求路径：%s\n", r.Method, path)
		switch r.Method {
		case http.MethodPost:
			handler.Create(w, r)
		case http.MethodPut:
			handler.Update(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 注册 /{path}/{id} 路由
	http.HandleFunc("/"+path+"/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("请求方法：%s，请求路径：%s\n", r.Method, path)
		switch r.Method {
		case http.MethodGet:
			handler.FindByID(w, r)
		case http.MethodDelete:
			handler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// ExtractIDFromPath 通用路径ID提取函数
func ExtractIDFromPath(path, resource string) (int, error) {
	segments := strings.Split(path, "/")
	if len(segments) < 3 || segments[1] != resource {
		return 0, fmt.Errorf("invalid path format")
	}
	return strconv.Atoi(segments[2])
}
