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
	go func() {
		for {
			f, err := ioutil.ReadFile("/var/log/secure")
			if err != nil {
				fmt.Println("Error to read file")
				return
			}
			str := string(f)
			r, err := regexp.Compile("sudo")
			if err != nil {
				fmt.Println("Error to compile regex")
				return
			}
			c := len(r.FindAllString(str, -1))

			//con, err := strconv.ParseInt(c, 10, 64)
			fmt.Println(c)
			// Incrementart
			//fmt.Println(len(c))
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	c = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_sudo_count",
		Help: "The total number of processed sudo events",
	})
)

func main() {
	getsudo()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
