package idgenerator

import "sync/atomic"

type idEagerGenerator struct {
	id uint64
}

var iEG *idEagerGenerator

func init() {
	iEG = &idEagerGenerator{
		id: 0,
	}
}

func GetEagerInstance() *idEagerGenerator {
	return iEG
}

func (self *idEagerGenerator) GetId() uint64 {
	return atomic.AddUint64(&self.id, 1)
}
