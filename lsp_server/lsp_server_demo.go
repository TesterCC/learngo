package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// ref: https://aistudio.google.com/prompts/12wQ2vjicDFhZkctE3xuO3CeZO3S72LNQ
// 1. 定义 Semgrep/OpenGrep JSON 结构 (简化版)
type OpenGrepOutput struct {
	Results []struct {
		Path    string `json:"path"`
		CheckID string `json:"check_id"`
		Start   struct {
			Line int `json:"line"`
			Col  int `json:"col"`
		} `json:"start"`
		End struct {
			Line int `json:"line"`
			Col  int `json:"col"`
		} `json:"end"`
		Extra struct {
			Message  string `json:"message"`
			Severity string `json:"severity"`
		} `json:"extra"`
	} `json:"results"`
}

// 2. 定义 LSP 结构 (简化版)
type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity"` // 1: Error, 2: Warning
	Code     string `json:"code"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type PublishDiagnosticsParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// JSON-RPC 基础消息
type Request struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  json.RawMessage  `json:"params"`
	ID      *json.RawMessage `json:"id,omitempty"`
}

func main() {
	// 设置日志，输出到 stderr，以免污染 stdout (LSP 通信通道)
	log.SetOutput(os.Stderr)
	log.Println("OpenGrep LSP Server Started...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(splitRPC) // 自定义分割函数处理 Content-Length

	for scanner.Scan() {
		msg := scanner.Bytes()
		handleMessage(msg)
	}
}

// 简单处理 JSON-RPC 消息
func handleMessage(msg []byte) {
	var req Request
	if err := json.Unmarshal(msg, &req); err != nil {
		return
	}

	// 只处理 didSave 事件 (MVP)
	if req.Method == "textDocument/didSave" {
		var params struct {
			TextDocument struct {
				URI string `json:"uri"`
			} `json:"textDocument"`
		}
		json.Unmarshal(req.Params, &params)

		// 移除 file:// 前缀获取本地路径
		path := strings.TrimPrefix(params.TextDocument.URI, "file://")
		// Windows 兼容性处理略...

		runScan(path, params.TextDocument.URI)
	}
}

func runScan(filePath string, fileURI string) {
	log.Printf("Scanning file: %s", filePath)

	// 调用 CLI
	// 假设 opengrep_cli 在 PATH 中
	cmd := exec.Command("opengrep_cli", "scan", "--json", filePath)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Scan error: %v", err)
		return
	}

	var grepOut OpenGrepOutput
	if err := json.Unmarshal(output, &grepOut); err != nil {
		log.Printf("JSON parse error: %v", err)
		return
	}

	diagnostics := []Diagnostic{}
	for _, res := range grepOut.Results {
		// Semgrep 行号从 1 开始，LSP 从 0 开始
		diag := Diagnostic{
			Source:  "OpenGrep",
			Code:    res.CheckID,
			Message: res.Extra.Message,
			Range: Range{
				Start: Position{Line: res.Start.Line - 1, Character: res.Start.Col - 1},
				End:   Position{Line: res.End.Line - 1, Character: res.End.Col - 1},
			},
		}

		// 映射 Severity
		if res.Extra.Severity == "ERROR" {
			diag.Severity = 1
		} else {
			diag.Severity = 2 // Warning
		}
		diagnostics = append(diagnostics, diag)
	}

	// 发送 publishDiagnostics 通知
	notification := struct {
		JSONRPC string                   `json:"jsonrpc"`
		Method  string                   `json:"method"`
		Params  PublishDiagnosticsParams `json:"params"`
	}{
		JSONRPC: "2.0",
		Method:  "textDocument/publishDiagnostics",
		Params: PublishDiagnosticsParams{
			URI:         fileURI,
			Diagnostics: diagnostics,
		},
	}

	sendJson(notification)
}

func sendJson(v interface{}) {
	bytes, _ := json.Marshal(v)
	// LSP 协议头: Content-Length: ...\r\n\r\nJSON
	fmt.Printf("Content-Length: %d\r\n\r\n%s", len(bytes), bytes)
}

// splitRPC 是一个 bufio.SplitFunc，用于解析 LSP 头部 (Header + Body)
// 实际开发建议直接用 go-lsp 库处理底层协议
func splitRPC(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 省略底层协议解析代码，实际项目中请使用库
	// 这里的占位符意味着你需要引入像 "github.com/sourcegraph/jsonrpc2" 这样的库来处理流
	return 0, nil, nil
}
