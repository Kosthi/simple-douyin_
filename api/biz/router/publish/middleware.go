// Code generated by hertz generator.

package publish

import (
	mw "simple-douyin/api/biz/middleware"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishactionMw() []app.HandlerFunc {
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc()}
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc()}
}
