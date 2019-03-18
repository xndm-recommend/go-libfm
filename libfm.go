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

type LibFMClient struct {
	ModelPath string
	FMModel   *Model
}

type LibLROptions struct {
	Model_path string // 模型路径
}

type Model struct {
	cModel *C.struct_tag_fm_model
}

func (c *LibFMClient) Predict(data []string) []float64 {
	p := make([]float64, 0, len(data))
	for _, s := range data {
		var cmsg *C.char = C.CString(s)
		defer C.free(unsafe.Pointer(cmsg))
		p = append(p, float64(C.call_predict(c.FMModel.cModel, cmsg)))
	}
	return p
}

func (c *LibFMClient) LoadModel(filename string) {
	var cmsg *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(cmsg))
	if ok := int(C.call_loadModel(c.FMModel.cModel, cmsg)); ok != 1 {
		errors_.CheckErrSendEmail(errors.New(fmt.Sprintf("Can't load model from %v", filename)))
	}
}

func (c *LibFMClient) InitModel(filename string) error {
	model := C.call_create()
	m := &Model{
		cModel: model,
	}
	m.cModel.LoadModel(filename)
	if model == nil {
		return errors.New(fmt.Sprintf("Can't load model from %v", filename))
	}
	return nil
}

func (c *LibFMClient) init() {
	// 模型初始化
	errors_.CheckFatalErr(c.InitModel(c.ModelPath))
}

func NewLibFMClient(opt *LibLROptions) (*LibFMClient, error) {
	c := &LibFMClient{
		ModelPath: opt.Model_path,
		FMModel:   new(Model),
	}
	c.init()
	if c.FMModel == nil {
		errors_.CheckFatalErr(fmt.Errorf("Can't load model from %v", c.ModelPath))
	}
	return c, nil
}
