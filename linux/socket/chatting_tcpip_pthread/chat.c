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
 * file: chat.c
 * desciption: implementations for error handling, operation of thread, and signal handler.
 */

#include <stdio.h>	//scanf(), fprintf(), perror()
#include <stdlib.h>
#include <signal.h>	//SIGINT
#include <string.h>	//memset()
#include <unistd.h> //write()
#include <pthread.h>//pthread_cancel()
#include "chat.h"

extern int state;
extern pthread_t talkThread;

void error(char* msg) {
	perror(msg);
	exit(0);
}


void *talking(void *socket) {
	int sock = *(int *)socket;
	int n;
	char messageBuffer[1024];
	while(1) {
		memset(messageBuffer, 0, 1024);
		scanf("%s", messageBuffer);
		//fprintf(stdout, "me:%s\n",messageBuffer);

		n = write(sock, messageBuffer, sizeof(messageBuffer));
		if(n<0) 
			error("ERROR on write");
	}
}


void sig_killThread(int signo) {
	if(signo == SIGINT){
		 if ((pthread_cancel(talkThread)) != 0) {
		  	      /* Handle Error */
		  	 error("Error on pthread_cancel()");
		 }
		 exit(0);
	}
}
