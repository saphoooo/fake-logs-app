package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	format := flag.String("o", "custom", "log format syle (nginx|custom)")
	frequency := flag.Int("f", 5, "log frequency in seconds")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	switch *format {
	case "nginx":
		// exemaple output: '10.0.1.182 - - [25/Nov/2021:08:54:53 +0000] "GET / HTTP/1.1" 200 615 "-" "curl/7.74.0" "-"'
		for {

			fmt.Printf("172.17.0.3 - - [%v] \"%s %s HTTP/1.1\" %s %s \"http://example.com/\" \"%s\" \"%s\"\n", time.Now().Format("2006/01/02:15:04:05 +0000"), randMethod(), randPath(), randReturnCode(), strconv.Itoa(rand.Intn(999-200)+200), randClient(), randAddress())
			time.Sleep(time.Second * time.Duration(*frequency))
		}
	case "custom":
		// exemaple output: '[INFO] 2021/11/25 15:12:49 172.17.0.4 POST http://example.com/ 401 -- 218.214.93.28 Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0'
		for {
			log.New(os.Stdout, randPrefix(), log.Ldate|log.Ltime).Printf("172.17.0.4 %s http://example.com%s %s -- %s %s", randMethod(), randPath(), randReturnCode(), randAddress(), randClient())
			time.Sleep(time.Second * time.Duration(*frequency))
		}
	default:
		log.Fatal("format must be custom or nginx")
	}
}

func randPrefix() string {
	prefix := []string{"[EMERG] ", "[WARN] ", "[ERROR] ", "[ERROR] ", "[TRACE] ", "[INFO] ", "[INFO] ", "[INFO] ", "[INFO] ", "[DEBUG] "}
	n := rand.Intn(10)
	return prefix[n]
}

func randClient() string {
	client := []string{"curl/7.74.0", "curl/7.64.1", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0"}
	n := rand.Intn(10)
	return client[n]
}
func randReturnCode() string {
	methods := []string{"200", "200", "200", "200", "200", "200", "301", "401", "404", "500"}
	n := rand.Intn(10)
	return methods[n]
}

func randMethod() string {
	methods := []string{"GET", "GET", "GET", "GET", "POST", "POST", "POST", "PUT", "DELETE", "PATCH"}
	n := rand.Intn(10)
	return methods[n]
}

func randPath() string {
	path := []string{"/", "/", "/", "/api", "/api", "/status", "/api/items", "/api/checkout", "/login", "/logout"}
	n := rand.Intn(10)
	return path[n]
}

func randAddress() string {
	firstDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	secondDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	thirdDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	fourthDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	return firstDigit + "." + secondDigit + "." + thirdDigit + "." + fourthDigit
}
