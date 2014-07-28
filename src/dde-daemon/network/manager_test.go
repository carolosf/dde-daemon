/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package network

import (
	// liblogger "pkg.linuxdeepin.com/lib/logger"
	. "launchpad.net/gocheck"
	// "testing"
	// "time"
)

func init() {
	manager := &Manager{}
	Suite(manager)
}

// func (m *Manager) TestRemoveDevice(c *C) {
// 	devs := make([]*device, 0)
// 	devs = append(devs, &device{Path: "path1", State: 0})
// 	devs = append(devs, &device{Path: "path2", State: 0})
// 	devs = m.doRemoveDevice(devs, "path1")
// 	c.Check(len(devs), Equals, 1)
// }