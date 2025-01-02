package backend

import (
	"embed"
	"errors"
	"flag"
	"log"
	"net/http"
)

func Server(fs embed.FS) {
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", 8080, "Server port")
	docker := flag.Bool("docker", false, "Running in docker")
	flag.Parse()

	server := Utilities{}.SetupServer(fs, docker, host, port)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	Utilities{}.ShutdownServer(server)
}
