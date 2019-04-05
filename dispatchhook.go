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

func runDispatchHook(msg *win.MSG) bool {
	var hooks []func(msg *win.MSG) bool
	dispatchHook.m.RLock()
	hooks = append(hooks, dispatchHook.hooks...)
	dispatchHook.m.RUnlock()
	for _, h := range hooks {
		next := h(msg)
		if !next {
			return false
		}
	}
	return true
}
