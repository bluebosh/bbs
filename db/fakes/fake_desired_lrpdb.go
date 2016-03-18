// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeDesiredLRPDB struct {
	DesiredLRPsStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	DesiredLRPByProcessGuidStub        func(logger lager.Logger, processGuid string) (*models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	DesiredLRPSchedulingInfosStub        func(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error)
	desiredLRPSchedulingInfosMutex       sync.RWMutex
	desiredLRPSchedulingInfosArgsForCall []struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}
	desiredLRPSchedulingInfosReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	DesireLRPStub        func(logger lager.Logger, desiredLRP *models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	UpdateDesiredLRPStub        func(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (previousInstanceCount int32, err error)
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 int32
		result2 error
	}
	RemoveDesiredLRPStub        func(logger lager.Logger, processGuid string) error
	removeDesiredLRPMutex       sync.RWMutex
	removeDesiredLRPArgsForCall []struct {
		logger      lager.Logger
		processGuid string
	}
	removeDesiredLRPReturns struct {
		result1 error
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPs(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(logger, filter)
	} else {
		return fake.desiredLRPsReturns.result1, fake.desiredLRPsReturns.result2
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPsArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return fake.desiredLRPsArgsForCall[i].logger, fake.desiredLRPsArgsForCall[i].filter
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 error) {
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuid(logger lager.Logger, processGuid string) (*models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(logger, processGuid)
	} else {
		return fake.desiredLRPByProcessGuidReturns.result1, fake.desiredLRPByProcessGuidReturns.result2
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidArgsForCall(i int) (lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return fake.desiredLRPByProcessGuidArgsForCall[i].logger, fake.desiredLRPByProcessGuidArgsForCall[i].processGuid
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 error) {
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfos(logger lager.Logger, filter models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	fake.desiredLRPSchedulingInfosArgsForCall = append(fake.desiredLRPSchedulingInfosArgsForCall, struct {
		logger lager.Logger
		filter models.DesiredLRPFilter
	}{logger, filter})
	fake.desiredLRPSchedulingInfosMutex.Unlock()
	if fake.DesiredLRPSchedulingInfosStub != nil {
		return fake.DesiredLRPSchedulingInfosStub(logger, filter)
	} else {
		return fake.desiredLRPSchedulingInfosReturns.result1, fake.desiredLRPSchedulingInfosReturns.result2
	}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosCallCount() int {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return len(fake.desiredLRPSchedulingInfosArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosArgsForCall(i int) (lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return fake.desiredLRPSchedulingInfosArgsForCall[i].logger, fake.desiredLRPSchedulingInfosArgsForCall[i].filter
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.DesiredLRPSchedulingInfosStub = nil
	fake.desiredLRPSchedulingInfosReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesireLRP(logger lager.Logger, desiredLRP *models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		logger     lager.Logger
		desiredLRP *models.DesiredLRP
	}{logger, desiredLRP})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(logger, desiredLRP)
	} else {
		return fake.desireLRPReturns.result1
	}
}

func (fake *FakeDesiredLRPDB) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesireLRPArgsForCall(i int) (lager.Logger, *models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return fake.desireLRPArgsForCall[i].logger, fake.desireLRPArgsForCall[i].desiredLRP
}

func (fake *FakeDesiredLRPDB) DesireLRPReturns(result1 error) {
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRP(logger lager.Logger, processGuid string, update *models.DesiredLRPUpdate) (previousInstanceCount int32, err error) {
	fake.updateDesiredLRPMutex.Lock()
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		update      *models.DesiredLRPUpdate
	}{logger, processGuid, update})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(logger, processGuid, update)
	} else {
		return fake.updateDesiredLRPReturns.result1, fake.updateDesiredLRPReturns.result2
	}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPArgsForCall(i int) (lager.Logger, string, *models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return fake.updateDesiredLRPArgsForCall[i].logger, fake.updateDesiredLRPArgsForCall[i].processGuid, fake.updateDesiredLRPArgsForCall[i].update
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPReturns(result1 int32, result2 error) {
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 int32
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRP(logger lager.Logger, processGuid string) error {
	fake.removeDesiredLRPMutex.Lock()
	fake.removeDesiredLRPArgsForCall = append(fake.removeDesiredLRPArgsForCall, struct {
		logger      lager.Logger
		processGuid string
	}{logger, processGuid})
	fake.removeDesiredLRPMutex.Unlock()
	if fake.RemoveDesiredLRPStub != nil {
		return fake.RemoveDesiredLRPStub(logger, processGuid)
	} else {
		return fake.removeDesiredLRPReturns.result1
	}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPCallCount() int {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return len(fake.removeDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPArgsForCall(i int) (lager.Logger, string) {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return fake.removeDesiredLRPArgsForCall[i].logger, fake.removeDesiredLRPArgsForCall[i].processGuid
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPReturns(result1 error) {
	fake.RemoveDesiredLRPStub = nil
	fake.removeDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

var _ db.DesiredLRPDB = new(FakeDesiredLRPDB)
