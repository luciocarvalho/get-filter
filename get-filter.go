package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getsudo() {

}

var (
	//sudoCount = promauto.NewCounter(prometheus.CounterOpts{
	sudoCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "get_sudo_count",
		Help: "The total number of processed sudo events",
		//	Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	})
)

func main() {
	//getsudo()
	go func() {
		for {
			sudoCount.Set(0)
			f, err := ioutil.ReadFile("/var/log/secure")
			//f, err := os.OpenFile("/var/log/secure", os.O_RDONLY, os.ModePerm)
			if err != nil {
				fmt.Println("Error to read file")
				return
			}
			//defer f.Close()
			str := string(f)
			//str := bufio.NewScanner(f)

			/*for str.Scan() {
				_ = str.Text()
				fmt.Println("teste")
			}*/
			r, err := regexp.Compile("sudo")
			if err != nil {
				fmt.Println("Error to compile regex")
				return
			}
			//fmt.Println(str)
			//c := len(r.FindAllString(str, -1))
			c := len(r.FindAllStringSubmatchIndex(str, -1))

			flow := float64(c)

			sudoCount.Add(flow)
			//sudoCount.SetToCurrentTime()

			//con, err := strconv.ParseInt(c, 10, 64)
			//			fmt.Println(c)
			// Incrementart
			//fmt.Println(len(c))
			time.Sleep(2 * time.Second)
		}

	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
