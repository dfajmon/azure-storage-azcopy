// This file was autogenerated using go run mkcode.go -- io.go
// DO NOT EDIT.

package ntdll

// The IoPriorityHint constants have been derived from the IO_PRIORITY_HINT enum definition.
type IoPriorityHint uint32

const (
	IoPriorityVeryLow  IoPriorityHint = 0
	IoPriorityLow                     = 1
	IoPriorityNormal                  = 2
	IoPriorityHigh                    = 3
	IoPriorityCritical                = 4
	MaxIoPriorityTypes                = 5
)
