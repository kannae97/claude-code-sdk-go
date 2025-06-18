# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go SDK for Claude Code that provides programmatic access to Claude's agentic coding capabilities. The SDK wraps the Claude Code CLI and provides a Go-native interface compatible with the TypeScript and Python SDKs.

## Architecture

The SDK follows a layered architecture:

- **Core Types** (`types.go`): Defines interfaces (`Message`, `ContentBlock`) and structs (`Options`, `QueryRequest`) that mirror the TypeScript/Python SDK APIs. Includes `OutputFormat` constants and all CLI option fields.
- **CLI Communication** (`claude.go`): Contains dual API functions:
  - New API: `QueryWithRequest` and `QueryStreamWithRequest` (TypeScript/Python compatible)
  - Traditional API: `Query` and `QueryStream` (backward compatible)
  - Helper functions: `buildCommandArgs`, `readTextOutput`, `readMessages`, `findCLIExecutable`
- **Error Handling** (`errors.go`): Custom error types for different failure scenarios (CLI not found, connection errors, process errors, JSON decode errors)

## Key Implementation Details

### CLI Integration Strategy
The SDK communicates with Claude Code by:
1. Setting `CLAUDE_CODE_ENTRYPOINT=sdk-go` environment variable to identify SDK usage
2. Finding the Claude CLI executable (`claude` command, tries PATH and npm global paths)  
3. Building command arguments from `Options` struct using `buildCommandArgs()`
4. Executing with `--print` and dynamically determined flags (including automatic `--verbose` for stream-json)
5. Handling multiple output formats: text (plain text), json (structured), stream-json (real-time)
6. Parsing responses via format-specific readers (`readTextOutput` vs `readMessages`)

### Dual API Design
- **New API**: `QueryRequest{Prompt, Options}` pattern matching TypeScript/Python SDKs
- **Traditional API**: `Query(ctx, prompt, options)` for Go-style backward compatibility
- **Streaming**: Real-time message processing via Go channels, with true streaming for `stream-json` format

### Advanced Features Support
- **Session Management**: Resume/Continue via session IDs
- **MCP Integration**: Model Context Protocol server configuration and tool permissions
- **Tool Restrictions**: AllowedTools/DisallowedTools with proper CLI argument formatting
- **Model Selection**: Model parameter for choosing specific Claude variants
- **Debug Mode**: Debug flag for verbose CLI output
- **Permission Control**: PermissionMode (hidden CLI option) and DangerouslySkipPermissions
- **Directory Access**: AddDir for specifying additional allowed directories
- **Input Format**: Support for different input formats (text, stream-json)
- **Output Format Handling**: Automatic verbose flag injection for stream-json, different parsing strategies per format

## Development Commands

```bash
# Run all tests
go test -v

# Run tests with coverage
go test -v -cover

# Run specific test
go test -v -run TestOptions

# Run only unit tests (skip integration tests)
go test -v -short

# Build the module
go build

# Format and vet
go fmt ./...
go vet ./...

# Lint with golangci-lint (recommended for Go Report Card compliance)
golangci-lint run

# Install golangci-lint (if not already installed)
# macOS: brew install golangci-lint
# Linux: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# Windows: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Test examples (requires authenticated claude CLI)
cd examples/basic && go run main.go
cd examples/streaming && go run main.go
cd examples/advanced && go run main.go

# Verify Claude CLI availability
claude --help
```

## Prerequisites for Development

- Go 1.21 or later
- Claude CLI (`npm install -g @anthropic-ai/claude-code`) for integration tests
- Claude CLI must be authenticated for integration tests (run `claude` to verify)
- golangci-lint for comprehensive linting (optional but recommended for Go Report Card compliance)

## Critical Architecture Concepts

### Output Format Handling
The SDK automatically handles CLI requirements:
- `text`: Plain text output, parsed into single `ResultMessage`
- `json`/`stream-json`: JSON parsing via `readMessages()`
- `stream-json` automatically adds `--verbose` flag (CLI requirement)

### Message Type Hierarchy
- `Message` interface with `Type()`, `Content()`, `Timestamp()` methods
- Concrete types: `AssistantMessage`, `UserMessage`, `SystemMessage`, `ResultMessage`
- Message fields aligned with TypeScript/Python SDKs: `parent_tool_use_id`, `session_id`, usage data
- `ContentBlock` interface with `TextBlock`, `ToolUseBlock`, `ToolResultBlock` implementations

### CLI Option Translation
`buildCommandArgs()` translates Go struct fields to CLI arguments:
- Pointer fields (`*string`, `*int`, `*bool`) represent optional CLI flags
- Array fields (`[]string`) become space-separated arguments
- New CLI flags: `--model`, `--permission-mode`, `--debug`, `--input-format`, `--dangerously-skip-permissions`, `--add-dir`
- Field name mapping: `Cwd` → working directory, `SystemPrompt` → `--system-prompt`
- Special handling for output format and verbose flag dependencies

### Error Context Strategy
- `CLINotFoundError`: Claude CLI not available
- `ProcessError`: CLI execution failure with exit codes and stderr
- `CLIConnectionError`: I/O pipe failures
- `CLIJSONDecodeError`: JSON parsing failures with raw data context

## Testing Strategy

### Unit Tests (`claude_test.go`)
- Type construction and interface compliance
- Error handling scenarios and error type behavior  
- JSON parsing logic with realistic CLI output
- CLI executable discovery across different environments

### Integration Tests
- Require authenticated Claude CLI
- Test actual CLI communication and response parsing
- Skipped automatically when CLI unavailable (`testing.Short()`)

### Example Validation
- `basic/`: Demonstrates both new and traditional APIs
- `streaming/`: Real-time processing and tool restrictions
- `advanced/`: MCP, sessions, and comprehensive option usage

## API Compatibility

The SDK maintains strict compatibility with TypeScript and Python SDKs:
- `QueryRequest` struct mirrors the `{prompt, options}` pattern
- `Options` fields use Go pointer types to distinguish unset from zero values
- Field names and ordering match official SDKs: `AllowedTools`, `SystemPrompt`, `Cwd`, etc.
- All CLI options from official SDKs are supported, plus additional CLI-documented options
- `Message`/`ContentBlock` interfaces provide identical type hierarchy with matching field names
- Message structures include all official SDK fields: `parent_tool_use_id`, `session_id`, usage data
- Environment variable `CLAUDE_CODE_ENTRYPOINT=sdk-go` matches official SDK pattern
- Error handling patterns mirror Python SDK error classes
- Function signatures support both Go patterns and cross-language consistency