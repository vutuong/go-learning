#Learning Go-routines concepts through inllustrations
https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3

Go-routines: có thể coi như là lightweight threads. Tạo một go routine bằng các thêm cú pháp "go" vào trước một function.

Chú ý rằng: Mỗi khi chạy một chương trình go ( khi ta compile nó) chúng ta đã tự động tạo ra một go routine. Có thể tưởng tượng rằng Go-routine sẽ chỉ excecutes code từng dòng một.
Bằng cách thêm cú pháp "go" vào trước khi gọi một function ta có thể tạo ra các go-routine chạy song song được điều phối bởi go-main routine. Mỗi go-routine sẽ chạy trên nền một CPU core. Các go-routine sẽ được lập lịch để gán vào mỗi CPU core thông qua Go-scheduler.

Mặc dù main routine cần phải được chạy để điều khiển các child routine nhưng main routine sẽ không quan tâm đến việc các child routine có thực hiện xong tác vụ hay chưa. Main routine sẽ vẫn kết thúc khi thực hiện xong các nhiệm vụ của nó (dễ hiểu: chạy qua hết các dòng code).
=> Chúng ta cần cơ chế để main routine có thể biết được các child routine đã thực hiện xong hay chưa. Nói cách khác cần có cơ chế để các go-routines có thể nói chuyện với nhau.

Vì vậy chúng ta cần khái niệm khác gọi là Channel.