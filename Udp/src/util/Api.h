#ifndef UDP_UTIL_API
#define UDP_UTIL_API

#include <netinet/in.h>
#include <sys/socket.h>

bool Inet_pton(int family, const char *strptr, void *addrptr);
int  Socket(int family, int type, int protocol);
void Connect(int sockfd, const struct sockaddr *servaddr, socklen_t addrlen);

size_t SendTo(int sockfd, const void *buff, size_t len, int flags, 
		   const struct sockaddr *to, socklen_t addrlen);

size_t RecvFrom(int sockfd, void *buff, size_t len, int flags,
			struct sockaddr *from, socklen_t *addrlen);


#endif // #ifndef UDP_UTIL_API
