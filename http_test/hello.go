package main 

import (
    "net/http"
    "router"
    "log"
    "fmt"
)

//初始化
func init () {
    var config = & http.Server {
        Addr: ":8085",
        ReadTimeout: 100,
        WriteTimeout: 100,
        MaxHeaderBytes: 1 << 20,
    };

    for path,controller := range router.Router {
        fmt.Println(path,controller);
        http.HandleFunc(path, controller);
    }

    log.Fatal(config.ListenAndServe());
}

func main() {

}