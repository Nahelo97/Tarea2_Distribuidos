package main

import (
  "github.com/Nahelo97/Tarea2_Distribuidos/Code/comms"
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
)
func ver_libros_para_subir(){
  var files []string
  root := "../Books/"
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    panic(err)
  }
  var x int
  x=0
  for _, file := range files {
    x+=1
    s := strconv.Itoa(x)
    fmt.Println(s+".-"+file)
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
    panic(err)
  }
  var x int
  x=0
  for _, file := range files {
    x+=1
    if(x==y){
      return file
    }
  }
  return ""
}

func read_chunk(archivo string,numero int)(int){
  s := strconv.Itoa(numero)
  file, err := os.Open(archivo+"_"+s)
  if err != nil {
    fmt.Println(err)
    return 0
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
      return bytesread
  }
  return 0
}

func subir_libro(conn *grpc.ClientConn){
  c:=comms.NewCommsClient(conn)
  ver_libros_para_subir()
  var libro int

  log.Printf("escoja el libro que quiere subir")
  fmt.Scanln(&libro)

  var archivo string
  archivo=find_book_index(libro)

  var chunks int
  chunks=splitter(archivo)
  for i:=1;i<=chunks;i++{
    c.UploadBook(context.Background(),&comms.Request_UploadBook{
    Chunks:read_chunk(archivo,i),
    Nombre:archivo,
    Cantidad:int32(chunks),
    Chunk_id:int32(i)})
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

  const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

  // calculate total number of parts the file will be chunked into

  totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

  fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

  for i := uint64(0); i < totalPartsNum; i++ {

          partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
          partBuffer := make([]byte, partSize)

          file.Read(partBuffer)

          // write to disk
          fileName := "../temp/"+archivo+"_" + strconv.FormatUint(i+1, 10)
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

func joiner(archivo string){
  newFileName := archivo
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
  for j := uint64(0); j < totalPartsNum; j++ {
          //read a chunk
          currentChunkFileName := "../temp/" + strconv.FormatUint(j, 10)

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

func main(){
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  var accion int
  for{
    log.Printf("Bienvenido! Ingrese una opción")
    log.Printf("1-Subir Libro")
    log.Printf("2-Descargar Libro")
    log.Printf("3-Salir")
    fmt.Scanln(&accion)
    switch accion {
    case 1:
      subir_libro(conn)
    case 2:
      nil
    case 3:
      break
    }
  }
}
