package v1

import (
	"context"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/go-project-structure/internal/gen/server/model"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/product"
	"github.com/ysomad/go-project-structure/internal/service"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

type productHandlers struct {
	tracer trace.Tracer
}

var apiCounter metric.Int64Counter

func NewProductHandlers(t trace.Tracer, m metric.Meter, s *service.Product) *productHandlers {
	var err error

	apiCounter, err = m.Int64Counter(
		"api.counter",
		metric.WithDescription("Number of API calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		panic(err)
	}

	apiCounter.Add(context.Background(), 111)

	return &productHandlers{
		tracer: t,
	}
}

func (h *productHandlers) List(p product.ListProductsV1Params) product.ListProductsV1Responder {
	ctx := p.HTTPRequest.Context()

	slogx.TraceContext(ctx, "list products", "gender", p.Gender)

	apiCounter.Add(ctx, 1)

	ctx, span := h.tracer.Start(ctx, "handler")
	defer span.End()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func(ctx context.Context) {
		defer wg.Done()
		_, span := h.tracer.Start(ctx, "get from cache")
		defer span.End()
		time.Sleep(time.Second / 4)
	}(ctx)

	go func(ctx context.Context) {
		defer wg.Done()
		_, span := h.tracer.Start(ctx, "db call")
		defer span.End()
		time.Sleep(time.Second)
	}(ctx)

	go func(ctx context.Context) {
		defer wg.Done()
		_, span := h.tracer.Start(ctx, "save cache")
		defer span.End()
		time.Sleep(time.Second / 2)
	}(ctx)

	wg.Wait()

	_ = swag.Int32Value(p.MaxPrice)

	return product.NewListProductsV1OK().WithPayload(&model.ListProductsResponse{
		NextPageToken: "YEEET",
		Products: []*model.Product{
			{
				Brand:         "gucci",
				ID:            "fds678adsfasd",
				Images:        []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Label:         "epic product",
				MerchantCount: 2,
				MerchantIcons: []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Price: &model.Price{
					Old:                12500,
					Current:            16500,
					CurrencyCode:       "RUB",
					DiscountPercentage: 5,
				},
				Sizes: []string{"35", "42", "43"},
				Title: "epic gucci product",
			},
			{
				Brand:         "gucci",
				ID:            "fds678adsfasd",
				Images:        []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Label:         "epic product",
				MerchantCount: 2,
				MerchantIcons: []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Price: &model.Price{
					Old:                12500,
					Current:            16500,
					CurrencyCode:       "RUB",
					DiscountPercentage: 5,
				},
				Sizes: []string{"35", "42", "43"},
				Title: "epic gucci product",
			},
			{
				Brand:         "gucci",
				ID:            "fds678adsfasd",
				Images:        []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Label:         "epic product",
				MerchantCount: 2,
				MerchantIcons: []strfmt.URI{"https://google.com", "https://yandex.ru"},
				Price: &model.Price{
					Old:                12500,
					Current:            16500,
					CurrencyCode:       "RUB",
					DiscountPercentage: 5,
				},
				Sizes: []string{"35", "42", "43"},
				Title: "epic gucci product",
			},
		},
		TotalPages: 1,
	})
}
