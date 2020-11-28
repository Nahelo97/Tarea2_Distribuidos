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
  "unicode"
)

type Server struct {

}
var candado bool =false
func revisar_copia(nombre string)(bool){
  file, err := os.Open("./temp/nameNode/log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      cosa:=strings.Split(scanner.Text()," ")
      if(cosa[0]==nombre){
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
func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}
func catalogo()(string){
  var libros string
  libros=""
  file, err := os.Open("./temp/nameNode/log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      lineas:=strings.Split(scanner.Text()," ")
      if(isInt(lineas[1])) {
        libros+=lineas[0]+"\n"
      }
    }
    if(libros!=""){
      libros=libros[:len(libros)-1]
    }
  return libros
}
func encontrar_libro(numero int)(string){
  var libros string
  var contador int
  contador=0
  libros=""
  file, err := os.Open("./temp/nameNode/log.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lineas:=strings.Split(scanner.Text()," ")
    if(isInt(lineas[1])){
      contador+=1
    }
    if(contador==numero){
        libros+=scanner.Text()+"\n"
    }
  }
  return libros[:len(libros)-1]
}
func (s* Server) Catalogo(ctx context.Context, request *comms2.Request_Catalogo) (*comms2.Response_Catalogo, error) {
  return &comms2.Response_Catalogo{Libros:catalogo(),}, nil
}
func (s* Server) Pedir_Libro(ctx context.Context, request *comms2.Request_Libro) (*comms2.Response_Libro, error) {
  return &comms2.Response_Libro{Ubicaciones:encontrar_libro(int(request.Numero)),}, nil
}
func verificar_maquinas(propuesta string)(bool){
  lineas:=strings.Split(propuesta,"\n")
  cantidad,_:=strconv.Atoi(strings.Split(lineas[0]," ")[1])
  mensajes:= 0
  for i:=0;i<cantidad;i++{
    maquina:=strings.Split(lineas[i+1]," ")[1]
    conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    response,error:=c.EstadoMaquina(context.Background(),&comms.Request_Estado_M{})
    mensajes += 1
    // log.Printf("respuesta de maquina %s: %+v",maquina,response)
    if(error!=nil || int(response.Estado)!=7734){
      log.Printf("Mensajes NameNode-DataNode: %s", mensajes)
      return true
    }
  }
  log.Printf("Mensajes NameNode-DataNode: %s", mensajes)
  return false
}
func (s* Server) Propuesta(ctx context.Context, request *comms2.Request_Propuesta) (*comms2.Response_Propuesta, error) {
  for;candado;{
    log.Printf("ocupadito")
  }
  candado=true
  tasa := rand.Intn(10)
  cosa:=strings.Split(request.Propuesta,"\n")[0]
  cosa=strings.Split(cosa," ")[0]
  if(revisar_copia(cosa)){
    candado=false
    return &comms2.Response_Propuesta{Estado:int32(2),}, nil
  }
  condicion:=verificar_maquinas(request.Propuesta)
  if (tasa < 2 ||condicion) {
    candado=false
    return &comms2.Response_Propuesta{Estado:int32(0),}, nil
  }
  file, err := os.OpenFile("./temp/nameNode/log.txt", os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  defer file.Close()
  _, err = file.WriteString(request.Propuesta)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  }
  candado=false
  return &comms2.Response_Propuesta{Estado:int32(1),}, nil
}
func (s* Server) Propuesta_D(ctx context.Context, request *comms2.Request_Propuesta) (*comms2.Response_Propuesta, error){
  cosa:=strings.Split(request.Propuesta,"\n")[0]
  cosa=strings.Split(cosa," ")[0]
  if(revisar_copia(cosa)){
    return &comms2.Response_Propuesta{Estado:int32(2),}, nil
  }
  file, err := os.OpenFile("./temp/nameNode/log.txt", os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  defer file.Close()
  _, err = file.WriteString(request.Propuesta)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  }
  return &comms2.Response_Propuesta{Estado:int32(1),}, nil
}

func remover(){
  var files []string
  root := "./temp/nameNode/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("remover")
    panic(err)
  }
  for i:=1;i<len(files);i++{
    if(strings.Contains(files[i], ".pdf") || strings.Contains(files[i], "log.txt") ){
      os.Remove(files[i])
    }
  }
  os.Create("./temp/nameNode/log.txt")
}

func main(){
  log.Printf("corriendo")
  remover()
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
