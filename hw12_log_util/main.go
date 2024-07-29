package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

type settings struct {
	fileName, level, outputFile string
}

func main() {
	logsSettings := getSettings()

	logs, err := ReadLogFile(logsSettings.fileName)
	if err != nil {
		log.Fatal()
	}
	foundLogs := AnalizeTheLog(logs, logsSettings.level)
	WriteOutput(foundLogs, logsSettings.outputFile)
}

func getSettings() settings {
	var fileName, level, outputFile string
	pflag.StringVarP(&fileName, "file", "f", "", "указывает путь к анализируемому лог-файлу")
	pflag.StringVarP(&level, "level", "l", "", "указывает уровень логов для анализа")
	pflag.StringVarP(&outputFile, "output", "o", "", "указывает путь к файлу, в который будет записана статистика")
	configPath := pflag.String("config", ".env", "path to config file")
	pflag.Parse()

	fileNameEnv := os.Getenv("LOG_ANALYZER_FILE")
	levelEnv := os.Getenv("LOG_ANALYZER_LEVEL")
	outputEnv := os.Getenv("LOG_ANALYZER_OUTPUT")

	err := godotenv.Load(*configPath)
	if err != nil {
		log.Printf("Can't parse config: %v \n", err)
	}

	fileName = checkFlagAndPullFromConfigFile(fileName, fileNameEnv, "LOG_ANALYZER_FILE")
	level = checkFlagAndPullFromConfigFile(level, levelEnv, "LOG_ANALYZER_LEVEL")
	outputFile = checkFlagAndPullFromConfigFile(outputFile, outputEnv, "LOG_ANALYZER_OUTPUT")

	logsSettings := settings{
		fileName:   fileName,
		level:      level,
		outputFile: outputFile,
	}
	return logsSettings
}

func checkFlagAndPullFromConfigFile(flag string, variableEnv string, kindVarEnv string) string {
	if flag == "" && variableEnv == "" {
		variableConfEnv, ok := os.LookupEnv(kindVarEnv)
		if ok {
			flag = variableConfEnv
		} else if kindVarEnv == "LOG_ANALYZER_FILE" {
			log.Fatalf("The name of the log file is not specified. It must be specified using the -f (or --file) flag,"+
				" either in the environment variable %s or in the \".env\" configuration file\n", kindVarEnv)
		}
	} else if flag == "" {
		flag = variableEnv
	}
	return flag
}

func ReadLogFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Can't open file \"%s\": %v", fileName, err)
		return nil, err
	}
	defer file.Close()

	fileReader := csv.NewReader(file)
	lines, err := fileReader.ReadAll()
	if err != nil {
		log.Printf("Can't read file \"%s\": %v", fileName, err)
		return nil, err
	}

	return lines, nil
}

func AnalizeTheLog(logs [][]string, level string) []string {
	var foundLogs []string
	for _, log := range logs {
		if log[0] == level {
			for j, elem := range log {
				if j == 0 && elem != level {
					break
				} else if j > 0 {
					foundLogs = append(foundLogs, elem)
				}
			}
		}
	}
	return foundLogs
}

func WriteOutput(foundLogs []string, outputFile string) {
	var (
		file *os.File
		w    *bufio.Writer
		err  error
	)

	if outputFile == "" {
		file = os.Stdout
	} else {
		file, err = os.Create(outputFile)
		if err != nil {
			log.Printf("Can't create output file \"%s\": %v", outputFile, err)
		}
		defer file.Close()
	}

	w = bufio.NewWriter(file)
	for _, log := range foundLogs {
		w.Write([]byte(fmt.Sprintf("%s\n", log)))
	}
	w.Flush()
}
