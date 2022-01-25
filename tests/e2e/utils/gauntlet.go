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
	"github.com/smartcontractkit/helmenv/environment"
)

type Gauntlet struct {
	exec string
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

func (g Gauntlet) ExecCommand(args []string, errHandling []string) (string, error) {
	output := ""
	printArgs(args)
	cmd := exec.Command(g.exec, args...)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return output, err
	}
	stdin, _ := cmd.StdinPipe()
	reader := bufio.NewReader(stdout)
	line, err := reader.ReadString('\n')
	for err == nil {
		log.Info().Str("stdout", line).Msg("Gauntlet")
		output += line + "\n"
		line, err = reader.ReadString('\n')
		rerr := respondToErrors(errHandling, line, stdin)
		if rerr != nil {
			return output, rerr
		}
	}

	if strings.Compare("EOF", err.Error()) > 0 {
		return output, err
	}
	return output, nil
}

func respondToErrors(errHandling []string, line string, stdin io.WriteCloser) error {
	for _, e := range errHandling {
		if strings.Contains(line, e) {
			log.Debug().Str("Error", line).Msg("Gauntlet Error Found")
			return fmt.Errorf("found a gauntlet error")
		}
	}
	return nil
}

func printArgs(args []string) {
	out := "gauntlet"
	for _, arg := range args {
		out = fmt.Sprintf("%s %s", out, arg)

	}
	log.Info().Str("Command", fmt.Sprintf("%s\n", out)).Msg("Gauntlet")
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

func (g Gauntlet) ExecuteAndRead(args []string, errHandling []string, retryCount int) (GReport, string, error) {
	var output string
	var report GReport
	var err error
	for i := 0; i < retryCount; i++ {
		log.Info().Msg(fmt.Sprintf("Gauntlet Command Attempt: %v", i+1))
		output, err = g.ExecCommand(args, errHandling)
		if err != nil {
			continue
			// return GReport{}, output, err
		}

		report, err = g.ReadCommandReport()
		if err != nil {
			continue
			// return GReport{}, output, err
		}
	}
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Failed to exucute Gauntlet command after %v attempts", retryCount))
	}
	return report, output, err
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

// WriteNetworkConfigMap write a network config file for gauntlet testing
func WriteNetworkConfigMap(file string, config map[string]string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for k, v := range config {
		fmt.Printf("key[%s] value[%s]\n", k, v)
		_, err = f.WriteString(fmt.Sprintf("\n%s=%s", k, v))
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDefaultGauntletConfig gets and writes the default config gauntlet will need to start making commands
// 	against the environment
func GetDefaultGauntletConfig(network string, env *environment.Environment) (map[string]string, error) {
	solUrls, err := env.Charts.Connections("solana-validator").LocalURLsByPort("http-rpc", environment.HTTP)
	if err != nil {
		return nil, err
	}

	networkConfig := map[string]string{
		"NETWORK":                      "local",
		"NODE_URL":                     solUrls[0].String(),
		"PROGRAM_ID_OCR2":              "CF13pnKGJ1WJZeEgVAtFdUi4MMndXm9hneiHs8azUaZt",
		"PROGRAM_ID_ACCESS_CONTROLLER": "2F5NEkMnCRkmahEAcQfTQcZv1xtGgrWFfjENtTwHLuKg",
		"PROGRAM_ID_STORE":             "A7Jh2nb1hZHwqEofm4N8SXbKTj82rx7KUfjParQXUyMQ",
		"PRIVATE_KEY":                  "[82,252,248,116,175,84,117,250,95,209,157,226,79,186,119,203,91,102,11,93,237,3,147,113,49,205,35,71,74,208,225,183,24,204,237,135,197,153,100,220,237,111,190,58,211,186,148,129,219,173,188,168,137,129,84,192,188,250,111,167,151,43,111,109]",
		"SECRET":                       "[only,unfair,fiction,favorite,sudden,strategy,rotate,announce,rebuild,keep,violin,nuclear]",
	}

	err = WriteNetworkConfigMap(fmt.Sprintf("networks/.env.%s", network), networkConfig)
	if err != nil {
		return nil, err
	}

	return networkConfig, nil
}
