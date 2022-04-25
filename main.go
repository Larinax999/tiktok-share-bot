package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	shared int
	failed int

	itemid string
)

func main() {
	//clear()
	fmt.Println("[*] Hello open src by larina")
	fmt.Print("[!] id > ")
	fmt.Scanln(&itemid)
	fmt.Println("[*] recommend 500-900 thread if over 1000 thread will fail than sent")
	var threads int
	fmt.Print("[!] thread > ")
	fmt.Scanln(&threads)
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}

	// go background()
	for i := 0; i < threads; i++ {
		wg.Add(1)
		//fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] Successfully started \x1b[38;5;63m%.0f\x1b[0m threads.", time.Now().Format("15:04:05"), float64(i+1))
		go share()
	}
	for {
		fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] sent : %d fail : %d", time.Now().Format("15:04:05"), shared, failed)
		time.Sleep(2e7)
	}
	//wg.Wait()
}

func random_device_id() string {
	device := ""
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 15+rand.Intn(4); i++ {
		num := rand.Intn(9-1) + 9
		device += strconv.Itoa(num)
	}

	return device
}

func share() {
	packet := []byte(fmt.Sprintf("item_id=%s&share_delta=1", itemid))

	for {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		url := fmt.Sprintf("https://api16-core-useast5.us.tiktokv.com/aweme/v1/aweme/stats/?device_id=%s&channel=googleplay&aid=1988&app_name=trill&version_code=210605&item_id=%s&share_delta=1", random_device_id(), itemid)
		req.SetRequestURI(url)
		req.Header.SetMethod("POST")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.SetBodyRaw(packet)

		fasthttp.Do(req, resp)

		if strings.Contains(string(resp.Body()), "status_code\":0") {
			shared += 1
		} else {
			//fmt.Println(string(resp.Body()))
			failed += 1
		}
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}
}
