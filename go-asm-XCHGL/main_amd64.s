
#include "textflag.h"

TEXT ·Xchg(SB), NOSPLIT, $0-20
    MOVQ	ptr+0(FP), BX   // 注释：接收第一个参数（ptr）放到BX寄存器里
    MOVL	new+8(FP), AX   // 注释：把第二个参数（new）放到寄存器AX里
    XCHGL	AX, 0(BX)       // 注释：原子操作，把0(BX)值和AX值交换
    MOVL	AX, ret+16(FP)  // 注释：把AX值放到返回值ret+16(FP)中（第三个参数（返回值参数））里
    RET
