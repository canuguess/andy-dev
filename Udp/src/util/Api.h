#ifndef UDP_UTIL_API
#define UDP_UTIL_API

#include <netinet/in.h>

bool Inet_pton(int family, const char *strptr, void *addrptr);
int  Socket(int family, int type, int protocol);
void Connect(int sockfd, const struct sockaddr *servaddr, socklen_t addrlen);


#endif // #ifndef UDP_UTIL_API
