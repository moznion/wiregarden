package routes

type IPRouter interface {
	AddRoute(ip string, deviceName string) error
	DelRoute(ip string, deviceName string) error
	Name() IPRoutingPolicyName
}

type IPRoutingPolicyName = string

const (
	IPRoutingPolicyIpcmd IPRoutingPolicyName = "ipcmd"
)

func IPRouterFrom(policyName IPRoutingPolicyName) IPRouter {
	switch policyName {
	case IPRoutingPolicyIpcmd:
		return &IpcmdIPRouter{
			name: IPRoutingPolicyIpcmd,
		}
	default:
		return nil
	}
}
