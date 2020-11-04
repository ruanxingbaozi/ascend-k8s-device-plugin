/*************************************************************************
 * Copyright (C) [2019] by Cambricon, Inc. All rights reserved
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *************************************************************************/

package dcmi

// #include "dcmi_dl.h"
import "C"

type Device struct {
	handle

	UUID string
	Path string
}

type DeviceMemory struct {
	Used *uint64
	Free *uint64
}

type UtilizationInfo struct {
	Board *uint
}

type DeviceStatus struct {
	Memory      DeviceMemory
	Utilization UtilizationInfo
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func Init() error {
	return dcmiInit()
}

func GetCardList() ([]uint, error) {
	return cardGetList()
}

func GetDeviceInCard(cardId uint) (uint, error) {
	return deviceInCard(cardId)
}
func GetDeviceHBMInfo(cardId uint) (DeviceMemory, error) {
	return deviceGetHBMInfo(cardId)
}
func GetDeviceHealth(cardId uint) (uint, error) {
	return deviceGetHealth(cardId)
}

//func NewDeviceLite(idx uint) (device *Device, err error) {
//	device = nil
//	h, err := deviceGetHandleByIndex(idx)
//
//	if err != nil {
//		return
//	}
//
//	uuid, err := h.deviceGetUUID()
//
//	if err != nil {
//		return
//	}
//
//	path := h.PATH
//	device = &Device{
//		handle: h,
//		UUID:   uuid,
//		Path:   path,
//	}
//	return
//}
//
//func (d *Device) DeviceHealthCheckState(delayTime int) (int, error) {
//	return d.handle.deviceHealthCheckState(delayTime)
//}
//
//func (d *Device) DeviceAllRunningProcessInfo() []*ProcessInfo {
//	pids, mems, err := d.handle.deviceProcessInfo()
//	assert(err)
//
//	processInfos := []*ProcessInfo{}
//	for i := 0; i < len(pids); i++ {
//		name, err := processName(pids[i])
//
//		if err != nil {
//			return processInfos
//		}
//		p := &ProcessInfo{
//			Pid:                pids[i],
//			Command:            name,
//			PhysicalMemoryUsed: mems[i],
//		}
//		processInfos = append(processInfos, p)
//	}
//	return processInfos
//}
//
//func (d *Device) Status() (status *DeviceStatus, err error) {
//	board_u, err := d.handle.deviceGetBoardUtilization()
//	assert(err)
//
//	_, devMem, err := d.handle.deviceGetMemoryInfo()
//	assert(err)
//
//	status = &DeviceStatus{
//		Memory: devMem,
//		Utilization: UtilizationInfo{
//			Board: board_u,
//		},
//	}
//	return
//}

func Shutdown() error {
	return Release()
}
