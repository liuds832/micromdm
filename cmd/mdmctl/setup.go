package main

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/liuds832/micromdm/platform/appstore"
	"github.com/liuds832/micromdm/platform/blueprint"
	"github.com/liuds832/micromdm/platform/config"
	"github.com/liuds832/micromdm/platform/dep"
	"github.com/liuds832/micromdm/platform/dep/sync"
	"github.com/liuds832/micromdm/platform/device"
	"github.com/liuds832/micromdm/platform/profile"
	"github.com/liuds832/micromdm/platform/remove"
	"github.com/liuds832/micromdm/platform/user"
)

type remoteServices struct {
	profilesvc   profile.Service
	blueprintsvc blueprint.Service
	blocksvc     remove.Service
	usersvc      user.Service
	devicesvc    device.Service
	configsvc    config.Service
	appsvc       appstore.Service
	depsvc       dep.Service
	depsyncsvc   sync.Service
}

func setupClient(logger log.Logger) (*remoteServices, error) {
	cfg, err := LoadServerConfig()
	if err != nil {
		return nil, err
	}

	profilesvc, err := profile.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	blueprintsvc, err := blueprint.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	blocksvc, err := remove.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	usersvc, err := user.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	devicesvc, err := device.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	configsvc, err := config.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	appsvc, err := appstore.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	depsvc, err := dep.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	depsyncsvc, err := sync.NewHTTPClient(
		cfg.ServerURL, cfg.APIToken, logger,
		httptransport.SetClient(skipVerifyHTTPClient(cfg.SkipVerify)))
	if err != nil {
		return nil, err
	}

	return &remoteServices{
		profilesvc:   profilesvc,
		blueprintsvc: blueprintsvc,
		blocksvc:     blocksvc,
		usersvc:      usersvc,
		devicesvc:    devicesvc,
		configsvc:    configsvc,
		appsvc:       appsvc,
		depsvc:       depsvc,
		depsyncsvc:   depsyncsvc,
	}, nil
}
