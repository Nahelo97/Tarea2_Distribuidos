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
  "io"
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
func createChunk_v (chunk_id int, chunk []byte, bookName string) {
  s:=strconv.Itoa(chunk_id)
  name := bookName+"_"+s
  file, err := os.Create("../Chunks/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("../Chunks/" + name, chunk, os.ModeAppend)
}

func createChunk (chunk_id int, chunk []byte, bookName string) {
  s:=strconv.Itoa(chunk_id)
  name := bookName+"_"+s
  file, err := os.Create("../temp/node/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("../temp/node/" + name, chunk, os.ModeAppend)
}

func proponer (conn *grpc.ClientConn, chunks int, name string) (int,string) {
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
  estado,_ := c.Propuesta(context.Background(),&comms2.Request_Propuesta{
    Propuesta: propuesta,})
  aux:=int(estado.Estado)
  log.Printf("holi")
  return aux,propuesta
}

func read_chunk(archivo string)([]byte) {
  file, err := os.Open("../temp/node/"+archivo)
  if err != nil {
    fmt.Println(err)
    return []byte("0")
  }
  defer file.Close()
  buffer := make([]byte,100)
  for {
    bytesread, err := file.Read(buffer)
    if err != nil {
      if err != io.EOF {
        fmt.Println(err)
      }
      break
    }
    bs := []byte(strconv.Itoa(bytesread))
      return bs
  }
  return []byte("0")
}

func distribuidor(propuesta string){
  lineas:=strings.Split(propuesta,"\n")
  nombre:=strings.Split(lineas[0]," ")[0]
  cantidad,_:=strconv.Atoi(strings.Split(lineas[0]," ")[1])
  for i:=0;i<cantidad;i++{
    maquina:=strings.Split(lineas[i+1]," ")[1]
    chunk:=read_chunk(strings.Split(lineas[i+1]," ")[0])
    log.Printf(maquina)
    conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
    log.Printf(err)
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    c.DistribuirChunks(context.Background(),&comms.Request_Distribuir{
      Id:int32(i),Chunks:chunk,Nombre:nombre})
  }

}

func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
  log.Printf("Receive Book from client")
  tempChunk (int(request.Id), request.Nombre, int(request.Cantidad))
  createChunk (int(request.Id), request.Chunks, request.Nombre)
  if (request.Id != request.Cantidad) {
    return &comms.Response_UploadBook{State: int32(0)}, nil
  } else {
    var conn *grpc.ClientConn
    conn, err := grpc.Dial("dist96:9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    estado,prop := proponer(conn, int(request.Cantidad), request.Nombre)
    for ; estado == 0 ; {
      estado,prop = proponer(conn, int(request.Cantidad), request.Nombre)
    }
    if(estado == 1){
      log.Printf(prop)
      distribuidor(prop)
    }
    remover()
    return &comms.Response_UploadBook{State: int32(estado)}, nil
  }
}
func (s* Server) DownloadBook(ctx context.Context, request *comms.Request_DownloadBook) (*comms.Response_DownloadBook, error){
  return &comms.Response_DownloadBook{},nil
}
func (s* Server) DistribuirChunks(ctx context.Context, request *comms.Request_Distribuir) (*comms.Response_Distribuir, error){
  log.Printf("guardar chunk:")
  log.Printf(request.Nombre)
  createChunk_v(int(request.Id), request.Chunks, request.Nombre)
  return &comms.Response_Distribuir{}, nil
}

func remover(){
  var files []string
  root := "../temp/node/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("remover")
    panic(err)
  }
  for i:=1;i<len(files);i++{
    wea:=strings.Split(files[i],"/")
    if(wea[len(wea)-1]!=""){
      os.Remove(files[i])
    }
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
