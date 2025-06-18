# Streaming Example

This example demonstrates real-time streaming capabilities of the Claude Code SDK for Go, showing how to process messages as they arrive from Claude Code.

## Features Demonstrated

- **QueryStreamWithRequest**: TypeScript/Python SDK compatible streaming API
- **QueryStream**: Traditional Go streaming API
- **Real-time processing**: Messages processed as they arrive
- **Channel-based communication**: Go-native async patterns
- **Tool restrictions**: AllowedTools and DisallowedTools
- **Verbose logging**: Detailed streaming output

## Usage

```bash
go run main.go
```

## What it does

1. **Example 1**: Uses `QueryStreamWithRequest` to create a README.md file with streaming output
2. **Example 2**: Uses traditional `QueryStream` API to analyze project structure

## Key Features

- **True Streaming**: Messages are processed in real-time as they arrive
- **Tool Configuration**: Shows how to allow/disallow specific tools
- **System Prompts**: Custom and appended system prompts
- **Output Formats**: Stream-JSON format for real-time processing
- **Error Handling**: Proper error handling in streaming context

## Output

The example will:
1. Stream messages showing file creation and editing process
2. Display tool usage in real-time (LS, Read, Write, etc.)
3. Show project structure analysis with live updates
4. Provide message count and performance statistics

## Streaming Flow

```
System Message (init) → Assistant Messages → Tool Use → Tool Results → Final Result
```

Each message type is displayed with appropriate emoji indicators:
- 📝 Text content
- 🔧 Tool usage
- ✅ Tool results
- ❌ Error results