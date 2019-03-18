//
// Created by DELL on 2019/3/16.
//
#include "hello.h"
#include <assert.h>
#include <stdio.h>

int main(void) {
    char *a = "model.out";
    struct tag_fm_model *pApple;
    pApple = create();
//    SetColor(pApple, 1);
//    int color = GetColor(pApple);
    int b = call_loadModel(pApple, a);
    printf("loadModel = %d\n", b);
    printf("color = %d\n", 1);
    double d = call_predict(pApple, "-1 1:0.2 3:0.1 8:1 7:0.256312");
    printf("call_predict = %f\n", d);
//    ReleaseInstance(&pApple);
//    assert(b == 0);
    return 0;
}


