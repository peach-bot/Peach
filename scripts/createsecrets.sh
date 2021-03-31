kubectl create secret generic database --from-file=./secrets/DATABASE
kubectl create secret generic clustersecret --from-file=./secrets/CLUSTERSECRET