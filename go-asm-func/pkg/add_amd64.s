#include "textflag.h"

TEXT ·Add(SB), NOSPLIT, $0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    ADDQ AX, BX
    MOVQ BX, r+16(FP)
    RET;