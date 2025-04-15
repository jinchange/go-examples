package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// addRoute 添加路由 + 处理器
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// 如果方法根节点不存在就创建
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// getRoute 返回路由节点+路径参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// 处理http请求的入口
func (r *router) handle(c *Context) {
	// 解析获得路径参数":"
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		// 设置context中的路径参数
		c.Params = params
		// 查找处理函数，处理context
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.JSON(http.StatusNotFound, H{
			"status": http.StatusNotFound,
			"msg":    c.Path + " not find",
		})
	}
}
