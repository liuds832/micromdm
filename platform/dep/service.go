package dep

import (
	"context"
	"sync"

	"github.com/liuds832/micromdm/dep"
	"github.com/liuds832/micromdm/platform/pubsub"
)

type Service interface {
	DefineProfile(ctx context.Context, p *dep.Profile) (*dep.ProfileResponse, error)
	AssignProfile(ctx context.Context, uuid string, serials ...string) (*dep.ProfileResponse, error)
	RemoveProfile(ctx context.Context, serials ...string) (map[string]string, error)
	GetAccountInfo(ctx context.Context) (*dep.Account, error)
	GetDeviceDetails(ctx context.Context, serials []string) (*dep.DeviceDetailsResponse, error)
	FetchProfile(ctx context.Context, uuid string) (*dep.Profile, error)
	SetActivationLock(ctx context.Context, p *dep.ActivationLockRequest) (*dep.ActivationLockResponse, error)
}

type DEPClient interface {
	DefineProfile(*dep.Profile) (*dep.ProfileResponse, error)
	AssignProfile(string, ...string) (*dep.ProfileResponse, error)
	RemoveProfile(...string) (map[string]string, error)
	FetchProfile(string) (*dep.Profile, error)
	Account() (*dep.Account, error)
	DeviceDetails(...string) (*dep.DeviceDetailsResponse, error)
	ActivationLock(*dep.ActivationLockRequest) (*dep.ActivationLockResponse, error)
}

type DEPService struct {
	mtx        sync.RWMutex
	client     DEPClient
	subscriber pubsub.Subscriber
}

func (svc *DEPService) Run() error {
	return svc.watchTokenUpdates(svc.subscriber)
}

func New(client DEPClient, subscriber pubsub.Subscriber) *DEPService {
	return &DEPService{client: client, subscriber: subscriber}
}
