#include "test.h"

#include <stdio.h>

#ifdef __linux__

#include <sys/types.h>
#include <sys/ptrace.h>
#include <sys/wait.h>
#include <unistd.h>
#include <linux/user.h>

#endif

void run_in_c(char * p_message) {
	if(p_message) {
		fprintf(stdout, "IN C : %s\n", p_message);
	}
}

#ifdef __linux__

void test_ptrace(char * pid_str) {
	pid_t pid;
	struct user_regs_struct regs;

	long eip;

	pid = atoi(pid_str);

	ptrace(PTRACE_ATTACH, pid, NULL, NULL);

	wait(NULL);

    ptrace(PTRACE_GETREGS, pid, NULL, &regs);

    eip = ptrace(PTRACE_PEEKTEXT, pid, regs.eip, NULL);

    printf("EIP VALUE : %lx, INSTRUCTION : %lx\n", regs.eip, eip)

    ptrace(PTRACE_DETACH, pid, NULL, NULL);
}

#endif