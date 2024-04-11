#include "textflag.h"

// 把函数名作为变量，然后调用变量函数
TEXT ·MyInfo(SB), NOSPLIT, $8
    MOVQ $·MyLike(SB), AX // 把函数指针放到AX里
    CALL AX
RET;

// 函数返回字符串
TEXT ·GetName(SB), NOSPLIT, $128
    GLOBL name<>(SB), NOPTR|RODATA, $8
    DATA name<>+0(SB)/4, $"fire"

    MOVQ $name<>(SB), AX
    MOVQ $4, BX
    MOVQ AX,  ret+0(FP)
    MOVQ BX,  ret+8(FP)
RET;