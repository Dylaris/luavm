TARGET = $(PWD)/go/bin/luago

all: $(TARGET)
$(TARGET):
	export GOPATH=$(PWD)/go
	mkdir -p $(PWD)/go/bin
	go build -o $@ $(PWD)/go/src/luago/main.go	

clean: 
	rm -rf $(PWD)/go/bin

exec: clean $(TARGET)
	$(TARGET)

.PHONY: all clean
