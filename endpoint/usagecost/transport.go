package usagecost

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	kitlogrus "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/middleware"
	"joi-energy-golang/http/serveroption"
)

func MakeCalculateLastWeekUsageCostHandler(
	s Service,
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeCalculateLastWeekUsageCostEndpoint(s)),
		decodeSmartMeterIdFromRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeSmartMeterIdFromRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return strings.Split(r.URL.Path, "/")[2], nil
}
