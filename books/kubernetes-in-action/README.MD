# package
docker run busybox echo "Hello world"
docker build -t kubik .
docker run --name kubik-container -p 8080:8080 -d kubik
docker exec -it kubik-container bash
docker tag kubik proz/kubik
docker run -p 8080:8080 -d proz/kubik

# run
minikube start
kubectl run kubik --image=proz/kubik --port=8080 --generator=run-pod/v1
kubectl expose rc kubik --type=LoadBalancer --name kubik-http
minikube service kubik-http

# examine
minikube dashboard

kubectl api-resources
kubectl scale rc kubik --replicas=3
kubectl get pods

kubectl get po kubia-zxzij -o yaml
kubectl logs kubik-manual

kubectl explain pods
kubectl explain pod.spec

# create manual
kubectl create -f kubik-manual.yaml
kubectl get po kubik-manual -o yaml

# port forwarding
kubectl port-forward kubik-manual 8888:8080

kubectl get po -L creation_method,env
kubectl label po kubia-manual creation_method=manual
kubectl get po -l creation_method=manual

kubectl annotate pod kubik-manual mycompany.com/someannotation="foo bar"