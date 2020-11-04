DCMI_LIB_DIR = /user/local/dcmi

GCC      :=  gcc
CFLAGS   := -fPIC -fstack-protector-all -I$(DCMI_LIB_DIR)
LDFLAGS  := -ldl -L${DCMI_LIB_DIR} -ldcmi

SRC = $(wildcard *.c)
OBJ = $(patsubst %.c,%.o,$(SRC))

.PHONY: clean all

all: dcmi-demo

$(OBJ):
    ${GCC} -c $(SRC) ${CFLAGS}

dcmi-demo: $(OBJ)
    ${GCC} ${LDFLAGS} -o $@ $^

clean:
    rm -f *.o dcmi-demo