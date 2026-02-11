CLUSTER_NAME=pdp
NAMESPACE=pdp
IMAGE=pdp:local

.PHONY: up down install-rollouts build load deploy curl v2

up:
	kind create cluster --config deploy/kind/kind.yaml || true
	kubectl apply -f deploy/k8s/namespace.yaml

install-rollouts:
	kubectl create namespace argo-rollouts 2>/dev/null || true
	kubectl apply -n argo-rollouts -f https://raw.githubusercontent.com/argoproj/argo-rollouts/stable/manifests/install.yaml

build:
	docker build -t $(IMAGE) .

load:
	kind load docker-image $(IMAGE) --name $(CLUSTER_NAME)

deploy: install-rollouts build load
	kubectl apply -f deploy/k8s/service.yaml
	kubectl apply -f deploy/k8s/rollout.yaml

curl:
	curl -sS http://localhost:30080/ && echo
	curl -sS http://localhost:30080/healthz && echo

# v2に更新（失敗させたい場合は /error?rate= を叩く）
v2: build load
	kubectl -n $(NAMESPACE) set env rollout/pdp APP_VERSION=v2

down:
	kind delete cluster --name $(CLUSTER_NAME) || true
