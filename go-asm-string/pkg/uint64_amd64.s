#include "textflag.h"

// uint64用4个十六进制数表示 // 结果是：0x4000300020001
// 呈现的结果与定义是相反的，实际操作可以反着定义
GLOBL ·Id(SB), NOPTR | RODATA, $8
DATA ·Id+0(SB)/2, $0x1 // 十六进制数
DATA ·Id+2(SB)/2, $0x2
DATA ·Id+4(SB)/2, $0x3
DATA ·Id+6(SB)/2, $0x4

// 结果：0x1000200030004
// 注意：定义时必须是从0(SB)开始，否则提示重复定义错误
GLOBL ·UserId(SB), NOPTR | RODATA, $8
DATA ·UserId+0(SB)/2, $0x4
DATA ·UserId+2(SB)/2, $0x3
DATA ·UserId+4(SB)/2, $0x2
DATA ·UserId+6(SB)/2, $0x1
