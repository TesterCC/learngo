package util

import (
	"fmt"
	"strings"
	"time"
)

// Bar 进度条结构体
// 用于在命令行中显示文件下载的进度和速度
// 提供了直观的下载进度反馈，帮助用户了解下载状态
type Bar struct {
	Percent int64     // 百分比，表示当前下载进度的百分比值(0-100)
	Current int64     // 当前进度位置，表示已下载的字节数
	Total   int64     // 总进度，表示文件的总字节数
	Rate    string    // 进度条字符串，用于存储进度条的显示内容
	Graph   string    // 显示符号，用于构建进度条的字符（默认为"█"）
	start   time.Time // 开始时间，记录下载开始的时间点，用于计算下载速度
}

// NewOption 初始化并创建一个新的进度条
// 参数:
//   - start: 起始位置（字节数），通常是已下载的文件大小，用于断点续传
//   - total: 总大小（字节数），文件的总大小
//
// 该方法会设置进度条的初始状态，包括当前位置、总大小、显示符号和开始时间
// 同时会根据起始位置计算初始百分比并构建相应长度的进度条
func (bar *Bar) NewOption(start, total int64) {
	// 设置当前位置为起始位置（已下载的字节数）
	bar.Current = start
	// 设置总大小
	bar.Total = total
	// 设置进度条显示符号为方块
	bar.Graph = "█"
	// 记录开始时间为当前时间
	bar.start = time.Now()
	// 计算初始百分比
	bar.Percent = bar.getPercent()
	// 根据百分比构建初始进度条
	// 每2%的进度增加一个显示符号，这样50个符号表示100%
	for i := 0; i < int(bar.Percent); i += 2 {
		bar.Rate += bar.Graph
	}
}

// getPercent 计算并返回当前下载进度的百分比值
// 返回值:
//   - int64: 表示当前下载进度的百分比(0-100)
//
// 该方法通过当前位置与总大小的比例计算百分比
// 将浮点数结果转换为整数返回
func (bar *Bar) getPercent() int64 {
	// 将当前位置和总大小转换为float32以进行除法运算
	// 然后乘以100得到百分比，最后转换为int64类型返回
	return int64(float32(bar.Current) / float32(bar.Total) * 100)
}

// getSpeed 计算并返回当前下载速度（KB/s）
// 返回值:
//   - float64: 表示当前下载速度，单位为KB/s
//
// 该方法通过已下载的字节数除以经过的时间计算下载速度
// 如果经过的时间为0或负数，则返回0，避免除以零错误
func (bar *Bar) getSpeed() float64 {
	// 计算从开始时间到现在经过的秒数
	elapsed := time.Since(bar.start).Seconds()
	// 防止除以零错误
	if elapsed <= 0 {
		return 0
	}
	// 计算下载速度：已下载字节数 / 经过时间 / 1024 = KB/s
	return float64(bar.Current) / elapsed / 1024
}

// Play 更新并显示进度条
// 参数:
//   - current: 当前已下载的总字节数
//
// 返回值:
//   - int: 本次更新新增的字节数（当前实现中似乎有计算错误，应该总是返回0）
//
// 该方法会更新进度条的当前位置和百分比，并在控制台显示进度条
// 进度条格式为：[████████████████████████        ]48% 1254.67 KB/s
func (bar *Bar) Play(current int64) int {
	// 更新当前位置
	bar.Current = current
	// 保存上一次的百分比，用于判断是否需要更新进度条显示
	last := bar.Percent
	// 重新计算当前百分比
	bar.Percent = bar.getPercent()

	// 如果百分比发生变化且是偶数（每增加2%更新一次进度条显示）
	if bar.Percent != last && bar.Percent%2 == 0 {
		// 在进度条字符串中添加一个新的显示符号
		bar.Rate += bar.Graph
	}

	// 在控制台打印进度条
	// \r表示回车但不换行，使进度条能够在同一行更新
	// %-50s表示左对齐的50字符宽度字符串
	// %3d表示右对齐的3字符宽度整数
	// %.2f表示保留2位小数的浮点数
	fmt.Printf("\r[%-50s]%3d%% %.2f KB/s",
		// 使用Graph符号重复Percent/2次构建进度条
		// 因为进度条总宽度为50，所以每2%对应一个符号
		strings.Repeat(bar.Graph, int(bar.Percent/2)),
		// 显示百分比
		bar.Percent,
		// 显示下载速度
		bar.getSpeed())

	// 注意：这里的返回值计算可能有误
	// current - bar.Current 应该总是等于0，因为前面已经设置了bar.Current = current
	// 正确的实现应该是保存旧值，然后计算差值
	return int(current - bar.Current)
}
