package team

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type PythonCaller struct {
	pythonPath string
	timeout    time.Duration
}

func NewPythonCaller() *PythonCaller {
	return &PythonCaller{
		pythonPath: "python3",
		timeout:    30 * time.Second,
	}
}

type pyRequest struct {
	Module   string         `json:"module"`
	Function string         `json:"function"`
	Args     []interface{}  `json:"args,omitempty"`
	Kwargs   map[string]any `json:"kwargs,omitempty"`
}

type pyResponse struct {
	Result json.RawMessage `json:"result"`
	Error  string          `json:"error,omitempty"`
}

func (pc *PythonCaller) Call(ctx context.Context, module, function string, kwargs map[string]any) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, pc.timeout)
	defer cancel()

	req := pyRequest{
		Module:   module,
		Function: function,
		Kwargs:   kwargs,
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	cmd := exec.CommandContext(ctx, pc.pythonPath, "-c", fmt.Sprintf(`
import sys
import json

req = json.loads(%q)
module = __import__(req['module'])
func = getattr(module, req['function'])

# Build kwargs
kwargs = req.get('kwargs', {})
args = req.get('args', [])

result = func(**kwargs) if kwargs else func(*args)

# Handle non-JSON serializable results
if hasattr(result, '__dict__'):
    result = vars(result) if not hasattr(result, 'model_dump') else result.model_dump(by_alias=True)
elif isinstance(result, (list, tuple)):
    result = list(result)

print(json.dumps(result))
`, string(reqBytes)))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("python call failed: %w, stderr: %s", err, stderr.String())
	}

	var resp pyResponse
	if err := json.Unmarshal(stdout.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resp.Error != "" {
		return nil, fmt.Errorf("python error: %s", resp.Error)
	}

	return resp.Result, nil
}

// CallSimple calls a Python function with keyword arguments and returns parsed JSON
func (pc *PythonCaller) CallSimple(ctx context.Context, module, function string, kwargs map[string]any) error {
	_, err := pc.Call(ctx, module, function, kwargs)
	return err
}
