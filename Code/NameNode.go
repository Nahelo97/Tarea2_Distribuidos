package main

import (
  "./comms"
  "./comms2"
  "math/rand"
  "strconv"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "log"
  "os"
  "net"
  "path/filepath"
  "bufio"
  "strings"
)

type Server struct {

}

func revisar_copia(nombre string)(bool){
  file, err := os.Open("../temp/nameNode/log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    log.Printf(nombre)
    for scanner.Scan() {
      wea:=strings.Split(scanner.Text()," ")
      log.Printf(wea[0])
      if(wea[0]==nombre){
        return true
      }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return false
}

func (s* Server) Log(ctx context.Context, request *comms2.Request_Log) (*comms2.Response_Log, error) {
  return &comms2.Response_Log{}, nil
}

func verificar_maquinas(propuesta string)(bool){
  lineas:=strings.Split(propuesta,"\n")
  cantidad,_:=strconv.Atoi(strings.Split(lineas[0]," ")[1])
  for i:=0;i<cantidad;i++{
    maquina:=strings.Split(lineas[i+1]," ")[1]
    conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    response,_:=c.EstadoMaquina(context.Background(),&comms.Request_Estado_M{})
    log.Printf("estado recibido:")
    log.Printf(response.Estado)
    if(int(response.Estado)!=7734){
      return true
    }
  }
  return false
}
func (s* Server) Propuesta(ctx context.Context, request *comms2.Request_Propuesta) (*comms2.Response_Propuesta, error) {
  tasa := rand.Intn(10)
  wea:=strings.Split(request.Propuesta,"\n")[0]
  wea=strings.Split(wea," ")[0]
  if(revisar_copia(wea)){
    return &comms2.Response_Propuesta{Estado:int32(2),}, nil
  }
  if (tasa < 2 || verificar_maquinas(request.Propuesta)) {
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

func remover(){
  var files []string
  root := "../temp/nameNode/"
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
  comms2.RegisterComms2Server(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
