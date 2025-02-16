docker build -t fastfood-order-production-app:latest .



kubectl create secret generic aws-secrets \
  --from-literal=access-key-id=$(aws configure get aws_access_key_id) \
  --from-literal=secret-access-key=$(aws configure get aws_secret_access_key) \
  --from-literal=access-session-token=$(aws configure get aws_session_token)
kubectl apply -f kubernetes/fastfood-order-production-deployment.yaml
kubectl apply -f kubernetes/fastfood-order-production-service.yaml
kubectl apply -f kubernetes/fastfood-order-production-hpa.yaml