package domain

type HealthState int32

const (
	Up HealthState = iota
	Down
	OutOfService
	Unknown
)

type HealthStatus struct {
	State HealthState
}
