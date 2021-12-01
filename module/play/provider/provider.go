package provider

import (
	"github.com/bytelang/kplayer/module"
	kptypes "github.com/bytelang/kplayer/types"
	"github.com/bytelang/kplayer/types/config"
	kpproto "github.com/bytelang/kplayer/types/core/proto"
	svrproto "github.com/bytelang/kplayer/types/server"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	defaultRPCAddress string = "127.0.0.1"
	defaultRPCPort    uint32 = 4156
)

type ProviderI interface {
	GetStartPoint() uint32
	GetPlayModel() string
	GetRPCParams() config.Rpc
	PlayStop(args *svrproto.PlayStopArgs) (*svrproto.PlayStopReply, error)
	PlayPause(args *svrproto.PlayPauseArgs) (*svrproto.PlayPauseReply, error)
	PlaySkip(args *svrproto.PlaySkipArgs) (*svrproto.PlaySkipReply, error)
	PlayContinue(args *svrproto.PlayContinueArgs) (*svrproto.PlayContinueReply, error)
	PlayDuration(args *svrproto.PlayDurationArgs) (*svrproto.PlayDurationReply, error)
	PlayInformation(args *svrproto.PlayInformationArgs) (*svrproto.PlayInformationReply, error)
}

var _ ProviderI = &Provider{}

// Provider play module provider
type Provider struct {
	config *config.Play
	module.ModuleKeeper

	// module member
	startTime time.Time
}

// NewProvider return provider
func NewProvider() *Provider {
	return &Provider{
		config: &config.Play{},
	}
}

func (p *Provider) GetConfig() *config.Play {
	return p.config
}

func (p *Provider) setConfig(config config.Play) {
	p.config = &config
}

// InitConfig set module config on kplayer started
func (p *Provider) InitModule(ctx *kptypes.ClientContext, cfg config.Play) {
	if cfg.Rpc == nil {
		cfg.Rpc = &config.Rpc{}
	}

	if cfg.Rpc.Address == "" {
		cfg.Rpc.Address = defaultRPCAddress
	}
	if cfg.Rpc.Port == 0 {
		cfg.Rpc.Port = defaultRPCPort
	}

	p.setConfig(cfg)
}

func (p *Provider) ParseMessage(message *kpproto.KPMessage) {
	switch message.Action {
	case kpproto.EVENT_MESSAGE_ACTION_PLAYER_STARTED:
		log.Info("kplayer start success")
		p.startTime = time.Now()
	}
}

func (p *Provider) ValidateConfig() error {
	return nil
}

func (p *Provider) GetStartPoint() uint32 {
	return p.config.StartPoint
}

func (p *Provider) GetPlayModel() string {
	return p.config.PlayModel
}
