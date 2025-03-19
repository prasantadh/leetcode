# this will always rebuild what is sent in as Q
# but that's alright. it's just a github code
all: 
	g++ --std=c++23 $(Q)*

clean:
	rm -rf a.out
