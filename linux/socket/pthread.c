#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>


void *myfunc (void *myvar);

int main(void)
{

	pthread_t thread1=NULL, thread2=NULL;
	char *msg1 = "First thread";
	char *msg2 = "Second thread";
	int ret1, ret2;

	

	ret1 = pthread_create(&thread1, NULL, myfunc, (void *) msg1);
	ret2 = pthread_create(&thread2, NULL, myfunc, (void *) msg2);


	printf("Main function after pthread_create\n");



	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);
	
	printf("First thread ret1 = %d\n", ret1);
	printf("Second thread ret2 = %d\n", ret2);

	

	return 0;
}


void *myfunc (void *myvar) {
	char *msg;
	msg = (char *) myvar;

	int i;
	for (i = 0; i < 10; i++) {
		printf("%s %d\n", msg, i);
		sleep(1);
	}
	return NULL;
}
