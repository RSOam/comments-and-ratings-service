package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RSOam/comments-and-ratings-service/commrat"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	httpAddr := flag.String("http", ":8080", "http listen addr")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "comments and ratings",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "comments and ratings service started")
	defer level.Info(logger).Log("msg", "comments and ratings service ended")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://amAdmin:"+os.Getenv("DBpw2")+"@cluster0.lkbf1.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	level.Info(logger).Log("msg", "DB Connected")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("RSOam2")

	flag.Parse()
	ctx2 := context.Background()
	var srv commrat.CommRatService
	{
		database := commrat.NewDatabase(collection, logger)
		srv = commrat.NewService(database, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := commrat.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := commrat.NewHttpServer(ctx2, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
func test() string {
	return "Ok"
}
