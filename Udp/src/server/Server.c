#include "server/Server.h"

#include "util/Api.h"

UdpServer::UdpServer(const string addr, const size_t port)
{
}

UdbServer::~UdpServer()
{
	close();
}

int
UdpServer::sendTo(char *buff, const size_t len);
{
}

int 
UdpServer::recvFrom(char *buff, const size_t len)
{
}

int 
UdpServer::close()
{
	::close(mSockFd);
}
