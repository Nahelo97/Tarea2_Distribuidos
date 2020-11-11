package main

import (
  "github.com/Nahelo97/Tarea2_Distribuidos/Code/comms"
  "google.golang.org/grpc"
  "os"
  "ioutil"
  "strings"
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
    len, err := file.WriteString(bookName + "\n" + ctdad_chunk + "\nchunk_id")
    if err != nil {
      log.Fatalf("failed writing to file: %s", err)
    }
  }
}


func createChunk (chunk_id int, chunk bytes, bookName string) {
  name := strings.Split(bookName, ".pdf")[0]
  file, err := os.Create("../Chunks" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile(file, chunk, os.ModeAppend)
}

func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
  log.Printf("Receive Book from client")
  tempChunk (request.Chunk_id, request.BookName, request.Ctdad_chunk)
  createChunk (request.Chunk_id, request.Chunk, request.BookName)
  if (request.Chunk_id != request-Ctdad_chunk) {
    return &comms.Response_UploadBook{State: int32(0)}, nil
  } else {
    //mandar propuesta
    return &comms.Response_UploadBook{State: int32(1)}, nil
  }
}
