#include "textflag.h"

TEXT ·MyInfo(SB), NOSPLIT, $8
    MOVQ $·MyLike(SB), AX
    CALL AX
RET;