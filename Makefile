build:
	go build
zip:
	exifio exiftool.tar.gz
	tar -zxvf exiftool.tar.gz
	rm -rf exiftool
	mv Image-ExifTool-11.93 exiftool
	zip -r exifio.zip exifio exiftool
update:
 	make zip
	aws lambda update-function-code --function-name exifio --zip-file fileb://exifio.zip --region ap-northeast-1