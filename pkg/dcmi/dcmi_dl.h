// Copyright (c) 2015-2018, CAMBRICON CORPORATION. All rights reserved.

#ifndef _CNDEV_DL_H_
#define _CNDEV_DL_H_

#include "./include/dcmi_interface_api.h"

#define CNDEV_DL(x) x##_dl

extern cndevRet_t  CNDEV_DL(dcmi_init)(void);
extern cndevRet_t  CNDEV_DL(cndevShutdown)(void);
extern cndevRet_t  CNDEV_DL(cndevRelease)(void);
#endif // _CNDEV_DL_H_

