package middleware

import (
	"context"
	"fmt"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/micro/go-micro/v2/server"
)

func InitSentinel() {
	// 务必先进行初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "rand-service",
			Threshold:              2,
			StatIntervalInMs:       1000,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func LimitingWrapper() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			entry, err := sentinel.Entry("rand-service")
			if err != nil {
				log.Println(err)
				// return errors.BadRequest("rand-service", " try again")
			}
			/* else {
				//允许通过正常操作
			} */
			defer entry.Exit()
			return handlerFunc(ctx, req, rsp)
		}
	}
}
