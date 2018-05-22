// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeEvacuationDB struct {
	RemoveEvacuatingActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error
	removeEvacuatingActualLRPMutex       sync.RWMutex
	removeEvacuatingActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}
	removeEvacuatingActualLRPReturns struct {
		result1 error
	}
	removeEvacuatingActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	EvacuateActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, uint64) (*models.ActualLRP, error)
	evacuateActualLRPMutex       sync.RWMutex
	evacuateActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 *models.ActualLRPNetInfo
		arg5 uint64
	}
	evacuateActualLRPReturns struct {
		result1 *models.ActualLRP
		result2 error
	}
	evacuateActualLRPReturnsOnCall map[int]struct {
		result1 *models.ActualLRP
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEvacuationDB) RemoveEvacuatingActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey) error {
	fake.removeEvacuatingActualLRPMutex.Lock()
	ret, specificReturn := fake.removeEvacuatingActualLRPReturnsOnCall[len(fake.removeEvacuatingActualLRPArgsForCall)]
	fake.removeEvacuatingActualLRPArgsForCall = append(fake.removeEvacuatingActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.recordInvocation("RemoveEvacuatingActualLRP", []interface{}{arg1, arg2, arg3})
	fake.removeEvacuatingActualLRPMutex.Unlock()
	if fake.RemoveEvacuatingActualLRPStub != nil {
		return fake.RemoveEvacuatingActualLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeEvacuatingActualLRPReturns.result1
}

func (fake *FakeEvacuationDB) RemoveEvacuatingActualLRPCallCount() int {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return len(fake.removeEvacuatingActualLRPArgsForCall)
}

func (fake *FakeEvacuationDB) RemoveEvacuatingActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return fake.removeEvacuatingActualLRPArgsForCall[i].arg1, fake.removeEvacuatingActualLRPArgsForCall[i].arg2, fake.removeEvacuatingActualLRPArgsForCall[i].arg3
}

func (fake *FakeEvacuationDB) RemoveEvacuatingActualLRPReturns(result1 error) {
	fake.RemoveEvacuatingActualLRPStub = nil
	fake.removeEvacuatingActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationDB) RemoveEvacuatingActualLRPReturnsOnCall(i int, result1 error) {
	fake.RemoveEvacuatingActualLRPStub = nil
	if fake.removeEvacuatingActualLRPReturnsOnCall == nil {
		fake.removeEvacuatingActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeEvacuatingActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationDB) EvacuateActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey, arg4 *models.ActualLRPNetInfo, arg5 uint64) (*models.ActualLRP, error) {
	fake.evacuateActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateActualLRPReturnsOnCall[len(fake.evacuateActualLRPArgsForCall)]
	fake.evacuateActualLRPArgsForCall = append(fake.evacuateActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 *models.ActualLRPNetInfo
		arg5 uint64
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("EvacuateActualLRP", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.evacuateActualLRPMutex.Unlock()
	if fake.EvacuateActualLRPStub != nil {
		return fake.EvacuateActualLRPStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.evacuateActualLRPReturns.result1, fake.evacuateActualLRPReturns.result2
}

func (fake *FakeEvacuationDB) EvacuateActualLRPCallCount() int {
	fake.evacuateActualLRPMutex.RLock()
	defer fake.evacuateActualLRPMutex.RUnlock()
	return len(fake.evacuateActualLRPArgsForCall)
}

func (fake *FakeEvacuationDB) EvacuateActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, uint64) {
	fake.evacuateActualLRPMutex.RLock()
	defer fake.evacuateActualLRPMutex.RUnlock()
	return fake.evacuateActualLRPArgsForCall[i].arg1, fake.evacuateActualLRPArgsForCall[i].arg2, fake.evacuateActualLRPArgsForCall[i].arg3, fake.evacuateActualLRPArgsForCall[i].arg4, fake.evacuateActualLRPArgsForCall[i].arg5
}

func (fake *FakeEvacuationDB) EvacuateActualLRPReturns(result1 *models.ActualLRP, result2 error) {
	fake.EvacuateActualLRPStub = nil
	fake.evacuateActualLRPReturns = struct {
		result1 *models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationDB) EvacuateActualLRPReturnsOnCall(i int, result1 *models.ActualLRP, result2 error) {
	fake.EvacuateActualLRPStub = nil
	if fake.evacuateActualLRPReturnsOnCall == nil {
		fake.evacuateActualLRPReturnsOnCall = make(map[int]struct {
			result1 *models.ActualLRP
			result2 error
		})
	}
	fake.evacuateActualLRPReturnsOnCall[i] = struct {
		result1 *models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeEvacuationDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	fake.evacuateActualLRPMutex.RLock()
	defer fake.evacuateActualLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEvacuationDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.EvacuationDB = new(FakeEvacuationDB)
