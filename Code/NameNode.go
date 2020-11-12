package main

import (
  "./comms"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
)

type Server struct {

}

func (s* Server) DownloadBook(ctx context.Context, request *comms.Request_DownloadBook) (*comms.Response_DownloadBook) {

}

func (s* Server) Log(ctx context.Context, request *comms.Request_Log) (*comms.Response_Log) {

}

func (s* Server) Propuesta(ctx context.Context, request *comms.Request_Propuesta) (*comms.Response_Propuesta) {

}

func (s* Server) Distribuir(ctx context.Context, request *comms.Request_Distribuir) (*comms.Response_Distribuir) {

}

func main(){
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := Server{}
  grpcServer := grpc.NewServer()
  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
