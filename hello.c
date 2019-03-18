#include <stdio.h>
#include "wrapper.h"
#include "hello.h"

#ifdef __cplusplus
extern "C" {
#endif

struct tag_fm_model *call_create() {
    return create();
}

int call_loadModel(struct tag_fm_model *fm, char *model_file_path) {
    return loadModel(fm, model_file_path);
}

double call_predict(struct tag_fm_model *fm, const char *data) {
    return predict(fm, data);
}

#ifdef __cplusplus
}
#endif