# Basic Example

This example demonstrates the basic usage of the Claude Code SDK for Go, showing both the new TypeScript/Python-compatible API and the traditional Go API.

## Features Demonstrated

- **QueryWithRequest**: TypeScript/Python SDK compatible API
- **Query**: Traditional Go API (backward compatibility)
- **Multiple output formats**: JSON and text
- **Basic options**: MaxTurns, AllowedTools, SystemPrompt
- **Error handling**: Detailed error information

## Usage

```bash
go run main.go
```

## What it does

1. **Example 1**: Uses `QueryWithRequest` with JSON output format to create a Go file
2. **Example 2**: Uses traditional `Query` API with text output format to list files

## Key Features

- Shows API compatibility between different SDK versions
- Demonstrates proper error handling
- Uses helper functions for pointer types
- Shows message processing and content extraction

## Output

The example will:
1. Create a simple hello.go file (if not exists)
2. List files in the current directory
3. Display detailed message information including content blocks and tool usage