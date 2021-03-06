/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package inputdevices

import (
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/log"
)

//go:generate dbusutil-gen -type Mouse,Touchpad,TrackPoint,Wacom mouse.go touchpad.go trackpoint.go wacom.go

var (
	_manager *Manager
	logger   = log.NewLogger("daemon/inputdevices")
)

type Daemon struct {
	*loader.ModuleBase
}

func init() {
	loader.Register(NewInputdevicesDaemon(logger))
}
func NewInputdevicesDaemon(logger *log.Logger) *Daemon {
	var d = new(Daemon)
	d.ModuleBase = loader.NewModuleBase("inputdevices", d, logger)
	return d
}

func (*Daemon) GetDependencies() []string {
	return []string{}
}

func (*Daemon) Start() error {
	if _manager != nil {
		return nil
	}

	service := loader.GetService()
	_manager = NewManager(service)

	err := service.Export(dbusPath, _manager)
	if err != nil {
		return err
	}

	err = service.Export(kbdDBusPath, _manager.kbd)
	if err != nil {
		return err
	}

	err = service.Export(wacomDBusPath, _manager.wacom)
	if err != nil {
		return err
	}

	err = service.Export(touchPadDBusPath, _manager.tpad)
	if err != nil {
		return err
	}

	err = service.Export(mouseDBusPath, _manager.mouse, _manager.trackPoint)
	if err != nil {
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		return err
	}

	go func() {
		_manager.init()
		startDeviceListener()
	}()
	return nil
}

func (*Daemon) Stop() error {
	if _manager == nil {
		return nil
	}
	_manager = nil

	if _manager.wacom != nil {
		_manager.wacom.destroy()
		_manager.wacom = nil
	}

	// TODO endDeviceListener will be stuck
	endDeviceListener()
	return nil
}
