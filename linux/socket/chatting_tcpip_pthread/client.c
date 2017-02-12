/*
 * Goal: implement chat program using pthread and tcp/ip socket
 * files: server.c client.c chat.h chat.c
 * date: 2016-10-20 ~ 2016-10-21
 * author: yubi Lee (SKKU undergraduate)
 * github: https://github.com/eubnara
 * e-mail: eubnara@gmail.com
 *
 */

/*
 * file: client.c
 * desciption: client program in TCP/IP socket programming using pthread
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>	//pthread
#include <signal.h>		//signal
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include "chat.h"	//handler for error, thread, signal

//global variable also used in function "sig_killThread()" in "chat.c"
pthread_t talkThread = NULL;

int main(int argc, char* argv[]) {
	/* variables for socket programming */
	int sockfd, portno, n;
	struct sockaddr_in serv_addr;
	struct hostent *server; //Description of data base entry for a single host.
	char buffer[1024];

	/* register signal */
	signal(SIGINT, (void*)sig_killThread);

	if(argc != 3) {
		fprintf(stderr, "usage %s hostname port\n", argv[0]);
		exit(0);
	}
	portno = atoi(argv[2]);

	sockfd = socket(AF_INET, SOCK_STREAM, 0); 

	if(sockfd < 0) {
		error("ERROR opening socket");
	}

	server = gethostbyname(argv[1]);
	if(server == NULL) {
		fprintf(stderr, "ERROR, no such host\n");
		exit(0);
	}
	/* initialize struct sockaddr_in */
	memset(&serv_addr, 0, sizeof(serv_addr));
	//deprecated: bzero((char*)&serv_addr, sizeof(serv_addr));


	/* set struct sockaddr_in */
	serv_addr.sin_port = htons(portno);
	serv_addr.sin_family = AF_INET;
	memcpy(&serv_addr.sin_addr.s_addr, server->h_addr, server->h_length);
	//deprecated: bcopy((char*)server->h_addr, (char*)&serv_addr.sin_addr.s_addr,server->h_length);

	/* connect to server */
	if(connect(sockfd, (const struct sockaddr*) &serv_addr, sizeof(serv_addr)) <0){
		error("ERROR connecting");
	}
	/* make talk thread */
	pthread_create(&talkThread, NULL, talking, (void*)&sockfd);

	/* keep reading messages from server until this program interrupted. */
	while(1) {
		memset(buffer, 0, sizeof(buffer));
		n = read(sockfd, buffer, 1023);
		if(n < 0) {
			error("ERROR on read");
		}
		if(strlen(buffer) !=0)
			fprintf(stdout, "server:%s\n", buffer);
	}
	return 0;
}

