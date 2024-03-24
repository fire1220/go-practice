#include "textflag.h"

TEXT ·Add(SB), NOSPLIT, $0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    ADDQ AX, BX
    MOVQ BX, r+16(FP)
    RET; // 文件最后要有换行或在最后一行末尾写个分号（;），否则编译报错