package rapi

import (
	"context"
	"os"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"realworld-go/conf"
	"realworld-go/infra"
	"realworld-go/internal"
)

type Presenter struct {
	rapiConf             *conf.RapiConf
	Dependency           *internal.Dependency
	app                  *fiber.App
	traceProviderCloseFn []gcommon.CloseFn
}

func (p *Presenter) InitProviderAndStart(name string) {
	p.rapiConf = conf.LoadEnvRapiConf()

	otelAgentAddr, ok := os.LookupEnv("OTEL_RECEIVER_OTLP_ENDPOINT")
	if !ok {
		log.Fatal().Msg("OTEL_RECEIVER_OTLP_ENDPOINT is not set")
	}

	spanExporter := infra.NewOTLP(otelAgentAddr)
	traceProvider, traceProviderCloseFn, err := infra.NewTraceProviderBuilder(name).SetExporter(spanExporter).Build()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed initializing the tracer provider")
	}
	p.traceProviderCloseFn = append(p.traceProviderCloseFn, traceProviderCloseFn)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(traceProvider)

	err = p.app.Listen(p.rapiConf.ListenerAddr())
	if err != nil {
		log.Fatal().Err(err).Msgf("failed starting server")
	}
}

func (p *Presenter) Closed(ctx context.Context) {
	for _, closeFn := range p.traceProviderCloseFn {
		err := closeFn(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("Unable to close trace provider")
		}
	}

	if err := p.app.ShutdownWithContext(ctx); err != nil {
		log.Fatal().Err(err).Msgf("failed graceful shutdown app fiber")
	}
}

func NewPresenter(dependency *internal.Dependency) *Presenter {
	presenter := &Presenter{
		Dependency: dependency,
	}

	app := fiber.New()
	app.Use(otelfiber.Middleware())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Post("/article", presenter.CreateArticle)
	app.Put("/article/:id", presenter.UpdateArticle)
	app.Delete("/article/:id", presenter.DeletedArticle)
	app.Get("/article/:id", presenter.FindOneArticle)
	app.Get("article", presenter.FindAllArticle)

	presenter.app = app
	return presenter
}
