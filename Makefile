build:
	for dir in functions/*/ ; do \
		echo "Building" $${dir} ; \
		go get ./... ; \
		go build -o ../../$${dir} ./... ; \
	done

	hugo