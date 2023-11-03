// main_amd64.s
#include "textflag.h"
TEXT ·main(SB), $16-0
    MOVQ ·helloWorld+0(SB), AX; MOVQ AX, 0(SP)
    MOVQ ·helloWorld+8(SB), BX; MOVQ BX, 8(SP)
    // CALL ·printString(SB) // 可以直接调用当前表的方法printString
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET
