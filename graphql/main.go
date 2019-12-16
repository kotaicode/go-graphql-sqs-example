package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	graphqlRequests = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "graphql_requests_total",
			Help: "total number of graphql requests served",
		},
	)

	graphiqlRequests = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "graphiql_requests_total",
			Help: "total number of requests to graphiql ui served",
		},
	)
)

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default  log level is info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
	// increase prometheus request count
	graphiqlRequests.Inc()

	graphiql, err := graphiqlHtmlBytes()
	if err != nil {
		log.Fatal().Err(err)
	}
	w.Write(graphiql)
}

func init() {
	setupLogger()

	sqsQueueName := MustGetEnv("SQS_QUEUE_NAME")

	err := initSqsService(sqsQueueName)
	if err != nil {
		log.Fatal().Err(err)
	}
}

func main() {

	schema, err := parseSchema()
	if err != nil {
		log.Fatal().Err(err)
	}

	r := mux.NewRouter()
	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.HandleFunc("/", graphiqlHandler)

	//provide prometheus endpoint
	r.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	// Start Server
	go func() {
		log.Info().Msg("Starting Server, listening on port http://localhost:3000/")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}
