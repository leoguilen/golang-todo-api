#-------------------
# Makefile Commands
#-------------------

all: set-env-vars build-docker apply-k8s

drop: remove-k8s unset-env-vars

#-------------------
# Steps
#-------------------

set-env-vars: ## Set environments variables
	@echo "Setting environment variables"
	@export DOCKER_REGISTRY=localhost:5000
	@export DOCKER_IMAGE=todo-api
	@export DOCKER_IMAGETAG=latest-dev
	@echo ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${DOCKER_IMAGETAG}

unset-env-vars: ## Unset environments variables
	@echo "Unsetting environment variables"
	@unset DOCKER_REGISTRY DOCKER_IMAGE DOCKER_IMAGETAG

build-docker: ## Build docker image and push to registry
	@echo "Building docker image"
	@docker builder build -t ${DOCKER_IMAGE}:${DOCKER_IMAGETAG} .
	@docker tag ${DOCKER_IMAGE}:${DOCKER_IMAGETAG} ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${DOCKER_IMAGETAG}
	@docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${DOCKER_IMAGETAG}

apply-k8s: ## Deploy kube artefacts
	@echo "Applying kube files"
	@kubectl apply -f ./.k8s/configmap.yaml
	@envsubst < ./.k8s/deployment.yaml | kubectl apply -f -
	@kubectl apply -f ./.k8s/service.yaml
	@kubectl apply -f ./.k8s/ingress.yaml
	@kubectl get po -l "app=todo-api"

remove-k8s: ## Remove kube deployments
	@echo "Removing kube deployments"
	@kubectl delete -f ./.k8s