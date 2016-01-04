#include "util/Api.h"

#include <arpa/inet.h>
#include <cstring>
#include <errno.h>
#include <iostream>
#include <sys/socket.h>

using namespace std;

bool 
Inet_pton(int family, const char *strptr, void *addeptr)
{
	int rslt = inet_pton(family, strptr, addeptr);
	if (rslt == 0)
	{
		std::cout << "[ERROR] bad addrs : " << strptr << endl;
		return false;
	}
	if (rslt != 1)
	{
		std::cout << "[ERROR] " << strerror(errno)<< endl;
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
		std::cout << "[ERROR] " << strerror(errno)<< endl;
		return -1;
	}
	return sock;
}

size_t
SendTo(int sockfd, const void *buff, size_t len, int flags, 
		   const struct sockaddr *to, socklen_t addrlen)
{
	return sendto(sockfd, buff, len, flags, to, addrlen);
}

size_t
RecvFrom(int sockfd, void *buff, size_t len, int flags,
			struct sockaddr *from, socklen_t *addrlen)
{
	return recvfrom(sockfd, buff, len, flags, from, addrlen);
}



