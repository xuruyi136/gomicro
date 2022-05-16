package handler

import (
	"context"
	"gomicro/proto/rand"

	"gomicro/rand-service/service"
)

type handler struct {
}

func (h handler) GetRand(ctx context.Context, request *rand.RandRequest, response *rand.RandResponse) error {
	max := request.GetMax()
	response.Max = service.GetRand(max)
	return nil
}

func Handler() rand.RandHandler {
	return handler{}
}
