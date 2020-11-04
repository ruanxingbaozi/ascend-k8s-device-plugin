#!/bin/bash
gcc get_device_health.c /usr/local/Ascend/driver/lib64/libdrvdsmi_host.so -L. -o get_device_health
./get_device_health -p 1
