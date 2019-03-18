package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lwrapper -lhello -lstdc++ -lm
#include <stdio.h>
#include <stdlib.h>
#include "hello.h"
#include "wrapper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/xndm-recommend/go-utils/errors_"
)

type Model struct {
	cModel *C.struct_tag_fm_model
}

func (c *Model) Predict(data []string) []float64 {
	fmt.Println("leb data", len(data))
	p := make([]float64, 0, len(data))
	for _, s := range data {
		var cmsg *C.char = C.CString(s)
		defer C.free(unsafe.Pointer(cmsg))
		p = append(p, float64(C.call_predict(c.cModel, cmsg)))
	}
	return p
}

func (c *Model) LoadModel(filename string) {
	var cmsg *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(cmsg))
	if ok := int(C.call_loadModel(c.cModel, cmsg)); ok != 1 {
		errors_.CheckFatalErr(errors.New(fmt.Sprintf("Can't load model from %v", filename)))
	}
}

func InitModel(filename string) *Model {
	model := C.call_create()
	if model == nil {
		errors_.CheckFatalErr(errors.New(fmt.Sprintf("Can't load model from %v", filename)))
	}
	m := &Model{
		cModel: model,
	}
	m.LoadModel(filename)
	return m
}

//func main() {
//	//router := gin.Default()
//	//router.GET("/aaaa/aaaa", get_word)
//	//router.Run(":8116")
//
//	m := InitModel("model.out")
//	a := []string{"-1 1:0.2 3:0.1 8:1 7:0.256312"}
//	fmt.Println(m.Predict(a))
//}
