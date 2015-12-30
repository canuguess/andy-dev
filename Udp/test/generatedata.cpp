#include <stdio.h>
#include <stdlib.h>
#include <string>

using namespace std;

#define u64 unsigned long long

u64 FILELEN = 1024*1024*1024; // 1G 
string FILENAME = "./udptestdata_10g";

int main()
{
	u64 count = 0;
	FILE *fd = ::fopen(FILENAME.data(), "w+b");
	string data = "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111";
	while (1)
	{
		::fwrite(data.data(), 1, data.length(), fd);
		count+= data.length();
		if (count > FILELEN) break;
	}
	::fclose(fd);
	return 0;
}
