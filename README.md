the plan 


<b>Start a server (Node)</b></br>
this will discove other node in the cluter according to the config file

```sh 
go-cluster node start --conif /etc/config/go-cluster.cfg 
```

<b>Write you go lang app</b></br>
Use the api for accessing files, the scheduler and exec environment</br>
Example: open a file read -> content modify content -> write back


```go
//main.go
package main

import (
    fs "github.com/leulshawell/go-cluster/api/fs"
    log "github.com/leulshawell/go-cluster/api/log"
)


func logError(err error, message string){
    if err != nil {
        log.Fatal(err.Error())
        panic(message)
    }
}
func main(){

    content := make([]byte, 1024)
    newContent := []byte("Hello  file system") 

    //other options for selecting volumes and stuff will be there
    file, err := fs.Open("path/to/file", fs.O_RDWR) //open file from the dist. file system
    logError(err, "panic when opening")

    n, err := file.Read(content)
    logError(err, "panic when reading")



    n, err := file.Write()
    logError(err, "panic when writing")

    log.Success("")

}

 ```

 <b>Build and Run the app</b></br>
 options for the scheduler like setting priority, cpu util and max-memry size 

```sh
go build main.go -o main
go-cluster run main  --stdout="$" --stderr="$"  
```

<b>Wait for the result to come to the node you are scheduling from</b>
