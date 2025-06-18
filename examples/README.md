# Claude Code SDK Go Examples

This directory contains comprehensive example programs demonstrating the usage of the Claude Code SDK for Go, from basic usage to advanced features.

## Directory Structure

```
examples/
├── README.md          # This file
├── go.mod            # Go module for examples
├── quick_start.go    # Quick start example (Python SDK compatible format)
├── basic/            # Basic usage examples
│   ├── README.md     # Detailed basic example documentation
│   └── main.go       # Basic API usage with both new and traditional APIs
├── streaming/        # Real-time streaming examples
│   ├── README.md     # Detailed streaming documentation
│   └── main.go       # Streaming API with real-time message processing
└── advanced/         # Advanced features and configuration
    ├── README.md     # Detailed advanced features documentation
    └── main.go       # MCP, sessions, tool restrictions, custom prompts
```

## Examples Overview

### ⚡ [Quick Start](./quick_start.go)
**Perfect first example (Python SDK compatible format)**
- Three focused examples: basic, with options, with tools
- Simple structure matching Python SDK examples
- Clean, minimal code for learning

```bash
go run quick_start.go
```

### 🚀 [Basic Examples](./basic/)
**Basic usage demonstration**
- TypeScript/Python SDK compatible API (`QueryWithRequest`)
- File creation with tools
- Error handling

```bash
cd basic && go run main.go
```

### 📡 [Streaming Examples](./streaming/)
**Real-time message processing**
- Live streaming with `QueryStreamWithRequest` and `QueryStream`
- Channel-based async processing
- Tool restrictions and verbose logging
- Real-time project analysis

```bash
cd streaming && go run main.go
```

### ⚡ [Advanced Examples](./advanced/)
**Full feature demonstration**
- MCP (Model Context Protocol) integration
- Session management (resume/continue)
- Advanced tool configurations
- Custom system prompts
- Output format comparisons

```bash
cd advanced && go run main.go
```


## Quick Start

1. **Install Prerequisites**:
```bash
# Install Go (1.21+)
# Install Claude Code CLI
npm install -g @anthropic-ai/claude-code

# Verify Claude CLI is working
claude --help
```

2. **Run Quick Start Example**:
```bash
go run quick_start.go
```

3. **Try Streaming**:
```bash
cd streaming  
go run main.go
```

## API Compatibility

These examples demonstrate both API styles:

### New API (TypeScript/Python Compatible)
```go
request := claudecode.QueryRequest{
    Prompt: "Create a function",
    Options: &claudecode.Options{...},
}
messages, err := claudecode.QueryWithRequest(ctx, request)
```

### Traditional API (Backward Compatible)
```go
messages, err := claudecode.Query(ctx, prompt, options)
```

## Features Demonstrated

- ✅ **Full CLI Option Support**: All Claude Code CLI options
- ✅ **Multiple Output Formats**: text, json, stream-json
- ✅ **Session Management**: Resume and continue conversations
- ✅ **Tool Configuration**: Allow/disallow specific tools
- ✅ **MCP Integration**: Model Context Protocol support
- ✅ **Custom Prompts**: System prompts and extensions
- ✅ **Error Handling**: Comprehensive error management
- ✅ **Streaming**: Real-time message processing

## Prerequisites

- **Go 1.21 or later**
- **Claude Code CLI**: `npm install -g @anthropic-ai/claude-code`
- **Authenticated Claude CLI**: Run `claude` to verify authentication
- **Internet connection** for Claude API access

## Troubleshooting

### Common Issues

1. **"claude command not found"**
   ```bash
   npm install -g @anthropic-ai/claude-code
   ```

2. **Authentication errors**
   ```bash
   claude  # Follow authentication prompts
   ```

3. **Go module issues**
   ```bash
   go mod tidy
   ```

## Learning Path

1. **Start with [Quick Start](./quick_start.go)** - Learn fundamental concepts
2. **Try [Basic Examples](./basic/)** - Understand detailed usage
3. **Explore [Streaming Examples](./streaming/)** - Real-time processing  
4. **Master [Advanced Examples](./advanced/)** - All features

## Contributing

When adding new examples:
1. Create a new directory with descriptive name
2. Include a detailed README.md
3. Add entry to this main README.md
4. Ensure examples are well-commented
5. Test with `go run` and verify output