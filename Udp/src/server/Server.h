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
	int					mSockFd;

public:
	UdpServer(const string addr, const size_t port);
	~UdpServer();

	int sendTo(char *buff, const size_t len);
	int recvFrom(char *&buff, const size_t len);

	void close();

}; 

#endif;
