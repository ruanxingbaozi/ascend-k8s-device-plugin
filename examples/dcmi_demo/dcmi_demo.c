#include <stdio.h>
#include "dcmi_interface_api.h"

int main()
{
    int ret;
    int errcnt = 0;
    int temperature = 0;
    int device_id;
    int device_num = 0;
    int card_id;
    int card_num = 0;
    int card_list[8] = {0};

    dcmi_init();

    ret = dcmi_get_card_num_list(&card_num, card_list, sizeof(card_list));
    if(ret != 0) {
        printf("call dcmi_get_card_num_list fail!\n");
        return ret;
    }

    for (card_id = 0; card_id < card_num; card_id++) {
        ret = dcmi_get_device_num_in_card(card_list[card_id], &device_num);
        if(ret != 0) {
            printf("Get card %d device num fail!\n");
            errcnt++;
            continue;
        }

        for (device_id = 0; device_id < device_num; device_id++) {
            ret = dcmi_get_device_temperature(card_list[card_id], device_id, &temperature);
            if(ret != 0) {
                printf("Get card %d device %d temperature fail!\n");
                errcnt++;
                continue;
            }
            printf("card%d device%d temperature: %d C\n", card_list[card_id], device_id, temperature);
        }
    }

    return errcnt;
}