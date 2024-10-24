package main

// package zrecorder

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Step is a struct that contains the details of a step
type Step struct {
	Action      Action  `json:"action"`
	ElapsedTime float64 `json:"elapsedTime"`
	Message     string  `json:"message"`
	Sequence    int     `json:"sequence"`
	Service     string  `json:"service"`
	Severity    string  `json:"severity"`
	Detail      any     `json:"detail,omitempty"`
	Code        string  `json:"code,omitempty"`
	Caller      string  `json:"caller,omitempty"`

	callerDepth int `json:"-"`
}

// Recorder is a struct that contains the steps of a process
type Recorder struct {
	StartTime time.Time `json:"startTime"`
	Steps     []Step    `json:"steps"`
}

// NewRecorder creates a new recorder
func NewRecorder() *Recorder {
	return &Recorder{
		StartTime: time.Now(),
	}
}

// H is a map of strings to any values
type H map[string]any

// Action is a string that represents the action of a step
type Action string

// Option is a function that can be used to set options for a step
type Option func(*Step)

// Set the caller depth
func WithCallerDepth(depth int) Option {
	return func(step *Step) {
		step.callerDepth = depth
	}
}

// Append a step to the recorder
func (ss *Recorder) AppenStep(action Action, detail any, opts ...Option) {
	// Generate a default file path and line where the
	// function is called, you can override it with the opts
	defaultCallerDepth := 1
	step := &Step{
		Action:      action,
		Sequence:    len(ss.Steps) + 1,
		Severity:    "INFO",
		Message:     "success",
		Detail:      detail,
		ElapsedTime: 0, // TODO: update this
		callerDepth: defaultCallerDepth,
	}

	for _, opt := range opts {
		opt(step)
	}

	// Set the caller
	_, filename, line, _ := runtime.Caller(step.callerDepth)
	step.Caller = fmt.Sprintf("%s:%d", trimmedPath(filename), line)

	ss.Steps = append(ss.Steps, *step)
}

// Get all steps grouped by action
func (ss *Recorder) GetActionSteps() map[Action][]Step {
	actionSteps := make(map[Action][]Step)
	for _, step := range ss.Steps {
		actionSteps[step.Action] = append(actionSteps[step.Action], step)
	}
	return actionSteps
}

func main() {
	rec := NewRecorder()
	rec.AppenStep("MAIN", Example{KPAZA: "kpa_za", WASHOQL: "washoql"})

	// Add rec in context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "rec", rec)

	ExampleDepth1(ctx)

	// print in json
	j, _ := json.MarshalIndent(rec, "", "  ")
	fmt.Println(string(j))
}

func ExampleDepth1(ctx context.Context) {
	// Get rec from context
	rec := ctx.Value("rec").(*Recorder)
	rec.AppenStep("EXAMPLE_DEPTH_1", H{"message": "the"}, WithCallerDepth(2))

	ExampleDepth2(ctx)
}

func ExampleDepth2(ctx context.Context) {
	// Get rec from context
	rec := ctx.Value("rec").(*Recorder)
	rec.AppenStep("EXAMPLE_DEPTH_2", H{"message": "police"}, WithCallerDepth(3))
}

func trimmedPath(fullPath string) string {
	// Split the path into components
	parts := strings.Split(fullPath, "/")

	// If there are less than 2 parts, return the full path
	if len(parts) < 2 {
		return fullPath
	}

	// Extract the last two parts (leaf directory and file name)
	leafDir := parts[len(parts)-2]
	fileName := parts[len(parts)-1]

	// Combine the leaf directory and file name
	trimmedPath := fmt.Sprintf("%s/%s", leafDir, fileName)

	return trimmedPath
}

type Example struct {
	KPAZA   string `json:"kpa_za"`
	WASHOQL string `json:"washoql"`
}
