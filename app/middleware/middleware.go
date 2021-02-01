package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	. "github.com/woaijssss/common-golib/app/context"
	"github.com/woaijssss/common-golib/app/logger"
	"github.com/woaijssss/common-golib/client"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Timeout(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			//cancel to clear resources after finished
			cancel()
		}()
		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

//解决前端跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-data_type", "application/json")                                                                                                                                                         // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next() //  处理请求
	}
}

//Metric metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info(c, "Metric start")
		tBegin := time.Now()
		c.Next()
		logger.Info(c, "Metric end")
		duration := float64(time.Since(tBegin)) / float64(time.Second)
		communityId := GetCommunityID(c)
		path := c.Request.URL.Path
		errorCode := GetErrorCode(c)
		logger.Infof(c, "communityId=[%s] uri=[%s] duration=[%f] code=[%d]", communityId, path, duration, errorCode)
		// 请求数加1
		client.HttpRequestQps.With(prometheus.Labels{
			client.CommunityLabel: communityId,
			client.UrlLabel:       path,
			client.ErrorCodeLabel: strconv.Itoa(int(errorCode)),
		}).Inc()

		// 记录本次请求处理时间
		client.HttpRequestLatency.With(prometheus.Labels{
			client.CommunityLabel: communityId,
			client.UrlLabel:       path,
		}).Set(duration)
	}
}
