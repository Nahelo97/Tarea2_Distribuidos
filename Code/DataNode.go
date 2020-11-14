package main

import (
  "./comms"
  "./comms2"
  "google.golang.org/grpc"
  "os"
  "io/ioutil"
  "strings"
  "log"
  "strconv"
  "golang.org/x/net/context"
  "net"
  "path/filepath"
  "math/rand"
  "fmt"
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

  if (fileExists("../temp/node/" + bookName)) {
    file, err := os.OpenFile("../temp/node/" + bookName, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
      log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(chunk_id)
    _, err = file.WriteString("\n"+s)
    if err != nil {
      // log.Printf("aqui 1")
      log.Fatalf("failed writing to file: %s", err)
    }
  } else {
    file, err := os.Create("../temp/node/" + bookName)
    if err != nil {
      // log.Printf("aqui 2")
      log.Fatalf("failed writing to file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(ctdad_chunk)
    _, err = file.WriteString(bookName + "\n" + s + "\nchunk_id")
    if err != nil {
      // log.Printf("aqui 3")
      log.Fatalf("failed writing to file: %s", err)
    }
  }
}

func createChunk (chunk_id int, chunk []byte, bookName string) {
  name := strings.Split(bookName, ".pdf")[0]
  file, err := os.Create("../Chunks/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("../Chunks/" + name, chunk, os.ModeAppend)
}

func proponer (conn *grpc.ClientConn, chunks int, name string) ( int) {
  c:=comms2.NewComms2Client(conn)
  var propuesta string
  ctdad_chunks := strconv.Itoa(chunks)
  propuesta = name + " " + ctdad_chunks + "\n"
  for i:=0; i<chunks; i++ {
    num := strconv.Itoa(rand.Intn(2) + 93)
    aux := strconv.Itoa(i + 1)
    propuesta += name + "_" + aux + " " + "dist" + num + "\n"
  }
  fmt.Println( "propuesta terminada")
  log.Printf("hla1")
  estado,_ := c.Propuesta(context.Background(),&comms2.Request_Propuesta{
    Propuesta: propuesta,})
  log.Printf("hla2")
  return int(estado.Estado)

}

func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
  log.Printf("Receive Book from client")
  tempChunk (int(request.Id), request.Nombre, int(request.Cantidad))
  createChunk (int(request.Id), request.Chunks, request.Nombre)
  if (request.Id != request.Cantidad) {
    return &comms.Response_UploadBook{State: int32(0)}, nil
  } else {
    var conn *grpc.ClientConn
    conn, err := grpc.Dial("dist96", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    estado := proponer(conn, int(request.Cantidad), request.Nombre)
    log.Printf("hla3")
    for ; estado == 0 ; {
      log.Printf("hla4")
      estado = proponer(conn, int(request.Cantidad), request.Nombre)
    }
    return &comms.Response_UploadBook{State: int32(1)}, nil
  }
}

func (s* Server) DownloadBook(ctx context.Context, request *comms.Request_DownloadBook) (*comms.Response_DownloadBook, error){
  return &comms.Response_DownloadBook{},nil
}

func (s* Server) DistribuirChunks(ctx context.Context,request *comms.Request_Distribuir) (*comms.Response_Distribuir, error){
  return &comms.Response_Distribuir{},nil
}
func remover(){
  var files []string
  root := "../temp/node/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    panic(err)
  }
  for i:=1;i<len(files);i++{
    os.Remove(files[i])
  }
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
