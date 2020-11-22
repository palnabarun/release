// Code generated by counterfeiter. DO NOT EDIT.
package anagofakes

import (
	"sync"

	"github.com/blang/semver"
	"k8s.io/release/pkg/announce"
	"k8s.io/release/pkg/build"
	"k8s.io/release/pkg/gcp/gcb"
	"k8s.io/release/pkg/release"
)

type FakeReleaseImpl struct {
	ArchiveReleaseStub        func(*release.ArchiverOptions) error
	archiveReleaseMutex       sync.RWMutex
	archiveReleaseArgsForCall []struct {
		arg1 *release.ArchiverOptions
	}
	archiveReleaseReturns struct {
		result1 error
	}
	archiveReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	BranchNeedsCreationStub        func(string, string, semver.Version) (bool, error)
	branchNeedsCreationMutex       sync.RWMutex
	branchNeedsCreationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 semver.Version
	}
	branchNeedsCreationReturns struct {
		result1 bool
		result2 error
	}
	branchNeedsCreationReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	CheckPrerequisitesStub        func() error
	checkPrerequisitesMutex       sync.RWMutex
	checkPrerequisitesArgsForCall []struct {
	}
	checkPrerequisitesReturns struct {
		result1 error
	}
	checkPrerequisitesReturnsOnCall map[int]struct {
		result1 error
	}
	CheckReleaseBucketStub        func(*build.Options) error
	checkReleaseBucketMutex       sync.RWMutex
	checkReleaseBucketArgsForCall []struct {
		arg1 *build.Options
	}
	checkReleaseBucketReturns struct {
		result1 error
	}
	checkReleaseBucketReturnsOnCall map[int]struct {
		result1 error
	}
	CopyStagedFromGCSStub        func(*build.Options, string, string) error
	copyStagedFromGCSMutex       sync.RWMutex
	copyStagedFromGCSArgsForCall []struct {
		arg1 *build.Options
		arg2 string
		arg3 string
	}
	copyStagedFromGCSReturns struct {
		result1 error
	}
	copyStagedFromGCSReturnsOnCall map[int]struct {
		result1 error
	}
	CreateAnnouncementStub        func(*announce.Options) error
	createAnnouncementMutex       sync.RWMutex
	createAnnouncementArgsForCall []struct {
		arg1 *announce.Options
	}
	createAnnouncementReturns struct {
		result1 error
	}
	createAnnouncementReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateReleaseVersionStub        func(string, string, string, bool) (*release.Versions, error)
	generateReleaseVersionMutex       sync.RWMutex
	generateReleaseVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}
	generateReleaseVersionReturns struct {
		result1 *release.Versions
		result2 error
	}
	generateReleaseVersionReturnsOnCall map[int]struct {
		result1 *release.Versions
		result2 error
	}
	NewGitPusherStub        func(*release.GitObjectPusherOptions) (*release.GitObjectPusher, error)
	newGitPusherMutex       sync.RWMutex
	newGitPusherArgsForCall []struct {
		arg1 *release.GitObjectPusherOptions
	}
	newGitPusherReturns struct {
		result1 *release.GitObjectPusher
		result2 error
	}
	newGitPusherReturnsOnCall map[int]struct {
		result1 *release.GitObjectPusher
		result2 error
	}
	PrepareWorkspaceReleaseStub        func(string, string) error
	prepareWorkspaceReleaseMutex       sync.RWMutex
	prepareWorkspaceReleaseArgsForCall []struct {
		arg1 string
		arg2 string
	}
	prepareWorkspaceReleaseReturns struct {
		result1 error
	}
	prepareWorkspaceReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	PublishVersionStub        func(string, string, string, string, string, []string, bool, bool) error
	publishVersionMutex       sync.RWMutex
	publishVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 []string
		arg7 bool
		arg8 bool
	}
	publishVersionReturns struct {
		result1 error
	}
	publishVersionReturnsOnCall map[int]struct {
		result1 error
	}
	PushBranchesStub        func(*release.GitObjectPusher, []string) error
	pushBranchesMutex       sync.RWMutex
	pushBranchesArgsForCall []struct {
		arg1 *release.GitObjectPusher
		arg2 []string
	}
	pushBranchesReturns struct {
		result1 error
	}
	pushBranchesReturnsOnCall map[int]struct {
		result1 error
	}
	PushMainBranchStub        func(*release.GitObjectPusher) error
	pushMainBranchMutex       sync.RWMutex
	pushMainBranchArgsForCall []struct {
		arg1 *release.GitObjectPusher
	}
	pushMainBranchReturns struct {
		result1 error
	}
	pushMainBranchReturnsOnCall map[int]struct {
		result1 error
	}
	PushTagsStub        func(*release.GitObjectPusher, []string) error
	pushTagsMutex       sync.RWMutex
	pushTagsArgsForCall []struct {
		arg1 *release.GitObjectPusher
		arg2 []string
	}
	pushTagsReturns struct {
		result1 error
	}
	pushTagsReturnsOnCall map[int]struct {
		result1 error
	}
	SubmitStub        func(*gcb.Options) error
	submitMutex       sync.RWMutex
	submitArgsForCall []struct {
		arg1 *gcb.Options
	}
	submitReturns struct {
		result1 error
	}
	submitReturnsOnCall map[int]struct {
		result1 error
	}
	ToFileStub        func(string) error
	toFileMutex       sync.RWMutex
	toFileArgsForCall []struct {
		arg1 string
	}
	toFileReturns struct {
		result1 error
	}
	toFileReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateImagesStub        func(string, string, string) error
	validateImagesMutex       sync.RWMutex
	validateImagesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	validateImagesReturns struct {
		result1 error
	}
	validateImagesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseImpl) ArchiveRelease(arg1 *release.ArchiverOptions) error {
	fake.archiveReleaseMutex.Lock()
	ret, specificReturn := fake.archiveReleaseReturnsOnCall[len(fake.archiveReleaseArgsForCall)]
	fake.archiveReleaseArgsForCall = append(fake.archiveReleaseArgsForCall, struct {
		arg1 *release.ArchiverOptions
	}{arg1})
	stub := fake.ArchiveReleaseStub
	fakeReturns := fake.archiveReleaseReturns
	fake.recordInvocation("ArchiveRelease", []interface{}{arg1})
	fake.archiveReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) ArchiveReleaseCallCount() int {
	fake.archiveReleaseMutex.RLock()
	defer fake.archiveReleaseMutex.RUnlock()
	return len(fake.archiveReleaseArgsForCall)
}

func (fake *FakeReleaseImpl) ArchiveReleaseCalls(stub func(*release.ArchiverOptions) error) {
	fake.archiveReleaseMutex.Lock()
	defer fake.archiveReleaseMutex.Unlock()
	fake.ArchiveReleaseStub = stub
}

func (fake *FakeReleaseImpl) ArchiveReleaseArgsForCall(i int) *release.ArchiverOptions {
	fake.archiveReleaseMutex.RLock()
	defer fake.archiveReleaseMutex.RUnlock()
	argsForCall := fake.archiveReleaseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) ArchiveReleaseReturns(result1 error) {
	fake.archiveReleaseMutex.Lock()
	defer fake.archiveReleaseMutex.Unlock()
	fake.ArchiveReleaseStub = nil
	fake.archiveReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ArchiveReleaseReturnsOnCall(i int, result1 error) {
	fake.archiveReleaseMutex.Lock()
	defer fake.archiveReleaseMutex.Unlock()
	fake.ArchiveReleaseStub = nil
	if fake.archiveReleaseReturnsOnCall == nil {
		fake.archiveReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.archiveReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) BranchNeedsCreation(arg1 string, arg2 string, arg3 semver.Version) (bool, error) {
	fake.branchNeedsCreationMutex.Lock()
	ret, specificReturn := fake.branchNeedsCreationReturnsOnCall[len(fake.branchNeedsCreationArgsForCall)]
	fake.branchNeedsCreationArgsForCall = append(fake.branchNeedsCreationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 semver.Version
	}{arg1, arg2, arg3})
	stub := fake.BranchNeedsCreationStub
	fakeReturns := fake.branchNeedsCreationReturns
	fake.recordInvocation("BranchNeedsCreation", []interface{}{arg1, arg2, arg3})
	fake.branchNeedsCreationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseImpl) BranchNeedsCreationCallCount() int {
	fake.branchNeedsCreationMutex.RLock()
	defer fake.branchNeedsCreationMutex.RUnlock()
	return len(fake.branchNeedsCreationArgsForCall)
}

func (fake *FakeReleaseImpl) BranchNeedsCreationCalls(stub func(string, string, semver.Version) (bool, error)) {
	fake.branchNeedsCreationMutex.Lock()
	defer fake.branchNeedsCreationMutex.Unlock()
	fake.BranchNeedsCreationStub = stub
}

func (fake *FakeReleaseImpl) BranchNeedsCreationArgsForCall(i int) (string, string, semver.Version) {
	fake.branchNeedsCreationMutex.RLock()
	defer fake.branchNeedsCreationMutex.RUnlock()
	argsForCall := fake.branchNeedsCreationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) BranchNeedsCreationReturns(result1 bool, result2 error) {
	fake.branchNeedsCreationMutex.Lock()
	defer fake.branchNeedsCreationMutex.Unlock()
	fake.BranchNeedsCreationStub = nil
	fake.branchNeedsCreationReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) BranchNeedsCreationReturnsOnCall(i int, result1 bool, result2 error) {
	fake.branchNeedsCreationMutex.Lock()
	defer fake.branchNeedsCreationMutex.Unlock()
	fake.BranchNeedsCreationStub = nil
	if fake.branchNeedsCreationReturnsOnCall == nil {
		fake.branchNeedsCreationReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.branchNeedsCreationReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) CheckPrerequisites() error {
	fake.checkPrerequisitesMutex.Lock()
	ret, specificReturn := fake.checkPrerequisitesReturnsOnCall[len(fake.checkPrerequisitesArgsForCall)]
	fake.checkPrerequisitesArgsForCall = append(fake.checkPrerequisitesArgsForCall, struct {
	}{})
	stub := fake.CheckPrerequisitesStub
	fakeReturns := fake.checkPrerequisitesReturns
	fake.recordInvocation("CheckPrerequisites", []interface{}{})
	fake.checkPrerequisitesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CheckPrerequisitesCallCount() int {
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	return len(fake.checkPrerequisitesArgsForCall)
}

func (fake *FakeReleaseImpl) CheckPrerequisitesCalls(stub func() error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = stub
}

func (fake *FakeReleaseImpl) CheckPrerequisitesReturns(result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	fake.checkPrerequisitesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CheckPrerequisitesReturnsOnCall(i int, result1 error) {
	fake.checkPrerequisitesMutex.Lock()
	defer fake.checkPrerequisitesMutex.Unlock()
	fake.CheckPrerequisitesStub = nil
	if fake.checkPrerequisitesReturnsOnCall == nil {
		fake.checkPrerequisitesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkPrerequisitesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CheckReleaseBucket(arg1 *build.Options) error {
	fake.checkReleaseBucketMutex.Lock()
	ret, specificReturn := fake.checkReleaseBucketReturnsOnCall[len(fake.checkReleaseBucketArgsForCall)]
	fake.checkReleaseBucketArgsForCall = append(fake.checkReleaseBucketArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.CheckReleaseBucketStub
	fakeReturns := fake.checkReleaseBucketReturns
	fake.recordInvocation("CheckReleaseBucket", []interface{}{arg1})
	fake.checkReleaseBucketMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CheckReleaseBucketCallCount() int {
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	return len(fake.checkReleaseBucketArgsForCall)
}

func (fake *FakeReleaseImpl) CheckReleaseBucketCalls(stub func(*build.Options) error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = stub
}

func (fake *FakeReleaseImpl) CheckReleaseBucketArgsForCall(i int) *build.Options {
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	argsForCall := fake.checkReleaseBucketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) CheckReleaseBucketReturns(result1 error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = nil
	fake.checkReleaseBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CheckReleaseBucketReturnsOnCall(i int, result1 error) {
	fake.checkReleaseBucketMutex.Lock()
	defer fake.checkReleaseBucketMutex.Unlock()
	fake.CheckReleaseBucketStub = nil
	if fake.checkReleaseBucketReturnsOnCall == nil {
		fake.checkReleaseBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReleaseBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CopyStagedFromGCS(arg1 *build.Options, arg2 string, arg3 string) error {
	fake.copyStagedFromGCSMutex.Lock()
	ret, specificReturn := fake.copyStagedFromGCSReturnsOnCall[len(fake.copyStagedFromGCSArgsForCall)]
	fake.copyStagedFromGCSArgsForCall = append(fake.copyStagedFromGCSArgsForCall, struct {
		arg1 *build.Options
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CopyStagedFromGCSStub
	fakeReturns := fake.copyStagedFromGCSReturns
	fake.recordInvocation("CopyStagedFromGCS", []interface{}{arg1, arg2, arg3})
	fake.copyStagedFromGCSMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSCallCount() int {
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	return len(fake.copyStagedFromGCSArgsForCall)
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSCalls(stub func(*build.Options, string, string) error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = stub
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSArgsForCall(i int) (*build.Options, string, string) {
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	argsForCall := fake.copyStagedFromGCSArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSReturns(result1 error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = nil
	fake.copyStagedFromGCSReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CopyStagedFromGCSReturnsOnCall(i int, result1 error) {
	fake.copyStagedFromGCSMutex.Lock()
	defer fake.copyStagedFromGCSMutex.Unlock()
	fake.CopyStagedFromGCSStub = nil
	if fake.copyStagedFromGCSReturnsOnCall == nil {
		fake.copyStagedFromGCSReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyStagedFromGCSReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CreateAnnouncement(arg1 *announce.Options) error {
	fake.createAnnouncementMutex.Lock()
	ret, specificReturn := fake.createAnnouncementReturnsOnCall[len(fake.createAnnouncementArgsForCall)]
	fake.createAnnouncementArgsForCall = append(fake.createAnnouncementArgsForCall, struct {
		arg1 *announce.Options
	}{arg1})
	stub := fake.CreateAnnouncementStub
	fakeReturns := fake.createAnnouncementReturns
	fake.recordInvocation("CreateAnnouncement", []interface{}{arg1})
	fake.createAnnouncementMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) CreateAnnouncementCallCount() int {
	fake.createAnnouncementMutex.RLock()
	defer fake.createAnnouncementMutex.RUnlock()
	return len(fake.createAnnouncementArgsForCall)
}

func (fake *FakeReleaseImpl) CreateAnnouncementCalls(stub func(*announce.Options) error) {
	fake.createAnnouncementMutex.Lock()
	defer fake.createAnnouncementMutex.Unlock()
	fake.CreateAnnouncementStub = stub
}

func (fake *FakeReleaseImpl) CreateAnnouncementArgsForCall(i int) *announce.Options {
	fake.createAnnouncementMutex.RLock()
	defer fake.createAnnouncementMutex.RUnlock()
	argsForCall := fake.createAnnouncementArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) CreateAnnouncementReturns(result1 error) {
	fake.createAnnouncementMutex.Lock()
	defer fake.createAnnouncementMutex.Unlock()
	fake.CreateAnnouncementStub = nil
	fake.createAnnouncementReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) CreateAnnouncementReturnsOnCall(i int, result1 error) {
	fake.createAnnouncementMutex.Lock()
	defer fake.createAnnouncementMutex.Unlock()
	fake.CreateAnnouncementStub = nil
	if fake.createAnnouncementReturnsOnCall == nil {
		fake.createAnnouncementReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createAnnouncementReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) GenerateReleaseVersion(arg1 string, arg2 string, arg3 string, arg4 bool) (*release.Versions, error) {
	fake.generateReleaseVersionMutex.Lock()
	ret, specificReturn := fake.generateReleaseVersionReturnsOnCall[len(fake.generateReleaseVersionArgsForCall)]
	fake.generateReleaseVersionArgsForCall = append(fake.generateReleaseVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.GenerateReleaseVersionStub
	fakeReturns := fake.generateReleaseVersionReturns
	fake.recordInvocation("GenerateReleaseVersion", []interface{}{arg1, arg2, arg3, arg4})
	fake.generateReleaseVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionCallCount() int {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	return len(fake.generateReleaseVersionArgsForCall)
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionCalls(stub func(string, string, string, bool) (*release.Versions, error)) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = stub
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionArgsForCall(i int) (string, string, string, bool) {
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	argsForCall := fake.generateReleaseVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionReturns(result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	fake.generateReleaseVersionReturns = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) GenerateReleaseVersionReturnsOnCall(i int, result1 *release.Versions, result2 error) {
	fake.generateReleaseVersionMutex.Lock()
	defer fake.generateReleaseVersionMutex.Unlock()
	fake.GenerateReleaseVersionStub = nil
	if fake.generateReleaseVersionReturnsOnCall == nil {
		fake.generateReleaseVersionReturnsOnCall = make(map[int]struct {
			result1 *release.Versions
			result2 error
		})
	}
	fake.generateReleaseVersionReturnsOnCall[i] = struct {
		result1 *release.Versions
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) NewGitPusher(arg1 *release.GitObjectPusherOptions) (*release.GitObjectPusher, error) {
	fake.newGitPusherMutex.Lock()
	ret, specificReturn := fake.newGitPusherReturnsOnCall[len(fake.newGitPusherArgsForCall)]
	fake.newGitPusherArgsForCall = append(fake.newGitPusherArgsForCall, struct {
		arg1 *release.GitObjectPusherOptions
	}{arg1})
	stub := fake.NewGitPusherStub
	fakeReturns := fake.newGitPusherReturns
	fake.recordInvocation("NewGitPusher", []interface{}{arg1})
	fake.newGitPusherMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseImpl) NewGitPusherCallCount() int {
	fake.newGitPusherMutex.RLock()
	defer fake.newGitPusherMutex.RUnlock()
	return len(fake.newGitPusherArgsForCall)
}

func (fake *FakeReleaseImpl) NewGitPusherCalls(stub func(*release.GitObjectPusherOptions) (*release.GitObjectPusher, error)) {
	fake.newGitPusherMutex.Lock()
	defer fake.newGitPusherMutex.Unlock()
	fake.NewGitPusherStub = stub
}

func (fake *FakeReleaseImpl) NewGitPusherArgsForCall(i int) *release.GitObjectPusherOptions {
	fake.newGitPusherMutex.RLock()
	defer fake.newGitPusherMutex.RUnlock()
	argsForCall := fake.newGitPusherArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) NewGitPusherReturns(result1 *release.GitObjectPusher, result2 error) {
	fake.newGitPusherMutex.Lock()
	defer fake.newGitPusherMutex.Unlock()
	fake.NewGitPusherStub = nil
	fake.newGitPusherReturns = struct {
		result1 *release.GitObjectPusher
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) NewGitPusherReturnsOnCall(i int, result1 *release.GitObjectPusher, result2 error) {
	fake.newGitPusherMutex.Lock()
	defer fake.newGitPusherMutex.Unlock()
	fake.NewGitPusherStub = nil
	if fake.newGitPusherReturnsOnCall == nil {
		fake.newGitPusherReturnsOnCall = make(map[int]struct {
			result1 *release.GitObjectPusher
			result2 error
		})
	}
	fake.newGitPusherReturnsOnCall[i] = struct {
		result1 *release.GitObjectPusher
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseImpl) PrepareWorkspaceRelease(arg1 string, arg2 string) error {
	fake.prepareWorkspaceReleaseMutex.Lock()
	ret, specificReturn := fake.prepareWorkspaceReleaseReturnsOnCall[len(fake.prepareWorkspaceReleaseArgsForCall)]
	fake.prepareWorkspaceReleaseArgsForCall = append(fake.prepareWorkspaceReleaseArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.PrepareWorkspaceReleaseStub
	fakeReturns := fake.prepareWorkspaceReleaseReturns
	fake.recordInvocation("PrepareWorkspaceRelease", []interface{}{arg1, arg2})
	fake.prepareWorkspaceReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCallCount() int {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	return len(fake.prepareWorkspaceReleaseArgsForCall)
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCalls(stub func(string, string) error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = stub
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseArgsForCall(i int) (string, string) {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	argsForCall := fake.prepareWorkspaceReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturns(result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	fake.prepareWorkspaceReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturnsOnCall(i int, result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	if fake.prepareWorkspaceReleaseReturnsOnCall == nil {
		fake.prepareWorkspaceReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.prepareWorkspaceReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PublishVersion(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 []string, arg7 bool, arg8 bool) error {
	var arg6Copy []string
	if arg6 != nil {
		arg6Copy = make([]string, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.publishVersionMutex.Lock()
	ret, specificReturn := fake.publishVersionReturnsOnCall[len(fake.publishVersionArgsForCall)]
	fake.publishVersionArgsForCall = append(fake.publishVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 []string
		arg7 bool
		arg8 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8})
	stub := fake.PublishVersionStub
	fakeReturns := fake.publishVersionReturns
	fake.recordInvocation("PublishVersion", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6Copy, arg7, arg8})
	fake.publishVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PublishVersionCallCount() int {
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	return len(fake.publishVersionArgsForCall)
}

func (fake *FakeReleaseImpl) PublishVersionCalls(stub func(string, string, string, string, string, []string, bool, bool) error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = stub
}

func (fake *FakeReleaseImpl) PublishVersionArgsForCall(i int) (string, string, string, string, string, []string, bool, bool) {
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	argsForCall := fake.publishVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8
}

func (fake *FakeReleaseImpl) PublishVersionReturns(result1 error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = nil
	fake.publishVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PublishVersionReturnsOnCall(i int, result1 error) {
	fake.publishVersionMutex.Lock()
	defer fake.publishVersionMutex.Unlock()
	fake.PublishVersionStub = nil
	if fake.publishVersionReturnsOnCall == nil {
		fake.publishVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushBranches(arg1 *release.GitObjectPusher, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.pushBranchesMutex.Lock()
	ret, specificReturn := fake.pushBranchesReturnsOnCall[len(fake.pushBranchesArgsForCall)]
	fake.pushBranchesArgsForCall = append(fake.pushBranchesArgsForCall, struct {
		arg1 *release.GitObjectPusher
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.PushBranchesStub
	fakeReturns := fake.pushBranchesReturns
	fake.recordInvocation("PushBranches", []interface{}{arg1, arg2Copy})
	fake.pushBranchesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PushBranchesCallCount() int {
	fake.pushBranchesMutex.RLock()
	defer fake.pushBranchesMutex.RUnlock()
	return len(fake.pushBranchesArgsForCall)
}

func (fake *FakeReleaseImpl) PushBranchesCalls(stub func(*release.GitObjectPusher, []string) error) {
	fake.pushBranchesMutex.Lock()
	defer fake.pushBranchesMutex.Unlock()
	fake.PushBranchesStub = stub
}

func (fake *FakeReleaseImpl) PushBranchesArgsForCall(i int) (*release.GitObjectPusher, []string) {
	fake.pushBranchesMutex.RLock()
	defer fake.pushBranchesMutex.RUnlock()
	argsForCall := fake.pushBranchesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseImpl) PushBranchesReturns(result1 error) {
	fake.pushBranchesMutex.Lock()
	defer fake.pushBranchesMutex.Unlock()
	fake.PushBranchesStub = nil
	fake.pushBranchesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushBranchesReturnsOnCall(i int, result1 error) {
	fake.pushBranchesMutex.Lock()
	defer fake.pushBranchesMutex.Unlock()
	fake.PushBranchesStub = nil
	if fake.pushBranchesReturnsOnCall == nil {
		fake.pushBranchesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushBranchesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushMainBranch(arg1 *release.GitObjectPusher) error {
	fake.pushMainBranchMutex.Lock()
	ret, specificReturn := fake.pushMainBranchReturnsOnCall[len(fake.pushMainBranchArgsForCall)]
	fake.pushMainBranchArgsForCall = append(fake.pushMainBranchArgsForCall, struct {
		arg1 *release.GitObjectPusher
	}{arg1})
	stub := fake.PushMainBranchStub
	fakeReturns := fake.pushMainBranchReturns
	fake.recordInvocation("PushMainBranch", []interface{}{arg1})
	fake.pushMainBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PushMainBranchCallCount() int {
	fake.pushMainBranchMutex.RLock()
	defer fake.pushMainBranchMutex.RUnlock()
	return len(fake.pushMainBranchArgsForCall)
}

func (fake *FakeReleaseImpl) PushMainBranchCalls(stub func(*release.GitObjectPusher) error) {
	fake.pushMainBranchMutex.Lock()
	defer fake.pushMainBranchMutex.Unlock()
	fake.PushMainBranchStub = stub
}

func (fake *FakeReleaseImpl) PushMainBranchArgsForCall(i int) *release.GitObjectPusher {
	fake.pushMainBranchMutex.RLock()
	defer fake.pushMainBranchMutex.RUnlock()
	argsForCall := fake.pushMainBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) PushMainBranchReturns(result1 error) {
	fake.pushMainBranchMutex.Lock()
	defer fake.pushMainBranchMutex.Unlock()
	fake.PushMainBranchStub = nil
	fake.pushMainBranchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushMainBranchReturnsOnCall(i int, result1 error) {
	fake.pushMainBranchMutex.Lock()
	defer fake.pushMainBranchMutex.Unlock()
	fake.PushMainBranchStub = nil
	if fake.pushMainBranchReturnsOnCall == nil {
		fake.pushMainBranchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushMainBranchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushTags(arg1 *release.GitObjectPusher, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.pushTagsMutex.Lock()
	ret, specificReturn := fake.pushTagsReturnsOnCall[len(fake.pushTagsArgsForCall)]
	fake.pushTagsArgsForCall = append(fake.pushTagsArgsForCall, struct {
		arg1 *release.GitObjectPusher
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.PushTagsStub
	fakeReturns := fake.pushTagsReturns
	fake.recordInvocation("PushTags", []interface{}{arg1, arg2Copy})
	fake.pushTagsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PushTagsCallCount() int {
	fake.pushTagsMutex.RLock()
	defer fake.pushTagsMutex.RUnlock()
	return len(fake.pushTagsArgsForCall)
}

func (fake *FakeReleaseImpl) PushTagsCalls(stub func(*release.GitObjectPusher, []string) error) {
	fake.pushTagsMutex.Lock()
	defer fake.pushTagsMutex.Unlock()
	fake.PushTagsStub = stub
}

func (fake *FakeReleaseImpl) PushTagsArgsForCall(i int) (*release.GitObjectPusher, []string) {
	fake.pushTagsMutex.RLock()
	defer fake.pushTagsMutex.RUnlock()
	argsForCall := fake.pushTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseImpl) PushTagsReturns(result1 error) {
	fake.pushTagsMutex.Lock()
	defer fake.pushTagsMutex.Unlock()
	fake.PushTagsStub = nil
	fake.pushTagsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PushTagsReturnsOnCall(i int, result1 error) {
	fake.pushTagsMutex.Lock()
	defer fake.pushTagsMutex.Unlock()
	fake.PushTagsStub = nil
	if fake.pushTagsReturnsOnCall == nil {
		fake.pushTagsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushTagsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) Submit(arg1 *gcb.Options) error {
	fake.submitMutex.Lock()
	ret, specificReturn := fake.submitReturnsOnCall[len(fake.submitArgsForCall)]
	fake.submitArgsForCall = append(fake.submitArgsForCall, struct {
		arg1 *gcb.Options
	}{arg1})
	stub := fake.SubmitStub
	fakeReturns := fake.submitReturns
	fake.recordInvocation("Submit", []interface{}{arg1})
	fake.submitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) SubmitCallCount() int {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	return len(fake.submitArgsForCall)
}

func (fake *FakeReleaseImpl) SubmitCalls(stub func(*gcb.Options) error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = stub
}

func (fake *FakeReleaseImpl) SubmitArgsForCall(i int) *gcb.Options {
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	argsForCall := fake.submitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) SubmitReturns(result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	fake.submitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) SubmitReturnsOnCall(i int, result1 error) {
	fake.submitMutex.Lock()
	defer fake.submitMutex.Unlock()
	fake.SubmitStub = nil
	if fake.submitReturnsOnCall == nil {
		fake.submitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.submitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ToFile(arg1 string) error {
	fake.toFileMutex.Lock()
	ret, specificReturn := fake.toFileReturnsOnCall[len(fake.toFileArgsForCall)]
	fake.toFileArgsForCall = append(fake.toFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ToFileStub
	fakeReturns := fake.toFileReturns
	fake.recordInvocation("ToFile", []interface{}{arg1})
	fake.toFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) ToFileCallCount() int {
	fake.toFileMutex.RLock()
	defer fake.toFileMutex.RUnlock()
	return len(fake.toFileArgsForCall)
}

func (fake *FakeReleaseImpl) ToFileCalls(stub func(string) error) {
	fake.toFileMutex.Lock()
	defer fake.toFileMutex.Unlock()
	fake.ToFileStub = stub
}

func (fake *FakeReleaseImpl) ToFileArgsForCall(i int) string {
	fake.toFileMutex.RLock()
	defer fake.toFileMutex.RUnlock()
	argsForCall := fake.toFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseImpl) ToFileReturns(result1 error) {
	fake.toFileMutex.Lock()
	defer fake.toFileMutex.Unlock()
	fake.ToFileStub = nil
	fake.toFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ToFileReturnsOnCall(i int, result1 error) {
	fake.toFileMutex.Lock()
	defer fake.toFileMutex.Unlock()
	fake.ToFileStub = nil
	if fake.toFileReturnsOnCall == nil {
		fake.toFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.toFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ValidateImages(arg1 string, arg2 string, arg3 string) error {
	fake.validateImagesMutex.Lock()
	ret, specificReturn := fake.validateImagesReturnsOnCall[len(fake.validateImagesArgsForCall)]
	fake.validateImagesArgsForCall = append(fake.validateImagesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ValidateImagesStub
	fakeReturns := fake.validateImagesReturns
	fake.recordInvocation("ValidateImages", []interface{}{arg1, arg2, arg3})
	fake.validateImagesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) ValidateImagesCallCount() int {
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	return len(fake.validateImagesArgsForCall)
}

func (fake *FakeReleaseImpl) ValidateImagesCalls(stub func(string, string, string) error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = stub
}

func (fake *FakeReleaseImpl) ValidateImagesArgsForCall(i int) (string, string, string) {
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	argsForCall := fake.validateImagesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) ValidateImagesReturns(result1 error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = nil
	fake.validateImagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) ValidateImagesReturnsOnCall(i int, result1 error) {
	fake.validateImagesMutex.Lock()
	defer fake.validateImagesMutex.Unlock()
	fake.ValidateImagesStub = nil
	if fake.validateImagesReturnsOnCall == nil {
		fake.validateImagesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateImagesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.archiveReleaseMutex.RLock()
	defer fake.archiveReleaseMutex.RUnlock()
	fake.branchNeedsCreationMutex.RLock()
	defer fake.branchNeedsCreationMutex.RUnlock()
	fake.checkPrerequisitesMutex.RLock()
	defer fake.checkPrerequisitesMutex.RUnlock()
	fake.checkReleaseBucketMutex.RLock()
	defer fake.checkReleaseBucketMutex.RUnlock()
	fake.copyStagedFromGCSMutex.RLock()
	defer fake.copyStagedFromGCSMutex.RUnlock()
	fake.createAnnouncementMutex.RLock()
	defer fake.createAnnouncementMutex.RUnlock()
	fake.generateReleaseVersionMutex.RLock()
	defer fake.generateReleaseVersionMutex.RUnlock()
	fake.newGitPusherMutex.RLock()
	defer fake.newGitPusherMutex.RUnlock()
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	fake.publishVersionMutex.RLock()
	defer fake.publishVersionMutex.RUnlock()
	fake.pushBranchesMutex.RLock()
	defer fake.pushBranchesMutex.RUnlock()
	fake.pushMainBranchMutex.RLock()
	defer fake.pushMainBranchMutex.RUnlock()
	fake.pushTagsMutex.RLock()
	defer fake.pushTagsMutex.RUnlock()
	fake.submitMutex.RLock()
	defer fake.submitMutex.RUnlock()
	fake.toFileMutex.RLock()
	defer fake.toFileMutex.RUnlock()
	fake.validateImagesMutex.RLock()
	defer fake.validateImagesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseImpl) recordInvocation(key string, args []interface{}) {
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
