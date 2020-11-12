package main

import (
  "./comms"
  "google.golang.org/grpc"
  "os"
  "io/ioutil"
  "strings"
  "log"
  "strconv"
  "golang.org/x/net/context"
  "net"
)

type Server struct{
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func tempChunk (chunk_id int, bookName string, ctdad_chunk int) {

  if (fileExists("../temp" + bookName)) {
    file, err := os.OpenFile("../temp" + bookName, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(chunk_id)
    _, err = file.WriteString("\n"+s)
    if err != nil {
      log.Printf("aqui 1")
      log.Fatalf("failed writing to file: %s", err)
    }
  } else {
    file, err := os.Create("../temp" + bookName)
    if err != nil {
      log.Printf("aqui 2")
      log.Fatalf("failed writing to file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(ctdad_chunk)
    _, err = file.WriteString(bookName + "\n" + s + "\nchunk_id")
    if err != nil {
      log.Printf("aqui 3")
      log.Fatalf("failed writing to file: %s", err)
    }
  }
}

func createChunk (chunk_id int, chunk []byte, bookName string) {
  name := strings.Split(bookName, ".pdf")[0]
  file, err := os.Create("../Chunks" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("../Chunks" + name, chunk, os.ModeAppend)
}

func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
  log.Printf("Receive Book from client")
  tempChunk (int(request.Id), request.Nombre, int(request.Cantidad))
  createChunk (int(request.Id), request.Chunks, request.Nombre)
  if (request.Id != request.Cantidad) {
    return &comms.Response_UploadBook{State: int32(0)}, nil
  } else {
    //mandar propuesta
    return &comms.Response_UploadBook{State: int32(1)}, nil
  }
}

func (s* Server) DownloadBook(ctx context.Context, request *comms.Request_DownloadBook) (*comms.Response_DownloadBook, error){
  return &comms.Response_DownloadBook{},nil
}
func (s* Server) Log(ctx context.Context,request *comms.Request_Log) (*comms.Response_Log, error){
  return &comms.Response_Log{},nil
}
func (s* Server) Propuesta(ctx context.Context,request *comms.Request_Propuesta) (*comms.Response_Propuesta, error){
  return &comms.Response_Propuesta{},nil
}
func (s* Server) DistribuirChunks(ctx context.Context,request *comms.Request_Distribuir) (*comms.Response_Distribuir, error){
  return &comms.Response_Distribuir{},nil
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
