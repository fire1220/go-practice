// pkg/pkg_amd64.s
#include "textflag.h"

// NOPTR : 报告函之战
// RODATA : 只读数据
// 此方法的gopher不是私有的变量，理论上可以尾部访问
GLOBL ·Name(SB), NOPTR | RODATA, $16
DATA ·Name+0(SB)/8, $·Name+16(SB)
DATA ·Name+8(SB)/8, $11
DATA ·Name+16(SB)/16, $"hello world"

//  内存分布
// ***************
// *   Name变量   *
// ***************
// *     data    *   0X123  <- ·Name+16(SB)的地址
// ***************
// *     len     *   8
// ***************
// * hello world *   0X123  <- ·Name+16(SB)的地址
// ***************

// 定义局部变量，变量区分大小写，（局部变量关键词是左右尖括号<>）
GLOBL title<>(SB), NOPTR | RODATA , $16      // 定义局部变量
DATA title<>+0(SB)/16, $"hello Golang"       // 定义局部变量的值
// 定义全局变量（全局变量关键词是重点·）
GLOBL ·Title(SB), NOPTR | RODATA, $16
DATA ·Title+0(SB)/8, $title<>(SB)           // 设置字符串指针指向局部变量
DATA ·Title+8(SB)/8, $12                    // 设置字符串长度

//  内存分布
// ****************
// * Title 全局变量 *
// ****************
// *     data     *   0X123   <- &title
// ****************
// *     len      *   16
// ****************
// *     ...      *
// ****************
// * title 局部变量 *   0X123   <-  &title
// ****************
// * hello Golang *
// ****************
