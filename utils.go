package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Liza-Developer/apiGO"
)

func init() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	Acc.LoadState()
	animIntro()

	_, err := os.Open("accounts.txt")
	if os.IsNotExist(err) {
		os.Create("accounts.txt")
	}

	if Acc.DiscordID == "" {
		fmt.Print("Enter a Discord ID: ")
		fmt.Scan(&Acc.DiscordID)

		Acc.SaveConfig()
		Acc.LoadState()

		fmt.Println()
	}
}

func isGC(bearer string) string {
	conn, _ := tls.Dial("tcp", "api.minecraftservices.com"+":443", nil)

	fmt.Fprintln(conn, "GET /minecraft/profile/namechange HTTP/1.1\r\nHost: api.minecraftservices.com\r\nUser-Agent: Dismal/1.0\r\nAuthorization: Bearer "+bearer+"\r\n\r\n")

	e := make([]byte, 12)
	conn.Read(e)

	switch string(e[9:12]) {
	case `404`:
		return "Giftcard"
	default:
		return "Microsoft"
	}
}

func animIntro() {
	//First Row
	fmt.Print("╭━━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━╮╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮╭╮")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╭━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━━┳")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━━━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮\n")
	//Second Row
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃╭━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┣╯")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰┫┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╱╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱┃╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━╮┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╭━╮")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃\n")
	time.Sleep(10 * time.Millisecond)
	//Third Row
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃┃╱")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃┣╮")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╭┫╰")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┳━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┳━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮╭━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┫┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╱╰┫")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃╱┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("\n")
	//Fourth Row
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃╰━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╯┃┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃┃╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┃┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┫╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┫╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┃┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╭━┫")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃╱┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("\n")
	time.Sleep(10 * time.Millisecond)
	//Fifth Row
	fmt.Print("┃╭━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┃┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰┫┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃┃┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┫┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃┃╭")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╮┃╰")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┻━┃")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰━╯")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("┃\n")
	time.Sleep(10 * time.Millisecond)
	//Sixth Row
	fmt.Print("╰╯ ")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰╯╰")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┻╯")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰┻━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━┻╯")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰┻╯")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╰┻━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━━┻")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("━━━")
	time.Sleep(10 * time.Millisecond)
	fmt.Print("╯\nVER - 2.1.5 / Beta\n\n")
}

func formatTime(t time.Time) string {
	return t.Format("05.00000")
}

func (Account Details) check(name, searches, AccType string) {
	var details checkDetails
	body, _ := json.Marshal(Data{Name: name, Bearer: Account.Bearer, Id: Acc.DiscordID, Unix: Account.UnixRecv, Config: string(jsonValue(embeds{Content: "<@" + Acc.DiscordID + ">", Embeds: []embed{{Description: fmt.Sprintf("[%v] Succesfully sniped %v with %v searches :bow_and_arrow:", AccType, name, searches), Color: 770000, Footer: footer{Text: "MCSN"}, Time: time.Now().Format(time.RFC3339)}}}))})

	req, _ := http.NewRequest("POST", "http://droptime.site/api/v2/webhook", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &details)

	if details.Error != "" {
		fmt.Printf("[ERROR] %v\n", details.Error)
	} else if details.Sent != "" {
		fmt.Printf("[200] %v\n", details.Sent)
	}

	removeDetails(Account)
}

func removeDetails(Account Details) {
	var new []apiGO.Bearers
	for _, Accs := range Acc.Bearers {
		if Account.Email != Accs.Email {
			new = append(new, Accs)
		}
	}

	Acc.Bearers = new

	var meow []apiGO.Info
	for _, Accs := range Acc.Bearers {
		for _, Acc := range Bearers.Details {
			if Acc.Email != Accs.Email {
				meow = append(meow, Acc)
			}
		}
	}

	Bearers.Details = meow

	var Accz []string
	file, _ := os.Open("accounts.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Split(scanner.Text(), ":")[0] != Account.Email {
			Accz = append(Accz, scanner.Text())
		}
	}

	rewrite("accounts.txt", strings.Join(Accz, "\n"))

	Acc.Logs = append(Acc.Logs, apiGO.Logs{
		Email:   Account.Email,
		Send:    Account.SentAt,
		Recv:    Account.RecvAt,
		Success: Account.Success,
	})

	Acc.SaveConfig()
	Acc.LoadState()
}

func jsonValue(f interface{}) []byte {
	g, _ := json.Marshal(f)
	return g
}
