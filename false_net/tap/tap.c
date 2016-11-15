#include <sys/ioctl.h>
#include <net/if.h>

#include "tap.h"

int tun_alloc(char* dev)
{
    struct ifreq ifr;
    int fd, err;
    fd = open("/dev/net/tap", O_RDWR)
    if(fd < 0)
        print_error("Cannot open TAP device");

    CLEAR(ifr);

    ifr.ifr_flags = IFF_TAP | IFF_NO_PI;


}
