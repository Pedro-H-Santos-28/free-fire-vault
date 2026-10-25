package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.8" )

func vdu5rE0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V89Aeu8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6ampskaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RxmzKCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVbsPPktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEpfJzXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kO3Y1TtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYtU4PaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NukOxLjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIvLTtcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEGFfAUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3znTM1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Soo9byM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q0TVmIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FNz1TSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBCcEw2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVOy1aXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBgFkAwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxkKN6yxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMtKucHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXTAh5AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7EvwcOrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsKdGbUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqsXJRJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZfhG5JNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iN9d8gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5r2W3DoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUo2WjZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqgZp4LOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuhxdb0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZv2pVXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjDozxdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7NHV84oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIT59cJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FePzHppDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDAcmXIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhxSN7QRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tPkBkDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAuOZfPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYNIhwobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iZMIgExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6xWLsXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4ejIlNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvLcjL2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZ2TiHJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJqSXe4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLsqJ25XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JzLRAmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xiw6Ni4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gB3FiUPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pO9Lx2RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGECmEZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHEjxes8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HOp9TW5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZGB0sYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhQTLZcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bt7DS5mxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYgcqQRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7vyXmPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbidFZ0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O8AaBzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omNA2z4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttIgKCLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OlO7frZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubpqHivTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I76Q1OniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5oH4b9X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgmkH8NFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mj277JRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BroxjdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkavWHlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hx7xBH0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qCoBrtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oadEWsFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyGmAkeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KeBlk7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SL2htFH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeB7CR2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v76RJv6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7sw6NfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8tY8iUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPf8EaYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAxbm3TxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhVMnVDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8g8dI4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3rfzl5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yqz6TMaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCoubqhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB71EA3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upgeTmuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bJVd1EVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjCPYZSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func velueiyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9czBLXQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhLdnMBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sg7hlAzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pd3H4nkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w54LHPsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gc4ir6vnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHXzmPC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Olv6Vgl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLjSF2n0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJydXs9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc6Ch5jvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYcu1tB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kjoQhVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGXSvOAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfWqfXqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xhr2eZ3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQwXRhmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEenA9guWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CL4sJNYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzdhvHVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTOaH1quWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AyQ8WRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NX7r78ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDTsG3XYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAdFcKw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jvyt5ss3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bY9jJqpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arsBADgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E667hEU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bd94XLGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4zzsqX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjLa1hXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usCqZSNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOav9F0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1TMdjyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thzrdDg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMDg5uIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2UKbHpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxLDY6wQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jgdfNLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbcIesLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DYJ6C7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJhFul7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE4Aop2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiKuUO1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhb3p7fyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gulnWRF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHvVwvLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsYmSHbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1LX4TBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBEiv3ALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4No0wf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2762kptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuFTxAy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0kWuNNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmmQOD9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NImVCG06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UkoOYBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02FEgfa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixt8QYPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYnPO9nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22qQpv1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2vivv8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euS5DymyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpUkwTGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYLQA9vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vW1nAydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5MXvfwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1FGzgYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nqD2RjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HjUG5yJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDgL7vpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 140I1wzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htctZ2wAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpdH25UEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mR3dOCVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BMUbOu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW1jeoF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0ILgo0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlVgNfIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVKfDeROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9sxkGnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtHTLt9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DORXqb5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGplEpC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUEINQxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2W0ymMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAjpWyXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UL5EZk5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cL30zuHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DisOVspAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvRfBcepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKR5cCQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcYhhd7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FniEzEpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0x8W1EqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5O3Nd1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZxsdCM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYyfo3vIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2eNinYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZBioPPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TarecdWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQDWTuuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzXTrgMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eo8vcc6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOsr9hUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPC5x0CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ct7hQGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLKHnxWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93qhZ9ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICdzCFwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94vEKCQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXQcrMd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gi6CdXLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjYWRntNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9Qn3qjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvKy1LCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGAnlx5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhRTB3v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BssUOgeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Swb4yot4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tk5FG1TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHiw7aPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtMMEzuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUqrcRbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBoCWLUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHsqbbAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9SLwILKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZDppOGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qq68Y3giWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAv63ojEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iC4SohEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWCfXq0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqNwUDM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdT4p5xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AahvU9lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEDh92NuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQkXgiyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddjeLrEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyeAB9ZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYb5DRahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQpBxVWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNHAI1FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d99qwdcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB0crfYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evZDwp8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2QpzlDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func js1bmxEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU9jA21YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOIEfRToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHq31EOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNXO0pKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMUriiunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2ecLS8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OPMI4mjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQFGbXE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNEueMVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsfAeCQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksYsXNSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pjo8Ghe0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydJxSlnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0VqCNJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCAUOjEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TISzrSb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNv1xNgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M36zku9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGv6FsmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhYX9PRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8zXYR5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2SB8LPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtHqPtGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KlxEGitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stJ8njlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGSBu6K5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYyCpXpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYyvEKdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgQCtlC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEbLCQhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 153zbI7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gV9vp83CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9YLO62uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNRTeYvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9Hel80cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNdo2dhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gJRm1HEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2feXDKi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0S3AzYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZStI1SeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS06DqsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gP4q0VbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgFhndU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiakmUC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJSlpKEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuxlblJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUTRCOBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g88YKxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func es9pSQTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuurGikKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMsRWijEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoOx4ohGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qipPWvm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4D9YuiZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7mHLkQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEyAsavDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OhcrxjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZba0METWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTME5942Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbgZBqhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Sgu0S4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qk4akbxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2XjpERJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waJdPegVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4bKMv2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KThUfVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hD2DrOkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqaGCWAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4A2FXmrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaGZ5x49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdviOMIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwezGrW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvYJookAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMLJ4HIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SmlRhLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 520F1pqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APAKKE87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lokdbPNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mofOz5r7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XioCtlatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzIjXImkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHJaed12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jt2nOvlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCJy0uLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3SRKxcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBD9IlLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJrZdtdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pxgzn31mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26H58KW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Tp97iKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd438hYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vw2ZO86pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y7xlq5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nd2orFwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELWfBULaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxowhglCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVGOUlZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnTVzcUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCeeh4hzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFRAFxekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nBgtRUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NpgRK5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqlTUMvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihz5gVG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufW3wfWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUvGueXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0A7fKCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJddZVVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26YpPF0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxmDY4vLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYUwoAStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0lllkQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfcbznGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVBvaEd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBs6j2h1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfZxbx3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQhkmBjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1mZZJ5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htUDzwY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBfRRzI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HA3VMFwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfKSWpKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjhBUy0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5Z0TpzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnqBbLkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RExD8HtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1R5bj8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTHysnYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k47vbWAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eaY8fByzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQ1bSAgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQWKgZ76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ScB0Zi2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erzLlVxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvb1F3lJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Pyux23lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bxFKkj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0v9BTHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7dLml3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rnZBGtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func javQed4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Grconb6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kb4RYfhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YKWfHqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IOApJ6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Fi7kh5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJUCGKIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsEvmyV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwjIUH97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rS7OlHorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TitFjm2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJsXgmO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIAlM6IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRcQfrEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1tOE6FIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvMsfRlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hr70JZXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dW0NLR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gIIi69LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjkBbM6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dGC3Sj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCu9TTRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgBTTmAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NCOtfdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r10cKQzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oQf13x6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5mFMgEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rr43HVcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0YxFGZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBllkqSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL7HpPkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hi7a0C9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHrbbHuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKWqZ3M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ehh2FIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzLCakzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0xk5CyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYqEq1lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiS0hiuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpY5xcR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4O9RIMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djnQ1jbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciYWnx72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjWL4Z3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRpHN5gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMqDGq58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGCz8R7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWzZ1HjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQHRW9TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEy7GtrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKZ0DD9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqd36C2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTc6IOVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuqYfnhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwpNxzKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuTQSfjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEc7LFQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HofyhMa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWLrr8kpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxxmggfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g05QcGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXHievbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7TiCSwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06M2TmHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWrmteu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bukywfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSrXsj6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccSlJTuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdaSAnW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80AKx22MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2ptuQK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ov6u2dLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvIXcnxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68LVEy06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LBjmjReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hV4iGtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUvgacJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhDfii0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryomeyJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cZ9S3CEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qicBP63EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsMoa771Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdqpRWTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlmd5Z1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUdrbKtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jL1FznReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6UVbRBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apKc23nzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8QNGJX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WngB9Q8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2ZpUUcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cKtC3WXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjF3qBB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDY2KzhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkdo3U9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiTNC9r4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gphdDYvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzWj9wyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCFI0gHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d11gVqJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CP2lm2VZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH7ZhzmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHkPwHwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8S4CyRy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbHf7gRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcFBe4skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49XBnMPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADSILLbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twbVzsaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xz0OjZg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ij8sGAnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BgZAWzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3OeZGwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YboHQ4EcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oorCLvm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3o6BsiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZkvznprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nS4TSMorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnGxTTOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITlqGFJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwGZdYIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwcsQdWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHt1tIIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zQEVDYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQI9yrYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7uhgUfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTzpZVEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W05NPKwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rh8rdwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfziCRC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXriJLlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VhGXkewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pw0T1k5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjPlNzYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avTJVhmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NriZRVXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHeuR29FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQEDJVS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPnTihTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIyDfjfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yh0Tw1YEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMGgS6k9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCceyP19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8BcK5SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDa5b9PCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MB1b2PlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgqEIafiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MO0LeROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyaGJ6kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHypyO7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCb61Zd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY7xuEeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SX2Au4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E57JUF9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXZeDU01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElY1c5bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhXMsR4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yug4XwZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NrSnJYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zx1a7Y0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjD1bQPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i13Np9iUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2m72gIaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ozJBoRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83ih7rZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sC6DIA2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8XxAVMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7Djs2HHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6V4GrNMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaNDBsraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oiSKkQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n6jYQOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWTEELrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzufMdvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r32ouM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQ7MPCi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AugA7NY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtnxfCiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfCIhcUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9AcLwIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJP3jxJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQBjR9NvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP4GcjysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n97lPryJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func on79geyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEAWQvKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPIGSQFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odk4xroHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrTVmXsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ocns05MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahqMqZW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sIsdmxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErlUDwZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iwiz2rxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iC2JipRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzJWF1ORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAfXxAy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJKgtE54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfxpuyTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sc6WMHTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph7ETbDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dF38KqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBmYfUBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hv8k2Im2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtfS7nHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62HmZ7LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViBCy3F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sh4p46KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNgcQW58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGqamxOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ5FWPtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsMCFHEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2GsBlz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxG9946VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyXZWF6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2in0LpHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdaYK0TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5PKR3WzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTSpI6ymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THIMr7wjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14xoldq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVSgxNtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gJl5tzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nal4jFQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMhqmTOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axjoPg56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSCvdX9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZOQvYaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zh1XuMoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jg6MUYxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5a3xSG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7coAQcE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaQuLdYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7aAzDEruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmriGbiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtgBf11ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u3RIqRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0weK4lcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvV14sHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVsxGv3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzXkQt7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiQzi9XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEyCNr8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kj9O3kS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VprtGbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbS9ZLmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngQS6QXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8X358f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waBdJ9k1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OjcrxeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func preQs0ehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2Q3FmE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nZ7AQRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvUcoyCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9JE54lbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1494ljQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ogrF4dkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18S9IBQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBo0ib1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfnHJruQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mry9H0xzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXWrTvc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWFo6XopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func On4vlmzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3vPHEV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5c9xBnGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U61l7l4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7AUmRd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEJBDtCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xS650mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vw8lTFdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YR6ENGffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sS6ldTUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TF8Yt590Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rN2IYqrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKljitpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YajxvnXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX0ZdE4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tlo6PuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEx4Nx54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUY2BLWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guIfDviFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8MWHxwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jG3Kj32nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHNXsV2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSskkyyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzlkxTpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeJ2tMoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgDYusHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTCd7q33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYg6uJxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHDOR5tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3wFrgE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJBmSHNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Dlpyo2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFUnlnrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22RpIATeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iluWDXGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGG8uGNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpxZZDejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mr7GfecEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HF8aYp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cAfEDsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5D7JjwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5L0DOjJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lKkY1bmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03JbqnLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvXPXDd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L27uXEbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItLF4BkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmRJwOPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uq3hTZU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sc1vKTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fnibUlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gP01gYHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi3FYL2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpwewWDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70llHEKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThbqcsmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4VUjIUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H683PubxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6hybFszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UyKZJxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1JN24NeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCbZbKxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKaTVYHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TguQJbmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func na7fkM8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgKgB2GcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VloPstKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8nD5R2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjs6FL0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnKTAjZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKFdXCEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLeQ4m8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUcXyvYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rptTeHneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ly1DdwwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHVkbAyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPTxacr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcfQh6X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNOxX34EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jvdtnikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZ4iTWBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiZuGs3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gK7N5zBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpCZn5ssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CG57723EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc2Okq2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aACzfzopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJiCeuaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZuohVlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ic2x4IVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBSLIZlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxkRSpR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDfilwHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHmHHAXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12IeIdCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhL1X3kVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iT7mKlUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfopaKFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjUW3ZkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHtV6v1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKgEwZJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1P6cAQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFPnOfUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FE5L9xQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzVkcqnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZzFIF6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBvqkJo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAkXVISeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9DWC7lmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6X0XnuVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrl3RbZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KR4k6Fi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQIhj5HEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHNm4xQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXrm7nEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCEKMbgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0rO53bdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ns21BoyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLkeRV1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WH0QYCX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ab0BgzOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func et6BnpxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pn4VwU8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICRAtoQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvP4GM08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozL2z2icWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iE9xTStGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3Mp1esuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFs7g0x8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVk5IEQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fd24F9CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLmBO74mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROerHke2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2om891DWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6asrDBm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLrT7g37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzfFifCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NTmrTyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INTR9BfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2h36iLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owiRUrcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rt36K0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGmmXpRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7PnaO8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7meEIesUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zMOYv2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iZamYuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ap4oMM6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0mzFC9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ozbl3dfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzXsUMrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UvcbjKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0Y7FQY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmAKE0OuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEfgXZjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xW848ilAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GnnOrnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c77ag2teWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sauXGJFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpAe6eS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7yUH2HQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C26te948Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bygEmlTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPd2r7anWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZACPXEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3Ye8q0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VOlVUW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0S6QOtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQIVOucZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxFC80teWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9u24EHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHmzRs3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUW2781PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2VQSBxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgbpJ4fLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFIcHo7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCXrJqwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2czMImeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVtYL2a6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgBXzASJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0h2zI9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UfdsNIG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G37vWhxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4lhdZdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQxgFxr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEz7i9RPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TfZshUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOkc3Z9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJ7O23wLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwwKW8FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOmzp4MAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqLUUVO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJQ5JPlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgdoaWuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhFfxKylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnA1SpL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35NiKmjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZ1zMe4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4y4nqTAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1vrjF97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNKPd7FtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7m5rnE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmnjjtPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvyVMrK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KsfMsqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3d33WtQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQykaNzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6KLaqzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErDFebCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARV0VfxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3uDo7AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKDyYwG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvKqLrIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aH5VpGlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIEkkGEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JnVYGtVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2g1dze5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fV18KlFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKniAG7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2olwOUXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jz1mMh3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ksyga5VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0H8X8x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3MOhaIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCsoh0TiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9Ad6GaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8qiRHJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xISDFtJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paks469WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3db9DXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7OB6VpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WdZkaeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRb9FOAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ErJmWPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXaOmIXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgo7wPofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PAaypcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVBJ5oh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXwCdtquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TUy20lSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7nv50ViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjsvWIrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL3VVRjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPbrNTrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAdhNnRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIWjcvoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0jV3prlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9tWLQWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVogD1phWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBTZinR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RdDDkT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbc1RzKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBe807RRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gd86GAqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTRtvKbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9ldZiVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTaMMSnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naAgTYUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hqq7SZkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0esJLuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Is7AFdoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Usdb8o2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Do7BWqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aiF56MZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNMB57zcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func us2hKKBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWSOoU7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOQ8GQW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwvFRIPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7FCtdxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTNOwjRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfzPnNUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOfy9M0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpAaauf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUSoqX0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nb0YLRMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aU1akFClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMXDrw8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X81PRFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8DWAuQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fF8rBwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsG8DYL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O74T435yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVTG8UFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dG7LUpADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zzq0R0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GU8jbHmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eS1qmkuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDnCURlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRvCyuQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbKopCoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiVsJ0FGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xT4tE3LsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViR691L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPCWckIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WnKW421Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWxz8UT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDVudZs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sevdHetDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiXKN5jkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjS6r81kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfmTAfklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGSXRlEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjXx0zEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKxWbqVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqmYRR8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eu4cCTDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvUdWMPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCq3jF7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgcQNGZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbDUsdk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tk9I3K08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIut2G9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egS9imUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bsmg8GAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MLaA2mqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpWFyBSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func et1A7RT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ui2b3KmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YV0qA7dtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffGnsLaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMEgVTvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxdDbtTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SikzJHafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuWzXiugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHdXA5fgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IS0I0FKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2eDRfXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLTJazb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj9dljWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBswTGEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZI0lgLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiuPHD5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBgxY2v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUWAAOy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5J1rswLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kig7rLI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nch6e7uwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oni7yBU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWWU5xTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XqazZaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdC9FHOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vf4frEf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pa9sTN8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncJWNepYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sUBzmzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSpHo69RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2mlfm1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLxMOF5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T1LdktTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G891IiZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TS1YFRbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08zA8ADEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HAIRNMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Cr5CFQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6Lz3WGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMBi8C8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiGB7P28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ezweolQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKxTq7aMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odwmyH1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rd0Gu0iFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjCQCG4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2gHWjunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxEnO4DQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzmJNTREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wc3owmDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owJqqkMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYmMXEQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoYUHlGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f253kTcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0s8MWQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4nfF4wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3fLrlcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXxOrB1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISNUsP2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N7eviqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BF4ROXAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qt5l6ZQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jfMzzC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BmrvCwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AE3S1MJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLF6vX5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTqiXCVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxfdgKTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wtm796QRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sH57cZrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g12C4NqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtBGvxZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79D2g9NzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clJBZcpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnUFtYCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkUmku8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2bchRSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht5wMAvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75s45QAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UMJqNodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtI6kSgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbXaFhdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMm9g0nbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZJEQnfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5A3jkFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 810xXomLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6ivGUUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1wkCK11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogF4z3MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYQOnWTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTcrrvAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKRd3S1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCYLXG9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Cy54i93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxTgkdgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yjr9JWBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVLSRZyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wXzsnOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tv9csZjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2m39vXBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1TZbXGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYmHOBzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qR5fokJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

