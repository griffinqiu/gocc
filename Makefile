update\:data:
	# Fetch to update data from https://github.com/BYVoid/OpenCC
	mkdir -p ./tmp && rm -Rf tmp/OpenCC-master
	wget https://github.com/BYVoid/OpenCC/archive/refs/heads/master.zip -O tmp/opencc.zip
	unzip tmp/opencc.zip -d tmp/
	sh ./merge-data.sh

test:
	GO111MODULE=auto go test -cover ./...
