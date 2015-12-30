#include "util/Api.h"

#include <arpa/inet.h>
#include <sys/socket.h>

using namespace std;

bool 
Inet_pton(int family, const char *strptr, void *addeptr)
{
	int rslt = inet_pton(family, strptr, addeptr);
	if (rslt == 0)
	{
		std::cout << "[ERROR] bad addrs : " << strptr << end;
		return false;
	}
	if (rslt != 1)
	{
		std::cout << "[ERROR] " << strerror(errno)<< end;
		return false;
	}
	return true;
}

int 
Socket(int family, int type, int protocol)
{
	int sock = socket(family, type, protocol);
	if (sock < 0)
	{
		std::cout << "[ERROR] " << strerror(errno)<< end;
		return -1;
	}
	return sock;
}

