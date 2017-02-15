

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
