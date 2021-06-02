package main

import (
	"content_subsystem"
	"content_subsystem/cmd"
	"content_subsystem/content"
	"content_subsystem/inmem"
	"content_subsystem/local"
	"content_subsystem/postgres"
	"content_subsystem/server"
	"flag"
	"fmt"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort              = "8080"
	defaultSecret			 = "secret"
	defaultDBHost			 = "database"
	defaultDBUser			 = "livewithchat"
	defaultDBPassword		 = "default"
	defaultDBName			 = "livewithchat"
	defaultDBPort			 = "5432"
	defaultHlsScript		 = "/create-vod-hls.sh"
)

func main() {
	var (
		addr   = cmd.EnvString("PORT", defaultPort)
		// secretKey = cmd.EnvString("SECRET", defaultSecret)

		dbHost = cmd.EnvString("DB_HOST", defaultDBHost)
		dbPort = cmd.EnvString("DB_PORT", defaultDBPort)
		dbUser = cmd.EnvString("DB_USER", defaultDBUser)
		dbPassword = cmd.EnvString("DB_PASSWORD", defaultDBPassword)
		dbName = cmd.EnvString("DB_NAME", defaultDBName)

		script = cmd.EnvString("HLS_SCRIPT", defaultHlsScript)

		workerCount 	  = flag.Int("worker", 3, "number of worker")
		inmemory          = flag.Bool("inmem", false, "use in-memory repositories")
		localContent      = flag.Bool("local", false, "use local storage")
	)
	flag.Parse()

	var (
		videoDB content_subsystem.VideoRepository
		videoStorage content_subsystem.VideoStorage
		videoScheduler content_subsystem.VideoScheduler
		videoProcessFunc content_subsystem.ProcessVideoFunc
		// tokenManager stream_subsystem.TokenManager
	)

	// 	tokenManager = jwt.NewTokenManager(secretKey)

	if *localContent {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		root := fmt.Sprintf("%s/storage", path)
		videoProcessFunc = local.CreateProcessVideoFunc(root, script)
		videoStorage = local.NewVideoStorage(root)
	}

	if *inmemory {
		videoDB = inmem.NewVideoRepository()
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)

		var client *gorm.DB
		client, err := gorm.Open(psql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		defer func () {
			db, err := client.DB()
			if err != nil {
				print(err)
				return
			}
			db.Close()
		}()

		videoDB = postgres.NewVideoRepository(client)
	}

	videoScheduler = local.NewVideoScheduler()
	videoScheduler.Run()
	for i := 0; i < *workerCount; i++ {
		channel := make(chan string)
		go func() {
			for {
				videoScheduler.WorkerReady(channel)
				vid := <-channel
				println("Start to process vid#" + vid)
				err := videoProcessFunc(vid)
				if err != nil {
					println("Process vid#" + vid + ": " + err.Error())
					continue
				}
			}
		}()
	}

	var st content.Service
	st = content.NewService(videoDB, videoStorage, videoScheduler)

	srv := server.New(st)
	srv.Host.Server.Addr = fmt.Sprintf(":%s", addr)
	srv.Host.Logger.Fatal(
		gracehttp.Serve(srv.Host.Server))
}