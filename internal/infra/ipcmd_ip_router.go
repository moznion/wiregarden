package infra

import "os/exec"

type IpcmdIPRouter struct {
	name IPRoutingPolicyName
}

func (r *IpcmdIPRouter) AddRoute(ip string, deviceName string) error {
	return r.runIpcmd("add", ip, deviceName)
}

func (r *IpcmdIPRouter) DelRoute(ip string, deviceName string) error {
	return r.runIpcmd("del", ip, deviceName)
}

func (r *IpcmdIPRouter) Name() IPRoutingPolicyName {
	return r.name
}

func (r *IpcmdIPRouter) runIpcmd(op string, ip string, deviceName string) error {
	err := exec.Command("ip", "route", op, ip, "dev", deviceName).Run() // #nosec
	if err != nil {
		return err
	}
	return nil
}
