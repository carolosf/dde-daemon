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

package dock

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/log"
)

type Daemon struct {
	*loader.ModuleBase
}

const moduleName = "dock"

func NewDaemon(logger *log.Logger) *Daemon {
	daemon := new(Daemon)
	daemon.ModuleBase = loader.NewModuleBase(moduleName, daemon, logger)
	return daemon
}

func (d *Daemon) Stop() error {
	if dockManager != nil {
		dockManager.destroy()
		dockManager = nil
	}

	if XU != nil {
		XU.Conn().Close()
		XU = nil
	}

	return nil
}

func (d *Daemon) startFailed() {
	d.Stop()
}

func (d *Daemon) Start() error {
	if dockManager != nil {
		return nil
	}

	var err error
	// init x conn
	XU, err = xgbutil.NewConn()
	if err != nil {
		d.startFailed()
		return err
	}

	initAtom()
	initDir()
	initPathDirCodeMap()

	service := loader.GetService()

	dockManager, err = newManager(service)
	if err != nil {
		d.startFailed()
		return err
	}

	dockManager.listenRootWindowPropertyChange()
	go xevent.Main(XU)

	err = service.RequestName(dbusServiceName)
	if err != nil {
		d.startFailed()
		return err
	}

	service.Emit(dockManager, "ServiceRestarted")
	return nil
}

func (d *Daemon) GetDependencies() []string {
	return []string{}
}

func (d *Daemon) Name() string {
	return moduleName
}
