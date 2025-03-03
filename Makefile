TARGET = $(PWD)/go/bin/luago

all: $(TARGET)
$(TARGET):
	mkdir -p $(PWD)/go/bin
	go build -o $@ $(PWD)/go/src/luago/main.go	

clean: 
	rm -rf $(PWD)/go/bin

.PHONY: all clean
