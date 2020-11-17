#Learning Go-routines concepts through inllustrations
https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3

Go-routines: có thể coi như là lightweight threads. Tạo một go routine bằng các thêm cú pháp "go" vào trước một function.

Chú ý rằng: Mỗi khi chạy một chương trình go ( khi ta compile nó) chúng ta đã tự động tạo ra một go routine.

- Có thể tưởng tượng rằng Go-routine sẽ chỉ excecutes code từng dòng một.
- Bằng cách thêm cú pháp "go" vào trước khi gọi một function ta có thể tạo ra các go-routine chạy song song được điều phối bởi go-main routine. Mỗi go-routine sẽ chạy trên nền một CPU core. Các go-routine sẽ được lập lịch để gán vào mỗi CPU core thông qua Go-scheduler.
Example: 
```
go finder1(theMine)
```
- Mặc dù main routine cần phải được chạy để điều khiển các child routine nhưng main routine sẽ không quan tâm đến việc các child routine có thực hiện xong tác vụ hay chưa. Main routine sẽ vẫn kết thúc khi thực hiện xong các nhiệm vụ của nó (dễ hiểu: chạy qua hết các dòng code).

==> Chúng ta cần cơ chế để main routine có thể biết được các child routine đã thực hiện xong hay chưa. Nói cách khác cần có cơ chế để các go-routines có thể nói chuyện với nhau.

Vì vậy chúng ta cần khái niệm khác gọi là Channel.

Channel: Channel được sử dụng để giao tiếp giữa các routine khác nhau mà đang chạy.

- Chúng ta sử dụng channel để đảm bảo rằng main routine sẽ biết được khi nào thì mỗi child routine sẽ hoàn thành code.
- Data shared giữa các routine phải có cùng 1 kiểu. Theo đó, khi tạo một channel thì khai báo luôn kiểu dữ liệu cho channel đó, có nghĩa là channel sẽ chỉ share được 1 kiểu dữ liệu đã khai báo (Only one type)


Ví dụ, Khai báo một channel với kiểu dữ liệu ```string```:
```
c := make(chan string)
```
Sau đó ta có thể truyền channel vào trong hàm như là một biến số:
```
func test ( c chan string) {…}
```
Có thể coi một channel là một thiết bị 2 chiều
- Để gửi dữ liệu vào trong channel: ```channel <-5```
- Chờ dữ liệu từ channel và gán nó vào biến số khi nhận được:```myNumber <-channel```
- Chờ dữ liệu được gửi vào channel khi nhận được thì in nó ra:```fmt.Println(<-channel)```
- Vòng lặp chờ cho channel trả lại giá trị:
```
for l := range c {
    go checkLink(l, c)
}
```
- Tạo một anonymous function để chạy trong go-routine trong trường hợp ta chỉ cần chạy 1 lần mà ko muốn tạo một official function:
```
// Anonymous go routine
go func() {
 fmt.Println("I'm running in my own go routine")
}()
```



