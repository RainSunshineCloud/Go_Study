package controller

import (
    "fmt"
    "log"
    "net/http"
    "captcha"
)

type BaseController struct {}

func (BaseController * BaseController) Home (w http.ResponseWriter,r * http.Request) {
    str := "skdfjsk";
    data := []byte(str);
    _,err := w.Write(data)
    if err != nil {
        log.Fatal(err);
    }
}

func (BaseController * BaseController) Hello (w http.ResponseWriter,r * http.Request) {
    str := "1212";
    data := []byte(str);
    _,err := w.Write(data)
    if err != nil {
        fmt.Println(err);
        log.Fatal(err);
    }
}

func (BaseController * BaseController) Image (w http.ResponseWriter,r * http.Request) {
    captcha.New().Print(w);
}