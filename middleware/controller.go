package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.3" )

func biZIH7U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8BoETcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RoQvTpanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hghy2BuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFt7Zu1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtSnGo1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCB9FRDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLnYkL0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6MQJwaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwBMejF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEPO8iJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diT5ThpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyOiVHVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILLJcAAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZT44ew3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcoPQum0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzkFt6ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y25kPG7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F38erGrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtyhDuk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blmfP6RKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTz3Vc1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjIM9G1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJEmtbrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdfpjYPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95diXTGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzdGC57ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YVvHK7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sK2J1mnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrefBFD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ST2LM4LvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKWmKNK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DuB8PFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYkKPRUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIks1udjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKiDovCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7oBfoIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwKIqRCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxo01NPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43ydazN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5VroholWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCEWQIFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaA5FXDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6JCA81cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jcz8VLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yukAnpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PuR2lD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMNZpBrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVVnNdXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JEZHm7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny8RX452Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiDDX5xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYpgakf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vj2UATdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WuuMeibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5MwMPsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VIAXtO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twQqB07IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sCPzdlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfGpN9KpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pmz0i2O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lj2e5IatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxiumiRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tcTJiHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCUmJ1skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFEPZVypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oZCFF47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPBLbbaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k1roFcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92ynTHM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3FQBrqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHn4COvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dimVuSuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIJJkyABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Or8EOTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pj2JoGccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGFtnHqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzHSPZv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FuhxdEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfFWfAJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMOSjnjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmLs0xvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jF1p5j3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3a3NhSJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQECwjDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6IEwcutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYvhgKGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6Av1hSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vudkpkXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jD2Hlp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8X4X5nuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkabOLmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFXxifQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6q5B1QVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YjDKTwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0S7GPAApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGqmyPNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AB7aLdRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lE4haSmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8inyVuGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKXz6SLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74qqOMraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUbT7DF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAXXLbffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dY6CiUusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtBidKXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbQmCjOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFC2w2d4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yb5hXUvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLOEbrfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bvibhOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FjvpAf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASeFswM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlIzsOnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1XhM3oeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQERxY2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0S1TgGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YM5PiXwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08jsEPxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2yFNgUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHTLY0YYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3imXs6AsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByVZGmrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2562lXLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXxfoU6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31iG0ORtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fCQ5pJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDgRTGujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFH77dz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhUn4XgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woA2OYeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wMlxK2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20rY3ytUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3ODKfstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdMp771cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESM8L9NiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhP2Lv29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jw2C0kMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gO7RKWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0DiwnVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udD8K1UBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kV6CKKEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGpMLBwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARYc9fQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7v3orIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aytuqLQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 228UCpOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvLBJrtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrtVD7auWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FHYDAe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qZF7eGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrIh1tkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXkdW8CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2X6epjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbHojS4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4YgojI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xk4dh6bVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTLqztBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBa68z2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9TrHmiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqbnE09bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xu7yiU0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgomWPgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DaKJ4RFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Bw7wKj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiVEep1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4dr9CvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKVo0SpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8taVTzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9MgHkTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csnf35RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbBTaGFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juirL6zfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VOr2SKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh7wdia6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3BNOiK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJ8tHD7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHow0MnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEhsum7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OB21ItGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZ8Iriy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnF6Iwm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctyE8lDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbvOgw1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNvxakgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRtyJ5Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujloUbvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeZF5u3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQnxpMaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pY67rycTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0txUqbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zA7smBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD149hSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nx743m2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNd4I3k5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egDVPikNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Txa6XYufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uoFxlS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cH3zhkeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ym2KxQHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIwHiAp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8grY62GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hh3XP62YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txgzVv1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9KEYz0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rYREfxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7PTdczmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZGEey0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA4O5n7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jA48Y2x7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWFScRiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEXhz11MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9KckaQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjB2ffBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swahdoGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acV1wKIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func so5A6xH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZCcRx1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nuWPTVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G43TdZcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROvMnSEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RrlKW5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1I8WRYnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6Aqk7trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU9kndWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zNwmzXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hJqHe2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5BNrkp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQnzyeozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoRveTrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WraFKlv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kf4yzPWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Op2bXaqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVihd2SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 376rTKLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPdCG95oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiF9yXy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRVV9O48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7jIktuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiCh9sxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tK3GNi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nbOAU28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYEvRFYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkcsrrVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOYbhbUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S121hamVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhYHTTWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldwLN7D5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3c85TXEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HodEEsDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEMdEquDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVoFL6q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwRuOMR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwbkJzDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdWvwFecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSg1tjFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3HQdKczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C32vqgFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvMP1yFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMNO7ZvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiEUVPGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxhGaUseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEpBnNKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Duxy4ZihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRY9XNksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fK2iVy7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRKuDp1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgOwI7bvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5sKf5ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjucrudeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9AuVYTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Kf3ql4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHl8ddSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AIAnV3asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mk5UmuHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivfJhxb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K29brLwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gU9lFoxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ad1hCx4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A92JbNSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8y4qK6WPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8523NvBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9AGVZInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hytZsKWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BsjplrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9q3fqIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35fObX1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwtStFjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKNZFSwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0F1PXOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MafJuBr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38Q7YU7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfmjuLSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ah9sjpiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Qt1lBQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAVF5dedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWlzm0FqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCFI5NdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLWevaGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nefqVG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnW9ebUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luX0URZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRav1ybJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B55XmvG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23OQWmCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj85AZTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxwIXW6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUSWxue6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2kGN8EfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M66SghyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgD880DrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJk1WqFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgbvrCFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eabx4zKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RT4SyvonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osjMwqKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y03FPZx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzTidU5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNSXiPcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8feUcaaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGAfq3U7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKbBxm08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gU9oAz0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlkuicsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0AM6kxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNIrMFacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8DkuOurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Em3ORqozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKS69O7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEan2Q3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OP5Wy26TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guPDTCD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X87eY3f9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdUzOV3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiAvsK0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xntUdyqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lzxpt4xQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dz7HFYvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9Ot6qqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJ9OiXqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRfWDFDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OySQv1UiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t83tPmUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUW3ianOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXXBiS3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u86LbFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXrzJYoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEs4BkyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FX77w4O9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFbz8XJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfQZPh4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qfmyb2w7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QqOD8uGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YygIMhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pgko7dBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDRg49B7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FouDly2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvkhN6HDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2XYnW3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGciUFyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eap8uaTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3pBsYuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQktnODLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUrQC66oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIOP5hvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9dLv4hDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKAizpmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faNBYGsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Zt0DNgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zbeCNpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxRNdOuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IULTI1nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbkEBLu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiUTybA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdWnXGPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyY6sQZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuyoEPcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxaLcKa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdGbTijNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPI41IucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9enTcp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TR3Ai4klWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Moj0oOiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L40LvCYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AX6H5z4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vq7QROjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJr69jWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STwXg8H7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bug2XtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWo3hR2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWGBh192Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BM6aCdF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUixN4QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5lTu7HnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGFrC0vpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3PwHDUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5tF5wiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1N1VPStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXHNmiS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7DKm07zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdUnK1JxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7ZIkjnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fMnISEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfwXhtDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQpK0dDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bddLx14NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZFSoT0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOcUnN58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zeol7CEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pInfsSWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjXoY7f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4uNzAl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEU0K8jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tra7RWprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJvUBE5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzCOk9JPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRpharQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGgkj5c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4OAkttGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNI58UoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvdfXr2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STMlMCC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DHWmcMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqD3NtmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ThNcrHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAQLYd5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drTBjR8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKbThIgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STUPDal8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3h2IE0MvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdv700e8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T96Bg9HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BzWVif8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfyqxIz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9gbMaARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPb6pS9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qujkZwthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21NDsHeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFMCMKtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgsuklo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyVrFWpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GclLknF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAFGlsJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnMNoophWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqXId5biWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUB8R5JiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jK3UQfbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWL5MXGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHIjMNJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr2v5vnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JDqyKvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NfDGlPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8fXdDPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zF6TRhgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7QkNFs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCHU35QwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3m55J4zpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdsL5eUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41aeQtFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ckkIBE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAPhzvowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2AqN9vcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJpxX7hoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChgaTjGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bt1i6c3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CADPzlG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aS8aDvGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8s3D6m8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m255tnvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzmqJ9w3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q582c5asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjhKdcRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsyDKZRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hq5kgMiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIaZXFvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRVBEqsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvkZk9NoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3I9y3DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVin9varWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlA6cVsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zysKBb1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ui5WYN1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFvzVnqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8hxVksmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqJ8OF66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXyr3pqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBpErNNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4TDZybUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0E8ALzopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37FfMZY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCQievIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIyJqfW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bvHoIJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7n5RzDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5fUO62wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBHwg0h9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMvJmIZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19rAffxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htLwCaPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRb5ftWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9tWXFS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCMpzPYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkeajOaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvGU05UEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXKJBxtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NigzdKdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMGbvoT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4x7Y5BP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjFxbyHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaWA6nfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXByyzMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKtJUbAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1FzzzitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewN6qzgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiTs7eA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPk3Jb07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISneqZl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GT9PyLl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DX3jkGtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4XydIgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiXhiT5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LIzurOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39hM87MMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Kw3C4TBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yq1S4hgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDnzM0LcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3vJrjiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMTpTiRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZS4hmN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPh5Ix5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9J7AxF2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZPQN94iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NK9sAaMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CMXEbfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3l1nuKXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlXnvdrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1O2mPfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3tj28WcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wr5MwxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amflU05BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Htg2ZUTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfBS785KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ke00OsDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlgFYNd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2Ai6YAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58GRxXJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBOLWpueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wePmW5A0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OysagqoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhctOYi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qv5c7rvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ya3XkgBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 022oyzDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfrGfYJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fj3Ub8lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJY5G8kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ykaLiwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eRqhqHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEDGZ18PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJBHoS3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DDCbOLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4Wuh328Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qz7M5BR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0z3jclTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VnTbL4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUvLuluCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7RtMNJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOfSzsiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjfY1eonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dInifbyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYGH24gYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swStjMlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9klhFzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHsnNXgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8nNlYkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ds545PcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfmmitAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wki6xzX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzAjSkUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIDGeGhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhKt5lCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3O9hhXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g12KoeKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOm8zzD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8PHCQjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baSB2DEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrBCcZz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgG2C0GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieT7ylarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3mc35IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbDNG7loWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FxF0xy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WvSPWmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BubnUkKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tK26pbRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajSObbQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Td8nxxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xpxCEkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YopeWuD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLGvwbkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oS4Yv4YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqFEbor0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHjeRjBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRbtBJkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1BpKDzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMqaDBIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShGgoRBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUq5v9XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMAM3Y3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pl65GLX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQvirPKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hx7lRud8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GatknrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwPQ5GUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewxislsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbAlqyXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbaDjlKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Z6XfMwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJQd3trrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9573q0oOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4L7sjtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGukWvbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODij4yb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9W25IBVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4Uz1g6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThchAKY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvAPE1JEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtdGvV2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gr9zmykZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etO9RsA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jG1KO1gbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RkTUmlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkzfneN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2lmrr8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIv4C5LvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kiq7FAZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xOs1fPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnc25BV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfqcIP4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Evx2aH5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpOLICiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sr3u9YagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnPuzP0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r3EzCIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysTmVoktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNuRp2WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWNGAVArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XT3f1FuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlrkmWaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKjx3N57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuc6lG1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr3m4deqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dRU48ccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiWECAkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufKVyHAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1VnfGUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZoDfoQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FozC2SRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E37V0d9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtYLaGu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgc1snj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTi6ebatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZJUWOhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWGug7RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uCEbzsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqSrseilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwOtxLWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5hjPuYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nGu7eNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1b7EmRYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEactIwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JI0cs49LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HOjVRZymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLIze9aEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7JWimu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MANlQxmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zt2phO8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFreO21uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnUIAtEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIEIWAVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ftkl1DfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7Lb7rQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ze0CQyyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41IwAI62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QWweaeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRWifQvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwF6PTwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyO57zHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUSItFwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAeG9PXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5vSnlnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiEnUFVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vd7OJ1tKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZFidZ0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yq3efRpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jd6XjbQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVhFJjvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4aTmOqhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjLFGFF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBiPZnBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TBgUmrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmYcDSkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vtxC6XMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEG4evEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sC807Z6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7ba0uzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpzqFdp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPKY0spfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9z1bNLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GweqI3IjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wvTrwtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGg0oZhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGNP01gBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoVeMF11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbt8Q0BYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0wSCG5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVOstJ6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IofvfwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDg5rPoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2vCfCYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErZ5aHTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8QeoU2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0VfHFXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1ECV7AJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2x8Q7U9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCUt3FOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evPZCC7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKATVjDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjVxaJHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YnYFeC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7XWzb0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4FqW8Q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0K39QzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inwjGWmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qBYUxcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTAVTuJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiFU2u5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDElElvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMlGheK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTa7PAIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEfeaPKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LC57adVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEmcaivzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTGYVaRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8o1cAfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TK2uv2owWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vyeu9UwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ivfw7ZxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4lMSvnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6HaJsP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrTzkAz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCEZDTM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACqAaWtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeyZPh8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdLG0KXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkTDvPfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Gnc2lGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dfhyWoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmUvHHBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ax1NF6YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEiYPBwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STZPwYgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2II9ApPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLQGl3HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fccnZurcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGi7JZUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQTnH2PmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmViGoSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwjEt5KRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vhxsTKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P32LJ8d7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcLx8hetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4v5RRz93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeswtQOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func motaAxqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhdlvQFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6F5hcAxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0LxLyIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5t894By2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLzjTQ2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zR2dqYC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B20wdAluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TA1Zdbf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsp4ue1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDrI07djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5QuPIyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3d5PaC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9OntsTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func je1P2QuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2H22YXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w49gqvGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5NdueUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLQcBolLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jmeIGWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28Y45ZeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmcKv48WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xl4wf85qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vLReTuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg93MivHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNMAdZxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u92lQwjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l1s9TgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ea2TdIbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DziJkF4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8Acj6nsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x60A9eYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qw6Q6EsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogpxeibfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4J5NjTZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9uEPyxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oetX8IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68pObD4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EhAnmBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReeBqyxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfmsF3aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBBrRvVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FVgY0gyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNsmbp2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMMnEWh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDfW5WKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnw3fb0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSNgLE1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZj9lwJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIl87PiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbQqp4kIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eS1XSZDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFm9J5eKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0rdrHotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tx854PMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWKKl4EpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vJZXPiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOuJF1kDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phKwjEfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8526V9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jik7FPcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mp2QkBtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viWR2OQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhudkOMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FF99ivwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JHNufyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLGAj78fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tT47xDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2dYT4YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKFpudOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Am4nuyZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cE5YEB96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oI0KnxkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Gtb3rSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4I7H4An1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQn9EapBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG9ufaL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nh2zsFdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iiUtheTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b37NMSNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrVsbR7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKHdA6SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jESUAYYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ctbirf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3X0qcazNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCQjdKjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22bqX5rXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSsOy6brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJscN0QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCHGjEL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UIVW7s63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A44p4JOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyawqYEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3EpVLLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qs4ILi6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEdaziItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TUa1R4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIKzzYhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGyxp4E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtVRxdS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeBDWjJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJpQ76KsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AH00QihhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUkmQRp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BA7mponBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8L5urVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59omrVemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPzuEJfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgQyzg3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APQpJTfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaXDNWn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmGdnGpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rh67KBFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYIoAukqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4WMaUrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJhkC5NUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kvaL1plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ENHoO4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSHUc5m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H76v9F6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7uDTTcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgpJ67PBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZSwDXsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArmHbH3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUFXZLvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuIXxdLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vj5z9tPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvfo59YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KWlPevZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7pHH6irWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIHhRNm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcBuSt8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdow9gQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQX889uHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWfRMy4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlUoJyS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohOQyYeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuCT9GAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IXBr5HhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNTViNyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ybg47J9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeJ0PBaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcygeSGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKSansndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mdcO3OVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYWiVtG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jvitfr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCTrfpmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTvCdUkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2V6GYIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdJNYIJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJ5s5FeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9BCGhK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTPfKAIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qH84qdgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bK77ELEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tu9GVVvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxcslPATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVc9UHlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dcp5HiaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imxof6jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdcyJGISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tH5nCN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nySZDBzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4jAtCpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTTxUe0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AGIyw1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYjcdEaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Vy14tvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMadYCLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3cWsWvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dbvp5Z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1yez42iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zHATXrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNrF6UEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9ZAhbVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSpRCZaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6KAKpQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gd9fjiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmJmgC98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vE4UyDjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BplXMuTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdPFZMZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMZ7Tc53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnRRI5ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6C6uXnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pTE604BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Rmkk73jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwIITKVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXasDiNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iafwze9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMQy3eNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G53tfZKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyYwz7fvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QMTcecmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFzwnjC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGajVotBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pARKHrwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wnIwNuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGWokBZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6osguaUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DT8mKwMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2J9TyXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhkWL8unWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 254wl6JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akiF1iJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqxlKrphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0jVvI56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8a3KfXpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfjSYOicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Guf5kT89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpCzYSgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 799kY3AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXF1eYv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjAiLX3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSN7T64RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46zAKkHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1lmQ1K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ts6KzV06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dyDsh0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AqvPO75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngvcASV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPZMiLtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jP4v847VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEBx2MnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTLO3wxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EPAuYPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTofqtseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ymkpw5auWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umAWHDNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJ2ePHX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAbV8rCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDt9VGzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9X7fCCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xrk3Uo3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lOYPadvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3h44aVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBUvwnW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCJVqmhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHTiuXkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71c2w43tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SC3yNO35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOVlVeivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXuq2WFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WE8JSvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dRzAN45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVPWo7GIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZdiUDlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayuTymM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPaagh5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ChvTiuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svkjlk4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4KEHYMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QaYwP2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16kDPv4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl3T8xExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ugswfo45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5Cz1WPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoSMDESuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRTSxCXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBm8mqWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GEE9QAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSNtrPg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEBnYuPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kh3xTzVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUGJrNxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8CBYwbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujyneFecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuihXSZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJcuna3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFDxqOBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SOmaGS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXzU6jhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXk2QGDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func povCWlaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

