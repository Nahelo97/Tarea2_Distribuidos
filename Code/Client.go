package main

import (
  "./comms"
  "./comms2"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "os"
  "strconv"
  "path/filepath"
  "fmt"
  "io"
  "log"
  "math"
  "io/ioutil"
  "bufio"
  "strings"
  "math/rand"
)

func ver_libros_para_subir(){
  var files []string
  root := "../Books/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("ver_libros_para_subir")
    panic(err)
  }
  var x int
  x=0
  log.Printf("\n\n")
  for _, file := range files {
    s := strconv.Itoa(x)
    aux:=strings.Split(file,"/")[2]
    if(x!=0){
      fmt.Println(s+".-"+aux)
    }
    x+=1
  }
}

func ver_libros_descargados(){
  var files []string
  root := "../nbooks/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("ver_libros_para_subir")
    panic(err)
  }
  var x int
  x=0
  log.Printf("\n\n")
  for _, file := range files {
    s := strconv.Itoa(x)
    aux:=strings.Split(file,"/")[2]
    if(x!=0){
      fmt.Println(s+".-"+aux)
    }
    x+=1
  }
}


func find_book_index(y int )(string){
  var files []string
  root := "../Books/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    log.Printf("find_book_index")
    panic(err)
  }
  var x int
  x=0
  for _, file := range files {
    if(x==y){
      return file
    }
    x+=1
  }
  return ""
}

func read_chunk(archivo string,numero int)([]byte){
  s := strconv.Itoa(numero)
  file, err := os.Open("../temp/cliente/"+archivo+"_"+s)
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

func subir_libro(conn *grpc.ClientConn){
  c:=comms.NewCommsClient(conn)
  ver_libros_para_subir()
  var libro int

  log.Printf("escoja el libro que quiere subir")
  fmt.Scanln(&libro)

  var archivo string
  archivo=find_book_index(libro)
  archivo=strings.Split(archivo,"/")[2]
  var chunks int
  chunks=splitter(archivo)
  for i:=1;i<=chunks;i++{
    response,_:=c.UploadBook(context.Background(),&comms.Request_UploadBook{
      Chunks:[]byte(read_chunk(archivo,i)),
      Nombre:archivo,
      Cantidad:int32(chunks),
      Id:int32(i),})
    if(response.State==2){
      log.Printf("Este libro ya existe")
    }
  }
}

func splitter(archivo string)(int){
  fileToBeChunked := "../Books/"+archivo // change here!

  file, err := os.Open(fileToBeChunked)

  if err != nil {
          fmt.Println(err)
          os.Exit(1)
  }

  defer file.Close()

  fileInfo, _ := file.Stat()

  var fileSize int64 = fileInfo.Size()

  const fileChunk = 0.25 * (1 << 20) // 1 MB, change this to your requirement

  // calculate total number of parts the file will be chunked into

  totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

  fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

  for i := uint64(0); i < totalPartsNum; i++ {

          partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
          partBuffer := make([]byte, partSize)

          file.Read(partBuffer)

          // write to disk
          fileName := "../temp/cliente/"+archivo+"_" + strconv.FormatUint(i+1, 10)
          _, err := os.Create(fileName)

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          // write/save buffer to disk
          ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

          fmt.Println("Split to : ", fileName)
  }
  return int(totalPartsNum)
}

func joiner(archivo string,totalPartsNum int){
  newFileName := "../nbooks/"+archivo
  _, err := os.Create(newFileName)
  if err != nil {
          fmt.Println(err)
          os.Exit(1)
  }
  //set the newFileName file to APPEND MODE!!
  // open files r and w
  file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
          fmt.Println(err)
          os.Exit(1)
  }
  // IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
  // defer file.Close()

  // just information on which part of the new file we are appending
  var writePosition int64 = 0
  for j := uint64(0); j < uint64(totalPartsNum); j++ {
          //read a chunk
          currentChunkFileName := "../temp/cliente/"+archivo+"_"+strconv.FormatUint(j+1, 10)

          newFileChunk, err := os.Open(currentChunkFileName)

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          defer newFileChunk.Close()

          chunkInfo, err := newFileChunk.Stat()

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          // calculate the bytes size of each chunk
          // we are not going to rely on previous data and constant

          var chunkSize int64 = chunkInfo.Size()
          chunkBufferBytes := make([]byte, chunkSize)

          fmt.Println("Appending at position : [", writePosition, "] bytes")
          writePosition = writePosition + chunkSize

          // read into chunkBufferBytes
          reader := bufio.NewReader(newFileChunk)
          _, err = reader.Read(chunkBufferBytes)

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          // DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
          // write/save buffer to disk
          //ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

          n, err := file.Write(chunkBufferBytes)

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          file.Sync() //flush to disk

          // free up the buffer for next cycle
          // should not be a problem if the chunk size is small, but
          // can be resource hogging if the chunk size is huge.
          // also a good practice to clean up your own plate after eating

          chunkBufferBytes = nil // reset or empty our buffer

          fmt.Println("Written ", n, " bytes")

          fmt.Println("Recombining part [", j, "] into : ", newFileName)
  }
  file.Close()
}

func remover(){
  var files []string
  root := "../temp/cliente/"
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
  root = "../nbooks/"
  err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    panic(err)
  }
  for i:=1;i<len(files);i++{
    wea:=strings.Split(files[i],"/")
    if(wea[len(wea)-1]!=""){
      os.Remove(files[i])
    }
  }
}
func verificar_maquinas(maquina string)(bool){
    conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    response,error:=c.EstadoMaquina(context.Background(),&comms.Request_Estado_M{})
    log.Printf("respuesta de maquina %s: %+v",maquina,response)
    if(error!=nil || int(response.Estado)!=7734){
      return true
    }
  return false
}

func mostrar_catalogo(catalogo string)(int){
  var numero int
  libros:=strings.Split(catalogo,"\n")
  for i:=0;i<len(libros);i++{
    log.Printf("%d.-$s",i+1,libros[i])
  }
  log.Printf("Seleccione un libro")
  fmt.Scanln(&numero)
  if(numero<1 || numero>len(libros)){
    log.Printf("Libro invalido")
    return -1
  }
  return numero

}
func createChunk (chunk_id int, chunk []byte, bookName string) {
  s:=strconv.Itoa(chunk_id)
  name := bookName+"_"+s
  file, err := os.Create("../temp/cliente/" + name)
  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  defer file.Close()}
  ioutil.WriteFile("../temp/cliente/" + name, chunk, os.ModeAppend)
}
func request_chunks(ubicaciones string){
  lineas:=strings.Split(ubicaciones,"\n")
  titulo:=strings.Split(lineas[0]," ")
  super_ayuda, _ := strconv.Atoi(titulo[1])
  for i:=0;i<super_ayuda;i++{
    var conn *grpc.ClientConn
    conn, err := grpc.Dial(strings.Split(lineas[i+1]," ")[1]+":9000", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %s", err)
    }
    defer conn.Close()
    c:=comms.NewCommsClient(conn)
    response,error:=c.SolicitarChunk(context.Background(),&comms.Request_Chunk{Nombre:strings.Split(lineas[i+1]," ")[0]})
    if(error!=nil){
      return
    }
    createChunk(i+1,response.Chunks,titulo[0])
  }
  joiner(titulo[0],super_ayuda)
}
func bajar_libro(){
  var conn2 *grpc.ClientConn
  conn2, err := grpc.Dial("dist96:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn2.Close()
  c:=comms2.NewComms2Client(conn2)
  request,_:=c.Catalogo(context.Background(),&comms2.Request_Catlogo{})
  libro:=mostrar_catalogo(request.Libros)
  if(libro==-1){
    return
  }
  ubicaciones,_:=c.Pedir_Libro(context.Background(),&comms2.Request_Libro{Numero:int32(libro)})
  request_chunks(ubicaciones.Ubicaciones)
}

func main(){
  var conn *grpc.ClientConn
  maquina:="dist"+strconv.Itoa(rand.Intn(3) + 93)
  for;verificar_maquinas(maquina);{
    maquina="dist"+strconv.Itoa(rand.Intn(3) + 93)
  }
  conn, err := grpc.Dial(maquina+":9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  var accion int
  var flag bool

  flag=true
  for;flag;{
    remover()
    log.Printf("\n\nBienvenido! Ingrese una opción")
    log.Printf("1-Subir Libro")
    log.Printf("2-Descargar Libro")
    log.Printf("3-Ver Libros Descargados")
    log.Printf("4-Salir")
    fmt.Scanln(&accion)
    //remover()
    switch accion {
    case 1:
      subir_libro(conn)
    case 2:
      bajar_libro()
    case 3:
      ver_libros_descargados()
    case 4:
      flag=false
    }
  }
}
