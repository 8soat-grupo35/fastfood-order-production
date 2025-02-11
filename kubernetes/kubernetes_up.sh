docker build -t fastfood-app:latest .
kubectl apply -f kubernetes/fastfood-deployment.yaml
kubectl apply -f kubernetes/fastfood-service.yaml
kubectl apply -f kubernetes/fastfood-hpa.yaml