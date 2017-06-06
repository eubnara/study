#include <unistd.h>

int main() {

	execlp("/bin/ls", ".", (char*) 0);
	return 0;
}
