[调用示例](https://support.huawei.com/enterprise/zh/doc/EDOC1100156981/36ef7b1d)  
Ascend EP场景下使用调用示例
以HwHiAiUser用户将get_device_health.c、dsmi_common_interface.h传到Host侧服务器的同一个目录下。
以HwHiAiUser用户登录到Host侧服务器。
执行如下命令，编译get_device_health.c中的代码，生成可执行文件get_device_health。
gcc get_device_health.c /usr/local/Ascend/driver/lib64/libdrvdsmi_host.so -L. -o get_device_health

/usr/local/Ascend表示Driver组件的默认安装路径，请根据实际情况替换。
Driver组件的安装，请参见《驱动和开发环境安装指南CANN软件安装指南》。
执行可执行文件get_device_health。
./get_device_health -p 1