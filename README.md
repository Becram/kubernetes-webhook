#Kubernetes Webhook

USAGE:

##Create Cluster

First, we need to create a Kubernetes cluster:

```
❯ make cluster

🔧 Creating Kubernetes cluster...
kind create cluster --config dev/manifests/kind/kind.cluster.yaml
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.21.1) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Have a nice day! 👋
```

##Deploy Admission Webhook

```
❯ make deploy

📦 Building kubernetes-webhook Docker image...
docker build -t kubernetes-webhook:1.0 . 
                                        
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them

📦 Pushing admission-webhook image into Kind's Docker daemon...
kind load docker-image kubernetes-webhook:1.0
Image: "kubernetes-webhook:1.0" with ID "sha256:1358caa0f5bc712f4f431a58c033c84194e74971e7749122d97401942f38729b" not yet present on node "kind-control-plane", loading...

♻️  Deleting kubernetes-webhook deployment if existing...
kubectl delete -f dev/manifests/webhook/ || true
deployment.apps "kubernetes-webhook" deleted
service "kubernetes-webhook" deleted
secret "kubernetes-webhook-tls" deleted

⚙️  Applying cluster config...
kubectl apply -f dev/manifests/cluster-config/
namespace/apps unchanged
mutatingwebhookconfiguration.admissionregistration.k8s.io/kubernetes-webhook.acme.com configured
validatingwebhookconfiguration.admissionregistration.k8s.io/kubernetes-webhook.acme.com configured

🚀 Deploying kubernetes-webhook...
kubectl apply -f dev/manifests/webhook/
deployment.apps/kubernetes-webhook created
service/kubernetes-webhook created
secret/kubernetes-webhook-tls created
```

##Deploying Pods

```
❯ make pod
🚀 Deploying test pod...
kubectl apply -f dev/manifests/pods/test-pod.yaml
pod/test-pod created
```

#View Webhook service logs 

```
❯ make logs
```


#THANK YOU
