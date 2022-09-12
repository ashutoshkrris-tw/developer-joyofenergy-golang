package usagecost

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeCalculateLastWeekUsageCostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		r, err := s.CalculateLastWeekUsageCost(req)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
