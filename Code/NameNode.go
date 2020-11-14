package main

import (
  "./comms2"
  "math/rand"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "log"
  "os"
  "net"
)

type Server struct {

}


func (s* Server) Log(ctx context.Context, request *comms2.Request_Log) (*comms2.Response_Log, error) {
  return &comms2.Response_Log{}, nil
}

func (s* Server) Propuesta(ctx context.Context, request *comms2.Request_Propuesta) (*comms2.Response_Propuesta, error) {
  tasa := rand.Intn(10)
  if (tasa < 2) {
    log.Printf("le respondi")
    return &comms2.Response_Propuesta{Estado:int32(0),}, nil
  }
  file, err := os.OpenFile("../temp/nameNode/log.txt", os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  defer file.Close()
  _, err = file.WriteString(request.Propuesta)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  }
  log.Printf("le respondi")
  return &comms2.Response_Propuesta{Estado:int32(1),}, nil
}


func main(){
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := Server{}
  grpcServer := grpc.NewServer()
  comms2.RegisterComms2Server(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
