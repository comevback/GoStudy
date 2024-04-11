package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StartDocker() {
	// 获取操作系统类型
	osName := getOSName()

	// 测试是否存在sudo命令
	sudoCommand := ""
	if _, err := exec.LookPath("sudo"); err == nil {
		sudoCommand = "sudo"
	}

	// 获取本地IP地址和公共IP地址
	localIPAddress, _ := getIPAddresses(osName)

	// 默认API和Central Server URL
	defaultAPIURL := fmt.Sprintf("http://%s:4000", localIPAddress)
	defaultCentralServerURL := fmt.Sprintf("http://%s:8000", localIPAddress)

	// 用户输入API和Central Server URL
	apiURL := getUserInput("Please enter your API host URL (default: "+defaultAPIURL+"): ", defaultAPIURL)
	centralServerURL := getUserInput("Please enter your central register server URL (default: "+defaultCentralServerURL+"): ", defaultCentralServerURL)

	// 用户输入React端口号
	reactPort := getUserInput("Please enter your React PORT (default: 3000): ", "3000")

	// 清除docker中的悬空镜像
	if output, err := exec.Command(sudoCommand, "docker", "images", "-f", "dangling=true", "-q").Output(); err == nil {
		danglingImages := strings.TrimSpace(string(output))
		if len(danglingImages) > 0 {
			exec.Command(sudoCommand, "docker", "rmi", danglingImages).Run()
		}
	}

	// 启动docker容器
	imageName := "afterlifexx/upc-system:latest"
	dockerCommand := fmt.Sprintf("%s docker run -e API_URL=%s -e API_PORT=4000 -e CENTRAL_SERVER=%s -e REGI_PORT=8000 -e REACT_APP_INITIAL_API_URL=%s -e REACT_APP_INITIAL_CENTRAL_SERVER_URL=%s -e PORT=%s -v /var/run/docker.sock:/var/run/docker.sock -p 4000:4000 -p %s:%s -p 8000:8000 -it --rm %s",
		sudoCommand, apiURL, centralServerURL, apiURL, centralServerURL, reactPort, reactPort, reactPort, imageName)
	exec.Command("sh", "-c", dockerCommand).Run()
}

func getOSName() string {
	output, err := exec.Command("uname", "-s").Output()
	if err != nil {
		fmt.Println("Error getting OS name:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

func getIPAddresses(osName string) (string, string) {
	var localIPAddress, publicIPAddress string
	switch osName {
	case "Linux":
		output, err := exec.Command("hostname", "-I").Output()
		if err != nil {
			fmt.Println("Error getting local IP address:", err)
		}
		localIPAddress = strings.Fields(string(output))[0]
	case "Darwin":
		output, err := exec.Command("ifconfig").Output()
		if err != nil {
			fmt.Println("Error getting local IP address:", err)
		}
		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "inet ") && !strings.Contains(line, "127.0.0.1") {
				localIPAddress = strings.Fields(line)[1]
				break
			}
		}
	}
	publicIPAddressOutput, err := exec.Command("curl", "-s", "https://api.ipify.org").Output()
	if err != nil {
		fmt.Println("Error getting public IP address:", err)
	}
	publicIPAddress = strings.TrimSpace(string(publicIPAddressOutput))
	return localIPAddress, publicIPAddress
}

func getUserInput(prompt, defaultValue string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}
