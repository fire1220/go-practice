// main_amd64.s
#include "textflag.h"
TEXT ·myPrint(SB), NOSPLIT, $16-16
    MOVQ strp+0(FP), AX
    MOVQ AX, 0(SP)  // 第一个参数：数据的开始指针
    MOVQ size+8(FP), BX
    // 第二个参数：string的大小。改成 MOVQ $100, 8(SP) 试试，会发现打印了其他地方的数据
    MOVQ BX, 8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET // 这里一定要有换行或写个分号（;），否则编译报错
