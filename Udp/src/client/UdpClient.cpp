#include "client/UdpClient.h"

#include <cassert>
#include <unistd.h>
#include "util/Api.h"
#include <cstring>

UDPClient::UDPClient(
		const string &addr,
		const int port)
:
mStrServAddr(addr),
mSvrPort(port)
{
	bzero(&mServAddr, sizeof(mServAddr));
	mServAddr.sin_family = AF_INET;
	mServAddr.sin_port = htons(mSvrPort);

	int rslt = Inet_pton(AF_INET, mStrServAddr.data(), &mServAddr.sin_addr);
	assert(rslt);

	mSockfd = Socket(AF_INET, SOCK_DGRAM, 0);
	assert(mSockfd != -1);
}

UDPClient::~UDPClient()
{
	close();
}


int 
UDPClient::sendTo(char *data, const int64_t len)
{
	return SendTo(mSockfd, data, len, 0, (struct sockaddr *)&mServAddr, sizeof(mServAddr));
}


int 
UDPClient::recvFrom(char *&buff, const int64_t buff_len, int64_t data_len)
{
	return RecvFrom(mSockfd, buff, buff_len, 0, NULL, NULL);
}

int 
UDPClient::recvAck()
{
}

int 
UDPClient::close()
{
	::close(mSockfd);
}

