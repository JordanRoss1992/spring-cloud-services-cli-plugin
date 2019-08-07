// Code generated by counterfeiter. DO NOT EDIT.
package instancefakes

import (
	"sync"

	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/instance"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/serviceutil"
)

type FakeOperation struct {
	IsServiceBrokerOperationStub        func() bool
	isServiceBrokerOperationMutex       sync.RWMutex
	isServiceBrokerOperationArgsForCall []struct {
	}
	isServiceBrokerOperationReturns struct {
		result1 bool
	}
	isServiceBrokerOperationReturnsOnCall map[int]struct {
		result1 bool
	}
	RunStub        func(serviceutil.ManagementParameters, string) (string, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 serviceutil.ManagementParameters
		arg2 string
	}
	runReturns struct {
		result1 string
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOperation) IsServiceBrokerOperation() bool {
	fake.isServiceBrokerOperationMutex.Lock()
	ret, specificReturn := fake.isServiceBrokerOperationReturnsOnCall[len(fake.isServiceBrokerOperationArgsForCall)]
	fake.isServiceBrokerOperationArgsForCall = append(fake.isServiceBrokerOperationArgsForCall, struct {
	}{})
	fake.recordInvocation("IsServiceBrokerOperation", []interface{}{})
	fake.isServiceBrokerOperationMutex.Unlock()
	if fake.IsServiceBrokerOperationStub != nil {
		return fake.IsServiceBrokerOperationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isServiceBrokerOperationReturns
	return fakeReturns.result1
}

func (fake *FakeOperation) IsServiceBrokerOperationCallCount() int {
	fake.isServiceBrokerOperationMutex.RLock()
	defer fake.isServiceBrokerOperationMutex.RUnlock()
	return len(fake.isServiceBrokerOperationArgsForCall)
}

func (fake *FakeOperation) IsServiceBrokerOperationCalls(stub func() bool) {
	fake.isServiceBrokerOperationMutex.Lock()
	defer fake.isServiceBrokerOperationMutex.Unlock()
	fake.IsServiceBrokerOperationStub = stub
}

func (fake *FakeOperation) IsServiceBrokerOperationReturns(result1 bool) {
	fake.isServiceBrokerOperationMutex.Lock()
	defer fake.isServiceBrokerOperationMutex.Unlock()
	fake.IsServiceBrokerOperationStub = nil
	fake.isServiceBrokerOperationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) IsServiceBrokerOperationReturnsOnCall(i int, result1 bool) {
	fake.isServiceBrokerOperationMutex.Lock()
	defer fake.isServiceBrokerOperationMutex.Unlock()
	fake.IsServiceBrokerOperationStub = nil
	if fake.isServiceBrokerOperationReturnsOnCall == nil {
		fake.isServiceBrokerOperationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isServiceBrokerOperationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) Run(arg1 serviceutil.ManagementParameters, arg2 string) (string, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 serviceutil.ManagementParameters
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOperation) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeOperation) RunCalls(stub func(serviceutil.ManagementParameters, string) (string, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeOperation) RunArgsForCall(i int) (serviceutil.ManagementParameters, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOperation) RunReturns(result1 string, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOperation) RunReturnsOnCall(i int, result1 string, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOperation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isServiceBrokerOperationMutex.RLock()
	defer fake.isServiceBrokerOperationMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOperation) recordInvocation(key string, args []interface{}) {
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

var _ instance.Operation = new(FakeOperation)