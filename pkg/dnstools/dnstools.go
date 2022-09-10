package dnstools

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"time"

	screen "github.com/aditya43/clear-shell-screen-golang"
	p "github.com/dariubs/percent"
	t "github.com/jedib0t/go-pretty/table"
)

func Request(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		return nil, err
	}

	return ips, nil
}

func Run(domain string, total int) {
	requests := 0
	errors := 0
	answers := make(map[string]int)
	start := time.Now()

	for i := 0; i < total; i++ {

		ips, err := Request(domain)
		if err != nil {
			errors++
		}
		if len(ips) > 0 {
			ip := ips[0].String()
			answers[ip]++
			requests++
		} else {
			requests++
		}

		table := t.NewWriter()
		table.SetOutputMirror(os.Stdout)
		table.Style().Options.SeparateRows = true
		table.AppendHeader(t.Row{"IP Address", "Count", "Percent of Reqs. Made"})

		keys := make([]string, 0, len(answers))
		for k := range answers {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for key := range keys {
			ip := keys[key]
			table.AppendRow(t.Row{ip, answers[ip], strconv.Itoa(int(float64(answers[ip])/float64(requests)*100)) + "%"})
		}

		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("DNS Check shows how often DNS answers appear for a given domain")
		fmt.Println("")
		fmt.Println("Domain: " + domain)
		fmt.Println("Request Status: " + strconv.Itoa(requests) + " / " + strconv.Itoa(total) + " (" + strconv.FormatFloat(p.PercentOf(requests, total), 'f', 2, 64) + "%)")
		table.Render()

	}

	end := time.Now()
	time := end.Sub(start)
	fmt.Println("\nFinished in " + strconv.FormatFloat(time.Seconds(), 'f', 2, 64) + " Seconds with " + strconv.Itoa(errors) + " Errors")
}
