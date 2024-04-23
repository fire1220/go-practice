#include "textflag.h"

// 局部变量(数组数据,数组容量是6，目前有4个位置有值，后2个位置时初始化的0数据)
GLOBL arrData<>(SB), NOPTR, $48
DATA arrData<>+0(SB)/8, $1
DATA arrData<>+8(SB)/8, $2
DATA arrData<>+16(SB)/8, $3
DATA arrData<>+24(SB)/8, $4
DATA arrData<>+32(SB)/8, $0
DATA arrData<>+40(SB)/8, $0

// 全局slice变量
GLOBL ·Arr(SB), NOPTR, $24
DATA ·Arr+0(SB)/8, $arrData<>(SB)   // data指针
DATA ·Arr+8(SB)/8, $4               // len是4
DATA ·Arr+16(SB)/8, $6              // cap是6(不能超过底层数组容量，否则可能修改其他内存)
