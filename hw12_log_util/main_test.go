package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLogFile(t *testing.T) {
	wantLogs := [][]string{
		{"error", "a file system error has occurred"},
		{"info", "a USB device is connected"},
		{"info", "a USB device is disconnected"},
		{"warning", "a suspicious file has been found"},
		{"warning", "a suspicious file has been found"},
		{"warning", "a suspicious file has been found"},
	}
	gotlogs, err := ReadLogFile("log_test.csv")
	require.NoError(t, err)
	assert.Equal(t, wantLogs, gotlogs)
}

func TestAnalizeTheLog(t *testing.T) {
	logs := [][]string{
		{"error", "a file system error has occurred"},
		{"info", "a USB device is connected"},
		{"info", "a USB device is disconnected"},
		{"warning", "a suspicious file has been found"},
		{"warning", "a suspicious file has been found"},
		{"warning", "a suspicious file has been found"},
	}

	wantLogs := []string{
		"a file system error has occurred",
	}
	gotLogs := AnalizeTheLog(logs, "error")
	assert.Equal(t, wantLogs, gotLogs)
}

func TestAnalizeTheLogEmpty(t *testing.T) {
	var wantLogs []string
	logs := [][]string{}

	gotLogs := AnalizeTheLog(logs, "error")
	assert.Equal(t, wantLogs, gotLogs)
}

func TestWriteOutput(t *testing.T) {
	fileName := "output_test.txt"
	logs := []string{
		"a file system error has occurred",
	}
	WriteOutput(logs, fileName)

	file, err := os.Open(fileName)
	require.NoError(t, err, "Can't open output file \"%s\": %v", fileName, err)
	defer file.Close()

	b, err := io.ReadAll(file)
	require.NoError(t, err, "Can't read output file \"%s\": %v", fileName, err)

	lines := bytes.Split(b, []byte("\n"))
	for i, wantLog := range logs {
		gotLog := string(lines[i])
		assert.Equal(t, wantLog, gotLog)
	}
}

/*func TestWriteOutputFileEmpty(t *testing.T) {
	var gotLogs []string
	fileName := ""
	logs := []string{
		"a file system error has occurred",
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		reader := bufio.NewReader(os.Stdout)
		fmt.Println("wait 3 sec")
		time.Sleep(time.Second * 3)
		for {
			line, err := reader.ReadString('\n')
			fmt.Println(line)
			if err == io.EOF {
				fmt.Println("break, line = ", line)
				break
			} else {
				require.NoError(t, err, "Can't read os.Stdout: %v", err)
				gotLogs = append(gotLogs, line)
				fmt.Println("line = ", line)
				fmt.Printf("gotLogs = %#v", gotLogs)
			}
		}
		fmt.Println("wg.Done()")
		wg.Done()
	}()

	fmt.Println("wait 1 sec")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("start WriteOutputFile")
	WriteOutputFile(logs, fileName)

	fmt.Println("wg.Wait()")
	wg.Wait()

	assert.Equal(t, logs, gotLogs)
}*/

/*func TestWriteOutputFileEmpty(t *testing.T) {
	var gotLogs []string
	fileName := ""
	logs := []string{
		"a file system error has occurred",
	}

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {

	//	wg.Done()
	//}()

	//fmt.Println("wait 1 sec")
	//time.Sleep(time.Millisecond * 1000)
	fmt.Println("start WriteOutputFile")
	WriteOutputFile(logs, fileName)

	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("wait 2 sec")
	time.Sleep(time.Second * 2)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			fmt.Println("break, line = ", line)
			break
		} else {
			require.NoError(t, err, "Can't read os.Stdout: %v", err)
			gotLogs = append(gotLogs, line)
			fmt.Println("line = ", line)
			fmt.Printf("gotLogs = %#v", gotLogs)
		}
	}
	//fmt.Println("wg.Done()")

	//fmt.Println("wg.Wait()")
	//wg.Wait()

	assert.Equal(t, logs, gotLogs)
}*/
