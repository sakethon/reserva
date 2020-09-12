REGION = us-east-2
PROFILE = reserva

upload:
	GOOS=linux go build
	zip reserva.zip reserva
	aws lambda update-function-code --function-name reserva --zip-file fileb://reserva.zip --region ${REGION} --profile ${PROFILE}
