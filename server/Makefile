serve:
	dev_appserver.py ./app/app.yaml --port 9000

deploy:
	gcloud app deploy ./app/app.yaml --project=hcm-cafe --quiet

browse:
	gcloud app browse --project=hcm-cafe

test:
	go test -v ./...

admin-create:
	http http://localhost:9000/rpc @fixture/Admin.Create.json

admin-auth:
	http http://localhost:9000/rpc @fixture/Admin.Authenticate.json

admin-all:
	http http://localhost:9000/rpc @fixture/Admin.All.json

gcloud-set-project:
	gcloud config set project hcm-cafe
