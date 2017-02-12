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
 * file: chat.h
 * desciption: prototypes for error handling, operation of thread, and signal handler.
 */

void error(char* msg); 			//error handler
void* talking(void *socket);	//talking to receiver by thread
void sig_killThread(int signo);	//SIGINT handler for termination
