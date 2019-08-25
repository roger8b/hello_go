package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorig = 3
const monitoringDelay = 5

func main() {
	showIntroduction()
	for {
		showMenu()
		menu(readOptions())
	}
}

func menu(option int) {
	switch option {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exibingo logs ...")
		printLog()
	case 0:
		fmt.Println("Saindo do programa ...")
		os.Exit(0)
	default:
		fmt.Println("Opção invalida ...")
		os.Exit(-1)
	}
}

func readOptions() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("A opção escolhida foi", option)

	return option
}

func showIntroduction() {
	name, age := getNameAndAge()
	version := 1.1
	fmt.Println("Ola senhor,", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
}

func getNameAndAge() (string, int) {
	name := "Roger"
	age := 33
	return name, age
}

func showMenu() {
	fmt.Println("Escolha uma opção abaixo")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func startMonitoring() {
	fmt.Println("Monitorando ...")

	sites := readSitesFromFile()

	for i := 0; i < monitorig; i++ {
		for ii, site := range sites {
			fmt.Println(ii, "- Testando site", site)
			checkSite(site)
		}
		time.Sleep(monitoringDelay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func checkSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro na requisição:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println(" O site", site, "foi carregado com sucesso")
		logRegister(site, true)
	} else {
		fmt.Println(" O site", site, "não foi carregado erro:", response.StatusCode)
		logRegister(site, false)
	}

	fmt.Println("")
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	logError(err)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		sites = append(sites, line)
		logError(err)
	}
	file.Close()
	return sites
}

func logError(err error) {
	if err != nil {
		fmt.Println("Erro:", err)
	}
}

func logRegister(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		"- online:" + strconv.FormatBool(status) + "\n")

	logError(err)

	file.Close()
}

func printLog() {
	file, err := ioutil.ReadFile("log.txt")

	fmt.Println(string(file))

	logError(err)
}
