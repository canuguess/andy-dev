#include "client/UdpClient.h"

#include <unistd.h>
#include "util/Api.h"
#include <cassert>

UDPClient::UDPClient(
		const string &addr,
		const int port)
:
mStrServAddr(addr),
mSvrPort(port)
{
	bzeor(&mServAddr, sizeof(mServAddr));
	mServAddr.sin_family = AF_INET;
	mServAddr.sin_port = htons(mSvrPort);

	int rslt = Inet_pton(AF_INET, mStrServAddr.data(), mServAddr.sin_addr);
	assert(rslt);

	mSockfd = Sockte(AF_INET, SOCK_DGTAM, 0);
	assert(mSockfd != -1);
}

UDPClient::~UDPClient()
{
	close();
}


int 
UDPClient::close()
{
	close(mSockfd);
}

