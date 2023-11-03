// add_amd64.s
#include "textflag.h"

TEXT ·add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    ADDQ AX, BX
    MOVQ BX, res+16(FP)
    // 文件最后要有换行或在最后一行末尾写个分号（;），否则编译报错
    RET
