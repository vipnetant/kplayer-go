package provider

import (
	moduletypes "github.com/bytelang/kplayer/types/module"
	"sync"
)

const (
	ModuleName = "plugin"
)

const (
	PluginNotFound       ResourceError = "plugin not found"
	PluginUniqueHasExist ResourceError = "plugin unique has exist"
)

type ResourceError string

func (r ResourceError) Error() string {
	return string(r)
}

type Plugins struct {
	plugins []moduletypes.Plugin
	lock    sync.Mutex
}

func (p *Plugins) GetPluginByUnique(unique string) (*moduletypes.Plugin, int, error) {
	for key, item := range p.plugins {
		if item.Unique == unique {
			return &p.plugins[key], key, nil
		}
	}

	return nil, 0, PluginNotFound
}

func (p *Plugins) Exist(unique string) bool {
	for _, item := range p.plugins {
		if item.Unique == unique {
			return true
		}
	}

	return false
}

func (ps *Plugins) RemovePluginByUnique(unique string) (*moduletypes.Plugin, error) {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	res, index, err := ps.GetPluginByUnique(unique)
	if res == nil {
		return nil, err
	}

	var newPlugins []moduletypes.Plugin
	newPlugins = append(newPlugins, ps.plugins[:index]...)
	newPlugins = append(newPlugins, ps.plugins[index+1:]...)

	ps.plugins = newPlugins

	return res, nil
}

func (ps *Plugins) AppendPlugin(plugin moduletypes.Plugin) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	res, _, _ := ps.GetPluginByUnique(plugin.Unique)
	if res != nil {
		return PluginUniqueHasExist
	}

	ps.plugins = append(ps.plugins, plugin)
	return nil
}
