#!/bin/bash

PKG_NAME="sensu-go-basic-checks"
PGK_VERSION="0.0.6"

ASSET_FILE=${PKG_NAME}_${PGK_VERSION}_windows_amd64.tar.gz
CHECKSUM_FILE=${PKG_NAME}_${PGK_VERSION}_sha512-checksums.txt

WORK_DIR=$(pwd)

export GOOS=windows
export GOARCH=amd64

#cleanup
rm -rf $WORK_DIR/bin

mkdir -p $WORK_DIR/bin

error=0
# Building sources
for file in sensu-go-basic-metrics go-check-plugins sensu-go-cpu-check sensu-go-memory-check; do

	if [ $file == "go-check-plugins" ];
	then
		for i in cert-file disk file-age file-size http log ping procs ssl-cert tcp uptime windows-eventlog; do \
			check_name=check-$i
		    echo "Building $check_name"
		    cd $WORK_DIR/upstream/$file/$check_name/
		    go build -o $WORK_DIR/bin/$check_name.exe main.go
		    if [ $? -ne 0  ]; then
		    	error=1
		    fi
		    cd $WORK_DIR
		done
	elif [ $file == "sensu-go-basic-metrics" ];
	then
		for i in $(ls $WORK_DIR/upstream/$file/windows); do
			metric_name=metric-$i
			echo "Building $metric_name"
			cd $WORK_DIR/upstream/$file/windows/$i
			go build -o $WORK_DIR/bin/$metric_name.exe main.go
		    if [ $? -ne 0  ]; then
			    error=1
			 fi
			cd $WORK_DIR
		done
	else
		tmp=${file##sensu-go-}
		check_name=check-${tmp%-check}
		echo "Building $check_name"
		cd $WORK_DIR/upstream/$file
		go build -o $WORK_DIR/bin/$check_name.exe main.go
	    if [ $? -ne 0  ]; then
		    error=1
		 fi
		cd $WORK_DIR
	fi
done
echo
echo "Built checks count: $(ls -l bin/  | wc -l)"
echo
if [ $error -ne 0  ]
then
	echo -e "\e[31mSome errors happens during GO build !!"
	echo -e "Cannot build the asset.\e[0m"
	echo
	exit 1
fi

tar -czf $ASSET_FILE bin/

if [ $? -eq 0 ]
then
	echo -e "\e[32mSuccessfully build sensu asset.\e[0m"
	echo
	echo "=====> Asset location:"
	readlink -f $ASSET_FILE
	echo
	echo "=====> Asset content:"
	tar -tvzf $ASSET_FILE
	echo
	echo "=====> Asset sha512sum:"
	checksum=$(sha512sum $ASSET_FILE)
	echo $checksum
	echo $checksum >> $CHECKSUM_FILE
	echo
fi
