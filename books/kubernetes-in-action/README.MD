docker run busybox echo "Hello world"
docker build -t kubik .
docker run --name kubik-container -p 8080:8080 -d kubik
docker exec -it kubik-container bash
docker tag kubik proz/kubik
docker run -p 8080:8080 -d proz/kubik

kubectl run kubik --image=proz/kubik --port=8080 --generator=run/v1
kubectl expose rc kubik --type=LoadBalancer --name kubik-http
minikube service kubik-http

kubectl api-resources

kubectl scale rc kubik --replicas=3
kubectl get pods

minikube dashboard
