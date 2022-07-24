package idgenerator

import (
	"sync"
	"sync/atomic"
)

type idLazyGenerator struct {
	id uint64
}

var (
	IdLazyGenerator *idLazyGenerator
	once            = &sync.Once{}
)

func GetLazyInstance() *idLazyGenerator {
	if IdLazyGenerator == nil {
		once.Do(func() {
			IdLazyGenerator = &idLazyGenerator{
				id: 0,
			}
		})
	}
	return IdLazyGenerator
}

func (self *idLazyGenerator) GetId() uint64 {
	return atomic.AddUint64(&self.id, 1)
}
