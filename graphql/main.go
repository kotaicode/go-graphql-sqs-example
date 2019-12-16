package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default  log level is info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func graphiqlHandler(w http.ResponseWriter, r *http.Request) {
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
