#include <stdio.h>
#include <stdlib.h>
#include <getopt.h>
#include <unistd.h>
#include "dsmi_common_interface.h"
int main(int argc, char *argv[])
{
    int ret = 0;
    int opt = 0;
    int test_case_num = 0;
    char optstring[20] = "p:s:gr";
    static struct option long_options[] =
    {
        {"process", 1, NULL, 'p'},
        {0, 0, 0, 0}
    };
    while ((opt = getopt_long(argc, argv, optstring, long_options, NULL)) != -1)
    {
        switch (opt)
        {
            case 'p':
            {
                if (argc < 3)
                {
                    printf("process test the para num: %d test_error. Input num should not be smaller than 3.test_fail \n",argc);
                    return -1;
                }
                test_case_num = strtol(argv[2], NULL, 10);
                switch (test_case_num)
                {
                    case 1:
                    {
                        printf("begin query health\n");
                        unsigned int phealth = 0;
                        ret = dsmi_get_device_health(0, &phealth);
                        if (ret)
                        {
                            printf("call dsmi_get_device_health fail, ret = %d\n", ret);
                            return -1;
                        }
                        else
                        {
                            /*phealth type is unsigned int, printf value should use %u*/
                            printf("dsmi_get_device_health success,phealth:%u.\n", phealth);
                        }
                        break;
                    }
                    default:
                        break;
                }
            }
        }
    }
    return 0;
}