#include <stdio.h>
#include "fm_model.h"
#include<vector>
#include <iostream>
#include <iterator>
#include <algorithm>

using namespace std;

#ifdef __cplusplus
extern "C" {
#endif
struct tag_fm_model {
    fm_model fm;
};

struct fm_node {
    uint id;
    float value;
};

struct matrix {
    vector <fm_node> v;
    int len;
};

struct tag_fm_model *create(void) {
    tag_fm_model *fm;
    fm = new struct tag_fm_model;
    fm->fm.init();

    return fm;
}

int loadModel(struct tag_fm_model *fm, char *model_file_path) {
    string path = model_file_path;
    return fm->fm.loadModel(path);
}

double predictRow(struct tag_fm_model *fm, matrix &x) {
    sparse_row<FM_FLOAT> tmp;
    tmp.data = new sparse_entry<float>[x.len];
    for (int i = 0; i < x.len; ++i) {
        tmp.data[i].id = x.v[i].id;
        tmp.data[i].value = x.v[i].value;
    }
    tmp.size = x.len;
    return fm->fm.predict(tmp);
}


double predict(struct tag_fm_model *fm, const char *data_in) {
    string data = data_in;

    vector <fm_node> fm_nodes;
    matrix m_tmp;
    fm_nodes.clear();

    string space = " ";
    string colon = ":";
    size_t pos;

    string fm_row = data + space;
    size_t size = fm_row.size();
    for (size_t i = 3; i < size; ++i) {
        pos = fm_row.find(space, i); //pos为分隔符第一次出现的位置，从i到pos之前的字符串是分隔出来的字符串
        if (pos < size) { //如果查找到，如果没有查找到分隔符，pos为string::npos
            string single_ffm = fm_row.substr(i, pos - i);//*****从i开始长度为pos-i的子字符串
            i = pos;
            size_t j = 0;
            size_t inner_pos = 0;
            inner_pos = single_ffm.find(colon, j);
            string field = single_ffm.substr(j, inner_pos - j);
            j = inner_pos + 1;
            string value = single_ffm.substr(inner_pos + 1);

            // 放入ffm_nodes中
            fm_node node;
            node.id = strtol(field.c_str(), NULL, 10);
            node.value = strtof(value.c_str(), NULL);
            fm_nodes.push_back(node);
        }
    }

    m_tmp.v = fm_nodes;
    m_tmp.len = fm_nodes.size();

//    for (uint i = 0; i < fm_nodes.size(); i++) {
//        printf("%d ", fm_nodes[i].id);
//        printf("%f ", fm_nodes[i].value);
//    }
//    printf("\n");

    return predictRow(fm, m_tmp);
}

#ifdef __cplusplus
}
#endif
