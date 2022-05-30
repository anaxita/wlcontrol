package mikrotik

import "wlcontrol/intertnal/domain/entity"

type Device struct {
}

func New() *Device {
	return &Device{}
}

func (d *Device) AddIPToAddrList(m entity.Mikrotik, ip string) error {
	return nil
}

func (d *Device) DeleteIPFromAddrList(m entity.Mikrotik, ip string) error {
	return nil
}
