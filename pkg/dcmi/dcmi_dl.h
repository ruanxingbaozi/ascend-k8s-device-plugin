// Copyright (c) 2015-2018, CAMBRICON CORPORATION. All rights reserved.

#ifndef _DCMI_DL_H_
#define _DCMI_DL_H_

#include "./include/dcmi_interface_api.h"

#define DCMI_DL(x) x##_dl

extern dcmiRet_t  DCMI_DL(dcmiInit)(void);
extern dcmiRet_t  DCMI_DL(dcmiShutdown)(void);
extern dcmiRet_t  DCMI_DL(dcmiRelease)(void);
#endif // _DCMI_DL_H_

