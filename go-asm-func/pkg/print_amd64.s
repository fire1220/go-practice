#include "textflag.h"

//  $16表示本函数需要16字节，完成写法是：$16-16,后面的-16表示需要接收16字节的参数或返回值
// 本函数利用"真SP"准备数据
TEXT ·Print(SB), NOSPLIT, $16
    // 接收参数,因为是字符串，所有有两个参数
    MOVQ str+0(FP), AX // 第一个参数：数据的开始指针
    MOVQ str+8(FP), BX
    // 把两个参数放到打印函数的入参，开始打印
    MOVQ AX, 0(SP)  // 使用"真SP"是用加号（加号可以省略），因为"真SP"是在低地址上
    MOVQ BX, 8(SP) // 如果是用runtime·printstring(SB)函数可以不用传入该参数，因为该函数打印的是C字符串
    CALL runtime·printstring(SB) // 这个函数是打印C字符串，所以可以不传第二个长度参数
    CALL runtime·printnl(SB)
RET; // 文件最后要有换行或在最后一行末尾写个分号（;），否则编译报错

// 内存读取方向：低地址 -> 高地址
// 所以param-16(SP)相对param-8(SP)是低地址，应传入data指针
// 该函数是用"伪SP"准备数据
TEXT ·Print2(SB), NOSPLIT, $16
    // 接收入参
    MOVQ str+0(FP), AX
    MOVQ str+8(FP), BX
    // 准备打印的参数
    MOVQ AX, param-16(SP) // 注意：使用"伪SP"时是减号，因为"伪SP"是在高地址上
    MOVQ BX, param-8(SP) // 第二个参数：string的大小。改成 MOVQ $100, 8(SP) 试试，会发现打印了其他地方的数据
    // 调用打印函数
    CALL ·printStringData(SB)
    RET // 文件最后要有换行或在最后一行末尾写个分号（;），否则编译报错
