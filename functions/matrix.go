package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Go_matrix(w http.ResponseWriter, r *http.Request) {

	buf2, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}
	cpuinfo := string(buf2)

	buf3, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	meminfo := string(buf3)

	buf4, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatal(err)
	}
	uptime := string(buf4)

	n := 100

	if r.URL.Query().Get("n") != "" {
		n, err = strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		n = 100
	}

	start := time.Now()
	matrix(n)
	elapsed := time.Since(start)

	msg := &Message{
		Success: true,
		Payload: Payload{
			Test: "matrix test",
			N:    n,
			Time: strconv.FormatInt(int64(elapsed/time.Millisecond), 10),
		},
		Metrics: Metrics{
			MachineId:  "",
			InstanceId: "",
			Cpu:        cpuinfo,
			Mem:        meminfo,
			Uptime:     uptime,
		},
	}

	b, err := json.Marshal(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprint(w, string(b))
}

func randomTable(n int) [][]int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = r1.Intn(100)
		}
	}
	return m
}

func matrix(n int) [][]int {
	matrixA := randomTable(n)
	matrixB := randomTable(n)
	matrixMult := make([][]int, n)

	for i := 0; i < len(matrixA); i++ {
		matrixMult[i] = make([]int, n)
		for j := 0; j < len(matrixB); j++ {
			sum := 0
			for k := 0; k < len(matrixA); k++ {
				sum += matrixA[i][k] * matrixB[k][j]
			}
			matrixMult[i][j] = sum
		}
	}

	return matrixMult
}
