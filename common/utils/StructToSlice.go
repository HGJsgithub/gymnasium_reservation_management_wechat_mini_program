package utils

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"unsafe"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func BadmintonStructToSlice1(structure []venue.BadmintonVenueState1) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.BadmintonVenueState1{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func BadmintonStructToSlice2(structure []venue.BadmintonVenueState2) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.BadmintonVenueState1{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func TableTennisStructToSlice1(structure []venue.TableTennisVenueState1) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.TableTennisVenueState1{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func TableTennisStructToSlice2(structure []venue.TableTennisVenueState2) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.TableTennisVenueState2{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func TennisStructToSlice1(structure []venue.TennisVenueState1) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.TennisVenueState1{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}

func TennisStructToSlice2(structure []venue.TennisVenueState2) [][]byte {
	var rawState [][]byte
	var byteSlice []byte
	Len := unsafe.Sizeof(venue.TennisVenueState2{})
	for i, _ := range structure {
		sliceMock := &SliceMock{
			addr: uintptr(unsafe.Pointer(&structure[i])),
			cap:  int(Len),
			len:  int(Len),
		}
		byteSlice = *(*[]byte)(unsafe.Pointer(sliceMock))
		rawState = append(rawState, byteSlice)
	}
	return rawState
}
