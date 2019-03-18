#ifndef HELLO_H_
#define HELLO_H_

#include "wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

extern struct tag_fm_model *call_create();

extern int call_loadModel(struct tag_fm_model* fm, char* model_file_path);

extern double call_predict(struct tag_fm_model* fm, const char* data);


#ifdef __cplusplus
}
#endif

#endif // HELLO_H_
