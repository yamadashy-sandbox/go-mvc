.PHONY: all

project_id := ${PROJECT_ID}
version := ${GAE_VERSION}

## ローカルサーバ実行
serve:
	dev_appserver.py app/app.yaml

## ローカルサーバをDockerコンテナで実行する
serve-docker:
	$(eval src := $(shell pwd))
	@docker run --rm \
		-it  \
		-v ${src}:/work/src/github.com/ryutah/gae-structure-sample \
		-p 8080:8080 \
		-p 8000:8000 \
		ryutah/gcloud-gaego \
		dev_appserver.py --host 0.0.0.0 --admin_host 0.0.0.0 /work/src/github.com/ryutah/gae-structure-sample/app

## gaeへデプロイ OPTIONS: project_id=${PROJECT_ID} version=${VERSION}
deploy:
	gcloud app deploy -application ${project_id} -version ${version} app