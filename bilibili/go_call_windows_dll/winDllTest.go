package go_call_windows_dll

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

/*
golang最简单的调用DLL
ref: https://www.bilibili.com/video/BV1jw411N7oc/
*/

// Method 1: 立即加载
func LoadDll1() {
	u32Dll, err := syscall.LoadLibrary("user32.dll")

	if err != nil {
		panic(err)
	}

	proc, err := syscall.GetProcAddress(u32Dll, "MessageBoxW")

	if err != nil {
		panic(err)
	}
	// ref: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-messageboxw
	/*
		int MessageBoxW(
		  [in, optional] HWND    hWnd,
		  [in, optional] LPCWSTR lpText,
		  [in, optional] LPCWSTR lpCaption,
		  [in]           UINT    uType
		);
	*/

	fromString, err := syscall.UTF16PtrFromString("Hello, go操作dll") // 拿到宽字符指针

	if err != nil {
		panic(err)
	}

	r1, _, _ := syscall.SyscallN(proc,
		0,
		uintptr(unsafe.Pointer(fromString)), // 类型不匹配，强制转换
		uintptr(unsafe.Pointer(fromString)),
		0)

	if int(r1) == 1 {
		fmt.Println("User click confirm button")
	}

}

// Method 2: 懒加载 Lazy Load
func LoadDll2() {
	u32Dll := syscall.NewLazyDLL("user32.dll")
	proc := u32Dll.NewProc("MessageBoxW")

	fromString, err := syscall.UTF16PtrFromString("Hello, go懒加载dll") // 拿到宽字符指针

	if err != nil {
		panic(err)
	}

	r1, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(fromString)), // 类型不匹配，强制转换
		uintptr(unsafe.Pointer(fromString)),
		0)

	r0, _, _ := syscall.SyscallN(proc.Addr(),
		0,
		uintptr(unsafe.Pointer(fromString)), // 类型不匹配，强制转换
		uintptr(unsafe.Pointer(fromString)),
		0)

	if int(r1) == 1 {
		fmt.Println("[r1] User click confirm button")
	}

	if int(r0) == 1 {
		fmt.Println("[r0] User click confirm button")
	}
}


// Method 3: 对Method 1的封装
func LoadDll3() {
	u32Dll, err := syscall.LoadDLL("user32.dll")

	if err != nil {
		panic(err)
	}

	proc ,err := u32Dll.FindProc("MessageBoxW")


	fromString, err := syscall.UTF16PtrFromString("Hello, go操作dll, test3") // 拿到宽字符指针

	r1, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(fromString)), // 类型不匹配，强制转换
		uintptr(unsafe.Pointer(fromString)),
		0)

	if int(r1) == 1 {
		fmt.Println("[r1] User click confirm button")
	}
}

// Method 4: 对Method 2使用的NewLazyDLL的封装，直接调用 golang.org/x/sys/windows包
// go get golang.org/x/sys/windows   如果安装不成功记得手动设置下IDE的GOPROXY=:https://goproxy.cn
func LoadDll4() {
    u32Dll := windows.NewLazySystemDLL("user32.dll")
	proc := u32Dll.NewProc("MessageBoxW")

	fromString, err := syscall.UTF16PtrFromString("Hello, go操作dll, test4") // 拿到宽字符指针

	if err != nil {
		panic(err)
	}

	r1, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(fromString)), // 类型不匹配，强制转换
		uintptr(unsafe.Pointer(fromString)),
		0)

	if int(r1) == 1 {
		fmt.Println("[r1] User click confirm button")
	}
}