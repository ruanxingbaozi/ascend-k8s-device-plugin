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

// #cgo CFLAGS: -I ./
// #cgo LDFLAGS: -ldl -Wl,--unresolved-symbols=ignore-in-object-files
// #include "dcmi_dl.h"
import "C"

import (
	"errors"
	"fmt"
)

const (
	VERSION = 3

	szProcs = 32
)

type handle struct {
	UUID  string
	MINOR uint
	PATH  string
}

func errorString(ret C.dcmiRet_t) error {
	if ret == C.DCMI_SUCCESS {
		return nil
	}
	// todo dcmi error info func 参数不匹配
	//err := C.GoString(C.cndevGetErrorString(ret))
	return fmt.Errorf("dcmi: %v", ret)
}

func init_() error {
	r := C.dcmiInit_dl()
	if r == C.DCMI_ERROR_UNINITIALIZED {
		return errors.New("could not load DCMI library")
	}
	return errorString(r)
}

func dcmiInit() error {
	return init_()
}

func release_() error {
	r := C.dcmiRelease_dl()
	return errorString(r)
}

func Release() error {
	return release_()
}

//int dcmi_get_card_num_list(int *card_num, int *card_list, int list_length);
func deviceGetCount() (uint, error) {
	ret := C.int(0)
	card_num := C.int(0)
	var card_list [8]C.int
	ret = C.dcmi_get_card_num_list(&card_num, &card_list[0], C.int(len(card_list)))
	if ret != 0 {
		fmt.Errorf("get card num list failed")
	}
	fmt.Printf("card list: %v\n", card_list)
	fmt.Println("*******************")
	// card list: [0 1 2 3 4 5 6 7 0 0]
	return 0, nil
}

//func deviceGetCardName(devId uint) (C.cndevNameEnum_t, error) {
//	var cardName C.cndevCardName_t
//	cardName.version = C.int(VERSION)
//	r := C.cndevGetCardName(&cardName, C.int(devId))
//	cardType := cardName.id
//	return cardType, errorString(r)
//}
//
//func deviceGetCardSN(devId uint) (C.cndevCardSN_t, error) {
//	var cardSN C.cndevCardSN_t
//
//	cardSN.version = C.int(VERSION)
//	r := C.cndevGetCardSN(&cardSN, C.int(devId))
//
//	return cardSN, errorString(r)
//}
//
//func deviceGetHandleByIndex(idx_uint uint) (handle, error) {
//	var h handle
//	var path string
//
//	cardSN, err := deviceGetCardSN(idx_uint)
//	if err != nil {
//		return h, err
//	}
//
//	cardUUID := fmt.Sprintf("%x", int(cardSN.sn))
//
//	// Type A or type B card has no SN code, fake one.
//	if cardUUID == "0" {
//		cardUUID = uuid.New().String()
//	}
//	cardUUID = "MLU-" + cardUUID
//
//	cardName, err := deviceGetCardName(idx_uint)
//
//	if err != nil {
//		return h, err
//	}
//
//	if cardName == C.MLU100 {
//		path = fmt.Sprintf("/dev/cambricon_c10Dev%d", idx_uint)
//	} else {
//		path = fmt.Sprintf("/dev/cambricon_dev%d", idx_uint)
//	}
//
//	h = handle{
//		UUID:  cardUUID,
//		MINOR: idx_uint,
//		PATH:  path,
//	}
//
//	return h, nil
//}
//
//func (h handle) deviceGetUUID() (string, error) {
//	var ret C.dcmiRet_t = C.DCMI_SUCCESS
//	return h.UUID, errorString(ret)
//}
//
//func (h handle) deviceGetPath() (string, error) {
//	var ret C.dcmiRet_t = C.DCMI_SUCCESS
//	return h.PATH, errorString(ret)
//}
//
////dcmiRet_t cndevGetCardHealthState(cndevCardHealthState_t* cardHealthState, int devId);
//func (h handle) deviceHealthCheckState(delayTime int) (int, error) {
//	var ret C.dcmiRet_t
//	var cardHealthState C.cndevCardHealthState_t
//	var healthCode int
//	cardHealthState.version = C.int(VERSION)
//	devId := C.int(h.MINOR)
//	// sleep some seconds
//	time.Sleep(time.Duration(delayTime) * time.Second)
//	ret = C.cndevGetCardHealthState(&cardHealthState, devId)
//	healthCode = int(cardHealthState.health)
//	return healthCode, errorString(ret)
//}
//
////dcmiRet_t cndevGetMemoryUsage(cndevMemoryInfo_t *memInfo, int devId);
//func (h handle) deviceGetMemoryInfo() (totalMem *uint64, devMem DeviceMemory, err error) {
//	var ret C.dcmiRet_t
//	var cndevMemoryInfo C.cndevMemoryInfo_t
//	cndevMemoryInfo.version = C.int(VERSION)
//	devId := C.int(h.MINOR)
//	ret = C.cndevGetMemoryUsage(&cndevMemoryInfo, devId)
//	totalMem = uint64Ptr(cndevMemoryInfo.PhysicalMemoryTotal)
//	usedMem := uint64Ptr(cndevMemoryInfo.PhysicalMemoryUsed)
//	freeMem := *totalMem - *usedMem
//	devMem = DeviceMemory{
//		Used: usedMem,
//		Free: &freeMem,
//	}
//	return totalMem, devMem, errorString(ret)
//}
//
////dcmiRet_t cndevGetDeviceUtilizationInfo(cndevUtilizationInfo_t *utilInfo, int devId);
//func (h handle) deviceGetBoardUtilization() (u *uint, err error) {
//	var ret C.dcmiRet_t
//	var cndevUtilizationInfo C.cndevUtilizationInfo_t
//	cndevUtilizationInfo.version = C.int(VERSION)
//	devId := C.int(h.MINOR)
//	ret = C.cndevGetDeviceUtilizationInfo(&cndevUtilizationInfo, devId)
//	u = uintPtr(cndevUtilizationInfo.BoardUtilization)
//	return u, errorString(ret)
//}
//
////dcmiRet_t cndevGetProcessInfo(unsigned *infoCount, cndevProcessInfo_t *procInfo, int devId);
//func (h handle) deviceProcessInfo() ([]uint, []uint64, error) {
//	var ret C.dcmiRet_t
//	var cndevProcessInfo [szProcs]C.cndevProcessInfo_t
//	var infoCount = C.uint(szProcs)
//	cndevProcessInfo[0].version = C.int(VERSION)
//	devId := C.int(h.MINOR)
//	ret = C.cndevGetProcessInfo(&infoCount, &cndevProcessInfo[0], devId)
//	n := int(szProcs)
//	pids := make([]uint, n)
//	mems := make([]uint64, n)
//	for i := 0; i < n; i++ {
//		pids[i] = uint(cndevProcessInfo[i].pid)
//		// convert to MB
//		mems[i] = uint64(cndevProcessInfo[i].PhysicalMemoryUsed)/1024
//	}
//	return pids, mems, errorString(ret)
//}
//
//func processName(pid uint) (string, error) {
//
//	f := `/proc/` + strconv.FormatUint(uint64(pid), 10) + `/comm`
//	d, err := ioutil.ReadFile(f)
//
//	if err != nil {
//		if pid == 0{
//			return "", err
//		}
//		// TOCTOU: process terminated
//		if os.IsNotExist(err) {
//			return "", nil
//		}
//		return "", err
//	}
//	return strings.TrimSuffix(string(d), "\n"), err
//}

func uint64Ptr(c C.long) *uint64 {
	i := uint64(c)
	return &i
}
func uintPtr(c C.int) *uint {
	i := uint(c)
	return &i
}
