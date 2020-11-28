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
  "time"
)

type Server struct{
}

var tiempo_p =time.Now()
var state string = "RELEASED"

//revisa si el archivo existe
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
//almacena los chunks que envia el cliente en una carpeta temporal
func tempChunk (chunk_id int, bookName string, ctdad_chunk int) {

  if (fileExists("./temp/node/" + bookName)) {
    file, err := os.OpenFile("./temp/node/" + bookName, os.O_WRONLY|os.O_APPEND, 0644)
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
    file, err := os.Create("./temp/node/" + bookName)
    if err != nil {
      // log.Printf("aqui 2")
      log.Fatalf("failed writing to file: %s", err)
    }
    defer file.Close()
    s := strconv.Itoa(ctdad_chunk)
    sas := strconv.Itoa(chunk_id)
    _, err = file.WriteString(bookName + "\n" + s + "\n"+sas)
    if err != nil {
      // log.Printf("aqui 3")
      log.Fatalf("failed writing to file: %s", err)
    }
  }
}
//almacena el chunk que le corresponde segun la distribucion
func createChunk_v (chunk_id int, chunk []byte, bookName string) {
  s:=strconv.Itoa(chunk_id)
  name := bookName+"_"+s
  log.Printf(name)
  file, err := os.Create("./Chunks/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("./Chunks/" + name, chunk, os.ModeAppend)
}
//almacena los chunks en el temporal
func createChunk (chunk_id int, chunk []byte, bookName string) {
  s:=strconv.Itoa(chunk_id)
  name := bookName+"_"+s
  file, err := os.Create("./temp/node/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("./temp/node/" + name, chunk, os.ModeAppend)
}
//genera y gestiona la propuesta para el sistema centralizado
func proponer (conn *grpc.ClientConn, chunks int, name string) (int,string) {
  c:=comms2.NewComms2Client(conn)
  var propuesta string
  ctdad_chunks := strconv.Itoa(chunks)
  propuesta = name + " " + ctdad_chunks + "\n"
  for i:=0; i<chunks; i++ {
    num := strconv.Itoa(rand.Intn(3) + 93)
    aux := strconv.Itoa(i + 1)
    propuesta += name + "_" + aux + " " + "dist" + num + "\n"
  }
  estado,_ := c.Propuesta(context.Background(),&comms2.Request_Propuesta{
    Propuesta: propuesta,})
  aux:=int(estado.Estado)
  log.Printf(propuesta)
  log.Printf("+1 mensaje DataNode-NameNode")
  return aux,propuesta
}
//genera y gestiona la propuesta para el sistema distribuido
func ProponerD (chunks int, name string) (int,string) {
  var propuesta string
  var ctdad_chunks string
  var aux int
  mensajes := 0
  for i:=93;i<96;i++{
    if(i==93){
      ctdad_chunks = strconv.Itoa(chunks)
      propuesta = name + " " + ctdad_chunks + "\n"
      for f:=0; f<chunks; f++ {
        num := strconv.Itoa(rand.Intn(3) + 93)
        aux := strconv.Itoa(f + 1)
        propuesta += name + "_" + aux + " " + "dist" + num + "\n"
      }
      if(verificar_maquinas(propuesta)){
        i=92
        continue
      }
    }
    log.Printf("hola : %d",i)
    maquina:=strconv.Itoa(i)
    conn, err := grpc.Dial("dist"+maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    estado,supererror:= c.PropuestaD(context.Background(),&comms.Request_PropuestaD{Propuesta: propuesta,})
    log.Printf("hola : %s",supererror)
    log.Printf("hola : %+v",estado)
    if supererror != nil {
      //log.Fatalf("did not connect: %s", err)
    }else{
      mensajes += 1
      aux=int(estado.Estado)
      log.Printf(propuesta)
      if(aux!=1){
        i=92
      }
    }
  }
  log.Printf("+1 mensaje DataNode-NameNode")

  return aux,propuesta
}
//se encarga de verificar si las maquinas estan disponibles
func verificar_maquinas (propuesta string) (bool){
  lineas:=strings.Split(propuesta,"\n")
  cantidad,_:=strconv.Atoi(strings.Split(lineas[0]," ")[1])
  mensajes := 0
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
    log.Printf("respuesta de maquina %s: %+v",maquina,response)
    log.Printf("mensajes DataNode-DataNode: %s", mensajes)
    if(error!=nil || int(response.Estado)!=7734){
      return true
    }
  }
  return false
}
//recibe la propuesta (distribuido) y se encarga de aceptar o rechazar
func (s* Server) PropuestaD(ctx context.Context, request *comms.Request_PropuestaD) (*comms.Response_PropuestaD, error) {
  tasa := rand.Intn(10)
  if (tasa < 1) {
    return &comms.Response_PropuestaD{Estado:int32(0),}, nil
  }
  return &comms.Response_PropuestaD{Estado:int32(1),}, nil
}
//lee un chunk desde el temporal
func read_chunk(archivo string)([]byte){
  file, err := os.Open("./temp/node/"+archivo)
   if err != nil {
      log.Fatal(err)
   }
  content, _ := ioutil.ReadAll(file)
  if (err != nil){
    log.Fatal(err)
  }
  return content
}
//se encarga de distribuir los chunks almacenados en la carpeta temporal
func distribuidor(propuesta string){
  log.Printf("distribuidor: %s",propuesta)
  lineas:=strings.Split(propuesta,"\n")
  nombre:=strings.Split(lineas[0]," ")[0]
  cantidad,_:=strconv.Atoi(strings.Split(lineas[0]," ")[1])
  mensajes := 0
  for i:=0;i<cantidad;i++{
    maquina:=strings.Split(lineas[i+1]," ")[1]
    chunk:=read_chunk(strings.Split(lineas[i+1]," ")[0])
    log.Printf(maquina)
    conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    mensajes += 1
    c.DistribuirChunks(context.Background(),&comms.Request_Distribuir{
      Id:int32(i+1),Chunks:chunk,Nombre:nombre})
  }
  log.Printf("Mensajes DataNode-DataNode: %s", mensajes)
}
//funcion que responde a un "ping"
func (s* Server) EstadoMaquina(ctx context.Context, request *comms.Request_Estado_M) (*comms.Response_Estado_M,error) {
  return &comms.Response_Estado_M{Estado:int32(7734)},nil
}
//lee un chunk desde sus chunks
func read_chunk_to_send(archivo string)([]byte){
  file, err := os.Open("./Chunks/"+archivo)
   if err != nil {
      log.Fatal(err)
   }
  content, _ := ioutil.ReadAll(file)
  if err != nil {
    log.Fatal(err)
  }
  return content
}
//funcion que se encara de enviar chunk solicitado
func (s* Server) SolicitarChunk(ctx context.Context, request *comms.Request_Chunk) (*comms.Response_Chunk,error) {
  return &comms.Response_Chunk{Chunks:read_chunk_to_send(request.Nombre),},nil
}
//funcion que se encarga de manejar la subida de un libro por parte de un cliente (centalizado)
func (s* Server) UploadBook(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
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
      log.Printf("Aceptado!\n\n")
      distribuidor(prop)
    }
    remover(false)
    return &comms.Response_UploadBook{State: int32(estado)}, nil
  }
}
//lo mismo de arriba pero en version distribuido
func (s* Server) UploadBookD(ctx context.Context, request *comms.Request_UploadBook) (*comms.Response_UploadBook, error) {
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
    estado,prop := ProponerD(int(request.Cantidad), request.Nombre)
    log.Printf("Aceptado!\n\n")
    resultado:=permisos_d(prop);
    if(resultado){
      distribuidor(prop)
    }
    remover(false)
    return &comms.Response_UploadBook{State: int32(estado)}, nil
  }
}
//permisos para sistema distribuido
func permisos_d(propuesta string)(bool){
  tiempo_p= time.Now()
  state="WANTED"
  mensajes := 0
  for i:=93;i<96;i++{
    maquina:=strconv.Itoa(i)
    conn, err := grpc.Dial("dist"+maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Printf("%d",i)
    }else{
      defer conn.Close()
      c:=comms.NewCommsClient(conn)
      c.PedirRecurso(context.Background(),&comms.Request_RecursoD{Tiempo:tiempo_p.String()})
      mensajes += 1
    }
  }
  log.Printf("Mensajes  DataNode-DataNode: %s", mensajes)

  state="HELD"
  conn, err := grpc.Dial("dist96:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatal("ay :c")
  }else{
    defer conn.Close()
    c:=comms2.NewComms2Client(conn)
    respuesta,errores:=c.Propuesta_D(context.Background(),&comms2.Request_Propuesta{Propuesta:propuesta})
    log.Printf("+1 mensaje DataNode-NameNode")
    state  = "RELEASED"
    if(errores!=nil){
      log.Fatal("ay x2 :c")
    }
    if(int32(respuesta.Estado)==1){
      return true
    }
    return false
  }
  return false
}
//solicitar recursos distribuido
func (s* Server) PedirRecurso(ctx context.Context, request *comms.Request_RecursoD) (*comms.Response_RecursoD, error){
  layout := "Mon Jan 02 2006 15:04:05 GMT-0700"
	t, _ := time.Parse(layout, request.Tiempo)
  for ;(state=="HELD"||(state=="WANTED" && tiempo_p.Before(t))); {
    log.Printf(state)
  }
  return &comms.Response_RecursoD{Estado:int32(1)}, nil
}
//recibe un chunk desde otro data node
func (s* Server) DistribuirChunks(ctx context.Context, request *comms.Request_Distribuir) (*comms.Response_Distribuir, error){
  log.Printf("guardar chunk:")
  createChunk_v(int(request.Id), request.Chunks, request.Nombre)
  return &comms.Response_Distribuir{}, nil
}
//limpia archivos residuos
func remover(kkl bool){
  var files []string
  root := "./temp/node/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("remover")
    panic(err)
  }
  for i:=1;i<len(files);i++{
    if(strings.Contains(files[i], ".pdf")){
      os.Remove(files[i])
    }
  }
  if(kkl){
    var files []string
    root = "./Chunks/"
    err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
      files = append(files, path)
      return nil
    })
    if err != nil {
      log.Printf("remover")
      panic(err)
    }
    for i:=1;i<len(files);i++{
      if(strings.Contains(files[i], ".pdf")){
        os.Remove(files[i])
      }
    }
  }
}
func main(){
  log.Printf("corriendo")
  remover(true)
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
