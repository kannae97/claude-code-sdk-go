# Advanced Example

This example demonstrates advanced features of the Claude Code SDK for Go, including MCP integration, session management, and comprehensive tool configuration.

## Features Demonstrated

- **MCP Configuration**: Model Context Protocol server integration
- **Session Management**: Resume and continue conversations
- **Tool Restrictions**: Advanced AllowedTools and DisallowedTools usage
- **Custom System Prompts**: Both override and append system prompts
- **Output Format Comparison**: Testing different output formats
- **Permission Management**: MCP permission prompt tools
- **Verbose Logging**: Detailed operation logging

## Usage

```bash
go run main.go
```

**Note**: Some features require additional setup (MCP servers, etc.) and are commented out to avoid errors in the demo.

## What it does

1. **MCP Configuration Example**: Shows how to configure MCP servers (demonstration only)
2. **Session Management**: Demonstrates resume/continue functionality
3. **Tool Restrictions**: Real analysis with restricted tool set
4. **Output Format Testing**: Compares text, JSON, and stream-JSON formats

## Advanced Features

### MCP Integration
```go
Options: &claudecode.Options{
    MCPConfig:           stringPtr("mcp-servers.json"),
    AllowedTools:        []string{"mcp__filesystem__read_file"},
    PermissionPromptTool: stringPtr("mcp__permissions__approve"),
}
```

### Session Management
```go
Options: &claudecode.Options{
    Resume:   stringPtr("session-id"),     // Resume specific session
    Continue: boolPtr(true),               // Continue latest session
}
```

### Tool Restrictions
```go
Options: &claudecode.Options{
    AllowedTools:    []string{"Read", "LS", "Grep"},
    DisallowedTools: []string{"Write", "Bash"},
}
```

### Custom System Prompts
```go
Options: &claudecode.Options{
    SystemPrompt:       stringPtr("You are a senior Go architect..."),
    AppendSystemPrompt: stringPtr("Always provide specific recommendations."),
}
```

## Output

The example will:
1. Show MCP configuration examples (commented out)
2. Demonstrate session management concepts
3. Perform live project analysis with tool restrictions
4. Compare different output formats
5. Display real-time streaming analysis with detailed tool usage

## Requirements for Full Functionality

To use all features, you would need:

1. **MCP Server Configuration** (`mcp-servers.json`):
```json
{
  "mcpServers": {
    "filesystem": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "/path"]
    }
  }
}
```

2. **Previous Session IDs**: From actual Claude Code sessions

3. **Permission Tools**: MCP permission management tools

## Performance

The example includes real-time streaming analysis which may take several minutes to complete as it performs comprehensive project analysis.