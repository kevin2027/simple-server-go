package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"simple-server-go/init/global"

	"golang.org/x/sync/errgroup"
)

func Start() {
	g, _ := errgroup.WithContext(context.Background())
	server, err := InitServer()
	if err != nil {
		panic(err)
	}

	// 启动http 服务
	g.Go(func() error {
		err := server.Http.ListenAndServe()
		return err
	})

	//启动grpc服务
	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", global.Config.Server.Grpc.Port))
		if err != nil {
			return err
		}
		log.Printf(" [*] Grpc serve Waiting for connection:%d", global.Config.Server.Grpc.Port)

		err = server.Grpc.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
