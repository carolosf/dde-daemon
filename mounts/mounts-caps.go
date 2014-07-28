/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
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

package mounts

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	CMD_DF = "/bin/df"
)

func getDiskCap(path string) (int64, int64) {
	contents := []byte{}
	for i := 0; i < 10; i++ {
		bytes, err := exec.Command(CMD_DF).Output()
		if err != nil || len(string(bytes)) < 1 {
			if i == 9 {
				logger.Warning("Exec 'df -h' failed:", err)
				return 0, 0
			}
			<-time.After(time.Second * 1)
		} else {
			contents = bytes
			break
		}
	}

	usedSize := int64(0)
	totalSize := int64(0)
	outStrs := strings.Split(string(contents), "\n")
	for _, str := range outStrs {
		array := strings.Split(str, " ")
		rets := delSpaceElment(array)
		l := len(rets)
		if l <= 2 {
			break
		}

		isMatch := false
		for _, v := range rets {
			if path == v {
				isMatch = true
				usedSize, _ = strconv.ParseInt(rets[2], 10, 64)
				totalSize, _ = strconv.ParseInt(rets[1], 10, 64)
				break
			}
		}
		if isMatch {
			break
		}
	}

	return totalSize, usedSize
}

func delSpaceElment(strs []string) []string {
	rets := []string{}

	for _, v := range strs {
		if len(v) > 0 {
			rets = append(rets, v)
		}
	}

	return rets
}