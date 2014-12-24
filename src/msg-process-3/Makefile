##
## Makefile for server
##
## Made by Fares Sabbagh
##
## Started on  Tue Dec 23 16:58:05 2014 Fares Sabbagh
##
##

####################### - Project Config - #######################

NAME		= server
BINDIR		= bin
SR		= ./srcs/
CMDBUILD	= go build
PRINTF		= /usr/bin/printf
RM		= /bin/rm -rf

SRCS		= $(SR)server.go	\
		  $(SR)queries.go	\
		  $(SR)models.go

####################### - Makefile Rules - #######################

all		: $(NAME) ok

$(NAME)		:
		@ $(PRINTF) "Executable name is: \033[33m$(NAME)\033[m\n"
		  $(CMDBUILD)  -v -o $(BINDIR)/$(NAME) $(SRCS)

run		: $(NAME)
		@ $(PRINTF) "Executable : \033[33m$(NAME)\033[m is now running\n"
		 ./$(BINDIR)/$(NAME)


clean		:
		@ $(RM) $(BINDIR)/$(NAME)
		@ $(PRINTF) "\033[32m[OK]\033[m CLEAN $(NAME)\n\n"

re		: clean $(NAME) ok


####################### - Make Utilities - #######################

ok	:
	@ $(PRINTF) "\033[32m[OK]\033[m Compilation done. \n\n"


.PHONY: clean run test
