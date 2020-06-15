package checkntservice

import (
	"os/exec"
	"runtime"
	"syscall"
	"testing"

	"github.com/mackerelio/checkers"
	"github.com/stretchr/testify/assert"
)

func stopFaxService() error {
	_, err := exec.Command("net", "stop", "Fax").CombinedOutput()
	return err
}

func startFaxService() error {
	_, err := exec.Command("net", "start", "Fax").CombinedOutput()
	return err
}

func mockServiceState() {
	getServiceStateFunc = func() ([]Win32Service, error) {
		runningService := Win32Service{
			Caption: "running-service-caption",
			Name:    "running-service-name",
			State:   "Running",
		}
		stoppedService := Win32Service{
			Caption: "stopped-service-caption",
			Name:    "stopped-service-name",
			State:   "Stopped",
		}
		ss := []Win32Service{
			runningService,
			stoppedService,
		}
		return ss, nil
	}
}

func TestNtService(t *testing.T) {
	ss, err := getServiceState()
	if runtime.GOOS != "windows" {
		if err == nil || err != syscall.ENOSYS {
			t.Fatal(runtime.GOOS + " should fail because it's not Windows")
		}
		t.Skip(runtime.GOOS + " doesn't implement Windows NT service")
	}
	if err != nil {
		t.Errorf("failed to get service status: %v", err)
	}
	for _, s := range ss {
		if s.Name == "Fax" {
			if s.State != "Running" {
				t.Errorf("Fax service should be started in default: %v", s.State)
			}
		}
	}

	err = stopFaxService()
	if err != nil {
		t.Skipf("failed to stop Fax service. But ignore this: %v", err)
	}
	defer startFaxService()

	ss, err = getServiceState()
	if err != nil {
		t.Errorf("failed to get service status: %v", err)
	}
	for _, s := range ss {
		if s.Name == "Fax" {
			if s.State == "Running" {
				t.Error("Fax service should be stopped now")
			}
		}
	}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		casename      string
		cmdline       []string
		expectStatus  checkers.Status
		expectMessage string
	}{
		{
			casename:      "check about running service",
			cmdline:       []string{"-s", "running-service"},
			expectStatus:  checkers.OK,
			expectMessage: "",
		},
		{
			casename:      "check about stopped service",
			cmdline:       []string{"-s", "stopped-service"},
			expectStatus:  checkers.CRITICAL,
			expectMessage: "stopped-service-name: stopped-service-caption - Stopped",
		},
		{
			casename:      "check about running service with exclude option",
			cmdline:       []string{"-s", "service", "-x", "stopped"},
			expectStatus:  checkers.OK,
			expectMessage: "",
		},
	}

	originalFunc := getServiceStateFunc
	defer func() {
		getServiceStateFunc = originalFunc
	}()
	mockServiceState()

	for _, tc := range testCases {
		t.Run(tc.casename, func(t *testing.T) {
			result := run(tc.cmdline)
			assert.Equal(t, tc.expectStatus, result.Status, "something went wrong")
			assert.Equal(t, tc.expectMessage, result.Message, "something went wrong")
		})
	}
}
