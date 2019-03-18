#ifndef WRAPPER_H_
#define WRAPPER_H_


#ifdef __cplusplus
extern "C" {
#endif
struct tag_fm_model;
struct entry;
struct matrix;

extern struct tag_fm_model *create();

extern int loadModel(struct tag_fm_model *fm, char *model_file_path);

//extern double *call_predict(const struct model *model_, const double *x, int nCols, int nRows);

extern double predict(struct tag_fm_model *fm,const char *data1) ;

#ifdef __cplusplus
}
#endif

#endif // HELLO_H_
