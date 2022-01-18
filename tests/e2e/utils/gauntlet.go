package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/rs/zerolog/log"
)

type Gauntlet struct {
	exec string
}

type ExecError struct {
	WhatIsIt     string
	HowToRespond string
}

type FlowReport []struct {
	Name string `json:"name"`
	Txs  []struct {
		Contract string `json:"contract"`
		Hash     string `json:"hash"`
		Success  bool   `json:"success"`
	}
	Data   map[string]string `json:"data"`
	StepId int               `json:"stepId"`
}

type GReport struct {
	Responses []struct {
		Tx struct {
			Hash    string `json:"hash"`
			Address string `json:"address"`
		}
		Contract string `json:"contract"`
	} `json:"responses"`
	Data map[string]string `json:"data"`
}

func NewGauntlet(binPath string) (Gauntlet, error) {
	fmt.Println(binPath)
	log.Debug().Str("PATH", binPath).Msg("BinPath")
	_, err := exec.Command(binPath).Output()
	if err != nil {
		return Gauntlet{}, errors.New("gauntlet installation check failed")
	}
	return Gauntlet{
		exec: binPath,
	}, nil
}

func (g Gauntlet) Flag(flag string, value string) string {
	return fmt.Sprintf("--%s=%s", flag, value)
}

func (g Gauntlet) ExecCommand(args []string, errHandling []ExecError) (string, error) {
	output := ""
	cmd := exec.Command(g.exec, args...)
	stdout, _ := cmd.StdoutPipe()
	// errPipe, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return output, err
	}
	stdin, _ := cmd.StdinPipe()
	// go func() {
	reader := bufio.NewReader(stdout)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		output += line + "\n"
		line, err = reader.ReadString('\n')
		rerr := respondToErrors(errHandling, line, stdin)
		if rerr != nil {
			return output, rerr
		}
	}
	// }()

	if strings.Compare("EOF", err.Error()) > 0 {
		return output, err
	}
	return output, nil
}

func respondToErrors(errHandling []ExecError, line string, stdin io.WriteCloser) error {
	for _, e := range errHandling {
		if strings.Contains(line, e.WhatIsIt) {
			log.Debug().Str("Error", line).Msg("Gauntlet Error Found")
			// _, err := stdin.Write([]byte(fmt.Sprintln(e.HowToRespond)))
			// if err != nil {
			// 	return err
			// }
			return fmt.Errorf("found a gauntlet error")
		}
	}
	return nil
}

func (g Gauntlet) ReadCommandReport() (GReport, error) {
	jsonFile, err := os.Open("report.json")
	if err != nil {
		return GReport{}, err
	}

	var report GReport
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &report)

	return report, nil
}

func (g Gauntlet) ExecuteAndRead(args []string, errHandling []ExecError) (GReport, error) {
	_, err := g.ExecCommand(args, errHandling)
	if err != nil {
		return GReport{}, err
	}

	report, err := g.ReadCommandReport()
	if err != nil {
		return GReport{}, err
	}
	return report, nil
}

func (g Gauntlet) ReadCommandFlowReport() (FlowReport, error) {
	jsonFile, err := os.Open("flow-report.json")
	if err != nil {
		return FlowReport{}, err
	}

	var report FlowReport
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return FlowReport{}, err
	}
	err = json.Unmarshal(byteValue, &report)
	if err != nil {
		return FlowReport{}, err
	}

	return report, nil
}
