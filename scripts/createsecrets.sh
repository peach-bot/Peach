kubectl create secret generic discord --from-file=./secrets/BOTTOKEN --from-file=./secrets/CLIENTID
kubectl create secret generic clustersecret --from-file=./secrets/CLUSTERSECRET