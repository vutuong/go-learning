Interface
- Là các types mà định nghĩa một contract nhưng nó không có implementation
- Giúp decouple code from specific implementation

Cũng có thể hiểu:
- Interface: dùng để gom các type (struct) có cùng method vào cùng một nhóm.
- Và khai báo cho nhóm thành viên đó một method dùng chung, hàm dùng chung mà không cần phải viết lại hàm đó nhiều lần.

Ví dụ sử dụng interface:
```
type bot interface {
	getGreeting() string
}
```
- Toàn bộ các type trong chương trình mà có implement function getGreeting() mà trả về string sẽ thuộc về type bot
- Khi đó ta có thể lấy và truyền giá trị của các type đó vào trong các function của type bot (tức là các function của type bot là được dùng chung bởi các type còn lại)