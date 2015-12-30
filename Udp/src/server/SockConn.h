#ifndef UDP_SERVER_SOCKCONN
#define UDP_SERVER_SOCKCONN

#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

class UdpServer
{
private;
	struct sockaddr_in 	mToSock;
	struct sockaddr_in 	mFromSock;
	int					mSockFd;
public:
	UdpServer(const struct sockaddr_in to_sock, const struct sockaddr_in from_sock);
	~UdpServer();
	int sendTo(char *buff, const size_t len);
	int recvFrom()

}; 

#endif;
