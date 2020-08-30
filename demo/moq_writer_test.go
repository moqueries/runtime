// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package demo_test

import "github.com/myshkin5/moqueries/pkg/hash"

// mockWriter holds the state of a mock of the Writer type
type mockWriter struct {
	resultsByParams_Write map[mockWriter_Write_params]mockWriter_Write_results
	params_Write          chan mockWriter_Write_params
}

// mockWriter_mock isolates the mock interface of the Writer type
type mockWriter_mock struct {
	mock *mockWriter
}

// mockWriter_recorder isolates the recorder interface of the Writer type
type mockWriter_recorder struct {
	mock *mockWriter
}

// mockWriter_Write_params holds the params of the Writer type
type mockWriter_Write_params struct{ p hash.Hash }

// mockWriter_Write_results holds the results of the Writer type
type mockWriter_Write_results struct {
	n   int
	err error
}

// mockWriter_Write_fnRecorder routes recorded function calls to the mockWriter mock
type mockWriter_Write_fnRecorder struct {
	params mockWriter_Write_params
	mock   *mockWriter
}

// newMockWriter creates a new mock of the Writer type
func newMockWriter() *mockWriter {
	return &mockWriter{
		resultsByParams_Write: map[mockWriter_Write_params]mockWriter_Write_results{},
		params_Write:          make(chan mockWriter_Write_params, 100),
	}
}

// mock returns the mock implementation of the Writer type
func (m *mockWriter) mock() *mockWriter_mock {
	return &mockWriter_mock{
		mock: m,
	}
}

func (m *mockWriter_mock) Write(p []byte) (n int, err error) {
	params := mockWriter_Write_params{
		p: hash.DeepHash(p),
	}
	m.mock.params_Write <- params
	results, ok := m.mock.resultsByParams_Write[params]
	if ok {
		n = results.n
		err = results.err
	}
	return n, err
}

// onCall returns the recorder implementation of the Writer type
func (m *mockWriter) onCall() *mockWriter_recorder {
	return &mockWriter_recorder{
		mock: m,
	}
}

func (m *mockWriter_recorder) Write(p []byte) *mockWriter_Write_fnRecorder {
	return &mockWriter_Write_fnRecorder{
		params: mockWriter_Write_params{
			p: hash.DeepHash(p),
		},
		mock: m.mock,
	}
}

func (r *mockWriter_Write_fnRecorder) ret(n int, err error) {
	r.mock.resultsByParams_Write[r.params] = mockWriter_Write_results{
		n:   n,
		err: err,
	}
}