docker build -t fastfood-order-production-app:latest .
kubectl apply -f kubernetes/fastfood-order-production-deployment.yaml
kubectl apply -f kubernetes/fastfood-order-production-service.yaml
kubectl apply -f kubernetes/fastfood-order-production-fastfood-hpa.yaml