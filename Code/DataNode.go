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

type Server struct {

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

    len, err := file.WriteString("\nchunk_id")
    if err != nil {
      log.Fatalf("failed writing to file: %s", err)
    }
  } else {
    file, err := os.Create("../temp" + bookName)
    if err != nil {
      log.Fatalf("failed writing to file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(ctdad_chunk)
    len, err := file.WriteString(bookName + "\n" + s + "\nchunk_id")
    if err != nil {
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


func main(){
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := CommsServer{}
  grpcServer := grpc.NewServer()
  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
