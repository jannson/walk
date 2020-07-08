// +build windows

package walk

import (
	"sync"

	"github.com/lxn/win"
)

var dispatchHook struct {
	m     sync.RWMutex
	hooks []func(msg *win.MSG) bool
}

func AddDispatchHook(hook func(msg *win.MSG) bool) {
	defer dispatchHook.m.Unlock()
	dispatchHook.m.Lock()
	dispatchHook.hooks = append(dispatchHook.hooks, hook)
}
