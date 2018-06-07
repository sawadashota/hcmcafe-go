serve:
	dev_appserver.py ./app/app.yaml --port 9000

deploy:
	gcloud app deploy ./app/app.yaml --project=hcm-cafe --quiet

browse:
	gcloud app browse --project=hcm-cafe

test:
	go test -v ./...
