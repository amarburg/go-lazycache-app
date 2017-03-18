

# namespace :gcloud do
#   task :cred do
#     sh "gcloud config configurations activate lazycache-deis-dev"
#     sh "gcloud container  clusters get-credentials cluster-1"
#   end
# end

task :default => :build

task :build do
  sh(*%w( go build ))
end

task :test => :build do
  sh(*%w( go test -tags integration ))
end

task :run_local => :build do
  sh *%w( ./go-lazycache-app
  --port 8080
  --bind 127.0.0.1 )
end

task :run_google => :build do
  sh(*%w( ./go-lazycache-app
            --port 8080
            --image-store google
            --image-store-bucket ooi-camhd-analytics
            --bind 127.0.0.1 ))
end

tmp_image_store = '/tmp/image_store'

task :run_local_store => :build do
  mkdir(tmp_image_store) unless FileTest.directory? tmp_image_store

  sh(*%W( ./go-lazycache-app
            --port 8080
            --image-store local
            --image-local-root #{tmp_image_store}
            --image-url-root file://#{tmp_image_store}
            --bind 127.0.0.1 ))
end

task :run_local_redis => :build do
  mkdir(tmp_image_store) unless FileTest.directory? tmp_image_store

  sh(*%W( ./go-lazycache-app
            --port 8080
            --quicktime-store redis
            --directory-store redis
            --image-store local
            --image-local-root #{tmp_image_store}
            --image-url-root file://#{tmp_image_store}
            --bind 127.0.0.1 )
end)


namespace :docker do

  namespace :local do
    tag = "lazycache-app-local:latest"

    task :build => "test" do
      sh(*%W( docker build --no-cache --tag #{tag} . ))
    end

    task :run => :build do
      sh(*%W( docker run --rm --attach STDOUT --attach STDERR
      --publish 8080:8080
      --env LAZYCACHE_IMAGESTORE=local
      --env LAZYCACHE_IMAGESTORE_LOCALROOT=/srv/image_store
      --env LAZYCACHE_IMAGESTORE_URLROOT=file:///tmp/image_store
      --volume /tmp/image_store:/srv/image_store
      #{tag} ))
    end

  end


end
