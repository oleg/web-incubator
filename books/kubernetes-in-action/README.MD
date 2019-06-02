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

# services
kubectl create -f kubik-svc.yaml
kubectl get pods

;;           <pod-name>                    <service cluster ip>
kubectl exec kubik-n9mqr -- curl -s http://10.96.64.184

kubectl delete po --all
kubectl exec kubia-3inly env


kubectl exec -it kubik-65spd bash
root@kubik-65spd:/# curl http://kubik.default.svc.cluster.local
root@kubik-65spd:/# curl http://kubik.default                  
You've hit kubik-ct45l
root@kubik-65spd:/# curl http://kubik        
You've hit kubik-ct45l


kubectl get endpoints kubik

minikube addons enable ingress
kubectl get po --all-namespaces
