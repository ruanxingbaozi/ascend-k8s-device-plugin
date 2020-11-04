/*
* Copyright (c) Huawei Technologies Co., Ltd. 2018-2020. All rights reserved.
* Description: dcmi interface api
* Author: wangchenglong w00345862
* Create: 2020-03-19
*/

#ifndef _DCMI_TRAINING_INTERFACE_H_
#define _DCMI_TRAINING_INTERFACE_H_

#define SENSOR_TEMP_LEN     2
#define SENSOR_NTC_TEMP_LEN 4
#define SENSOR_DATA_MAX_LEN 16
#define DIEID_INFO_LENTH    5

struct dsmi_ecc_info_stru {
    int enable_flag;
    unsigned int single_bit_error_count;
    unsigned int double_bit_error_count;
};

struct tag_pcie_idinfo {
    unsigned int deviceid;
    unsigned int venderid;
    unsigned int subvenderid;
    unsigned int subdeviceid;
    unsigned int bdf_deviceid;
    unsigned int bdf_busid;
    unsigned int bdf_funcid;
};

struct dm_flash_info_stru {
    unsigned long flash_id;         // combined device & manufacturer code
    unsigned short device_id;       // device id
    unsigned short vendor;          // the primary vendor id
    unsigned int state;             // flash health, 0x8:normal,0x10:abnormal
    unsigned long size;             // total size in bytes
    unsigned int sector_count;      // number of erase units
    unsigned short manufacturer_id; // manufacturer id
};

struct dsmi_memory_info_stru {
    unsigned long memory_size; // KB
    unsigned int freq;
    unsigned int utiliza;
};

struct dsmi_hbm_info_stru {
    unsigned long memory_size;  // HBM total size, KB
    unsigned int freq;          // HBM freq, MHZ
    unsigned long memory_usage; // HBM memory_usage, KB
    int temp;                   // HBM temperature
    unsigned int bandwith_util_rate;
};

struct dsmi_aicore_info_stru {
    unsigned int freq;          // normal freq
    unsigned int curfreq;       // current freq
};

struct dsmi_soc_die_stru {
    unsigned int soc_die[DIEID_INFO_LENTH];
};

struct dsmi_board_info_stru {
    unsigned int board_id;
    unsigned int pcb_id;
    unsigned int bom_id;
    unsigned int slot_id;
};

union tag_sensor_info {
    unsigned char uchar;
    unsigned short ushort;
    unsigned int uint;
    signed int iint;
    signed char temp[SENSOR_TEMP_LEN];
    signed int ntc_tmp[SENSOR_NTC_TEMP_LEN];
    unsigned int data[SENSOR_DATA_MAX_LEN];
};

enum {
    DCMI_MEMORY_TYPE = 1,     // memory
    DCMI_AICORE_TYPE,         // aicore
    DCMI_AICPU_TYPE,          // ai cpu
    DCMI_CTRLCPU_TYPE,        // ctrl cpu
    DCMI_MEMORY_BANDWIDTH,    // memory bandwidth
    DCMI_HBM_TYPE,            // hbm
    DCMI_HBM_BANDWIDTH = 10,  // hbm bandwidth
};

// ecc check type
enum {
    DCMI_DEVICE_TYPE_DDR,
    DCMI_DEVICE_TYPE_SRAM,
    DCMI_DEVICE_TYPE_HBM,
    DCMI_DEVICE_TYPE_NPU,
    DCMI_DEVICE_TYPE_NONE = 0xff
};

typedef enum {
  DCMI_SUCCESS = 0,                       /**< No error */
  DCMI_ERROR_NO_DRIVER = 1,               /**< No Driver */
  DCMI_ERROR_LOW_DRIVER_VERSION = 2,      /**< Driver Version Low*/
  DCMI_ERROR_UNSUPPORTED_API_VERSION = 3, /**< API Version is not support*/
  DCMI_ERROR_UNINITIALIZED = 4,           /**< API not Initial*/
  DCMI_ERROR_INVALID_ARGUMENT = 5,        /**< Invalid pointer*/
  DCMI_ERROR_INVALID_DEVICE_ID = 6,       /**< Invalid device id*/
  DCMI_ERROR_UNKNOWN = 7,                 /**< Unknown error*/
  DCMI_ERROR_MALLOC = 8,                  /**< internal malloc fail*/
  DCMI_ERROR_NOT_SUPPORTED = 10           /**< not supported*/
} dcmiRet_t;

int dcmi_init(void);
int dcmi_get_card_num_list(int *card_num, int *card_list, int list_length);
int dcmi_get_device_num_in_card(int card_id, int *device_num);
int dcmi_get_device_health(int card_id, int device_id, unsigned int *phealth);
int dcmi_get_device_errorcode(int card_id, int device_id, int *errorcount, unsigned int *perrorcode,
                              int *errorwidth);
int dcmi_get_device_errorinfo(int card_id, int device_id, int errorcode, unsigned char *perrorinfo, int buffsize);
int dcmi_get_ecc_info(int card_id, int device_id, int device_type, struct dsmi_ecc_info_stru *pdevice_ecc_info);
int dcmi_get_device_temperature(int card_id, int device_id, int *tmprature);
int dcmi_get_device_power_info(int card_id, int device_id, int *power);
int dcmi_get_pcie_info(int card_id, int device_id, struct tag_pcie_idinfo *pcie_idinfo);
int dcmi_get_device_voltage(int card_id, int device_id, unsigned int *pvoltage);
int dcmi_get_device_utilization_rate(int card_id, int device_id, int subdevice_type,
                                     unsigned int *putilization_rate);
int dcmi_get_device_frequency(int card_id, int device_id, int device_type, unsigned int *pfrequency);
int dcmi_get_device_flash_count(int card_id, int device_id, unsigned int *pflash_count);
int dcmi_get_device_flash_info(int card_id, int device_id, unsigned int flash_index,
                               struct dm_flash_info_stru *pflash_info);
int dcmi_get_memory_info(int card_id, int device_id, struct dsmi_memory_info_stru *pdevice_memory_info);
int dcmi_get_hbm_info(int card_id, int device_id, struct dsmi_hbm_info_stru *pdevice_hbm_info);
int dcmi_get_aicore_info(int card_id, int device_id, struct dsmi_aicore_info_stru *pAicore_info);
int dcmi_get_device_die(int card_id, int device_id, struct dsmi_soc_die_stru *pdevice_die);
int dcmi_get_device_ndie(int card_id, int device_id, struct dsmi_soc_die_stru *pdevice_ndie);
int dcmi_get_soc_sensor_info(int card_id, int device_id, int sensor_id, union tag_sensor_info *sensor_info);
int dcmi_get_system_time(int card_id, int device_id, unsigned int *time);
int dcmi_get_board_info(int card_id, int device_id, struct dsmi_board_info_stru *pboard_info);
int dcmi_get_version(int card_id, int device_id, char *verison_str, unsigned int version_len, int *len);
#endif /* _DCMI_TRAINING_INTERFACE_H_ */