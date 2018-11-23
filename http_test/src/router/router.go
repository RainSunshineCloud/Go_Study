package router

import (
    "net/http"
    "controller"
)

var base = & controller.BaseController{};
//路由
var Router map [string] http.HandlerFunc =  map [string] http.HandlerFunc {
    "/image" : base.Image,
    "/hello" : base.Hello,
};
