// pkg/pkg_amd64.s
#include "textflag.h"

// 此方法的gopher不是私有的变量，理论上可以尾部访问
GLOBL ·Name(SB), NOPTR | RODATA, $24
DATA ·Name+0(SB)/8, $·Name+16(SB)
DATA ·Name+8(SB)/8, $6
DATA ·Name+16(SB)/8, $"gopher"

//  内存分布
// **********
// * Name变量*
// **********
// *  data  *   0X123
// **********
// *  len   *   8
// **********
// * gopher *   0X123
// **********

// 设置变量Title
// 局部变量不用中点（·）
GLOBL sharedText<>(SB),NOPTR,$8      // 定义局部变量
DATA sharedText<>+0(SB)/8,$"Golang"  // 定义局部变量的值

// 此方法的Golang字符串时局部变量，不能尾部访问
GLOBL ·Title(SB),NOPTR,$16
DATA ·Title+0(SB)/8, $sharedText<>(SB)   // 设置字符串指针指向局部变量
DATA ·Title+8(SB)/8, $8                  // 设置字符串长度

//  内存分布
// **********
// * Golang *   0X123
// **********
// *  ...   *
// **********
// * Name变量*
// **********
// *  data  *   0X123
// **********
// *  len   *   8
// **********
