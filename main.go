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

var (
	prefix  = []string{"[EMERG] ", "[WARN] ", "[ERROR] ", "[ERROR] ", "[TRACE] ", "[INFO] ", "[INFO] ", "[INFO] ", "[INFO] ", "[DEBUG] "}
	agent   = []string{"curl/7.74.0", "curl/7.64.1", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Safari/605.1.15", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0"}
	status  = []string{"200", "200", "200", "200", "200", "200", "301", "401", "404", "500"}
	methods = []string{"GET", "GET", "GET", "GET", "POST", "POST", "POST", "PUT", "DELETE", "PATCH"}
	path    = []string{"/", "/", "/", "/api", "/api", "/status", "/api/items", "/api/checkout", "/login", "/logout"}
)

func main() {
	format := flag.String("f", "custom", "log format syle (nginx|custom)")
	interval := flag.Int("i", 5, "log interval in seconds")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	switch *format {
	case "nginx":
		// exemaple output: '10.0.1.182 - - [25/Nov/2021:08:54:53 +0000] "GET / HTTP/1.1" 200 615 "-" "curl/7.74.0" "-"'
		for {

			fmt.Printf("172.17.0.3 - - [%v] \"%s %s HTTP/1.1\" %s %s \"http://example.com/\" \"%s\" \"%s\"\n", time.Now().Format("02/Jan/2006:15:04:05 +0000"), randomize(methods), randomize(path), randomize(status), strconv.Itoa(rand.Intn(999-200)+200), randomize(agent), randAddress())
			time.Sleep(time.Second * time.Duration(*interval))
		}
	case "custom":
		// exemaple output: '[INFO] 2021/11/25 15:12:49 172.17.0.4 POST http://example.com/ 401 -- 218.214.93.28 Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:94.0) Gecko/20100101 Firefox/94.0'
		for {
			log.New(os.Stdout, randomize(prefix), log.Ldate|log.Ltime).Printf("172.17.0.4 %s http://example.com%s %s -- %s %s", randomize(methods), randomize(path), randomize(status), randAddress(), randomize(agent))
			time.Sleep(time.Second * time.Duration(*interval))
		}
	default:
		log.Fatal("format must be custom or nginx")
	}
}

func randomize(elements []string) string {
	n := rand.Intn(len(elements))
	return elements[n]
}

func randAddress() string {
	firstDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	secondDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	thirdDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	fourthDigit := strconv.Itoa(rand.Intn(254-2) + 2)
	return firstDigit + "." + secondDigit + "." + thirdDigit + "." + fourthDigit
}
