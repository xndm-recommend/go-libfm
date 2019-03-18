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

type LibFMOptions struct {
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

func (c *LibFMClient) LoadModel() error {
	var cmsg *C.char = C.CString(c.ModelPath)
	defer C.free(unsafe.Pointer(cmsg))
	if ok := int(C.call_loadModel(c.FMModel.cModel, cmsg)); ok != 1 {
		err := errors.New(fmt.Sprintf("Can't load model from %v", c.ModelPath))
		errors_.CheckErrSendEmail(err)
		return err
	}
	return nil
}

func initModel() *Model {
	return &Model{
		cModel: C.call_create(),
	}
}

func NewLibFMClient(opt *LibFMOptions) (*LibFMClient, error) {
	c := &LibFMClient{
		ModelPath: opt.Model_path,
		FMModel:   initModel(),
	}
	if err := c.LoadModel(); err == nil {
		errors_.CheckFatalErr(fmt.Errorf("Can't init fm model from %v", c.ModelPath))
	}
	return c, nil
}
