package main

import (
  "github.com/Nahelo97/Tarea2_Distribuidos/Code/comms"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
)

type Server struct {

}

func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {

}

func (s* Server) DownloadBook(ctx context.Context, request *comms.Request_DownloadBook) (*comms.Response_DownloadBook) {

}

func (s* Server) Log(ctx context.Context, request *comms.Request_Log) (*comms.Response_Log) {

}

func (s* Server) Propuesta(ctx context.Context, request *comms.Request_Propuesta) (*comms.Response_Propuesta) {

}

func (s* Server) Distribuir(ctx context.Context, request *comms.Request_Distribuir) (*comms.Response_Distribuir) {

}
