#ifndef UDP_CLIENT_CLIENT
#define UDP_CLIENT_CLIENT

#include <netinet/in.h>
#include <stdio.h>
#include <stdlib.h>
#include <string>

using namespace std;

class UDPClient 
{
public:
	int 					mSockfd;
	struct 	sockaddr_in		mServAddr;
	string					mStrServAddr;
	int						mSvrPort;
	
public:
	UDPClient(const string &addr, int port);
	~UDPClient();

	int sendTo(char *data, const int64_t len);
	int recvFrom(char *&buff, const int64_t buff_len, int64_t data_len);
	int recvAck();
	int close();
};

#endif //#ifndef UDP_CLIENT_CLIENT
