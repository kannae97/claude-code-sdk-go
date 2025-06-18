package claudecode

import "fmt"

// ClaudeSDKError represents a general SDK error
type ClaudeSDKError struct {
	Message string
	Cause   error
}

func (e *ClaudeSDKError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Claude SDK error: %s (caused by: %v)", e.Message, e.Cause)
	}
	return fmt.Sprintf("Claude SDK error: %s", e.Message)
}

func (e *ClaudeSDKError) Unwrap() error {
	return e.Cause
}

// CLINotFoundError is returned when the Claude Code CLI cannot be found
type CLINotFoundError struct {
	Path string
}

func (e *CLINotFoundError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("Claude Code not found at: %s", e.Path)
	}
	return "Claude Code not found or not installed.\n\n" +
		"Install Claude Code with:\n" +
		"  npm install -g @anthropic-ai/claude-code\n" +
		"\nIf already installed locally, try:\n" +
		"  export PATH=\"$HOME/node_modules/.bin:$PATH\"\n" +
		"\nOr specify the path when creating options:\n" +
		"  &Options{Executable: \"/path/to/claude\"}"
}

// CLIConnectionError is returned when there's an error connecting to the CLI
type CLIConnectionError struct {
	Message string
	Cause   error
}

func (e *CLIConnectionError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("CLI connection error: %s (caused by: %v)", e.Message, e.Cause)
	}
	return fmt.Sprintf("CLI connection error: %s", e.Message)
}

func (e *CLIConnectionError) Unwrap() error {
	return e.Cause
}

// ProcessError is returned when the CLI process encounters an error
type ProcessError struct {
	ExitCode int
	Stderr   string
	Stdout   string
}

func (e *ProcessError) Error() string {
	return fmt.Sprintf("CLI process error (exit code %d): %s", e.ExitCode, e.Stderr)
}

// CLIJSONDecodeError is returned when JSON from the CLI cannot be decoded
type CLIJSONDecodeError struct {
	Data  string
	Cause error
}

func (e *CLIJSONDecodeError) Error() string {
	return fmt.Sprintf("failed to decode CLI JSON response: %v (data: %s)", e.Cause, e.Data)
}

func (e *CLIJSONDecodeError) Unwrap() error {
	return e.Cause
}
