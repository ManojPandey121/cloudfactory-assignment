include src/.env

run:
	cd src && go build -o main *.go && \
	CSV=$(CSV) EMAIL_SENDER=$(EMAIL_SENDER) SENDER_PASSWORD=$(SENDER_PASSWORD) SMTP_HOST=$(SMTP_HOST) SMTP_PORT=$(SMTP_PORT) MAIL_TEMPLATE=$(MAIL_TEMPLATE) go run .

test:
	cd src && go test . -cover
