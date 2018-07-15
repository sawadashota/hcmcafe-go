serve:
	dev_appserver.py ./app/app.yaml --port 8080 --support_datastore_emulator true --datastore_path .datastore

deploy:
	gcloud app deploy ./app/app.yaml --project=hcm-cafe --quiet

browse:
	gcloud app browse --project=hcm-cafe

test:
	go test -v ./...

admin-create:
	http http://localhost:8080/rpc @fixture/Admin.Create.json

admin-auth:
	http http://localhost:8080/rpc @fixture/Admin.Authenticate.json

admin-all:
	http http://localhost:8080/rpc @fixture/Admin.All.json

gcloud-set-project:
	gcloud config set project hcm-cafe
