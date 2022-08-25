# Setup K8s cluster 

- on ACloudGuru

`kubectl --insecure-skip-tls-verify get nodes --kubeconfig ~/.kube/guru-con`

kubectl --insecure-skip-tls-verify exec --stdin --tty ambassador-test -c ambassador --kubeconfig ~/.kube/acloudguru -- /bin/bash


$ 172.31.113.134   k8s-control
$ 172.31.119.202   k8s-worker2
$ 172.31.126.131   k8s-worker1

vi hostpath-volume-test.yml

https://acloudguru-content-attachment-production.s3-accelerate.amazonaws.com/1642601759844-1077%20-%20S02L06%20Exploring%20Volumes.pdf

kubectl --insecure-skip-tls-verify scale deployment/nginx-deployment --replicas=4 --kubeconfig ~/.kube/guru-con

Exam Tips
A rolling update gradually rolls out changes to a Deployment’s Pod template by gradually replacing replicas with new ones.
Use kubectl rollout status to check the status of a rolling update.
Roll back the latest rolling update with: kubectl rollout undo .