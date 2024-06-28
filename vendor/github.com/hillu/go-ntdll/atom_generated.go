// This file was autogenerated using go run mkcode.go -- atom.go
// DO NOT EDIT.

package ntdll

import "unsafe"
import "reflect"

// The AtomInformationClass constants have been derived from the ATOM_INFORMATION_CLASS enum definition.
type AtomInformationClass uint32

const (
	AtomBasicInformation AtomInformationClass = 0
	AtomTableInformation                      = 1
)

var (
	procNtAddAtom              = modntdll.NewProc("NtAddAtom")
	procNtDeleteAtom           = modntdll.NewProc("NtDeleteAtom")
	procNtFindAtom             = modntdll.NewProc("NtFindAtom")
	procNtQueryInformationAtom = modntdll.NewProc("NtQueryInformationAtom")
)

// AtomBasicInformationT has been derived from the ATOM_BASIC_INFORMATION struct definition.
type AtomBasicInformationT struct {
	UsageCount uint16
	Flags      uint16
	NameLength uint16
	Name       [1]uint16
}

// NameSlice returns a slice over the elements of AtomBasicInformationT.Name.
//
// Beware: The data is not copied out of AtomBasicInformationT. The size can usually be taken from an other member of the struct (AtomBasicInformationT).
func (t *AtomBasicInformationT) NameSlice(size int) []uint16 {
	s := []uint16{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&t.Name[0]))
	hdr.Len = size
	hdr.Cap = size
	return s
}

// SetNameSlice copies s into the memory at AtomBasicInformationT.Name.
//
// Beware: No bounds check is performed. Another member of the struct (AtomBasicInformationT) usually has to be set to the array size.
func (t *AtomBasicInformationT) SetNameSlice(s []uint16) {
	s1 := []uint16{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	hdr.Data = uintptr(unsafe.Pointer(&t.Name[0]))
	hdr.Len = len(s)
	hdr.Cap = len(s)
	copy(s1, s)
}

// AtomTableInformationT has been derived from the ATOM_TABLE_INFORMATION struct definition.
type AtomTableInformationT struct {
	NumberOfAtoms uint32
	Atoms         [1]RtlAtom
}

// AtomsSlice returns a slice over the elements of AtomTableInformationT.Atoms.
//
// Beware: The data is not copied out of AtomTableInformationT. The size can usually be taken from an other member of the struct (AtomTableInformationT).
func (t *AtomTableInformationT) AtomsSlice(size int) []RtlAtom {
	s := []RtlAtom{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&t.Atoms[0]))
	hdr.Len = size
	hdr.Cap = size
	return s
}

// SetAtomsSlice copies s into the memory at AtomTableInformationT.Atoms.
//
// Beware: No bounds check is performed. Another member of the struct (AtomTableInformationT) usually has to be set to the array size.
func (t *AtomTableInformationT) SetAtomsSlice(s []RtlAtom) {
	s1 := []RtlAtom{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	hdr.Data = uintptr(unsafe.Pointer(&t.Atoms[0]))
	hdr.Len = len(s)
	hdr.Cap = len(s)
	copy(s1, s)
}

// OUT-parameter: Atom.
func NtAddAtom(
	AtomName *uint16,
	Length uint32,
	Atom *RtlAtom,
) NtStatus {
	r0, _, _ := procNtAddAtom.Call(uintptr(unsafe.Pointer(AtomName)),
		uintptr(Length),
		uintptr(unsafe.Pointer(Atom)))
	return NtStatus(r0)
}

func NtDeleteAtom(
	Atom RtlAtom,
) NtStatus {
	r0, _, _ := procNtDeleteAtom.Call(uintptr(Atom))
	return NtStatus(r0)
}

// OUT-parameter: Atom.
// *OPT-parameter: Atom.
func NtFindAtom(
	AtomName *uint16,
	Length uint32,
	Atom *RtlAtom,
) NtStatus {
	r0, _, _ := procNtFindAtom.Call(uintptr(unsafe.Pointer(AtomName)),
		uintptr(Length),
		uintptr(unsafe.Pointer(Atom)))
	return NtStatus(r0)
}

// OUT-parameter: AtomInformation, ReturnLength.
func NtQueryInformationAtom(
	Atom RtlAtom,
	AtomInformationClass AtomInformationClass,
	AtomInformation *byte,
	AtomInformationLength uint32,
	ReturnLength *uint32,
) NtStatus {
	r0, _, _ := procNtQueryInformationAtom.Call(uintptr(Atom),
		uintptr(AtomInformationClass),
		uintptr(unsafe.Pointer(AtomInformation)),
		uintptr(AtomInformationLength),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}
