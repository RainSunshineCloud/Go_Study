package captcha

import (
    "image"
    "image/color/palette"
    "image/png"
    "math/rand"
    "io"
)

type Image struct {
    img  *image.RGBA
    text []byte
}

type Conf struct {
    width  int
    height int
    text_index string
}

var text_list map[string]([]byte);

var conf = Conf {
    width : 200,
    height: 200,
    text_index: "case_num",
}

func init () {
    text_list = map[string]([]byte) {
        "case" : []byte{'A','B','C','D'},
        "case_num" :[]byte{'A','B','C','D','1','2'},
    };
}

func New () (*Image) {
    rect := image.Rect(0,0,conf.width,conf.height);

    pic := &Image {
        img:image.NewRGBA(rect),
    }

    return pic
} 

func SetConf (index string,value string ) {

}

func (image *Image) Line(x1 int,y1 int,x2 int,y2 int) (*image.RGBA) {
    
    if x1 > x2 {
        x1,x2 = x2,x1
        y1,y2 = y2,y1
    }
    color := palette.WebSafe[rand.Intn(216)]
    for ;x1 <= x2;x1++ {
        image.img.Set(x1,y1,color);
    }

    return image.img
}

func (image * Image) Print (w io.Writer) {
    png.Encode(w,image.img);
}