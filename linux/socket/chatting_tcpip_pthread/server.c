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
 * file: server.c
 * desciption: server program in TCP/IP socket programming using pthread
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
	int sockfd, newsockfd, portno;
	socklen_t clilen;
	char buffer[1024];
	struct sockaddr_in serv_addr, cli_addr;
	int n;

	/* register signal */
	signal(SIGINT, (void*)sig_killThread);

	if(argc != 2) {
		fprintf(stderr, "ERROR, no port provided\n");
		exit(1);
	}

	sockfd = socket(AF_INET, SOCK_STREAM, 0); 

	if(sockfd < 0) {
		error("ERROR opening socket");
	}
	/* initialize struct sockaddr_in */
	memset(&serv_addr, 0, sizeof(serv_addr));
	//deprecated: bzero((char*)&serv_addr, sizeof(serv_addr));

	/* set struct sockaddr_in */
	portno = atoi(argv[1]);
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_port = htons(portno);
	serv_addr.sin_addr.s_addr = htonl(INADDR_ANY); // Address to accept any incoming messages. 

	/* bind */
	if(bind(sockfd, (struct sockaddr *) &serv_addr, sizeof(serv_addr)) <0)
		error("ERROR on binding");
	

	/* listen */
	if(listen(sockfd, 5) == -1)
		error("ERROR on listen");

	/* accept */
	clilen = sizeof(cli_addr);
	newsockfd = accept(sockfd, (struct sockaddr*)&cli_addr, &clilen); 
	if(newsockfd < 0) 
		error("ERROR on accept");
	
	/* make talk thread */
	pthread_create(&talkThread, NULL, talking, (void *)&newsockfd);

	/* keep reading messages from client until this program interrupted. */
	while(1) {
		memset(buffer, 0, sizeof(buffer));
		n = read(newsockfd, buffer, 1023);
		if(n < 0) {
			error("ERROR on read");
		}
		if(strlen(buffer)!=0)
			fprintf(stdout, "client:%s\n",buffer);
	}
	return 0;
}

