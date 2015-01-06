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
CMDGET		= go get
PRINTF		= /usr/bin/printf
RM		= /bin/rm -rf

SRCS		= $(SR)server.go			\
		  $(SR)queries.go			\
		  $(SR)models.go			\
		  $(SR)insertKeyword.go			\
		  $(SR)postHandler.go			\
		  $(SR)topRequestHandler.go

####################### - Makefile Rules - #######################

all		: $(NAME) ok

$(NAME)		: mapping
		@ $(PRINTF) "Executable name is: \033[33m$(NAME)\033[m\n"
		  $(CMDBUILD)  -v -o $(BINDIR)/$(NAME) $(SRCS)

run		: $(NAME)
		@ $(PRINTF) "Executable : \033[33m$(NAME)\033[m is now running\n"
		 ./$(BINDIR)/$(NAME)

mapping		:
		@ $(PRINTF) "Mapping of index mrdrive... : "
		@ curl -XPUT 127.0.0.1:9200/mrdrive/ --data-binary @conf/mapping.json;


clean		:
		@ $(RM) $(BINDIR)/$(NAME)
		@ $(PRINTF) "\033[32m[OK]\033[m CLEAN $(NAME)\n\n"

clean_all	: clean vendor_clean
		@ $(PRINTF) "\033[32m[OK]\033[m CLEAN ALL $(NAME)\n\n"

re		: clean $(NAME) ok

vendor_clean	:
		rm -rf ./_vendor/src

vendor_get	: vendor_clean
		$(CMDGET) -d -u -v			\
		github.com/belogik/goes 		\
		github.com/go-martini/martini		\
		github.com/opesun/nested

####################### - Make Utilities - #######################

ok	:
	@ $(PRINTF) "\033[32m[OK]\033[m Compilation done. \n\n"


.PHONY: clean run test clean_all vendor_clean vendor_get mapping
