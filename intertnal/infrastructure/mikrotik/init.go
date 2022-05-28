package mikrotik

import "wlcontrol/intertnal/domain/entity"

type Device struct {
}

func NewDevice() Device {
	return Device{}
}

func (d Device) AddIPToAddrList(m entity.Mikrotik, ip, addrList string) error {
	return nil
}

func (d Device) DeleteIPFromAddrList(m entity.Mikrotik, ip, addrList string) error {
	return nil
}
