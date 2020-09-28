package main

import (
	"fmt"
)

var tcpPorts []int = []int{8000, 8001, 8002}
var httpPorts []int = []int{9001, 9002, 9003}
var grpcPorts []int = []int{7000, 7002, 7003}

func main() {

	error := checkForDuplicatePorts()
	if len(error) == 0 {
		fmt.Println("Good to go ")
	} else {
		fmt.Println(error)
	}

}

func checkForDuplicatePorts() []string {
	strSlc := make([]string, 0)
	strSlc = append(strSlc, checkItself(httpPorts, "httpPorts")...)
	strSlc = append(strSlc, checkItself(grpcPorts, "grpcPorts")...)
	strSlc = append(strSlc, checkItself(tcpPorts, "tcpPorts")...)

	for i := 0; i < len(httpPorts); i++ {
		for j := 0; j < len(tcpPorts); j++ {
			if httpPorts[i] == tcpPorts[j] {
				s := fmt.Sprintf("Duplicate Port http:%d and tcp:%d", httpPorts[i], tcpPorts[j])
				strSlc = append(strSlc, s)
			}
			if httpPorts[i] == grpcPorts[j] {
				s := fmt.Sprintf("Duplicate Port http:%d and grpc:%d", httpPorts[i], grpcPorts[j])
				strSlc = append(strSlc, s)
			}
		}

	}
	for i := 0; i < len(tcpPorts); i++ {
		for j := 0; j < len(grpcPorts); j++ {
			if tcpPorts[i] == grpcPorts[j] {
				s := fmt.Sprintf("Duplicate Port tcp:%d and grpc:%d", tcpPorts[i], grpcPorts[j])
				strSlc = append(strSlc, s)
			}
		}
	}

	return strSlc
}

func checkItself(arr []int, portType string) []string {
	slc := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] == arr[j] {
				slc = append(slc, fmt.Sprintf("Duplicate port in %s", portType))
			}
		}
	}
	return slc
}
