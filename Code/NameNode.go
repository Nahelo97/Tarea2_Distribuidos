package main

import (
  "./comms2"
  "math/rand"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
)

type Server struct {

}


func (s* Server) Log(ctx context.Context, request *comms2.Request_Log) (*comms2.Response_Log) {
  return &comms2.Response_Log{},nil
}

func (s* Server) Propuesta(ctx context.Context, request *comms2.Request_Propuesta) (*comms2.Response_Propuesta) {
  tasa := rand.Int(10)
  if (tasa < 2) {
    return &comms2.Response_Propuesta{0},nil
  }
  file, err := os.OpenFile("../temp/nameNode/log.txt", os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  defer file.Close()
  _, err = file.WriteString(propuesta)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  }
  return &comms2.Response_Propuesta{1},nil
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
