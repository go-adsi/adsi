package api

import "unsafe"

// IADsComputerVtbl represents the component object model virtual
// function table for the IADsComputer interface.
type IADsComputerVtbl struct {
	IADsVtbl
	ComputerID                uintptr
	Site                      uintptr
	Description               uintptr
	SetDescription            uintptr
	Location                  uintptr
	SetLocation               uintptr
	PrimaryUser               uintptr
	SetPrimaryUser            uintptr
	Owner                     uintptr
	SetOwner                  uintptr
	Division                  uintptr
	SetDivision               uintptr
	Department                uintptr
	SetDepartment             uintptr
	Role                      uintptr
	SetRole                   uintptr
	OperatingSystem           uintptr
	SetOperatingSystem        uintptr
	OperatingSystemVersion    uintptr
	SetOperatingSystemVersion uintptr
	Model                     uintptr
	SetModel                  uintptr
	Processor                 uintptr
	SetProcessor              uintptr
	ProcessorCount            uintptr
	SetProcessorCount         uintptr
	MemorySize                uintptr
	SetMemorySize             uintptr
	StorageCapacity           uintptr
	SetStorageCapacity        uintptr
	NetAddresses              uintptr
	SetNetAddresses           uintptr
}

// IADsComputer represents the component object model interface for
// active directory computers.
type IADsComputer struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// computer.
func (v *IADsComputer) VTable() *IADsComputerVtbl {
	return (*IADsComputerVtbl)(unsafe.Pointer(v.RawVTable))
}
