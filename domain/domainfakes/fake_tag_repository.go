// Code generated by counterfeiter. DO NOT EDIT.
package domainfakes

import (
	"context"
	"realworld-go/domain"
	"realworld-go/domain/model"
	"sync"
)

type FakeTagRepository struct {
	CreateStub        func(context.Context, model.Tag) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 model.Tag
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteByIDStub        func(context.Context, model.Tag) error
	deleteByIDMutex       sync.RWMutex
	deleteByIDArgsForCall []struct {
		arg1 context.Context
		arg2 model.Tag
	}
	deleteByIDReturns struct {
		result1 error
	}
	deleteByIDReturnsOnCall map[int]struct {
		result1 error
	}
	FindAllByIDSStub        func(context.Context, []string) ([]model.Tag, error)
	findAllByIDSMutex       sync.RWMutex
	findAllByIDSArgsForCall []struct {
		arg1 context.Context
		arg2 []string
	}
	findAllByIDSReturns struct {
		result1 []model.Tag
		result2 error
	}
	findAllByIDSReturnsOnCall map[int]struct {
		result1 []model.Tag
		result2 error
	}
	FindByIDStub        func(context.Context, string) (model.Tag, error)
	findByIDMutex       sync.RWMutex
	findByIDArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	findByIDReturns struct {
		result1 model.Tag
		result2 error
	}
	findByIDReturnsOnCall map[int]struct {
		result1 model.Tag
		result2 error
	}
	FindTagPopulerStub        func(context.Context, int64) ([]domain.FindTagPopulerResult, error)
	findTagPopulerMutex       sync.RWMutex
	findTagPopulerArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	findTagPopulerReturns struct {
		result1 []domain.FindTagPopulerResult
		result2 error
	}
	findTagPopulerReturnsOnCall map[int]struct {
		result1 []domain.FindTagPopulerResult
		result2 error
	}
	UpdateByIDStub        func(context.Context, model.Tag, []string) error
	updateByIDMutex       sync.RWMutex
	updateByIDArgsForCall []struct {
		arg1 context.Context
		arg2 model.Tag
		arg3 []string
	}
	updateByIDReturns struct {
		result1 error
	}
	updateByIDReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTagRepository) Create(arg1 context.Context, arg2 model.Tag) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 model.Tag
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTagRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeTagRepository) CreateCalls(stub func(context.Context, model.Tag) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeTagRepository) CreateArgsForCall(i int) (context.Context, model.Tag) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTagRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) DeleteByID(arg1 context.Context, arg2 model.Tag) error {
	fake.deleteByIDMutex.Lock()
	ret, specificReturn := fake.deleteByIDReturnsOnCall[len(fake.deleteByIDArgsForCall)]
	fake.deleteByIDArgsForCall = append(fake.deleteByIDArgsForCall, struct {
		arg1 context.Context
		arg2 model.Tag
	}{arg1, arg2})
	stub := fake.DeleteByIDStub
	fakeReturns := fake.deleteByIDReturns
	fake.recordInvocation("DeleteByID", []interface{}{arg1, arg2})
	fake.deleteByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTagRepository) DeleteByIDCallCount() int {
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	return len(fake.deleteByIDArgsForCall)
}

func (fake *FakeTagRepository) DeleteByIDCalls(stub func(context.Context, model.Tag) error) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = stub
}

func (fake *FakeTagRepository) DeleteByIDArgsForCall(i int) (context.Context, model.Tag) {
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	argsForCall := fake.deleteByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTagRepository) DeleteByIDReturns(result1 error) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = nil
	fake.deleteByIDReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) DeleteByIDReturnsOnCall(i int, result1 error) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = nil
	if fake.deleteByIDReturnsOnCall == nil {
		fake.deleteByIDReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteByIDReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) FindAllByIDS(arg1 context.Context, arg2 []string) ([]model.Tag, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.findAllByIDSMutex.Lock()
	ret, specificReturn := fake.findAllByIDSReturnsOnCall[len(fake.findAllByIDSArgsForCall)]
	fake.findAllByIDSArgsForCall = append(fake.findAllByIDSArgsForCall, struct {
		arg1 context.Context
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.FindAllByIDSStub
	fakeReturns := fake.findAllByIDSReturns
	fake.recordInvocation("FindAllByIDS", []interface{}{arg1, arg2Copy})
	fake.findAllByIDSMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTagRepository) FindAllByIDSCallCount() int {
	fake.findAllByIDSMutex.RLock()
	defer fake.findAllByIDSMutex.RUnlock()
	return len(fake.findAllByIDSArgsForCall)
}

func (fake *FakeTagRepository) FindAllByIDSCalls(stub func(context.Context, []string) ([]model.Tag, error)) {
	fake.findAllByIDSMutex.Lock()
	defer fake.findAllByIDSMutex.Unlock()
	fake.FindAllByIDSStub = stub
}

func (fake *FakeTagRepository) FindAllByIDSArgsForCall(i int) (context.Context, []string) {
	fake.findAllByIDSMutex.RLock()
	defer fake.findAllByIDSMutex.RUnlock()
	argsForCall := fake.findAllByIDSArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTagRepository) FindAllByIDSReturns(result1 []model.Tag, result2 error) {
	fake.findAllByIDSMutex.Lock()
	defer fake.findAllByIDSMutex.Unlock()
	fake.FindAllByIDSStub = nil
	fake.findAllByIDSReturns = struct {
		result1 []model.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) FindAllByIDSReturnsOnCall(i int, result1 []model.Tag, result2 error) {
	fake.findAllByIDSMutex.Lock()
	defer fake.findAllByIDSMutex.Unlock()
	fake.FindAllByIDSStub = nil
	if fake.findAllByIDSReturnsOnCall == nil {
		fake.findAllByIDSReturnsOnCall = make(map[int]struct {
			result1 []model.Tag
			result2 error
		})
	}
	fake.findAllByIDSReturnsOnCall[i] = struct {
		result1 []model.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) FindByID(arg1 context.Context, arg2 string) (model.Tag, error) {
	fake.findByIDMutex.Lock()
	ret, specificReturn := fake.findByIDReturnsOnCall[len(fake.findByIDArgsForCall)]
	fake.findByIDArgsForCall = append(fake.findByIDArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.FindByIDStub
	fakeReturns := fake.findByIDReturns
	fake.recordInvocation("FindByID", []interface{}{arg1, arg2})
	fake.findByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTagRepository) FindByIDCallCount() int {
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	return len(fake.findByIDArgsForCall)
}

func (fake *FakeTagRepository) FindByIDCalls(stub func(context.Context, string) (model.Tag, error)) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = stub
}

func (fake *FakeTagRepository) FindByIDArgsForCall(i int) (context.Context, string) {
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	argsForCall := fake.findByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTagRepository) FindByIDReturns(result1 model.Tag, result2 error) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = nil
	fake.findByIDReturns = struct {
		result1 model.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) FindByIDReturnsOnCall(i int, result1 model.Tag, result2 error) {
	fake.findByIDMutex.Lock()
	defer fake.findByIDMutex.Unlock()
	fake.FindByIDStub = nil
	if fake.findByIDReturnsOnCall == nil {
		fake.findByIDReturnsOnCall = make(map[int]struct {
			result1 model.Tag
			result2 error
		})
	}
	fake.findByIDReturnsOnCall[i] = struct {
		result1 model.Tag
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) FindTagPopuler(arg1 context.Context, arg2 int64) ([]domain.FindTagPopulerResult, error) {
	fake.findTagPopulerMutex.Lock()
	ret, specificReturn := fake.findTagPopulerReturnsOnCall[len(fake.findTagPopulerArgsForCall)]
	fake.findTagPopulerArgsForCall = append(fake.findTagPopulerArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.FindTagPopulerStub
	fakeReturns := fake.findTagPopulerReturns
	fake.recordInvocation("FindTagPopuler", []interface{}{arg1, arg2})
	fake.findTagPopulerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTagRepository) FindTagPopulerCallCount() int {
	fake.findTagPopulerMutex.RLock()
	defer fake.findTagPopulerMutex.RUnlock()
	return len(fake.findTagPopulerArgsForCall)
}

func (fake *FakeTagRepository) FindTagPopulerCalls(stub func(context.Context, int64) ([]domain.FindTagPopulerResult, error)) {
	fake.findTagPopulerMutex.Lock()
	defer fake.findTagPopulerMutex.Unlock()
	fake.FindTagPopulerStub = stub
}

func (fake *FakeTagRepository) FindTagPopulerArgsForCall(i int) (context.Context, int64) {
	fake.findTagPopulerMutex.RLock()
	defer fake.findTagPopulerMutex.RUnlock()
	argsForCall := fake.findTagPopulerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTagRepository) FindTagPopulerReturns(result1 []domain.FindTagPopulerResult, result2 error) {
	fake.findTagPopulerMutex.Lock()
	defer fake.findTagPopulerMutex.Unlock()
	fake.FindTagPopulerStub = nil
	fake.findTagPopulerReturns = struct {
		result1 []domain.FindTagPopulerResult
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) FindTagPopulerReturnsOnCall(i int, result1 []domain.FindTagPopulerResult, result2 error) {
	fake.findTagPopulerMutex.Lock()
	defer fake.findTagPopulerMutex.Unlock()
	fake.FindTagPopulerStub = nil
	if fake.findTagPopulerReturnsOnCall == nil {
		fake.findTagPopulerReturnsOnCall = make(map[int]struct {
			result1 []domain.FindTagPopulerResult
			result2 error
		})
	}
	fake.findTagPopulerReturnsOnCall[i] = struct {
		result1 []domain.FindTagPopulerResult
		result2 error
	}{result1, result2}
}

func (fake *FakeTagRepository) UpdateByID(arg1 context.Context, arg2 model.Tag, arg3 []string) error {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.updateByIDMutex.Lock()
	ret, specificReturn := fake.updateByIDReturnsOnCall[len(fake.updateByIDArgsForCall)]
	fake.updateByIDArgsForCall = append(fake.updateByIDArgsForCall, struct {
		arg1 context.Context
		arg2 model.Tag
		arg3 []string
	}{arg1, arg2, arg3Copy})
	stub := fake.UpdateByIDStub
	fakeReturns := fake.updateByIDReturns
	fake.recordInvocation("UpdateByID", []interface{}{arg1, arg2, arg3Copy})
	fake.updateByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTagRepository) UpdateByIDCallCount() int {
	fake.updateByIDMutex.RLock()
	defer fake.updateByIDMutex.RUnlock()
	return len(fake.updateByIDArgsForCall)
}

func (fake *FakeTagRepository) UpdateByIDCalls(stub func(context.Context, model.Tag, []string) error) {
	fake.updateByIDMutex.Lock()
	defer fake.updateByIDMutex.Unlock()
	fake.UpdateByIDStub = stub
}

func (fake *FakeTagRepository) UpdateByIDArgsForCall(i int) (context.Context, model.Tag, []string) {
	fake.updateByIDMutex.RLock()
	defer fake.updateByIDMutex.RUnlock()
	argsForCall := fake.updateByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTagRepository) UpdateByIDReturns(result1 error) {
	fake.updateByIDMutex.Lock()
	defer fake.updateByIDMutex.Unlock()
	fake.UpdateByIDStub = nil
	fake.updateByIDReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) UpdateByIDReturnsOnCall(i int, result1 error) {
	fake.updateByIDMutex.Lock()
	defer fake.updateByIDMutex.Unlock()
	fake.UpdateByIDStub = nil
	if fake.updateByIDReturnsOnCall == nil {
		fake.updateByIDReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateByIDReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTagRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	fake.findAllByIDSMutex.RLock()
	defer fake.findAllByIDSMutex.RUnlock()
	fake.findByIDMutex.RLock()
	defer fake.findByIDMutex.RUnlock()
	fake.findTagPopulerMutex.RLock()
	defer fake.findTagPopulerMutex.RUnlock()
	fake.updateByIDMutex.RLock()
	defer fake.updateByIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTagRepository) recordInvocation(key string, args []interface{}) {
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

var _ domain.TagRepository = new(FakeTagRepository)