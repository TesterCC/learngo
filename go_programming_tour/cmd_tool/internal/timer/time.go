package timer

import "time"

// 1.3 便捷的时间工具
// ref: https://golang2.eddycjy.com/posts/ch1/03-time2format/
// 用于返回本地时间的Time对象
func GetNowTime() time.Time {
	return time.Now()
}
// 在 GetNowTime 方法中对标准库 time 的 Now 方法进行了封装，用于返回当前本地时间的 Time 对象，此处的封装主要是为了便于后续对 Time 对象做进一步的统一处理，因为可能会涉及时区的一些问题处理。

// 推算时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// ParseDuration 方法用于在字符串中解析出 duration（持续时间），其支持的有效单位有"ns”, “us” (or “µ s”), “ms”, “s”, “m”, “h”，例如：“300ms”, “-1.5h” or “2h45m”
	duration, err := time.ParseDuration(d)  // 预处理，因为输入未知
	if err != nil {
		return time.Time{}, err
	}
	// Add 方法中，我们可以将其返回的 duration 传入，就可以得到当前 timer 时间加上 duration 后所得到的最终时间。
	return currentTimer.Add(duration), nil
}