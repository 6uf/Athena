package main

import (
	"crypto/tls"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/Liza-Developer/apiGO"
)

func main() {

	AuthAccs()

	var name string
	var dropDelay int64

	var data SentRequests

	fmt.Println("[INFO] Name to Snipe:")
	fmt.Print(">: ")
	fmt.Scanln(&name)
	fmt.Println("[INFO] Delay:")
	fmt.Print(">: ")
	fmt.Scanln(&dropDelay)

	drop := apiGO.DropTime(name)

	for time.Now().Before(time.Unix(drop, 0).Add(-time.Second * 5)) {
		fmt.Printf("Generating Payloads/TLS Connection In: %v      \r", time.Until(time.Unix(drop, 0).Add(-time.Second*5)).Round(time.Second).Seconds())
		time.Sleep(time.Second * 1)
	}

	payload := Bearers.CreatePayloads(name)
	conn, _ := tls.Dial("tcp", "api.minecraftservices.com:443", nil)
	var wg sync.WaitGroup

	fmt.Print("\n[+] Dropping at @ ", time.Unix(drop, 0), "\n\n")

	time.Sleep(time.Until(time.Unix(drop, 0)) + time.Duration(dropDelay)*time.Millisecond)

	for req, acc := range Bearers.Details {
		for i := 0; i < acc.Requests; i++ {
			wg.Add(1)
			go func(i int, Account apiGO.Info) {
				SendTime, recvTime, Status := payload.SocketSending(conn, payload.Payload[i])

				data.Requests = append(data.Requests, Details{
					Bearer:     Account.Bearer,
					SentAt:     SendTime,
					RecvAt:     recvTime,
					StatusCode: Status,
					Success:    Status == "200",
					UnixRecv:   recvTime.Unix(),
					Email:      Account.Email,
					Type:       Account.AccountType,
				})

				wg.Done()
			}(req, acc)
		}
	}

	wg.Wait()

	sort.Slice(data.Requests, func(i, j int) bool {
		return data.Requests[i].SentAt.Before(data.Requests[j].SentAt)
	})

	searches := apiGO.Search(name)

	for _, request := range data.Requests {
		if request.Success {
			fmt.Printf("[+] Sent @ %v ~> [%v] @ %v ~ %v\n", formatTime(request.SentAt), request.StatusCode, formatTime(request.RecvAt), request.Email)

			fmt.Println()

			if Acc.ChangeskinOnSnipe {
				SendInfo := apiGO.ServerInfo{
					SkinUrl: Acc.ChangeSkinLink,
				}

				resp, _ := SendInfo.ChangeSkin(jsonValue(skinUrls{Url: SendInfo.SkinUrl, Varient: "slim"}), request.Bearer)
				if resp.StatusCode == 200 {
					fmt.Printf("[%v] Succesfully Changed your Skin!\n", resp.StatusCode)
				} else {
					fmt.Print("[ERROR] Couldnt Change your Skin..\n")
				}
			}

			request.check(name, searches.Searches, request.Type)
			fmt.Println()
		} else {
			fmt.Printf("[-] Sent @ %v ~> [%v] @ %v ~ %v\n", formatTime(request.SentAt), request.StatusCode, formatTime(request.RecvAt), request.Email)
		}
	}
}
