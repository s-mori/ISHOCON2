package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var host = "http://127.0.0.1"
var totalScore = 0
var totalResp = map[bool]int{}
var finished = false

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: ./benchmark [option]
Options:
  --workload	N	run benchmark with N workloads (default: 3)
  --target	HOST	specify target HOST (default: 127.0.0.1:80)`)
	}

	var (
		workload = flag.Int("workload", 3, "")
		target   = flag.String("host", "127.0.0.1", "")
	)
	flag.Parse()
	host = "http://" + *target

	startBenchmark(*workload)
}

func startBenchmark(workload int) {
	getInitialize()
	log.Print("期日前投票を開始します")
	validateInitialize()
	log.Print("期日前投票が終了しました")
	log.Print("投票を開始します  Workload: " + strconv.Itoa(workload))
	voteTime := time.Now().Add(45 * time.Second)
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < workload+1; i++ {
		wg.Add(1)
		if i%5 == 0 {
			go loopInvalidVoteScenario(wg, m, voteTime)
		} else {
			go loopVoteScenario(wg, m, voteTime)
		}
	}
	wg.Wait()
	log.Print("投票が終了しました")
	finishTime := time.Now().Add(15 * time.Second)
	log.Print("投票者が結果を確認しています")

	for i := 0; i < workload+2; i++ {
		wg.Add(1)
		if i%4 == 0 || i%4 == 3 {
			go loopIndexScenario(wg, m, finishTime)
		} else if i%4 == 1 {
			go loopCandidateScenario(wg, m, finishTime)
		} else {
			go loopPoliticalPartyScenario(wg, m, finishTime)
		}
	}
	wg.Wait()
	printScore()
}

func loopInvalidVoteScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	finished := false
	for {
		if finished {
			break
		} else {
			finished = invalidVoteScenario(wg, m, finishTime)
		}
	}
}

func loopVoteScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	finished := false
	for {
		if finished {
			break
		} else {
			finished = voteScenario(wg, m, finishTime)
		}
	}
}

func loopIndexScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	finished := false
	for {
		if finished {
			break
		} else {
			finished = indexScenario(wg, m, finishTime)
		}
	}
}

func loopCandidateScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	finished := false
	for {
		if finished {
			break
		} else {
			finished = candidateScenario(wg, m, finishTime)
		}
	}
}

func loopPoliticalPartyScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	finished := false
	for {
		if finished {
			break
		} else {
			finished = politicalPartyScenario(wg, m, finishTime)
		}
	}
}

func printScore() {
	log.Print("投票者の感心がなくなりました")
	log.Print("{\"score\": " + strconv.Itoa(totalScore) + ", \"success\": " + strconv.Itoa(totalResp[true]) + ", \"failure\": " + strconv.Itoa(totalResp[false]) + "}")
}
