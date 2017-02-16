

namespace :gcloud do
  task :cred do
    sh "gcloud config configurations activate lazycache-deis-dev"
    sh "gcloud container  clusters get-credentials cluster-1"
  end
end

task :build do
  sh *%w( go build )
end

task :run_google => :build do
    sh *%w( ./go-lazycache-app
            --port 8080
            --image-store google
            --image-store-bucket ooi-camhd-analytics
            --bind 127.0.0.1 )

end

task :run_local => :build do
    sh *%w( ./go-lazycache-app
            --port 8080
            --bind 127.0.0.1 )

end

task :test => :build do
    sh *%w( go test -tags integration )
end

namespace :prom do
  task :server => "gcloud:cred" do
  # Get the server URL by running these commands in the same shell:
  #   export POD_NAME=$(kubectl get pods --namespace default -l "app=prom-prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  #   kubectl --namespace default port-forward $POD_NAME 9090


  pod_name = `kubectl get pods --namespace default -l "app=prom-prometheus,component=server" -o jsonpath="{.items[0].metadata.name}"`
  sh "kubectl --namespace default port-forward #{pod_name} 9090"

end
  # The Prometheus alertmanager can be accessed via port 80 on the following DNS name from within your cluster:
  # prom-alerts.default.svc.cluster.local
  #
  #
  # Get the server URL by running these commands in the same shell:
  #   export POD_NAME=$(kubectl get pods --namespace default -l "app=prom-prometheus,component=alertmanager" -o jsonpath="{.items[0].metadata.name}")
  #   kubectl --namespace default port-forward $POD_NAME 9093

end
