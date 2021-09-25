package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-ping/ping"
)

func pingo(address string) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		panic(err)
	}
	pinger.Count = 4
	pinger.SetPrivileged(true)
	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Benvenuto in Pingo")
	var address string
	fmt.Println("Inserisci il dominio da pingare")
	fmt.Scan(&address)
	pingo(address)
	pingo("www." + address)
	var subdomain string
	var resp string
	for {
		fmt.Println("Vuoi provare altri sottodomini? (s/n)")
		fmt.Scan(&resp)
		if resp == "s" {
			fmt.Println("Quale sottodominio?")
			fmt.Scan(&subdomain)
			pingo(subdomain + "." + address)
		} else if resp == "n" {
			fmt.Println("Uscita in corso")
			break
		} else {
			fmt.Println("Per favore, scrivi la risposta nel modo corretto:\ns se sì\nn se no")
		}
	}
}
