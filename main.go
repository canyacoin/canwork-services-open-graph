package main

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	logging "github.com/op/go-logging"
)

var (
	serviceID       = "open-graph"
	router          *gin.Engine
	logger          = logging.MustGetLogger("main")
	startedAt       = time.Now()
	firestoreClient *firestore.Client
)

func init() {
	var err error
	logFormatter := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.10s} %{id:03x}%{color:reset} %{message}`)
	logging.SetFormatter(logFormatter)
	consoleBackend := logging.NewLogBackend(os.Stdout, "", 0)
	consoleBackend.Color = true
	logging.SetLevel(logging.DEBUG, "main")

	// connect firestore
	firebaseServiceFile := mustGetenv("FIREBASE_SERVICE_FILE")
	gcpProjectID := mustGetenv("GCP_PROJECT_ID")

	firestoreClient, err = firestore.NewClient(
		context.Background(),
		gcpProjectID,
		option.WithServiceAccountFile(firebaseServiceFile))
	if err != nil {
		logger.Fatalf("unable to connect to firestore. error was: %s", err.Error())
	}

<<<<<<< HEAD
	router = gin.Default()
	router.Use(gin.Logger())
	router.HTMLRender = gintemplate.Default()
=======
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.HTMLRender = gintemplate.Default()
>>>>>>> go-testing

	router.GET("/job/:slug", jobHandler)
	router.GET("/profile/:slug", profileHandler)
	router.GET("/status/", statusHandler)
}

func main() {
	router.Run()
}
